
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek15690(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 32
}
pub fn parse15691(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 39)
}
#[derive(Debug, Clone)]
pub enum Outcome15692<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15692<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum15693 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15693(values: &[i32]) -> i32 {
    let literal = sum15693!(53, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
