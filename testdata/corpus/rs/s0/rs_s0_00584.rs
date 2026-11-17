
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum2080 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2080(values: &[i32]) -> i32 {
    let literal = sum2080!(90, 62);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum2081 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2081(values: &[i32]) -> i32 {
    let literal = sum2081!(97, 12);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse2082(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 7)
}
macro_rules! sum2083 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold2083(values: &[i32]) -> i32 {
    let literal = sum2083!(14, 34);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
