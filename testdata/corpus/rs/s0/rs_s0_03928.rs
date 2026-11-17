
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops14380(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 75;
        break;
    }
    total
}
#[derive(Debug, Clone)]
pub enum Outcome14381<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14381<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum14382 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14382(values: &[i32]) -> i32 {
    let literal = sum14382!(89, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum14383 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold14383(values: &[i32]) -> i32 {
    let literal = sum14383!(96, 43);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse14384(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 6)
}
pub fn loops14385(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 13;
        break;
    }
    total
}
