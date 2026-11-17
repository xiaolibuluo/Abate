
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify13870<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 91 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify13871<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 98 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum13872 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13872(values: &[i32]) -> i32 {
    let literal = sum13872!(8, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops13873(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 15;
        break;
    }
    total
}
macro_rules! sum13874 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13874(values: &[i32]) -> i32 {
    let literal = sum13874!(22, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
