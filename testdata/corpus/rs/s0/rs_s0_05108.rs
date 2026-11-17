
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum18440 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18440(values: &[i32]) -> i32 {
    let literal = sum18440!(18, 46);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum18441 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18441(values: &[i32]) -> i32 {
    let literal = sum18441!(25, 57);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify18442<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 32 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome18443<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record18443<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify18444<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 46 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum18445 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18445(values: &[i32]) -> i32 {
    let literal = sum18445!(53, 40);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum18446 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18446(values: &[i32]) -> i32 {
    let literal = sum18446!(60, 51);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek18447(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 67
}
pub fn loops18448(limit: i32) -> i32 {
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
