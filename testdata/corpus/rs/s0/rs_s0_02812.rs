
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode10370 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10370<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10370 for Packet10370<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum10371 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10371(values: &[i32]) -> i32 {
    let literal = sum10371!(10, 20);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse10372(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 17)
}
pub trait Encode10373 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10373<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10373 for Packet10373<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
