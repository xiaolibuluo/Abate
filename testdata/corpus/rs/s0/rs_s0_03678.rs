
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify13470<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 56 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome13471<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13471<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome13472<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13472<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse13473(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 77)
}
pub fn parse13474(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 84)
}
pub fn parse13475(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 91)
}
pub fn peek13476(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 98
}
pub fn loops13477(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 8;
        break;
    }
    total
}
macro_rules! sum13478 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13478(values: &[i32]) -> i32 {
    let literal = sum13478!(15, 10);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
