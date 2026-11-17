
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek17840(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 14
}
pub fn peek17841(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 21
}
#[derive(Debug, Clone)]
pub enum Outcome17842<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17842<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome17843<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17843<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops17844(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 42;
        break;
    }
    total
}
pub fn loops17845(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 49;
        break;
    }
    total
}
pub fn classify17846<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 56 => name,
        Some(_) => "other",
        None => "empty",
    }
}
