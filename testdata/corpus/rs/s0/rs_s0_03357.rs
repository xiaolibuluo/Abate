
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek12370(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 81
}
pub fn classify12371<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 88 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops12372(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 95;
        break;
    }
    total
}
pub fn classify12373<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 5 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum12374 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12374(values: &[i32]) -> i32 {
    let literal = sum12374!(12, 50);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
