
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse16640(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 6)
}
#[derive(Debug, Clone)]
pub enum Outcome16641<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16641<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek16642(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 20
}
pub fn parse16643(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 27)
}
