
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome21230<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21230<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify21231<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 63 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek21232(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 70
}
#[derive(Debug, Clone)]
pub enum Outcome21233<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21233<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify21234<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 84 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum21235 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21235(values: &[i32]) -> i32 {
    let literal = sum21235!(91, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops21236(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 98;
        break;
    }
    total
}
