
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20630<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20630<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops20631(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 59;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome20632<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20632<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify20633<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 73 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum20634 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20634(values: &[i32]) -> i32 {
    let literal = sum20634!(80, 37);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
