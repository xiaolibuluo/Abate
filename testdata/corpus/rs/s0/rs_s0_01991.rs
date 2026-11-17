
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome7460<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7460<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops7461(limit: i32) -> i32 {
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
#[derive(Debug, Clone)]
pub enum Outcome7462<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record7462<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum7463 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7463(values: &[i32]) -> i32 {
    let literal = sum7463!(24, 18);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
