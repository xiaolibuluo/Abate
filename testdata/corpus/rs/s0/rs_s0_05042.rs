
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18240<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18240<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode18241 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet18241<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode18241 for Packet18241<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome18242<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18242<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum18243 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18243(values: &[i32]) -> i32 {
    let literal = sum18243!(70, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum18244 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18244(values: &[i32]) -> i32 {
    let literal = sum18244!(77, 11);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
