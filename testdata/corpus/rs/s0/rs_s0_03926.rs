
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops14370(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 62;
        break;
    }
    total
}
pub fn parse14371(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 69)
}
#[derive(Debug, Clone)]
pub enum Outcome14372<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14372<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum14373 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14373(values: &[i32]) -> i32 {
    let literal = sum14373!(83, 36);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome14374<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14374<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum14375 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14375(values: &[i32]) -> i32 {
    let literal = sum14375!(97, 58);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
