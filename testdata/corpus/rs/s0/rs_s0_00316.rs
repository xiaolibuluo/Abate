
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify1080<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 51 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse1081(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 58)
}
macro_rules! sum1082 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold1082(values: &[i32]) -> i32 {
    let literal = sum1082!(65, 55);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
