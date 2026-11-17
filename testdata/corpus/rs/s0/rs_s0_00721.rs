
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome2620<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record2620<T> {
    pub value: T,
    pub tag: u32,
}
pub fn parse2621(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 23)
}
pub fn classify2622<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 30 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode2623 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2623<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2623 for Packet2623<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub trait Encode2624 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet2624<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode2624 for Packet2624<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
