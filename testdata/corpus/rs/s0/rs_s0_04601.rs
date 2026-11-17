
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16680<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16680<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse16681(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 65)
}
macro_rules! sum16682 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16682(values: &[i32]) -> i32 {
    let literal = sum16682!(72, 56);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek16683(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 79
}
