
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome5640<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record5640<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum5641 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5641(values: &[i32]) -> i32 {
    let literal = sum5641!(69, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify5642<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 76 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum5643 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5643(values: &[i32]) -> i32 {
    let literal = sum5643!(83, 25);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops5644(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 90;
        break;
    }
    total
}
pub fn classify5645<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 97 => name,
        Some(_) => "other",
        None => "empty",
    }
}
