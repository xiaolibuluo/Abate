
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops2610(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 100;
        break;
    }
    total
}
pub fn peek2611(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 10
}
macro_rules! sum2612 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2612(values: &[i32]) -> i32 {
    let literal = sum2612!(17, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify2613<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 24 => name,
        Some(_) => "other",
        None => "empty",
    }
}
