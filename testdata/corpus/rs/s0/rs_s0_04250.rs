
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum15610 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15610(values: &[i32]) -> i32 {
    let literal = sum15610!(25, 17);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode15611 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15611<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15611 for Packet15611<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
