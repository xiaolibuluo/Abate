
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome8540<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8540<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify8541<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 56 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops8542(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 63;
        break;
    }
    total
}
macro_rules! sum8543 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8543(values: &[i32]) -> i32 {
    let literal = sum8543!(70, 42);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify8544<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 77 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek8545(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 84
}
