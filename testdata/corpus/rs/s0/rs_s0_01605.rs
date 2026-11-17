
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum5930 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold5930(values: &[i32]) -> i32 {
    let literal = sum5930!(51, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse5931(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 58)
}
pub fn classify5932<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 65 => name,
        Some(_) => "other",
        None => "empty",
    }
}
