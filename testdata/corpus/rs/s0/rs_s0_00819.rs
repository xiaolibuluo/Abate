
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify3000<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 25 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek3001(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 32
}
macro_rules! sum3002 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold3002(values: &[i32]) -> i32 {
    let literal = sum3002!(39, 57);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
