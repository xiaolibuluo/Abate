
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome9980<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record9980<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify9981<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 85 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn parse9982(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 92)
}
pub fn parse9983(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 99)
}
macro_rules! sum9984 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold9984(values: &[i32]) -> i32 {
    let literal = sum9984!(9, 24);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek9985(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 16
}
pub fn parse9986(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 23)
}
