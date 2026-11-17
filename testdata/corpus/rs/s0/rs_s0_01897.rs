
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify7020<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 13 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum7021 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7021(values: &[i32]) -> i32 {
    let literal = sum7021!(20, 54);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek7022(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 27
}
macro_rules! sum7023 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7023(values: &[i32]) -> i32 {
    let literal = sum7023!(34, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify7024<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 41 => name,
        Some(_) => "other",
        None => "empty",
    }
}
