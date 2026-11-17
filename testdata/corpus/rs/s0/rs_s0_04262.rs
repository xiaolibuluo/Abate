
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum15640 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15640(values: &[i32]) -> i32 {
    let literal = sum15640!(64, 38);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse15641(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 71)
}
#[derive(Debug, Clone)]
pub enum Outcome15642<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15642<T> {
    pub value: T,
    pub tag: u32,
}
