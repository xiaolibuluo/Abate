
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn peek16650(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
pub fn classify16651<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 26 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops16652(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 33;
        break;
    }
    total
}
