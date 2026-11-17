
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub fn loops15060(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 86;
        break;
    }
    total
}
pub fn classify15061<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 93 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn classify15062<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 100 => name,
        Some(_) => "other",
        None => "empty",
    }
}
#[derive(Debug, Clone)]
pub enum Outcome15063<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15063<T> {
    pub value: T,
    pub tag: u32,
}
pub fn classify15064<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 17 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn loops15065(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 24;
        break;
    }
    total
}
pub fn peek15066(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 31
}
#[derive(Debug, Clone)]
pub enum Outcome15067<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15067<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops15068(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 45;
        break;
    }
    total
}
