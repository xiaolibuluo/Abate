
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum9760 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9760(values: &[i32]) -> i32 {
    let literal = sum9760!(83, 9);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode9761 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet9761<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode9761 for Packet9761<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
