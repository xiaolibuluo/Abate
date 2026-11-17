
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome5810<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5810<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse5811(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 96)
}
#[derive(Debug, Clone)]
pub enum Outcome5812<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5812<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek5813(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 13
}
pub fn loops5814(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 20;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome5815<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5815<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse5816(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 34)
}
pub fn peek5817(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 41
}
pub fn loops5818(limit: i32) -> i32 {
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
pub fn classify5819<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 55 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify58110<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 62 => name,
        Some(_) => "other",
        None => "empty",
    }
}
