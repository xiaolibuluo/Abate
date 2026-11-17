
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome4160<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4160<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum4161 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4161(values: &[i32]) -> i32 {
    let literal = sum4161!(85, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum4162 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4162(values: &[i32]) -> i32 {
    let literal = sum4162!(92, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome4163<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4163<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome4164<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4164<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops4165(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 16;
        break;
    }
    total
}
