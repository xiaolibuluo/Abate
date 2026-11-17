
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20690<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20690<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek20691(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 40
}
pub fn peek20692(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 47
}
#[derive(Debug, Clone)]
pub enum Outcome20693<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20693<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum20694 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20694(values: &[i32]) -> i32 {
    let literal = sum20694!(61, 18);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum20695 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20695(values: &[i32]) -> i32 {
    let literal = sum20695!(68, 29);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome20696<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20696<T> {
    pub value: T,
    pub tag: u32,
}
