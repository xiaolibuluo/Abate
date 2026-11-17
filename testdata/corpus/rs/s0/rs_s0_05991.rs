
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome21650<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21650<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum21651 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21651(values: &[i32]) -> i32 {
    let literal = sum21651!(27, 47);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum21652 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21652(values: &[i32]) -> i32 {
    let literal = sum21652!(34, 58);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome21653<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21653<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops21654(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 48;
        break;
    }
    total
}
