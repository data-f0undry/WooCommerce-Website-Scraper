"""Python - official apify-client."""
import os
from apify_client import ApifyClient

client = ApifyClient(os.environ["APIFY_TOKEN"])
run = client.actor("data_foundry/woocommerce-scraper").call(run_input={
    "url": "https://example-store.com",
    "resource": "products",
    "limit": 100
})
print("dataset:", run["defaultDatasetId"])
for item in client.dataset(run["defaultDatasetId"]).iterate_items():
    print(item)
