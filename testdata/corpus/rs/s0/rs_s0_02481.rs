
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome9320<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9320<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek9321(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 100
}
macro_rules! sum9322 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9322(values: &[i32]) -> i32 {
    let literal = sum9322!(10, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
