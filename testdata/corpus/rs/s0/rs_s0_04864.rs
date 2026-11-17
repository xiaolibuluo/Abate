
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome17640<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record17640<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse17641(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 52)
}
pub fn classify17642<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 59 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify17643<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 66 => name,
        Some(_) => "other",
        None => "empty",
    }
}
