
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode10480 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet10480<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode10480 for Packet10480<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum10481 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10481(values: &[i32]) -> i32 {
    let literal = sum10481!(56, 36);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum10482 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10482(values: &[i32]) -> i32 {
    let literal = sum10482!(63, 47);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum10483 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10483(values: &[i32]) -> i32 {
    let literal = sum10483!(70, 58);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
macro_rules! sum10484 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold10484(values: &[i32]) -> i32 {
    let literal = sum10484!(77, 69);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
