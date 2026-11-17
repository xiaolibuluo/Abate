
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum23170 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23170(values: &[i32]) -> i32 {
    let literal = sum23170!(56, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek23171(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 63
}
pub fn peek23172(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 70
}
pub fn peek23173(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 77
}
pub fn loops23174(limit: i32) -> i32 {
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
pub fn peek23175(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 91
}
#[derive(Debug, Clone)]
pub enum Outcome23176<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23176<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek23177(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 8
}
