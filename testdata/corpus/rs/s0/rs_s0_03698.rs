
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum13570 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13570(values: &[i32]) -> i32 {
    let literal = sum13570!(89, 53);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek13571(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 96
}
macro_rules! sum13572 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13572(values: &[i32]) -> i32 {
    let literal = sum13572!(6, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode13573 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13573<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13573 for Packet13573<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse13574(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 20)
}
pub fn classify13575<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 27 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode13576 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13576<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13576 for Packet13576<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode13577 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13577<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13577 for Packet13577<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek13578(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 48
}
