
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse18380(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 37)
}
pub fn classify18381<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 44 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek18382(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 51
}
pub fn parse18383(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 58)
}
#[derive(Debug, Clone)]
pub enum Outcome18384<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18384<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome18385<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18385<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek18386(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 79
}
pub fn parse18387(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 86)
}
#[derive(Debug, Clone)]
pub enum Outcome18388<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18388<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse18389(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 100)
}
pub fn parse183810(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 10)
}
