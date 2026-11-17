
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome25180<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25180<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify25181<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 57 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome25182<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25182<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek25183(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 71
}
pub fn peek25184(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 78
}
pub fn parse25185(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 85)
}
pub fn parse25186(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 92)
}
pub fn peek25187(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 99
}
#[derive(Debug, Clone)]
pub enum Outcome25188<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25188<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome25189<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25189<T> {
    pub value: T,
    pub tag: u32,
}
