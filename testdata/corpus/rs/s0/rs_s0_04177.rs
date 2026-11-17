
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15360<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15360<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode15361 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15361<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15361 for Packet15361<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse15362(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 5)
}
pub fn classify15363<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 12 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek15364(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
pub fn loops15365(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 26;
        break;
    }
    total
}
