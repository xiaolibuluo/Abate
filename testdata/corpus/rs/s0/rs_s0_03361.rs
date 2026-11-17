
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek12390(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 10
}
#[derive(Debug, Clone)]
pub enum Outcome12391<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12391<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek12392(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 24
}
#[derive(Debug, Clone)]
pub enum Outcome12393<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12393<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode12394 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12394<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12394 for Packet12394<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode12395 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12395<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12395 for Packet12395<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
