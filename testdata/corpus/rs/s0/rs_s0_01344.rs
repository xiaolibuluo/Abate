
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn parse4960(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 51)
}
macro_rules! sum4961 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4961(values: &[i32]) -> i32 {
    let literal = sum4961!(58, 15);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum4962 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold4962(values: &[i32]) -> i32 {
    let literal = sum4962!(65, 26);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
