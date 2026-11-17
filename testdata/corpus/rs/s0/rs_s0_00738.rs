
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops2700(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 23;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome2701<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2701<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome2702<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2702<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify2703<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 44 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse2704(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 51)
}
macro_rules! sum2705 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2705(values: &[i32]) -> i32 {
    let literal = sum2705!(58, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
