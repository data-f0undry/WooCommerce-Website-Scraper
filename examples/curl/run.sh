#!/usr/bin/env bash
# REST API - works anywhere curl exists
set -euo pipefail
ACTOR="data_foundry~woocommerce-scraper"
: "${APIFY_TOKEN:?set APIFY_TOKEN first}"

RUN=$(curl -s -X POST "https://api.apify.com/v2/acts/${ACTOR}/runs?token=${APIFY_TOKEN}" \
  -H "Content-Type: application/json" \
  -d '@input.json')

DATASET=$(printf '%s' "$RUN" | grep -o '"defaultDatasetId":"[^"]*"' | cut -d'"' -f4)
echo "dataset: $DATASET"

curl -s "https://api.apify.com/v2/datasets/${DATASET}/items?token=${APIFY_TOKEN}&limit=5"
