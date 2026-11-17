
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify22060<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 68 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum22061 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22061(values: &[i32]) -> i32 {
    let literal = sum22061!(75, 29);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse22062(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 82)
}
pub fn parse22063(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 89)
}
macro_rules! sum22064 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold22064(values: &[i32]) -> i32 {
    let literal = sum22064!(96, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
