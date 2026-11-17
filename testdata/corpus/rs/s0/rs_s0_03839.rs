
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome14040<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14040<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome14041<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14041<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek14042(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 35
}
pub fn classify14043<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 42 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops14044(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 49;
        break;
    }
    total
}
pub fn parse14045(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 56)
}
#[derive(Debug, Clone)]
pub enum Outcome14046<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14046<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek14047(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 70
}
pub fn loops14048(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 77;
        break;
    }
    total
}
pub fn loops14049(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 84;
        break;
    }
    total
}
