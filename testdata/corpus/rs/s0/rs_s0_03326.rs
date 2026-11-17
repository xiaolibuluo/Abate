
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek12270(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 48
}
#[derive(Debug, Clone)]
pub enum Outcome12271<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12271<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode12272 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12272<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12272 for Packet12272<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
