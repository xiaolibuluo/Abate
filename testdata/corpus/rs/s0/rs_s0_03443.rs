
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify12640<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 44 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek12641(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 51
}
#[derive(Debug, Clone)]
pub enum Outcome12642<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12642<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum12643 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12643(values: &[i32]) -> i32 {
    let literal = sum12643!(65, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome12644<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12644<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse12645(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 79)
}
macro_rules! sum12646 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12646(values: &[i32]) -> i32 {
    let literal = sum12646!(86, 17);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek12647(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 93
}
pub fn peek12648(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 100
}
macro_rules! sum12649 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12649(values: &[i32]) -> i32 {
    let literal = sum12649!(10, 50);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse126410(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 17)
}
