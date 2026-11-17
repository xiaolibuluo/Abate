
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek3870(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 89
}
pub fn classify3871<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 96 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome3872<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3872<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome3873<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3873<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse3874(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 20)
}
macro_rules! sum3875 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3875(values: &[i32]) -> i32 {
    let literal = sum3875!(27, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum3876 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3876(values: &[i32]) -> i32 {
    let literal = sum3876!(34, 39);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek3877(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 41
}
pub fn loops3878(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 48;
        break;
    }
    total
}
pub fn parse3879(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 55)
}
macro_rules! sum38710 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold38710(values: &[i32]) -> i32 {
    let literal = sum38710!(62, 22);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
