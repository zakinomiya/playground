import argv
import gleam/dict.{type Dict}
import gleam/int
import gleam/io
import gleam/list
import gleam/option.{type Option, None, Some}
import gleam/result
import gleam/string

pub type JsonNumber {
  Integer(value: Int)
  Float(value: Float)
}

type JsonField =
  Dict(String, JsonValue)

pub type JsonValue {
  JsonObject(key: String, value: JsonField)
  JsonArray(key: String, value: List(JsonValue))
  JsonString(key: String, value: String)
  JsonNumber(key: String, value: JsonNumber)
  JsonBool(key: String, value: Bool)
  JsonNull(key: String)
}

fn print_json(self: JsonValue) -> Nil {
  io.debug(self)
  Nil
}

pub type JsonParseError {
  UnexpectedCharacter(expected: Option(String), given: String, position: Int)
}

pub fn main() {
  case argv.load().arguments {
    ["parse", jsonv] ->
      case parse(jsonv, 0) {
        Ok(json) -> print_json(json)
        Error(err) -> show_error(err)
      }
    _ -> print_usage()
  }
}

fn print_usage() {
  io.print("Usage: <json>")
}

fn parse(str: String, position: Int) -> Result(JsonValue, JsonParseError) {
  case str {
    " " <> _ -> parse(str, position + 1)
    "{" <> rest ->
      case parse_field(dict.new(), rest, position + 1) {
        Ok(#(fields, _, _)) -> Ok(JsonObject("", fields))
        Error(err) -> Error(err)
      }
    "[" <> rest ->
      case parse_array(list.new(), rest, position + 1) {
        Ok(#(lists, _, _)) -> Ok(JsonArray("", lists))
        Error(err) -> Error(err)
      }
    _ -> Error(UnexpectedCharacter(None, str, position))
  }
}

fn parse_field(
  fields: JsonField,
  str: String,
  position: Int,
) -> Result(#(JsonField, String, Int), JsonParseError) {
  case str {
    " " <> rest -> parse_field(fields, rest, position + 1)
    "," <> rest -> parse_field(fields, rest, position + 1)
    "\n" <> rest -> parse_field(fields, rest, position + 1)
    "\r" <> rest -> parse_field(fields, rest, position + 1)

    "\"" <> rest -> {
      use #(key, rest, pos) <- result.try(parse_key(rest, [], position))
      use rest <- result.try(ensure_next(rest, ":"))
      use #(value, rest, pos) <- result.try(parse_value(key, rest, pos))
      dict.insert(fields, key, value)
      use rest <- result.map(ensure_next(rest, ","))
      #(fields, rest, pos)
    }
    _ -> Error(UnexpectedCharacter(None, str, position))
  }
}

fn ensure_next(str: String, character: String) -> Result(String, JsonParseError) {
  case string.pop_grapheme(str) {
    Ok(#(" ", rest)) -> ensure_next(rest, character)
    Ok(#("\n", rest)) -> ensure_next(rest, character)
    Ok(#(head, rest)) if head == character -> Ok(rest)
    Ok(#(head, _)) -> Error(UnexpectedCharacter(Some(character), head, 0))
    Error(_) -> Error(UnexpectedCharacter(Some(character), "", 0))
  }
}

pub fn parse_key(
  str: String,
  key: List(String),
  position: Int,
) -> Result(#(String, String, Int), JsonParseError) {
  case string.pop_grapheme(str) {
    Ok(#("\"", rest)) -> Ok(#(string.join(key, ""), rest, position))
    Ok(#(head, rest)) -> parse_key(rest, list.append(key, [head]), position + 1)
    Error(_) -> Error(UnexpectedCharacter(None, str, position))
  }
}

fn parse_value(
  key: String,
  str: String,
  position: Int,
) -> Result(#(JsonValue, String, Int), JsonParseError) {
  case str {
    " " <> rest -> parse_value(key, rest, position + 1)
    "{" <> rest ->
      case parse_field(dict.new(), rest, position + 1) {
        Ok(#(fields, rest, pos)) -> Ok(#(JsonObject(key, fields), rest, pos))
        Error(err) -> Error(err)
      }
    "[" <> rest ->
      case parse_array(list.new(), rest, position + 1) {
        Ok(#(lists, rest, pos)) -> Ok(#(JsonArray(key, lists), rest, pos))
        Error(err) -> Error(err)
      }
    "\"" <> rest ->
      case parse_string(list.new(), str, position) {
        Ok(value) -> Ok(#(JsonString(key, value), rest, string.length(value)))
        Error(err) -> Error(err)
      }
    "null" <> rest -> Ok(#(JsonNull(key), rest, position + 4))
    "false" <> rest -> Ok(#(JsonBool(key, False), rest, position + 5))
    "true" <> rest -> Ok(#(JsonBool(key, True), rest, position + 4))
    _ -> Error(UnexpectedCharacter(None, str, position))
  }
}

fn parse_string(
  value: List(String),
  str: String,
  position: Int,
) -> Result(String, JsonParseError) {
  case string.pop_grapheme(str) {
    Ok(#("\"", rest)) -> Ok(string.join(value, ""))
    Ok(#("\\", rest)) ->
      case string.pop_grapheme(rest) {
        Ok(#(head, rest)) ->
          parse_string(list.append(value, [head]), rest, position + 2)
        Error(_) -> Error(UnexpectedCharacter(None, str, position))
      }
    Ok(#(head, rest)) ->
      parse_string(list.append(value, [head]), rest, position + 1)
    Error(_) -> Error(UnexpectedCharacter(None, str, position))
  }
}

fn parse_array(
  list: List(JsonValue),
  str: String,
  position: Int,
) -> Result(#(List(JsonValue), String, Int), JsonParseError) {
  todo
}

fn show_error(err: JsonParseError) {
  case err {
    UnexpectedCharacter(None, c, p) ->
      io.print(
        "Unexpected character '"
        <> result.unwrap(string.first(c), "")
        <> "' at position "
        <> int.to_string(p),
      )
    UnexpectedCharacter(Some(s), c, p) ->
      io.print(
        "Unexpected character where '"
        <> s
        <> "' is expected. Got '"
        <> result.unwrap(string.first(c), "")
        <> "' at position "
        <> int.to_string(p),
      )
  }
}
