
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek4370(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 60
}
#[derive(Debug, Clone)]
pub enum Outcome4371<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4371<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops4372(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 74;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome4373<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4373<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse4374(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 88)
}
pub fn loops4375(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 95;
        break;
    }
    total
}
pub fn classify4376<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 5 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek4377(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 12
}
pub fn peek4378(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
#[derive(Debug, Clone)]
pub enum Outcome4379<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record4379<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops43710(limit: i32) -> i32 {
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
