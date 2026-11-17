
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek4130(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
pub fn loops4131(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 46;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome4132<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4132<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify4133<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 60 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome4134<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4134<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek4135(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 74
}
