
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16020<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16020<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse16021(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 80)
}
pub fn classify16022<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 87 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome16023<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16023<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify16024<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 101 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum16025 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16025(values: &[i32]) -> i32 {
    let literal = sum16025!(11, 54);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome16026<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16026<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify16027<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 25 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek16028(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 32
}
macro_rules! sum16029 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16029(values: &[i32]) -> i32 {
    let literal = sum16029!(39, 37);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
