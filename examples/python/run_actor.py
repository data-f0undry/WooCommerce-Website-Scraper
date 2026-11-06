"""Python - official apify-client. Reads input.json from the repo root."""
import json
import os
from pathlib import Path

from apify_client import ApifyClient

ACTOR = "data_foundry/woocommerce-scraper"
run_input = json.loads(Path("input.json").read_text(encoding="utf-8"))

client = ApifyClient(os.environ["APIFY_TOKEN"])
run = client.actor(ACTOR).call(run_input=run_input)
print("dataset:", run["defaultDatasetId"])
for item in client.dataset(run["defaultDatasetId"]).iterate_items():
    print(item)
