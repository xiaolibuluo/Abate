
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops12940(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 46;
        break;
    }
    total
}
pub fn classify12941<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 53 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify12942<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 60 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops12943(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 67;
        break;
    }
    total
}
macro_rules! sum12944 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12944(values: &[i32]) -> i32 {
    let literal = sum12944!(74, 22);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum12945 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold12945(values: &[i32]) -> i32 {
    let literal = sum12945!(81, 33);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
