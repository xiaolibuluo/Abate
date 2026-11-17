
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome19560<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19560<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify19561<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 26 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse19562(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 33)
}
