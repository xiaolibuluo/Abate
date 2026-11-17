
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome3470<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3470<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum3471 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3471(values: &[i32]) -> i32 {
    let literal = sum3471!(61, 9);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
