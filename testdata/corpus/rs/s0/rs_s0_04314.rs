
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek15810(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 91
}
#[derive(Debug, Clone)]
pub enum Outcome15811<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15811<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify15812<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 8 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse15813(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 15)
}
pub fn peek15814(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 22
}
#[derive(Debug, Clone)]
pub enum Outcome15815<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15815<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify15816<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 36 => name,
        Some(_) => "other",
        None => "empty",
    }
}
