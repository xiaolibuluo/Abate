
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops13750(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 32;
        break;
    }
    total
}
pub fn peek13751(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 39
}
pub fn classify13752<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 46 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum13753 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13753(values: &[i32]) -> i32 {
    let literal = sum13753!(53, 29);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
