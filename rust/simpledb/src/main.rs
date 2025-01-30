use std::{collections::HashMap, error::Error};

struct DB {
    inner: HashMap<String, String>,
}

impl DBTrait for DB {
    fn get(&self, key: &str) -> Option<&String> {
        self.inner.get(key)
    }

    fn set(&mut self, key: String, value: String) {
        self.inner.insert(key, value);
    }

    fn del(&mut self, key: &str) {
        self.inner.remove(key);
    }

    fn dump(&self) {
        for (key, value) in &self.inner {
            println!("{}: {}", key, value);
        }
    }
}

trait DBTrait {
    fn get(&self, key: &str) -> Option<&String>;
    fn set(&mut self, key: String, value: String);
    fn del(&mut self, key: &str);
    fn dump(&self);
}

fn exec<T: DBTrait>(db: &mut T, input: &str) -> Result<(), Box<dyn Error>> {
    let arr = input.split_whitespace().collect::<Vec<&str>>();
    let cmd = arr[0];
    match Command::from(cmd) {
        Command::Get => {
            if arr.len() != 2 {
                return Err("Invalid number of arguments".into());
            }
            let val = db.get(arr[1]);
            println!("{:?}", val);
        }
        Command::Set => {
            if arr.len() != 3 {
                return Err("Invalid number of arguments".into());
            }
            db.set(arr[1].to_string(), arr[2].to_string());
            db.dump();
        }
        Command::Del => {
            if arr.len() != 2 {
                return Err("Invalid number of arguments".into());
            }
            db.del(arr[1]);
            db.dump();
        }
        Command::Exit => {
            return Ok(());
        }
        Command::Unknown => {
            return Err("Unknown command".into());
        }
    }
    return Ok(());
}

fn main() {
    let stdin = std::io::stdin();
    let mut db = DB {
        inner: HashMap::new(),
    };

    loop {
        let mut input = String::new();
        match stdin.read_line(&mut input) {
            Ok(_) => match exec(&mut db, input.as_str().trim()) {
                Ok(_) => {}
                Err(e) => {
                    eprintln!("Error: {}", e);
                }
            },
            Err(e) => {
                eprintln!("Error: {}", e);
                std::process::exit(1);
            }
        }
    }
}

#[derive(Debug)]
enum Command {
    Unknown,
    Get,
    Set,
    Del,
    Exit,
}

impl Command {
    pub fn from(input: &str) -> Command {
        match input {
            "get" => Command::Get,
            "set" => Command::Set,
            "del" => Command::Del,
            "exit" => Command::Exit,
            _ => Command::Unknown,
        }
    }
}
