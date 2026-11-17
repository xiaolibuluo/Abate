
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops23800(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 99;
        break;
    }
    total
}
pub fn parse23801(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 9)
}
#[derive(Debug, Clone)]
pub enum Outcome23802<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23802<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse23803(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 23)
}
pub fn classify23804<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 30 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops23805(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 37;
        break;
    }
    total
}
pub fn loops23806(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 44;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome23807<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23807<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum23808 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold23808(values: &[i32]) -> i32 {
    let literal = sum23808!(58, 43);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
