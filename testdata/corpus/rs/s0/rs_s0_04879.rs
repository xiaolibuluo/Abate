
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode17690 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17690<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17690 for Packet17690<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode17691 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17691<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17691 for Packet17691<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome17692<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17692<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode17693 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17693<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17693 for Packet17693<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
