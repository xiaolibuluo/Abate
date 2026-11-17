
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome20850<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record20850<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify20851<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 54 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops20852(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 61;
        break;
    }
    total
}
macro_rules! sum20853 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20853(values: &[i32]) -> i32 {
    let literal = sum20853!(68, 58);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify20854<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 75 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify20855<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 82 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify20856<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 89 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum20857 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold20857(values: &[i32]) -> i32 {
    let literal = sum20857!(96, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
