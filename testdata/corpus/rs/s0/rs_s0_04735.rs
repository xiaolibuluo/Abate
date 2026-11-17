
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum17130 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17130(values: &[i32]) -> i32 {
    let literal = sum17130!(61, 44);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome17131<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17131<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek17132(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 75
}
