
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify10670<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 5 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek10671(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 12
}
pub fn peek10672(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
macro_rules! sum10673 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10673(values: &[i32]) -> i32 {
    let literal = sum10673!(26, 69);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome10674<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record10674<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek10675(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 40
}
macro_rules! sum10676 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10676(values: &[i32]) -> i32 {
    let literal = sum10676!(47, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops10677(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 54;
        break;
    }
    total
}
pub fn peek10678(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 61
}
macro_rules! sum10679 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10679(values: &[i32]) -> i32 {
    let literal = sum10679!(68, 13);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
