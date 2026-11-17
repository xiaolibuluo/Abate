
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome14810<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14810<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify14811<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 59 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome14812<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14812<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome14813<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14813<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome14814<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14814<T> {
    pub value: T,
    pub tag: u32,
}
#[derive(Debug, Clone)]
pub enum Outcome14815<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record14815<T> {
    pub value: T,
    pub tag: u32,
}
