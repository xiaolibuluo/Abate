
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse23010(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 42)
}
pub fn peek23011(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 49
}
pub trait Encode23012 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23012<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23012 for Packet23012<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum23013 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23013(values: &[i32]) -> i32 {
    let literal = sum23013!(63, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
