
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode1880 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1880<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1880 for Packet1880<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn peek1881(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 31
}
pub trait Encode1882 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet1882<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode1882 for Packet1882<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
#[derive(Debug, Clone)]
pub enum Outcome1883<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1883<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek1884(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 52
}
pub fn peek1885(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 59
}
macro_rules! sum1886 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1886(values: &[i32]) -> i32 {
    let literal = sum1886!(66, 49);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum1887 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1887(values: &[i32]) -> i32 {
    let literal = sum1887!(73, 60);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome1888<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1888<T> {
    pub value: T,
    pub tag: u32,
}
