
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome24530<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24530<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum24531 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24531(values: &[i32]) -> i32 {
    let literal = sum24531!(85, 50);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum24532 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24532(values: &[i32]) -> i32 {
    let literal = sum24532!(92, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse24533(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 99)
}
pub fn classify24534<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 9 => name,
        Some(_) => "other",
        None => "empty",
    }
}
