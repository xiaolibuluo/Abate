
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify3890<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 18 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify3891<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 25 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome3892<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3892<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek3893(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
pub fn classify3894<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 46 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse3895(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 53)
}
#[derive(Debug, Clone)]
pub enum Outcome3896<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3896<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek3897(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 67
}
pub fn parse3898(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 74)
}
#[derive(Debug, Clone)]
pub enum Outcome3899<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3899<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek38910(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 88
}
macro_rules! sum38911 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold38911(values: &[i32]) -> i32 {
    let literal = sum38911!(95, 47);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
