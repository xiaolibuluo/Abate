
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek23940(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 87
}
pub trait Encode23941 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23941<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23941 for Packet23941<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome23942<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23942<T> {
    pub value: T,
    pub tag: u32,
}
