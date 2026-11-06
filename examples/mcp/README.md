# MCP

Connect this Actor to an MCP client (Claude Desktop, Cursor, LibreChat, and others):

```json
{
  "mcpServers": {
    "apify": {
      "type": "http",
      "url": "https://mcp.apify.com/?tools=fetch-actor-details,data_foundry/woocommerce-scraper"
    }
  }
}
```

The hosted server signs you in with OAuth on first connect, so no token belongs in this config. Clients without OAuth can send `Authorization: Bearer <APIFY_TOKEN>`.

Learn more: https://docs.apify.com/agents
