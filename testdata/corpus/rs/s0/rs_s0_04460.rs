
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome16260<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16260<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum16261 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16261(values: &[i32]) -> i32 {
    let literal = sum16261!(101, 56);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome16262<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record16262<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops16263(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 18;
        break;
    }
    total
}
macro_rules! sum16264 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold16264(values: &[i32]) -> i32 {
    let literal = sum16264!(25, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
