
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek3610(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 42
}
pub fn peek3611(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 49
}
pub trait Encode3612 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet3612<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode3612 for Packet3612<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome3613<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3613<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum3614 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3614(values: &[i32]) -> i32 {
    let literal = sum3614!(70, 18);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
