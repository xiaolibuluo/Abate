
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome24540<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24540<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum24541 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24541(values: &[i32]) -> i32 {
    let literal = sum24541!(98, 57);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek24542(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 8
}
pub fn classify24543<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 15 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome24544<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24544<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse24545(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 29)
}
