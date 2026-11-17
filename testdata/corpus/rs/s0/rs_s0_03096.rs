
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome11340<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11340<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum11341 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11341(values: &[i32]) -> i32 {
    let literal = sum11341!(10, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome11342<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11342<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum11343 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11343(values: &[i32]) -> i32 {
    let literal = sum11343!(24, 50);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
