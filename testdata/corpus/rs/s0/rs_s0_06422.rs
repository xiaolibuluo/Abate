
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome23120<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23120<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome23121<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23121<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops23122(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 5;
        break;
    }
    total
}
pub fn loops23123(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 12;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome23124<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23124<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek23125(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 26
}
