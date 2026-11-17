
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse4940(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 25)
}
pub fn parse4941(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 32)
}
pub fn peek4942(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
#[derive(Debug, Clone)]
pub enum Outcome4943<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4943<T> {
    pub value: T,
    pub tag: u32,
}
