
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome18230<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18230<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek18231(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 43
}
macro_rules! sum18232 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18232(values: &[i32]) -> i32 {
    let literal = sum18232!(50, 43);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops18233(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 57;
        break;
    }
    total
}
macro_rules! sum18234 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18234(values: &[i32]) -> i32 {
    let literal = sum18234!(64, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
