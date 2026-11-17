
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek11520(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 43
}
pub fn loops11521(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 50;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome11522<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11522<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek11523(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 64
}
macro_rules! sum11524 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11524(values: &[i32]) -> i32 {
    let literal = sum11524!(71, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek11525(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 78
}
#[derive(Debug, Clone)]
pub enum Outcome11526<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11526<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum11527 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold11527(values: &[i32]) -> i32 {
    let literal = sum11527!(92, 37);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
