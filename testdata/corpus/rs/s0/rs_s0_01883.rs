
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum6960 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold6960(values: &[i32]) -> i32 {
    let literal = sum6960!(32, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek6961(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
pub fn parse6962(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 46)
}
pub trait Encode6963 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet6963<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode6963 for Packet6963<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
