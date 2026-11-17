
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops25110(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 56;
        break;
    }
    total
}
pub fn classify25111<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 63 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek25112(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 70
}
macro_rules! sum25113 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25113(values: &[i32]) -> i32 {
    let literal = sum25113!(77, 51);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome25114<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record25114<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum25115 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold25115(values: &[i32]) -> i32 {
    let literal = sum25115!(91, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
