
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify5610<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 23 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum5611 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5611(values: &[i32]) -> i32 {
    let literal = sum5611!(30, 43);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome5612<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5612<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify5613<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 44 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome5614<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5614<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse5615(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 58)
}
