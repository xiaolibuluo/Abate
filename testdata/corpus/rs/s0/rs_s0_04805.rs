
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome17390<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17390<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify17391<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 18 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome17392<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17392<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse17393(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 32)
}
macro_rules! sum17394 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17394(values: &[i32]) -> i32 {
    let literal = sum17394!(39, 26);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify17395<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 46 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops17396(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 53;
        break;
    }
    total
}
macro_rules! sum17397 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold17397(values: &[i32]) -> i32 {
    let literal = sum17397!(60, 59);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
