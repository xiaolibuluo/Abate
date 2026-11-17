
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome23200<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23200<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek23201(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 5
}
pub fn peek23202(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 12
}
pub fn classify23203<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 19 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse23204(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 26)
}
pub fn peek23205(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 33
}
pub fn loops23206(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 40;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome23207<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23207<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify23208<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 54 => name,
        Some(_) => "other",
        None => "empty",
    }
}
