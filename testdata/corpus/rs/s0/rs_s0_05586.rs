
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20180<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20180<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops20181(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 56;
        break;
    }
    total
}
macro_rules! sum20182 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20182(values: &[i32]) -> i32 {
    let literal = sum20182!(63, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek20183(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 70
}
macro_rules! sum20184 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20184(values: &[i32]) -> i32 {
    let literal = sum20184!(77, 27);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
