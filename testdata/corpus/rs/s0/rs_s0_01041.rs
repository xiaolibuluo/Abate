
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops3820(limit: i32) -> i32 {
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
#[derive(Debug, Clone)]
pub enum Outcome3821<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3821<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome3822<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3822<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum3823 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3823(values: &[i32]) -> i32 {
    let literal = sum3823!(45, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek3824(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 52
}
#[derive(Debug, Clone)]
pub enum Outcome3825<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3825<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum3826 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3826(values: &[i32]) -> i32 {
    let literal = sum3826!(66, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum3827 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3827(values: &[i32]) -> i32 {
    let literal = sum3827!(73, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops3828(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 80;
        break;
    }
    total
}
