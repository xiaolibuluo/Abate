
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse22600(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 91)
}
macro_rules! sum22601 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22601(values: &[i32]) -> i32 {
    let literal = sum22601!(98, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome22602<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22602<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek22603(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 15
}
pub fn peek22604(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 22
}
pub fn parse22605(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 29)
}
macro_rules! sum22606 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22606(values: &[i32]) -> i32 {
    let literal = sum22606!(36, 35);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
