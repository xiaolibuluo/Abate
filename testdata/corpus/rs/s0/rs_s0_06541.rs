
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek23560(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 78
}
macro_rules! sum23561 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23561(values: &[i32]) -> i32 {
    let literal = sum23561!(85, 42);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome23562<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23562<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops23563(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 99;
        break;
    }
    total
}
pub fn peek23564(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 9
}
