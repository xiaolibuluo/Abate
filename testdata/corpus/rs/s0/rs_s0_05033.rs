
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18200<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18200<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse18201(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 101)
}
pub fn peek18202(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 11
}
pub fn peek18203(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 18
}
pub fn parse18204(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 25)
}
#[derive(Debug, Clone)]
pub enum Outcome18205<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18205<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify18206<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 39 => name,
        Some(_) => "other",
        None => "empty",
    }
}
