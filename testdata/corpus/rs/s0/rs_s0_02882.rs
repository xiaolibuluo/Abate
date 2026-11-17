
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify10630<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 50 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome10631<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10631<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek10632(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 64
}
#[derive(Debug, Clone)]
pub enum Outcome10633<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10633<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome10634<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10634<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum10635 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10635(values: &[i32]) -> i32 {
    let literal = sum10635!(85, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum10636 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10636(values: &[i32]) -> i32 {
    let literal = sum10636!(92, 13);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse10637(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 99)
}
pub fn classify10638<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 9 => name,
        Some(_) => "other",
        None => "empty",
    }
}
