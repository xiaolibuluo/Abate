
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum19620 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold19620(values: &[i32]) -> i32 {
    let literal = sum19620!(97, 18);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops19621(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 7;
        break;
    }
    total
}
pub fn classify19622<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 14 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek19623(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 21
}
