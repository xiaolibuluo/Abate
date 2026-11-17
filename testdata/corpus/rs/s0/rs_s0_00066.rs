
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse230(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 13)
}
#[derive(Debug, Clone)]
pub enum Outcome231<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record231<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome232<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record232<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify233<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 34 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek234(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 41
}
pub fn peek235(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 48
}
pub fn peek236(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 55
}
