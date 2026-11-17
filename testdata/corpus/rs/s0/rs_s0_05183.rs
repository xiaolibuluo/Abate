
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek18680(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
pub fn peek18681(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 46
}
pub fn peek18682(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 53
}
#[derive(Debug, Clone)]
pub enum Outcome18683<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18683<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome18684<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18684<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops18685(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 74;
        break;
    }
    total
}
