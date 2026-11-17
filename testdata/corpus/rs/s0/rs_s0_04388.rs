
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16070<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16070<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum16071 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16071(values: &[i32]) -> i32 {
    let literal = sum16071!(48, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
