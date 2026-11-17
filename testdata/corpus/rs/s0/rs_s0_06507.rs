
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum23440 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23440(values: &[i32]) -> i32 {
    let literal = sum23440!(19, 69);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum23441 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23441(values: &[i32]) -> i32 {
    let literal = sum23441!(26, 19);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome23442<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23442<T> {
    pub value: T,
    pub tag: u32,
}
