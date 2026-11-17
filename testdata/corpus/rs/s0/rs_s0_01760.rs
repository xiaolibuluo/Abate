
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome6500<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record6500<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek6501(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 23
}
pub fn loops6502(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 30;
        break;
    }
    total
}
pub fn peek6503(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 37
}
#[derive(Debug, Clone)]
pub enum Outcome6504<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record6504<T> {
    pub value: T,
    pub tag: u32,
}
