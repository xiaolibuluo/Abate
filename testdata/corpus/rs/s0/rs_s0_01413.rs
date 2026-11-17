
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome5200<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5200<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode5201 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet5201<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode5201 for Packet5201<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum5202 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5202(values: &[i32]) -> i32 {
    let literal = sum5202!(86, 11);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek5203(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 93
}
#[derive(Debug, Clone)]
pub enum Outcome5204<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5204<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops5205(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 10;
        break;
    }
    total
}
macro_rules! sum5206 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5206(values: &[i32]) -> i32 {
    let literal = sum5206!(17, 55);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
