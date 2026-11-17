
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15330<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15330<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops15331(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 56;
        break;
    }
    total
}
macro_rules! sum15332 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15332(values: &[i32]) -> i32 {
    let literal = sum15332!(63, 26);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum15333 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15333(values: &[i32]) -> i32 {
    let literal = sum15333!(70, 37);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek15334(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 77
}
pub fn classify15335<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 84 => name,
        Some(_) => "other",
        None => "empty",
    }
}
