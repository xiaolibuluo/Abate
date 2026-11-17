
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome21870<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21870<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify21871<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 22 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops21872(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 29;
        break;
    }
    total
}
pub fn parse21873(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 36)
}
pub fn peek21874(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 43
}
pub fn peek21875(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 50
}
pub trait Encode21876 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21876<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21876 for Packet21876<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify21877<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 64 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek21878(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 71
}
pub fn peek21879(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 78
}
pub fn peek218710(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 85
}
