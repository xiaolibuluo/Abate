
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome22180<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22180<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse22181(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 37)
}
#[derive(Debug, Clone)]
pub enum Outcome22182<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record22182<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops22183(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 51;
        break;
    }
    total
}
pub fn peek22184(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 58
}
pub fn loops22185(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 65;
        break;
    }
    total
}
macro_rules! sum22186 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22186(values: &[i32]) -> i32 {
    let literal = sum22186!(72, 46);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
