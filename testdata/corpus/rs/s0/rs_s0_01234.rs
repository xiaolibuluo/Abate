
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome4510<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4510<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops4511(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 55;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome4512<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4512<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome4513<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4513<T> {
    pub value: T,
    pub tag: u32,
}
