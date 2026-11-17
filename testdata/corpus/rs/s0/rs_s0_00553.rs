
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode1980 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1980<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1980 for Packet1980<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum1981 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1981(values: &[i32]) -> i32 {
    let literal = sum1981!(64, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
