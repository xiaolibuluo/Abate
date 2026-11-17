
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse25910(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 29)
}
pub fn peek25911(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 36
}
pub fn parse25912(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 43)
}
pub fn peek25913(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 50
}
#[derive(Debug, Clone)]
pub enum Outcome25914<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25914<T> {
    pub value: T,
    pub tag: u32,
}
