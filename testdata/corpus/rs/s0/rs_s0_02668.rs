
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek9900(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 71
}
#[derive(Debug, Clone)]
pub enum Outcome9901<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9901<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum9902 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9902(values: &[i32]) -> i32 {
    let literal = sum9902!(85, 68);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek9903(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 92
}
pub fn classify9904<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 99 => name,
        Some(_) => "other",
        None => "empty",
    }
}
