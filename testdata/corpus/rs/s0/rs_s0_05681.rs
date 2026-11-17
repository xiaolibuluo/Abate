
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode20550 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20550<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20550 for Packet20550<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum20551 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20551(values: &[i32]) -> i32 {
    let literal = sum20551!(52, 9);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek20552(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 59
}
