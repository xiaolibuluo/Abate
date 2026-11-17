
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode12800 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12800<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12800 for Packet12800<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome12801<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12801<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome12802<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12802<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome12803<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12803<T> {
    pub value: T,
    pub tag: u32,
}
