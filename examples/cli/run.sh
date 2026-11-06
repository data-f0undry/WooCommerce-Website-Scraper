#!/usr/bin/env bash
# Apify CLI - install with: npm install -g apify-cli
set -euo pipefail
apify call data_foundry/woocommerce-scraper --input-file input.json --output-dataset
