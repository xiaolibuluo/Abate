
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome21860<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record21860<T> {
    pub value: T,
    pub tag: u32,
}
macro_rules! sum21861 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21861(values: &[i32]) -> i32 {
    let literal = sum21861!(9, 11);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek21862(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 16
}
pub fn peek21863(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 23
}
macro_rules! sum21864 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21864(values: &[i32]) -> i32 {
    let literal = sum21864!(30, 44);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn parse21865(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 37)
}
macro_rules! sum21866 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold21866(values: &[i32]) -> i32 {
    let literal = sum21866!(44, 66);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub fn peek21867(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 51
}
pub fn loops21868(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 58;
        break;
    }
    total
}
pub fn parse21869(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + 65)
}
#[derive(Debug, Clone)]
pub enum Outcome218610<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record218610<T> {
    pub value: T,
    pub tag: u32,
}
