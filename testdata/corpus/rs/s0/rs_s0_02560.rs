
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum9540 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9540(values: &[i32]) -> i32 {
    let literal = sum9540!(88, 38);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome9541<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9541<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum9542 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9542(values: &[i32]) -> i32 {
    let literal = sum9542!(5, 60);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify9543<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 12 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum9544 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9544(values: &[i32]) -> i32 {
    let literal = sum9544!(19, 21);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum9545 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9545(values: &[i32]) -> i32 {
    let literal = sum9545!(26, 32);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
