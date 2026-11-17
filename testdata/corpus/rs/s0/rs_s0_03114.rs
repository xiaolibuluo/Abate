
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify11400<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 81 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify11401<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 88 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify11402<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 95 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse11403(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 5)
}
pub fn classify11404<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 12 => name,
        Some(_) => "other",
        None => "empty",
    }
}
