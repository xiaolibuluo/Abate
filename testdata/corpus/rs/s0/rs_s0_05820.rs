
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek21100(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 81
}
macro_rules! sum21101 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21101(values: &[i32]) -> i32 {
    let literal = sum21101!(88, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome21102<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21102<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome21103<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21103<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse21104(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 12)
}
