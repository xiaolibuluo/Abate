
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum25870 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25870(values: &[i32]) -> i32 {
    let literal = sum25870!(74, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome25871<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25871<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek25872(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 88
}
