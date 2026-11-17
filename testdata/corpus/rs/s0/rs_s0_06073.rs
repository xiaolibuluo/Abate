
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum21950 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21950(values: &[i32]) -> i32 {
    let literal = sum21950!(22, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome21951<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21951<T> {
    pub value: T,
    pub tag: u32,
}
