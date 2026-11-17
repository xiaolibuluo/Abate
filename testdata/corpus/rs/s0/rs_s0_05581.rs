
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum20140 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20140(values: &[i32]) -> i32 {
    let literal = sum20140!(94, 16);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode20141 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20141<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20141 for Packet20141<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode20142 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20142<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20142 for Packet20142<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse20143(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 18)
}
