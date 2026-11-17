
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum14940 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14940(values: &[i32]) -> i32 {
    let literal = sum14940!(27, 36);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum14941 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14941(values: &[i32]) -> i32 {
    let literal = sum14941!(34, 47);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify14942<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 41 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek14943(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 48
}
#[derive(Debug, Clone)]
pub enum Outcome14944<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14944<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek14945(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 62
}
pub fn parse14946(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 69)
}
macro_rules! sum14947 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14947(values: &[i32]) -> i32 {
    let literal = sum14947!(76, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
