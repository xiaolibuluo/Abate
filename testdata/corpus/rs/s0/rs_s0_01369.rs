
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode5070 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet5070<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode5070 for Packet5070<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode5071 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet5071<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode5071 for Packet5071<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse5072(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 14)
}
pub fn peek5073(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 21
}
pub fn peek5074(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 28
}
