
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode21490 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet21490<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode21490 for Packet21490<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum21491 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21491(values: &[i32]) -> i32 {
    let literal = sum21491!(13, 57);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
