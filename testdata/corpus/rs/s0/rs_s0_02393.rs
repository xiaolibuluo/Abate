
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek9000(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 65
}
pub fn parse9001(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 72)
}
pub fn peek9002(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 79
}
pub trait Encode9003 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9003<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9003 for Packet9003<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode9004 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9004<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9004 for Packet9004<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek9005(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 100
}
