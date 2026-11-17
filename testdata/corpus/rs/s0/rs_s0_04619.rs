
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek16720(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 13
}
pub fn classify16721<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 20 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse16722(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 27)
}
pub trait Encode16723 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet16723<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode16723 for Packet16723<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn parse16724(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 41)
}
