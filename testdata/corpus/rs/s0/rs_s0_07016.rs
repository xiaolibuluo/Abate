
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome25240<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25240<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome25241<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25241<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify25242<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 45 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek25243(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 52
}
#[derive(Debug, Clone)]
pub enum Outcome25244<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25244<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek25245(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 66
}
pub fn loops25246(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 73;
        break;
    }
    total
}
pub fn classify25247<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 80 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome25248<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25248<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops25249(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 94;
        break;
    }
    total
}
