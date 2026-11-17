
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome3120<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record3120<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops3121(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 91;
        break;
    }
    total
}
pub fn peek3122(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 98
}
macro_rules! sum3123 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3123(values: &[i32]) -> i32 {
    let literal = sum3123!(8, 30);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify3124<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 15 => name,
        Some(_) => "other",
        None => "empty",
    }
}
