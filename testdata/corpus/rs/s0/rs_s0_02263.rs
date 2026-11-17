
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops8420(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 87;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome8421<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8421<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum8422 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8422(values: &[i32]) -> i32 {
    let literal = sum8422!(101, 69);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek8423(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 11
}
