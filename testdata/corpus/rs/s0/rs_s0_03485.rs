
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome12820<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12820<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode12821 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet12821<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode12821 for Packet12821<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek12822(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 98
}
macro_rules! sum12823 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12823(values: &[i32]) -> i32 {
    let literal = sum12823!(8, 49);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
