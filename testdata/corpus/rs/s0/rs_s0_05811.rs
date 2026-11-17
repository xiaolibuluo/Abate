
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum21030 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21030(values: &[i32]) -> i32 {
    let literal = sum21030!(87, 29);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome21031<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21031<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum21032 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21032(values: &[i32]) -> i32 {
    let literal = sum21032!(101, 51);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
