
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome12740<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12740<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum12741 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12741(values: &[i32]) -> i32 {
    let literal = sum12741!(84, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome12742<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12742<T> {
    pub value: T,
    pub tag: u32,
}
