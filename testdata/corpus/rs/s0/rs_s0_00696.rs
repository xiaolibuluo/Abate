
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek2510(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 67
}
pub fn parse2511(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 74)
}
#[derive(Debug, Clone)]
pub enum Outcome2512<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2512<T> {
    pub value: T,
    pub tag: u32,
}
