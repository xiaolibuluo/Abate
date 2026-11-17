
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum570 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold570(values: &[i32]) -> i32 {
    let literal = sum570!(67, 42);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode571 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet571<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode571 for Packet571<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
