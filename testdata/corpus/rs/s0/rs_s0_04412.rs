
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum16140 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16140(values: &[i32]) -> i32 {
    let literal = sum16140!(35, 22);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode16141 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet16141<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode16141 for Packet16141<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
