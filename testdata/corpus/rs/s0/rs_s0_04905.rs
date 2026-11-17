
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify17800<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 59 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum17801 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17801(values: &[i32]) -> i32 {
    let literal = sum17801!(66, 36);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek17802(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 73
}
#[derive(Debug, Clone)]
pub enum Outcome17803<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17803<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify17804<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 87 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops17805(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 94;
        break;
    }
    total
}
