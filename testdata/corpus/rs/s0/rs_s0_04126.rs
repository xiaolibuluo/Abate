
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15150<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15150<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse15151(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 16)
}
macro_rules! sum15152 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15152(values: &[i32]) -> i32 {
    let literal = sum15152!(23, 22);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
