
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode3940 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet3940<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode3940 for Packet3940<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum3941 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3941(values: &[i32]) -> i32 {
    let literal = sum3941!(90, 33);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
