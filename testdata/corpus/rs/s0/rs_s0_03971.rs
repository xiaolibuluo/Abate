
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum14560 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14560(values: &[i32]) -> i32 {
    let literal = sum14560!(18, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome14561<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14561<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek14562(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 32
}
pub fn peek14563(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
pub fn parse14564(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 46)
}
pub fn loops14565(limit: i32) -> i32 {
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
pub fn peek14566(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 60
}
pub fn peek14567(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 67
}
pub fn loops14568(limit: i32) -> i32 {
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
pub fn classify14569<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 81 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum145610 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold145610(values: &[i32]) -> i32 {
    let literal = sum145610!(88, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
