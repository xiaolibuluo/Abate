
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse10460(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 23)
}
pub fn peek10461(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 30
}
pub fn classify10462<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 37 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome10463<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10463<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse10464(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 51)
}
#[derive(Debug, Clone)]
pub enum Outcome10465<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10465<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek10466(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 65
}
