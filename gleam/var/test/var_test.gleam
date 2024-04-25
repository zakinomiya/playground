import gleeunit
import gleeunit/should
import var

pub fn main() {
  gleeunit.main()
}

pub fn parse_key_test() {
  var.parse_key("hello\": \"world\"", [], 0)
  |> should.equal(Ok(#("hello", ": \"world\"", 5)))
}
