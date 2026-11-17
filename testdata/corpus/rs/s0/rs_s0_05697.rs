
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek20620(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
pub trait Encode20621 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20621<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20621 for Packet20621<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum20622 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20622(values: &[i32]) -> i32 {
    let literal = sum20622!(53, 69);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum20623 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20623(values: &[i32]) -> i32 {
    let literal = sum20623!(60, 19);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
