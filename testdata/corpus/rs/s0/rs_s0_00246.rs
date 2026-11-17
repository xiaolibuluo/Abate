
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome880<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record880<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome881<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record881<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops882(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 96;
        break;
    }
    total
}
pub fn loops883(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 6;
        break;
    }
    total
}
pub fn classify884<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 13 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify885<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 20 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek886(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 27
}
macro_rules! sum887 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold887(values: &[i32]) -> i32 {
    let literal = sum887!(34, 31);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
