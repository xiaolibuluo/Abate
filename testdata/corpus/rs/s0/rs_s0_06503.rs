
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify23400<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 64 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome23401<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record23401<T> {
    pub value: T,
    pub tag: u32,
}
