
pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

pub trait Encode13830 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13830<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13830 for Packet13830<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum13831 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13831(values: &[i32]) -> i32 {
    let literal = sum13831!(46, 63);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
pub trait Encode13832 {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet13832<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode13832 for Packet13832<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
macro_rules! sum13833 {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold13833(values: &[i32]) -> i32 {
    let literal = sum13833!(60, 24);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
