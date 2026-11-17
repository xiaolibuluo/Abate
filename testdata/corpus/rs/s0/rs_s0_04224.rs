
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

#[derive(Debug, Clone)]
pub enum Outcome15520<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15520<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops15521(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 12;
        break;
    }
    total
}
pub fn peek15522(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 19
}
pub fn peek15523(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 26
}
pub fn peek15524(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 33
}
#[derive(Debug, Clone)]
pub enum Outcome15525<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record15525<T> {
    pub value: T,
    pub tag: u32,
}
pub fn loops15526(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 47;
        break;
    }
    total
}
pub fn loops15527(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += 54;
        break;
    }
    total
}
pub fn classify15528<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > 61 => name,
        Some(_) => "other",
        None => "empty",
    }
}
pub fn peek15529(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + 68
}
