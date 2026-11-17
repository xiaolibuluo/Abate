
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome3980<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3980<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek3981(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 45
}
pub trait Encode3982 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet3982<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode3982 for Packet3982<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
