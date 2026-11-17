
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify2280<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 59 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops2281(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 66;
        break;
    }
    total
}
pub fn classify2282<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 73 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum2283 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2283(values: &[i32]) -> i32 {
    let literal = sum2283!(80, 52);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek2284(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 87
}
macro_rules! sum2285 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2285(values: &[i32]) -> i32 {
    let literal = sum2285!(94, 13);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify2286<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 101 => name,
        Some(_) => "other",
        None => "empty",
    }
}
