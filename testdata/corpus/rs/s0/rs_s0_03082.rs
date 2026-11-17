
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome11300<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11300<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum11301 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11301(values: &[i32]) -> i32 {
    let literal = sum11301!(55, 61);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum11302 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11302(values: &[i32]) -> i32 {
    let literal = sum11302!(62, 11);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
