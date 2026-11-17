
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20300<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20300<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify20301<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 18 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse20302(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 25)
}
pub fn parse20303(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 32)
}
pub fn peek20304(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
pub trait Encode20305 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20305<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20305 for Packet20305<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse20306(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 53)
}
pub fn parse20307(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 60)
}
pub fn peek20308(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 67
}
#[derive(Debug, Clone)]
pub enum Outcome20309<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20309<T> {
    pub value: T,
    pub tag: u32,
}
