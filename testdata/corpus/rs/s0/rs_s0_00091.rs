
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops290(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 91;
        break;
    }
    total
}
macro_rules! sum291 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold291(values: &[i32]) -> i32 {
    let literal = sum291!(98, 40);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum292 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold292(values: &[i32]) -> i32 {
    let literal = sum292!(8, 51);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn classify293<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 15 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify294<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 22 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum295 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold295(values: &[i32]) -> i32 {
    let literal = sum295!(29, 23);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
