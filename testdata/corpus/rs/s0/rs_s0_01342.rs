
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek4950(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 38
}
pub trait Encode4951 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet4951<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode4951 for Packet4951<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify4952<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 52 => name,
        Some(_) => "other",
        None => "empty",
    }
}
