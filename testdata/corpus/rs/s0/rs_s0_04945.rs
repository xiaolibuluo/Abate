
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops17960(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 73;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome17961<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17961<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek17962(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 87
}
pub fn peek17963(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 94
}
pub fn peek17964(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 101
}
#[derive(Debug, Clone)]
pub enum Outcome17965<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17965<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum17966 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17966(values: &[i32]) -> i32 {
    let literal = sum17966!(18, 20);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify17967<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 25 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify17968<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 32 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops17969(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 39;
        break;
    }
    total
}
