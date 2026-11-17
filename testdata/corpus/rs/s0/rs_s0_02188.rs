
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome8170<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8170<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek8171(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 60
}
pub fn classify8172<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 67 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome8173<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8173<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify8174<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 81 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops8175(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 88;
        break;
    }
    total
}
pub fn classify8176<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 95 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum8177 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8177(values: &[i32]) -> i32 {
    let literal = sum8177!(5, 10);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
