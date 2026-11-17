
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse6400(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 80)
}
pub fn classify6401<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 87 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum6402 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold6402(values: &[i32]) -> i32 {
    let literal = sum6402!(94, 58);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
