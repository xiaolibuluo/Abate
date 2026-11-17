
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse13800(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 97)
}
macro_rules! sum13801 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13801(values: &[i32]) -> i32 {
    let literal = sum13801!(7, 42);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse13802(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 14)
}
#[derive(Debug, Clone)]
pub enum Outcome13803<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13803<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify13804<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 28 => name,
        Some(_) => "other",
        None => "empty",
    }
}
