
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16750<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16750<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek16751(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 59
}
macro_rules! sum16752 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16752(values: &[i32]) -> i32 {
    let literal = sum16752!(66, 44);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome16753<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16753<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum16754 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16754(values: &[i32]) -> i32 {
    let literal = sum16754!(80, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
