
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode10870 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10870<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10870 for Packet10870<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode10871 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10871<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10871 for Packet10871<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum10872 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10872(values: &[i32]) -> i32 {
    let literal = sum10872!(85, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
