
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse4390(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 86)
}
#[derive(Debug, Clone)]
pub enum Outcome4391<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4391<T> {
    pub value: T,
    pub tag: u32,
}
