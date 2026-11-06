"""Run the WooCommerce Website Scraper Actor and print the first few results.

    pip install apify-client
    export APIFY_TOKEN=...        # https://console.apify.com/settings/integrations
    python run_actor.py
"""

import os

from apify_client import ApifyClient

ACTOR = "woocommerce-scraper"
INPUT = {"url": "https://example-store.com", "resource": "products", "limit": 100}


def main() -> None:
    client = ApifyClient(os.environ["APIFY_TOKEN"])
    run = client.actor(ACTOR).call(run_input=INPUT)
    print("dataset:", run["defaultDatasetId"])

    for item in client.dataset(run["defaultDatasetId"]).list_items(limit=5).items:
        print(item)


if __name__ == "__main__":
    main()
