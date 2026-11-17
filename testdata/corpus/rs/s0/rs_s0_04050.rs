
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify14870<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 33 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify14871<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 40 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek14872(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 47
}
#[derive(Debug, Clone)]
pub enum Outcome14873<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14873<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse14874(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 61)
}
pub fn loops14875(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 68;
        break;
    }
    total
}
pub fn parse14876(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 75)
}
