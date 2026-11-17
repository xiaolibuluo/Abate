
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops26180(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 89;
        break;
    }
    total
}
pub fn parse26181(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 96)
}
pub fn loops26182(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 6;
        break;
    }
    total
}
macro_rules! sum26183 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold26183(values: &[i32]) -> i32 {
    let literal = sum26183!(13, 68);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse26184(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 20)
}
macro_rules! sum26185 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold26185(values: &[i32]) -> i32 {
    let literal = sum26185!(27, 29);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn loops26186(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 34;
        break;
    }
    total
}
pub fn parse26187(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 41)
}
