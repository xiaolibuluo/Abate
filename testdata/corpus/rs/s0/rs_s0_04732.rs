
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum17110 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17110(values: &[i32]) -> i32 {
    let literal = sum17110!(35, 30);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode17111 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17111<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17111 for Packet17111<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
