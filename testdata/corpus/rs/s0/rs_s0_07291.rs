
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome26300<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26300<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse26301(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 58)
}
