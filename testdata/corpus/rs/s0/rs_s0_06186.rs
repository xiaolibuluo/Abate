
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse22410(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 38)
}
#[derive(Debug, Clone)]
pub enum Outcome22411<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22411<T> {
    pub value: T,
    pub tag: u32,
}
