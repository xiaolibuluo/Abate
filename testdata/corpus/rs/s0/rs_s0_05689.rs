
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode20590 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20590<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20590 for Packet20590<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek20591(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 7
}
pub trait Encode20592 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20592<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20592 for Packet20592<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum20593 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20593(values: &[i32]) -> i32 {
    let literal = sum20593!(21, 59);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
