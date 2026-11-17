
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify24240<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 89 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum24241 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24241(values: &[i32]) -> i32 {
    let literal = sum24241!(96, 30);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify24242<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 6 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify24243<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 13 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum24244 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24244(values: &[i32]) -> i32 {
    let literal = sum24244!(20, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek24245(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 27
}
#[derive(Debug, Clone)]
pub enum Outcome24246<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record24246<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse24247(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 41)
}
macro_rules! sum24248 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold24248(values: &[i32]) -> i32 {
    let literal = sum24248!(48, 46);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
