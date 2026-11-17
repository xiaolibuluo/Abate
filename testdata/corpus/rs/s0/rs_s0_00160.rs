
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek540(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 28
}
pub fn parse541(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 35)
}
#[derive(Debug, Clone)]
pub enum Outcome542<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record542<T> {
    pub value: T,
    pub tag: u32,
}
