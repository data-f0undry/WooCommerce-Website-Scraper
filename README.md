[![Run on Apify](https://img.shields.io/badge/Run%20on-Apify-141414?logo=apify&logoColor=white&labelColor=141414)](https://apify.com/data_foundry/woocommerce-scraper)

# WooCommerce Website Scraper

**WooCommerce Website Scraper** is a high-performance, feature-rich Apify Actor designed to scrape product and category data from any WooCommerce store.

This is an **enhanced version** of standard scrapers, offering significantly deeper data extraction while maintaining full compatibility with existing workflows.

## What is WooCommerce?
WooCommerce is the world's most popular open-source e-commerce plugin for WordPress. It powers over 3 million active stores. 
While WooCommerce provides a REST API, accessing it often requires authentication keys (Consumer Key/Secret) which are not available for public data extraction.
This scraper bypasses the need for API keys, allowing you to extract public product data (prices, descriptions, images, stock status) from any WooCommerce-powered site without credentials.

## Features 🚀

*   **⚡ Enhanced Data Extraction**: detailed stock info, purchase limits, rich images, and deep taxonomy (slugs/IDs).
*   **📂 Full Resource Support**: Scrape `products`, `categories`, and `payment_methods`.
*   **🔍 Search & Filters**: Support for search queries, price ranges, and sale status.
*   **⚙️ Data Parity**: Output format is fully compatible with standard schemas but enriched with extra fields.
*   **⏩ Pagination**: Auto-traverses all pages to get complete datasets.
*   **🛡️ WAF Bypass**: Uses specialized HTTP clients to bypass Cloudflare and other protections.

## Usage 💡

### Input Parameters

| Parameter | Type | Description |
| :--- | :--- | :--- |
| `url` | Array | List of WooCommerce store URLs (e.g., `https://woocommerce.com/`). |
| `resource` | String | What to scrape: `products` (default), `categories`, or `payment_methods`. |
| `limit` | Integer | Max items to scrape (0 for unlimited). |
| `filters` | Object | Optional filters (`search`, `min_price`, `on_sale`, etc.). |
| `sort` | String | Sort by `date`, `price`, `popularity`. |
| `order` | String | Sort order: `asc` or `desc`. |

### Output

The scraper delivers results in a JSON format that extends the standard WooCommerce schema.

### Output Fields
*   `id`, `name`, `slug`, `permalink`, `sku`
*   `price`, `regular_price`, `sale_price`
*   `images` (List of URLs), `categories`
*   `attributes`, `variations`
*   `stock_availability`: Detailed text like "Only 2 left in stock!" or "Available on backorder".
*   `add_to_cart`: Purchase constraints (minimum, maximum, step).
*   `images_detailed`: Full list of image objects containing `src`, `thumbnail`, `alt` (SEO text), and `srcset`.
*   `brands`: Brand information (if available).
*   `categories_detailed` & `tags_detailed`: Full objects with IDs and slugs.

## Tips

- If the scraper returns 0 results, ensure the target site is actually built with WooCommerce.

## Examples 📋

Here are comprehensive, real-world examples using `https://woocommerce.com`.

<details>
<summary><strong>1. Scrape Products 📦</strong></summary>

**Input:**
```json
{
    "url": [{ "url": "https://woocommerce.com" }],
    "resource": "products",
    "limit": 1
}
```

**Output:**
```json
[
    {
        "id": 18734006072352,
        "name": "Multiple Shipping Addresses",
        "slug": "multiple-shipping-customer-addresses",
        "permalink": "https://woocommerce.com/products/multiple-shipping-customer-addresses/",
        "type": "simple",
        "status": null,
        "description": "<ul class=\"wccom-tick-list-primary\">\n<li><b>Ship to Multiple Addresses:</b> Ship individual items... (truncated for brevity) ...</li></ul>...",
        "short_description": "<p><b>Allow customers to ship individual items to multiple addresses in a single order.</b></p>",
        "sku": "",
        "price": "3900",
        "regular_price": "3900",
        "sale_price": "3900",
        "on_sale": null,
        "prices": {
            "price": "3900",
            "regular_price": "3900",
            "sale_price": "3900",
            "price_range": null,
            "currency_code": "USD",
            "currency_symbol": "USD $",
            "currency_minor_unit": 2,
            "currency_decimal_separator": ".",
            "currency_thousand_separator": ",",
            "currency_prefix": "USD $",
            "currency_suffix": ""
        },
        "price_html": "<span class=\"woocommerce-Price-amount amount\"><span class=\"woocommerce-Price-currencySymbol\">USD &#036;</span>39.00</span>",
        "average_rating": "0",
        "review_count": 0,
        "images": [
            "https://woocommerce.com/wp-content/uploads/2025/12/Multiple-Shipping-Addresses-Plugin.png",
            "https://woocommerce.com/wp-content/uploads/2025/12/Use-a-Dropdown-to-Show-Saved-Addresses-On-Checkout-1.png",
            "https://woocommerce.com/wp-content/uploads/2025/12/Ship-to-Multiple-Addresses-in-a-Single-Order-1.png",
            "https://woocommerce.com/wp-content/uploads/2025/12/Save-Multiple-Shipping-Addresses-For-Future-Checkouts-1.png",
            "https://woocommerce.com/wp-content/uploads/2025/12/Personalized-Shipping-Address-Form-1.png",
            "https://woocommerce.com/wp-content/uploads/2025/12/Custom-Emails-for-Shipping-Notifications-1.png",
            "https://woocommerce.com/wp-content/uploads/2025/12/Allow-Customers-To-Add-Shipping-Addresses-Directly-from-Checkout-1.png"
        ],
        "categories": [
            "WooCommerce extensions",
            "Shipping, delivery and fulfillment",
            "Delivery options and enhancements",
            "Store content and customizations",
            "Cart and checkout features"
        ],
        "tags": [],
        "attributes": [],
        "variations": [],
        "has_options": false,
        "is_purchasable": true,
        "in_stock": true,
        "stock_quantity": null,
        "stock_availability": {
            "text": "",
            "class": "in-stock"
        },
        "add_to_cart": {
            "text": "Add to cart",
            "description": "Add to cart: &ldquo;Multiple Shipping Addresses&rdquo;",
            "url": "/wp-json/wc/store/v1/products?page=1&#038;per_page=20&#038;orderby=date&#038;order=desc&#038;add-to-cart=18734006072352",
            "single_text": "Add to cart",
            "minimum": 1,
            "maximum": 9999,
            "multiple_of": 1
        },
        "images_detailed": [
            {
                "id": 18734006124909,
                "src": "https://woocommerce.com/wp-content/uploads/2025/12/Multiple-Shipping-Addresses-Plugin.png",
                "thumbnail": "https://woocommerce.com/wp-content/uploads/2025/12/Multiple-Shipping-Addresses-Plugin.png?w=160",
                "srcset": "https://woocommerce.com/wp-content/uploads/2025/12/Multiple-Shipping-Addresses-Plugin.png 160w, ...",
                "sizes": "(max-width: 160px) 100vw, 160px",
                "name": "Multiple Shipping Addresses Plugin",
                "alt": "Multiple Shipping Addresses Plugin"
            }
        ],
        "brands": [],
        "categories_detailed": [
            {
                "id": 1021,
                "name": "WooCommerce extensions",
                "slug": "woocommerce-extensions",
                "link": "https://woocommerce.com/product-category/woocommerce-extensions/"
            },
            {
                "id": 28685,
                "name": "Shipping, delivery and fulfillment",
                "slug": "shipping-delivery-and-fulfillment",
                "link": "https://woocommerce.com/product-category/woocommerce-extensions/shipping-delivery-and-fulfillment/"
            }
        ],
        "tags_detailed": []
    }
]
```
</details>

<details>
<summary><strong>2. Scrape Categories 📂</strong></summary>

**Input:**
```json
{
    "url": [{ "url": "https://woocommerce.com" }],
    "resource": "categories"
}
```

**Output:**
```json
[
    {
        "id": 1028,
        "name": "Accounting",
        "slug": "accounting-extensions",
        "description": "Accounting Extensions for WooCommerce",
        "count": 23,
        "parent": 1888,
        "image": null,
        "link": "https://woocommerce.com/product-category/accounting-extensions/"
    }
]
```
</details>

<details>
<summary><strong>3. Scrape Payment Methods 💳</strong></summary>

**Input:**
```json
{
    "url": [{ "url": "https://woocommerce.com" }],
    "resource": "payment_methods"
}
```

**Output:**
```json
[
    {
        "url": "https://woocommerce.com",
        "payment_methods": [
            "woocommerce_payments",
            "ppcp-gateway"
        ]
    }
]
```
</details>

[![Run on Apify](https://img.shields.io/badge/Run%20on-Apify-141414?logo=apify&logoColor=white&labelColor=141414)](https://apify.com/data_foundry/woocommerce-scraper)

## License

Apache-2.0
