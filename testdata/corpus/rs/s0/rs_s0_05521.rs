
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum19910 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold19910(values: &[i32]) -> i32 {
    let literal = sum19910!(86, 38);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome19911<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19911<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify19912<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 100 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek19913(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 10
}
pub fn parse19914(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 17)
}
pub fn classify19915<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 24 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek19916(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 31
}
#[derive(Debug, Clone)]
pub enum Outcome19917<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19917<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops19918(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 45;
        break;
    }
    total
}
macro_rules! sum19919 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold19919(values: &[i32]) -> i32 {
    let literal = sum19919!(52, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
