
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify10390<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 29 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops10391(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 36;
        break;
    }
    total
}
pub fn parse10392(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 43)
}
macro_rules! sum10393 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10393(values: &[i32]) -> i32 {
    let literal = sum10393!(50, 56);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
