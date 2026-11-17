
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek15940(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 66
}
pub trait Encode15941 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15941<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15941 for Packet15941<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome15942<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15942<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode15943 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15943<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15943 for Packet15943<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode15944 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15944<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15944 for Packet15944<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome15945<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15945<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode15946 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15946<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15946 for Packet15946<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
