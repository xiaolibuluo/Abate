
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode15410 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15410<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15410 for Packet15410<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum15411 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15411(values: &[i32]) -> i32 {
    let literal = sum15411!(63, 10);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
#[derive(Debug, Clone)]
pub enum Outcome15412<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15412<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome15413<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15413<T> {
    pub value: T,
    pub tag: u32,
}
pub trait Encode15414 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet15414<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode15414 for Packet15414<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum15415 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15415(values: &[i32]) -> i32 {
    let literal = sum15415!(91, 54);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum15416 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold15416(values: &[i32]) -> i32 {
    let literal = sum15416!(98, 65);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
