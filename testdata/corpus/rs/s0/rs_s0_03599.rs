
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode13230 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13230<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13230 for Packet13230<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum13231 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13231(values: &[i32]) -> i32 {
    let literal = sum13231!(42, 9);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
