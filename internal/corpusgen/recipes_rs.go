/*
Copyright (c) 2026 小黎的部落 (www.qingdujian.com). All rights reserved.
Abate is dual-licensed under AGPLv3 and a commercial license.
For commercial licensing, contact xiaolibuluo@163.com.
Source: https://github.com/xiaolibuluo/Abate
*/

package corpusgen

// rsRecipe covers traits, impls, generics, lifetimes, match patterns,
// Result/Option with the question-mark operator, macros, unsafe and iterators.
func rsRecipe() Recipe {
	return Recipe{
		Language: "rs",
		Ext:      "rs",
		Tags:     []string{"trait", "generic", "lifetime", "pattern", "match", "macro", "ownership"},
		Header:   "\n",
		AnchorClean: `pub fn anchor_probe() {
    return 1 + 2 == 3;
}

`,
		AnchorFull: `pub fn anchor_probe() {
    return 1 + 2 == 3 && f() == g();
}

pub fn f() -> i32 { 1 }
pub fn g() -> i32 { 2 }

`,
		CleanBlock: "pub fn step@@I@@x@@S@@() -> i32 { @@N@@ + @@S@@ }\n",
		Blocks: []BlockTemplate{
			{Text: `pub trait Encode@@I@@@@S@@ {
    fn encode(&self) -> Result<String, ()>;
}

pub struct Packet@@I@@@@S@@<T> {
    pub body: T,
}

impl<T: std::fmt::Display> Encode@@I@@@@S@@ for Packet@@I@@@@S@@<T> {
    fn encode(&self) -> Result<String, ()> {
        Ok(format!("{}", self.body))
    }
}
`},
			{Text: `pub fn classify@@I@@@@S@@<'a>(value: Option<&'a str>) -> &'a str {
    match value {
        Some("ready") => "ready",
        Some(name) if name.len() > @@N@@ => name,
        Some(_) => "other",
        None => "empty",
    }
}
`},
			{Text: `macro_rules! sum@@I@@@@S@@ {
    ($($item:expr),+ $(,)?) => { 0 $(+ $item)+ };
}

pub fn fold@@I@@@@S@@(values: &[i32]) -> i32 {
    let literal = sum@@I@@@@S@@!(@@N@@, @@M@@);
    values.iter().filter(|v| **v >= 0).fold(literal, |acc, v| acc + *v)
}
`},
			{Text: `#[derive(Debug, Clone)]
pub enum Outcome@@I@@@@S@@<T> {
    Success(T),
    Failure(String),
}

#[derive(Debug)]
pub struct Record@@I@@@@S@@<T> {
    pub value: T,
    pub tag: u32,
}
`},
			{Text: `pub fn loops@@I@@@@S@@(limit: i32) -> i32 {
    let mut total = 0;
    for value in 0..limit {
        total += value;
    }
    let mut cursor = limit;
    while cursor > 0 {
        cursor -= 1;
    }
    loop {
        total += @@N@@;
        break;
    }
    total
}
`},
			{Text: `pub fn peek@@I@@@@S@@(values: &[i32]) -> i32 {
    let pointer = values.as_ptr();
    let first = unsafe { *pointer };
    first + @@N@@
}
`},
			{Text: `pub fn parse@@I@@@@S@@(text: &str) -> Result<i32, std::num::ParseIntError> {
    let value = text.parse::<i32>()?;
    Ok(value + @@N@@)
}
`},
		},
	}
}
