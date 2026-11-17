
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome14970<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14970<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek14971(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 73
}
pub fn classify14972<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 80 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum14973 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14973(values: &[i32]) -> i32 {
    let literal = sum14973!(87, 29);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse14974(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 94)
}
pub fn classify14975<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 101 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify14976<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 11 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome14977<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14977<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome14978<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14978<T> {
    pub value: T,
    pub tag: u32,
}
