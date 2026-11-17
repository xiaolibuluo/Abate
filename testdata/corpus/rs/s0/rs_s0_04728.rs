
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode17090 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet17090<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode17090 for Packet17090<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify17091<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 16 => name,
        Some(_) => "other",
        None => "empty",
    }
}
