
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum11990 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11990(values: &[i32]) -> i32 {
    let literal = sum11990!(72, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops11991(limit: i32) -> i32 {
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
#[derive(Debug, Clone)]
pub enum Outcome11992<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11992<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome11993<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11993<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek11994(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 100
}
pub fn peek11995(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 10
}
pub fn peek11996(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 17
}
pub fn loops11997(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 24;
        break;
    }
    total
}
