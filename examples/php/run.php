<?php
// PHP - REST API, no dependencies. Reads input.json.
$token = getenv("APIFY_TOKEN");
$actor = "data_foundry~woocommerce-scraper";
$input = file_get_contents("input.json");

$context = stream_context_create([
    "http" => [
        "method" => "POST",
        "header" => "Content-Type: application/json\r\n",
        "content" => $input,
        "ignore_errors" => true,
    ],
]);

$run = json_decode(file_get_contents(
    "https://api.apify.com/v2/acts/$actor/runs?token=$token",
    false,
    $context
), true);

$datasetId = $run["data"]["defaultDatasetId"] ?? "";
echo "dataset: $datasetId\n";
echo file_get_contents("https://api.apify.com/v2/datasets/$datasetId/items?token=$token&limit=5");
