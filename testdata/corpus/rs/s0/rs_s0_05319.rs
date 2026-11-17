
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome19160<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19160<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome19161<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record19161<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify19162<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 95 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse19163(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 5)
}
pub fn classify19164<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 12 => name,
        Some(_) => "other",
        None => "empty",
    }
}
