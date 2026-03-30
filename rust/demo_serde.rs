pub fn compute_hello(json: &str) -> String {
    let v: serde_json::Value = serde_json::from_str(json).unwrap();
    let greeting = v["greet"].as_str().unwrap();
    
    format!("Hello, {}!", greeting)
}

#[cfg(test)]
mod tests {
    use super::*;

    #[test]
    fn test_compute_hello_success() {
        let json = r#"{"greet": "Rust World"}"#;
        assert_eq!(compute_hello(json), "Hello, Rust World!");
    }

    #[test]
    fn test_compute_hello_accepts_empty_value() {
        let json = r#"{"greet": ""}"#;
        assert_eq!(compute_hello(json), "Hello, !");
    }

    #[test]
    #[should_panic]
    fn test_compute_hello_panics_on_invalid_json() {
        compute_hello("not json");
    }

    #[test]
    #[should_panic]
    fn test_compute_hello_panics_on_missing_key() {
        compute_hello(r#"{"wrong_key": "val"}"#);
    }
}
