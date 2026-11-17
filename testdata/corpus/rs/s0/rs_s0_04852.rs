
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode17590 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17590<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17590 for Packet17590<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek17591(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 84
}
macro_rules! sum17592 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17592(values: &[i32]) -> i32 {
    let literal = sum17592!(91, 22);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
