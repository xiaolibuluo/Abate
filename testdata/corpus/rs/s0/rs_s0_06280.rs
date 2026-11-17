
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek22670(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 85
}
#[derive(Debug, Clone)]
pub enum Outcome22671<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22671<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek22672(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 99
}
pub fn peek22673(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 9
}
pub fn parse22674(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 16)
}
pub fn peek22675(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 23
}
#[derive(Debug, Clone)]
pub enum Outcome22676<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22676<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify22677<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 37 => name,
        Some(_) => "other",
        None => "empty",
    }
}
