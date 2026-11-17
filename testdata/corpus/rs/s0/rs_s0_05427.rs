
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode19520 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet19520<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode19520 for Packet19520<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum19521 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold19521(values: &[i32]) -> i32 {
    let literal = sum19521!(71, 20);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
