
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome23110<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23110<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops23111(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 82;
        break;
    }
    total
}
pub fn loops23112(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 89;
        break;
    }
    total
}
pub fn loops23113(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 96;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome23114<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23114<T> {
    pub value: T,
    pub tag: u32,
}
