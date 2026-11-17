
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode10290 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10290<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10290 for Packet10290<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify10291<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 100 => name,
        Some(_) => "other",
        None => "empty",
    }
}
