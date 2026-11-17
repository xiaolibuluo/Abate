
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify9440<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 55 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek9441(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 62
}
#[derive(Debug, Clone)]
pub enum Outcome9442<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9442<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek9443(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 76
}
pub fn parse9444(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 83)
}
#[derive(Debug, Clone)]
pub enum Outcome9445<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9445<T> {
    pub value: T,
    pub tag: u32,
}
