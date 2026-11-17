
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops2540(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 9;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome2541<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2541<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek2542(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 23
}
