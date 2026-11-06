// JavaScript / TypeScript - official apify-client. Reads input.json.
import { readFileSync } from 'node:fs';
import { ApifyClient } from 'apify-client';

const ACTOR = 'data_foundry/woocommerce-scraper';
const runInput = JSON.parse(readFileSync('input.json', 'utf8'));

const client = new ApifyClient({ token: process.env.APIFY_TOKEN });
const run = await client.actor(ACTOR).call(runInput);
console.log('dataset:', run.defaultDatasetId);

const { items } = await client.dataset(run.defaultDatasetId).listItems();
items.forEach((item) => console.log(item));
