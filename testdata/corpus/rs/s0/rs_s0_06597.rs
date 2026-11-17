
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek23810(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 15
}
pub fn peek23811(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 22
}
#[derive(Debug, Clone)]
pub enum Outcome23812<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23812<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome23813<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23813<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum23814 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23814(values: &[i32]) -> i32 {
    let literal = sum23814!(43, 67);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify23815<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 50 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome23816<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23816<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops23817(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 64;
        break;
    }
    total
}
pub fn loops23818(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 71;
        break;
    }
    total
}
pub fn classify23819<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 78 => name,
        Some(_) => "other",
        None => "empty",
    }
}
