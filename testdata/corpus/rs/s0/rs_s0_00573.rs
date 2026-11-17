
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum2040 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2040(values: &[i32]) -> i32 {
    let literal = sum2040!(38, 34);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode2041 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2041<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2041 for Packet2041<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
