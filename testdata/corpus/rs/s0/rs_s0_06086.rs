
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode22020 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet22020<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode22020 for Packet22020<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek22021(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 23
}
#[derive(Debug, Clone)]
pub enum Outcome22022<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22022<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum22023 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22023(values: &[i32]) -> i32 {
    let literal = sum22023!(37, 23);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
