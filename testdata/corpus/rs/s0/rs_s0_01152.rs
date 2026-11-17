
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops4200(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 33;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome4201<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4201<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops4202(limit: i32) -> i32 {
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
pub fn classify4203<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 54 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse4204(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 61)
}
#[derive(Debug, Clone)]
pub enum Outcome4205<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4205<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum4206 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4206(values: &[i32]) -> i32 {
    let literal = sum4206!(75, 26);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
