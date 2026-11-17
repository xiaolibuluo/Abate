
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify23390<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 51 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub trait Encode23391 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet23391<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode23391 for Packet23391<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
