
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum7560 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7560(values: &[i32]) -> i32 {
    let literal = sum7560!(36, 55);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome7561<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7561<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek7562(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 50
}
pub fn peek7563(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 57
}
#[derive(Debug, Clone)]
pub enum Outcome7564<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7564<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum7565 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7565(values: &[i32]) -> i32 {
    let literal = sum7565!(71, 49);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek7566(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 78
}
