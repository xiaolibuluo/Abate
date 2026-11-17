
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode15730 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15730<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15730 for Packet15730<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
pub fn classify15731<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 91 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek15732(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 98
}
macro_rules! sum15733 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15733(values: &[i32]) -> i32 {
    let literal = sum15733!(8, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
