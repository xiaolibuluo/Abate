
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18650<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18650<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek18651(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 7
}
pub fn loops18652(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 14;
        break;
    }
    total
}
pub fn peek18653(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 21
}
#[derive(Debug, Clone)]
pub enum Outcome18654<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18654<T> {
    pub value: T,
    pub tag: u32,
}
