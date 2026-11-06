// Rust - no crates. Runs the Actor through the REST API with curl.
//
//   rustc main.rs && ./main
//   # or: cargo run
//
// Keeping this dependency-free means it builds offline and in any CI image
// that has Rust and curl, without waiting on a crate registry.

use std::process::Command;

const ACTOR: &str = "data_foundry~woocommerce-scraper";

fn main() {
    let token = std::env::var("APIFY_TOKEN").expect("set APIFY_TOKEN first");
    let input = std::fs::read_to_string("input.json").expect("input.json not found");

    let run_url = format!(
        "https://api.apify.com/v2/acts/{}/runs?token={}",
        ACTOR, token
    );

    let output = Command::new("curl")
        .args([
            "-s",
            "-X",
            "POST",
            &run_url,
            "-H",
            "Content-Type: application/json",
            "-d",
            &input,
        ])
        .output()
        .expect("curl is required");

    let body = String::from_utf8_lossy(&output.stdout);
    println!("{}", body);

    let dataset_id = body
        .split("\"defaultDatasetId\":\"")
        .nth(1)
        .and_then(|rest| rest.split('"').next())
        .expect("could not read defaultDatasetId");

    let items_url = format!(
        "https://api.apify.com/v2/datasets/{}/items?token={}&limit=5",
        dataset_id, token
    );

    let items = Command::new("curl")
        .args(["-s", &items_url])
        .output()
        .expect("curl is required");

    println!("{}", String::from_utf8_lossy(&items.stdout));
}
