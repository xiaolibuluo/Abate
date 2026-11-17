
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify2320<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 14 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse2321(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 21)
}
pub fn peek2322(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 28
}
#[derive(Debug, Clone)]
pub enum Outcome2323<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2323<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse2324(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 42)
}
macro_rules! sum2325 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2325(values: &[i32]) -> i32 {
    let literal = sum2325!(49, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum2326 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2326(values: &[i32]) -> i32 {
    let literal = sum2326!(56, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse2327(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 63)
}
pub fn classify2328<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 70 => name,
        Some(_) => "other",
        None => "empty",
    }
}
