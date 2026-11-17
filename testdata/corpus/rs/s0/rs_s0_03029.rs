
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome11130<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11130<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome11131<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11131<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek11132(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 35
}
#[derive(Debug, Clone)]
pub enum Outcome11133<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11133<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome11134<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11134<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome11135<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11135<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome11136<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11136<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome11137<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11137<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek11138(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 77
}
macro_rules! sum11139 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11139(values: &[i32]) -> i32 {
    let literal = sum11139!(84, 30);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode111310 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet111310<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode111310 for Packet111310<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
