
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek13850(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 65
}
pub fn classify13851<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 72 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops13852(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 79;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome13853<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13853<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome13854<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13854<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome13855<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13855<T> {
    pub value: T,
    pub tag: u32,
}
