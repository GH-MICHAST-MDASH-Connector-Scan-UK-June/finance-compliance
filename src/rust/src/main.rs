use std::collections::HashMap;

fn main() {
    println!("finance-compliance - Rust Application");
    process_data();
}

fn process_data() {
    let mut data: HashMap<&str, &str> = HashMap::new();
    data.insert("app", "finance-compliance");
    data.insert("status", "running");
    
    for (key, value) in &data {
        println!("{}: {}", key, value);
    }
}
