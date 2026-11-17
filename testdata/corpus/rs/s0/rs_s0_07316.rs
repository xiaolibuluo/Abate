
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify26400<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 84 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify26401<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 91 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse26402(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 98)
}
#[derive(Debug, Clone)]
pub enum Outcome26403<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record26403<T> {
    pub value: T,
    pub tag: u32,
}
