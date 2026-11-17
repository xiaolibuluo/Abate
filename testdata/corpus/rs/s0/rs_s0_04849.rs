
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome17570<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17570<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek17571(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 58
}
macro_rules! sum17572 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17572(values: &[i32]) -> i32 {
    let literal = sum17572!(65, 69);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
