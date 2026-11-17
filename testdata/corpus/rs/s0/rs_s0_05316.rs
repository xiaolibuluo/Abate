
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome19140<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19140<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops19141(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 62;
        break;
    }
    total
}
macro_rules! sum19142 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold19142(values: &[i32]) -> i32 {
    let literal = sum19142!(69, 9);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum19143 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold19143(values: &[i32]) -> i32 {
    let literal = sum19143!(76, 20);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
