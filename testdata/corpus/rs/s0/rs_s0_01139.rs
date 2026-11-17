
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify4150<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 65 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum4151 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4151(values: &[i32]) -> i32 {
    let literal = sum4151!(72, 58);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops4152(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 79;
        break;
    }
    total
}
pub fn loops4153(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 86;
        break;
    }
    total
}
macro_rules! sum4154 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4154(values: &[i32]) -> i32 {
    let literal = sum4154!(93, 30);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
