
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

macro_rules! sum18520 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold18520(values: &[i32]) -> i32 {
    let literal = sum18520!(25, 41);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse18521(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 32)
}
