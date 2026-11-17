
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse25900(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 16)
}
#[derive(Debug, Clone)]
pub enum Outcome25901<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25901<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify25902<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 30 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse25903(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 37)
}
