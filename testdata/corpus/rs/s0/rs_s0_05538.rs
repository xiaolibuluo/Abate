
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome19980<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19980<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse19981(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 87)
}
