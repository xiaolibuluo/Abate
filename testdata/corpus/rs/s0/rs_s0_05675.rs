
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode20520 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet20520<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode20520 for Packet20520<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify20521<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 13 => name,
        Some(_) => "other",
        None => "empty",
    }
}
