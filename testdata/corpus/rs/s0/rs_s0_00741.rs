
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome2720<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2720<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode2721 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2721<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2721 for Packet2721<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse2722(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 63)
}
pub fn parse2723(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 70)
}
pub fn parse2724(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 77)
}
macro_rules! sum2725 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2725(values: &[i32]) -> i32 {
    let literal = sum2725!(84, 16);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum2726 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2726(values: &[i32]) -> i32 {
    let literal = sum2726!(91, 27);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
