
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16740<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16740<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek16741(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 46
}
macro_rules! sum16742 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16742(values: &[i32]) -> i32 {
    let literal = sum16742!(53, 37);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse16743(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 60)
}
#[derive(Debug, Clone)]
pub enum Outcome16744<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16744<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify16745<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 74 => name,
        Some(_) => "other",
        None => "empty",
    }
}
