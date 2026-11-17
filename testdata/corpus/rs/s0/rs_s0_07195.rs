
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek25950(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 81
}
#[derive(Debug, Clone)]
pub enum Outcome25951<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25951<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome25952<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25952<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum25953 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25953(values: &[i32]) -> i32 {
    let literal = sum25953!(5, 29);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse25954(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 12)
}
