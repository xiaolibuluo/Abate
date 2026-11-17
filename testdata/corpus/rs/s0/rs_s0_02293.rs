
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify8580<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 101 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum8581 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8581(values: &[i32]) -> i32 {
    let literal = sum8581!(11, 48);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify8582<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 18 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify8583<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 25 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify8584<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 32 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum8585 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold8585(values: &[i32]) -> i32 {
    let literal = sum8585!(39, 31);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
