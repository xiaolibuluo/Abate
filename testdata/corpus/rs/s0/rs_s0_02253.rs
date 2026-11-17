
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode8390 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet8390<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode8390 for Packet8390<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum8391 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8391(values: &[i32]) -> i32 {
    let literal = sum8391!(55, 37);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse8392(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 62)
}
