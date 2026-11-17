
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18950<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18950<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse18951(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 9)
}
pub fn parse18952(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 16)
}
pub fn parse18953(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 23)
}
macro_rules! sum18954 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18954(values: &[i32]) -> i32 {
    let literal = sum18954!(30, 20);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome18955<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18955<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify18956<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 44 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum18957 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18957(values: &[i32]) -> i32 {
    let literal = sum18957!(51, 53);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum18958 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18958(values: &[i32]) -> i32 {
    let literal = sum18958!(58, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum18959 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18959(values: &[i32]) -> i32 {
    let literal = sum18959!(65, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
