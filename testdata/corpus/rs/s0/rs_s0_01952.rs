
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn classify7280<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 60 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify7281<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 67 => name,
        Some(_) => "other",
        None => "empty",
    }
}
macro_rules! sum7282 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7282(values: &[i32]) -> i32 {
    let literal = sum7282!(74, 64);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum7283 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7283(values: &[i32]) -> i32 {
    let literal = sum7283!(81, 14);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse7284(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 88)
}
pub fn parse7285(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 95)
}
pub fn loops7286(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 5;
        break;
    }
    total
}
macro_rules! sum7287 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7287(values: &[i32]) -> i32 {
    let literal = sum7287!(12, 58);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse7288(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 19)
}
macro_rules! sum7289 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold7289(values: &[i32]) -> i32 {
    let literal = sum7289!(26, 19);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
