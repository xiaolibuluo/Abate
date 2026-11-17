
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20840<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20840<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome20841<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20841<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify20842<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 48 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops20843(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 55;
        break;
    }
    total
}
pub fn peek20844(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 62
}
pub fn loops20845(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 69;
        break;
    }
    total
}
pub fn parse20846(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 76)
}
pub fn peek20847(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 83
}
macro_rules! sum20848 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20848(values: &[i32]) -> i32 {
    let literal = sum20848!(90, 45);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
