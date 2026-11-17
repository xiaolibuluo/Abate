
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome13340<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record13340<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops13341(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 88;
        break;
    }
    total
}
macro_rules! sum13342 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13342(values: &[i32]) -> i32 {
    let literal = sum13342!(95, 36);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse13343(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 5)
}
pub fn classify13344<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 12 => name,
        Some(_) => "other",
        None => "empty",
    }
}
