
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum15720 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15720(values: &[i32]) -> i32 {
    let literal = sum15720!(71, 33);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome15721<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15721<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome15722<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15722<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify15723<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 92 => name,
        Some(_) => "other",
        None => "empty",
    }
}
