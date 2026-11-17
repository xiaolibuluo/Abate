
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode20160 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20160<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20160 for Packet20160<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum20161 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20161(values: &[i32]) -> i32 {
    let literal = sum20161!(30, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome20162<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20162<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode20163 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20163<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20163 for Packet20163<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
