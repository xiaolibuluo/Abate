
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome8970<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record8970<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum8971 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8971(values: &[i32]) -> i32 {
    let literal = sum8971!(33, 16);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops8972(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 40;
        break;
    }
    total
}
pub fn loops8973(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 47;
        break;
    }
    total
}
pub fn classify8974<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 54 => name,
        Some(_) => "other",
        None => "empty",
    }
}
