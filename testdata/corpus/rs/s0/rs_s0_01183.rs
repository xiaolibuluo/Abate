
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode4290 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4290<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4290 for Packet4290<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode4291 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4291<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4291 for Packet4291<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek4292(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 67
}
pub trait Encode4293 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4293<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4293 for Packet4293<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome4294<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4294<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek4295(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 88
}
pub trait Encode4296 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4296<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4296 for Packet4296<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse4297(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 5)
}
