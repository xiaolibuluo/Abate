
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek25210(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 89
}
pub fn peek25211(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 96
}
pub fn parse25212(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 6)
}
pub fn parse25213(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 13)
}
pub trait Encode25214 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet25214<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode25214 for Packet25214<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode25215 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet25215<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode25215 for Packet25215<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode25216 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet25216<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode25216 for Packet25216<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse25217(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 41)
}
macro_rules! sum25218 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25218(values: &[i32]) -> i32 {
    let literal = sum25218!(48, 54);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
