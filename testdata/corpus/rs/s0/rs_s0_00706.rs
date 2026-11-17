
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum2560 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2560(values: &[i32]) -> i32 {
    let literal = sum2560!(35, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome2561<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2561<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode2562 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2562<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2562 for Packet2562<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
