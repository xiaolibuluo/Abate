
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome4550<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4550<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum4551 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4551(values: &[i32]) -> i32 {
    let literal = sum4551!(10, 33);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek4552(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 17
}
macro_rules! sum4553 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4553(values: &[i32]) -> i32 {
    let literal = sum4553!(24, 55);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum4554 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4554(values: &[i32]) -> i32 {
    let literal = sum4554!(31, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
