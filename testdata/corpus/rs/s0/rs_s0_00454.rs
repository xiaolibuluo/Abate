
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek1590(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 35
}
#[derive(Debug, Clone)]
pub enum Outcome1591<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record1591<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek1592(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 49
}
pub fn loops1593(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 56;
        break;
    }
    total
}
