
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum21140 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21140(values: &[i32]) -> i32 {
    let literal = sum21140!(36, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek21141(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 43
}
pub fn peek21142(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 50
}
pub fn parse21143(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 57)
}
#[derive(Debug, Clone)]
pub enum Outcome21144<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21144<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify21145<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 71 => name,
        Some(_) => "other",
        None => "empty",
    }
}
