
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify26020<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 75 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify26021<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 82 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify26022<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 89 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek26023(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 96
}
macro_rules! sum26024 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold26024(values: &[i32]) -> i32 {
    let literal = sum26024!(6, 28);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
