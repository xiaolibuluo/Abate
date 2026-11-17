
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek9580(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 43
}
pub fn parse9581(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 50)
}
pub fn parse9582(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 57)
}
#[derive(Debug, Clone)]
pub enum Outcome9583<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9583<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse9584(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 71)
}
pub fn classify9585<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 78 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek9586(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 85
}
pub fn parse9587(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 92)
}
macro_rules! sum9588 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9588(values: &[i32]) -> i32 {
    let literal = sum9588!(99, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
