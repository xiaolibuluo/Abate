
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse15200(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 74)
}
pub trait Encode15201 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15201<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15201 for Packet15201<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome15202<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15202<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum15203 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15203(values: &[i32]) -> i32 {
    let literal = sum15203!(95, 68);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
