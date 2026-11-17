
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome11310<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record11310<T> {
    pub value: T,
    pub tag: u32,
}
pub fn peek11311(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 68
}
pub fn classify11312<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 75 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse11313(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 82)
}
