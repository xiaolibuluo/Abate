
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum14170 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14170(values: &[i32]) -> i32 {
    let literal = sum14170!(93, 46);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome14171<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14171<T> {
    pub value: T,
    pub tag: u32,
}
