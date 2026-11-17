
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome11030<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11030<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify11031<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 92 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops11032(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 99;
        break;
    }
    total
}
pub fn parse11033(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 9)
}
pub fn classify11034<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 16 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome11035<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11035<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek11036(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 30
}
#[derive(Debug, Clone)]
pub enum Outcome11037<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11037<T> {
    pub value: T,
    pub tag: u32,
}
