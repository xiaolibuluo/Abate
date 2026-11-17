
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum7830 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7830(values: &[i32]) -> i32 {
    let literal = sum7830!(96, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome7831<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7831<T> {
    pub value: T,
    pub tag: u32,
}
