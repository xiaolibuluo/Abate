
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse1260(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 91)
}
pub fn parse1261(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 98)
}
#[derive(Debug, Clone)]
pub enum Outcome1262<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1262<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse1263(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 15)
}
macro_rules! sum1264 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1264(values: &[i32]) -> i32 {
    let literal = sum1264!(22, 20);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify1265<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 29 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome1266<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1266<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse1267(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 43)
}
