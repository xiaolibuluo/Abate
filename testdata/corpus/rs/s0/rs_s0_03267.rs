
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome12030<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12030<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome12031<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12031<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse12032(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 41)
}
macro_rules! sum12033 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12033(values: &[i32]) -> i32 {
    let literal = sum12033!(48, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome12034<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12034<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome12035<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record12035<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum12036 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12036(values: &[i32]) -> i32 {
    let literal = sum12036!(69, 17);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops12037(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 76;
        break;
    }
    total
}
