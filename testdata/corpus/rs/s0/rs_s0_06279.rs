
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome22660<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22660<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops22661(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 79;
        break;
    }
    total
}
pub fn loops22662(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 86;
        break;
    }
    total
}
macro_rules! sum22663 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22663(values: &[i32]) -> i32 {
    let literal = sum22663!(93, 44);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum22664 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22664(values: &[i32]) -> i32 {
    let literal = sum22664!(100, 55);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops22665(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 10;
        break;
    }
    total
}
