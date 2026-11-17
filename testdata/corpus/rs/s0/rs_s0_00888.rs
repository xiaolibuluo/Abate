
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek3260(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 72
}
pub fn peek3261(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 79
}
pub fn classify3262<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 86 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify3263<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 93 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse3264(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 100)
}
pub trait Encode3265 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet3265<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode3265 for Packet3265<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify3266<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 17 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode3267 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet3267<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode3267 for Packet3267<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
