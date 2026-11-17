
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse18120(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 87)
}
#[derive(Debug, Clone)]
pub enum Outcome18121<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18121<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek18122(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 101
}
pub fn peek18123(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 11
}
#[derive(Debug, Clone)]
pub enum Outcome18124<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18124<T> {
    pub value: T,
    pub tag: u32,
}
