
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum1700 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1700(values: &[i32]) -> i32 {
    let literal = sum1700!(81, 40);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome1701<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1701<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum1702 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1702(values: &[i32]) -> i32 {
    let literal = sum1702!(95, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome1703<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1703<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum1704 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1704(values: &[i32]) -> i32 {
    let literal = sum1704!(12, 23);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse1705(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 19)
}
