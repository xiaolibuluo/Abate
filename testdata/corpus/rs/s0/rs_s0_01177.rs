
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse4270(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 27)
}
pub fn classify4271<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 34 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum4272 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4272(values: &[i32]) -> i32 {
    let literal = sum4272!(41, 31);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome4273<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4273<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops4274(limit: i32) -> i32 {
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
pub fn classify4275<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 62 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum4276 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4276(values: &[i32]) -> i32 {
    let literal = sum4276!(69, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek4277(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 76
}
