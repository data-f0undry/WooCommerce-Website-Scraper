[![Run on Apify](https://img.shields.io/badge/Run%20on-Apify-141414?logo=apify&logoColor=white&labelColor=141414)](https://apify.com/data_foundry/woocommerce-scraper)

# WooCommerce Website Scraper

Scrape any WooCommerce store's products, categories and payment methods through the store's own public Store API. No plugin, no key, no browser.

WooCommerce ships a public Store API that its own front-end calls. This Actor reads that endpoint, so you get prices, sale prices, stock state, images and taxonomy without a Consumer Key or Secret. Pagination, filters and sorting are handled for you.

## Features

- **Products**: prices, sale prices, SKU, stock availability, categories, tags, images, description
- **Categories**: ids, slugs, counts, descriptions
- **Payment methods**: the gateways the store advertises
- **Filters**: category, min/max price, on sale
- **Sort and order**: whatever the Store API supports
- **Pagination** handled, with a hard `limit` you control

## Usage

### Input

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `url` | string | Store base URL. The Actor appends the Store API path. |
| `limit` | integer | Maximum items to return. Set it - a large catalogue is a large dataset. |
| `resource` | string | `products`, `categories` or `payment_methods`. |
| `filters` | object | `category`, `min_price`, `max_price`, `on_sale`. |
| `sort` | string | Sort field, e.g. `price`, `date`, `popularity`. |
| `order` | string | `asc` or `desc`. |

```json
{
  "url": "https://example-store.com",
  "resource": "products",
  "limit": 500,
  "filters": { "category": "hoodies", "on_sale": true },
  "sort": "price",
  "order": "asc"
}
```

### Output

```json
{
  "id": 1234,
  "name": "Organic Hoodie",
  "price": "45.00",
  "regular_price": "55.00",
  "sale_price": "45.00",
  "currency": "USD",
  "is_in_stock": true,
  "on_sale": true,
  "categories_detailed": [{"id": 12, "name": "Hoodies", "slug": "hoodies"}]
}
```

## Run it

**Apify Console** - open [data_foundry/woocommerce-scraper](https://apify.com/data_foundry/woocommerce-scraper) and click Start.

**API (Python)**

```bash
pip install apify-client
```

```python
from apify_client import ApifyClient

client = ApifyClient("<YOUR_API_TOKEN>")
run = client.actor("woocommerce-scraper").call(run_input={"url": "https://example-store.com", "limit": 100})
for item in client.dataset(run["default_dataset_id"]).iterate_items():
    print(item)
```

Or run the included example: `python run_actor.py`.

## Pricing

**$1.50 per 1,000 items** ($0.0015 each). One item is one product, category or payment method. Failed items are not billed.

## Limitations

- **Public Store API only.** Stores that disabled it, or block REST access, return an error rather than data.
- **Hidden and draft products** are not public and will not appear.
- **Variations are embedded** in the product record, not emitted as separate rows.
- **One store per run** - the input is single-store by design.

These are documented deliberately: an Actor that hides its limits generates refund requests.

## Links

- [Actor on Apify](https://apify.com/data_foundry/woocommerce-scraper)
- [All Data Foundry Actors](https://apify.com/data_foundry)
- [Issues](https://github.com/data-f0undry/WooCommerce-Website-Scraper/issues)

## License

Apache-2.0. See [LICENSE](LICENSE).
