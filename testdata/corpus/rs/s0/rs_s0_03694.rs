
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek13550(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 63
}
pub fn loops13551(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 70;
        break;
    }
    total
}
pub fn loops13552(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 77;
        break;
    }
    total
}
pub fn loops13553(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 84;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome13554<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13554<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops13555(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 98;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome13556<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13556<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome13557<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13557<T> {
    pub value: T,
    pub tag: u32,
}
