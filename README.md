# Mungfali Image Search API

[![License: MIT](https://img.shields.io/badge/License-MIT-blue.svg)](LICENSE)
[![API Version](https://img.shields.io/badge/API-v1.0.0-green.svg)](CHANGELOG.md)
[![Docs](https://img.shields.io/badge/docs-mungfali.net-0ea5e9.svg)](https://mungfali.net/dashboard/docs)
[![Pricing](https://img.shields.io/badge/pricing-from%20%240-8b5cf6.svg)](https://mungfali.net/pricing)
[![Status](https://img.shields.io/badge/status-production-success.svg)](https://mungfali.net)

**Affordable, developer-first image search** — up to **150 results per request**, SafeSearch, transparent-PNG filtering, and fast JSON over a single REST endpoint.

> **Keywords:** image search API, Google Images API alternative, Bing Image Search alternative, SerpAPI alternative, REST image search, JSON image API, affordable image search for developers.

[Get started free](https://mungfali.net/signup) · [Documentation](https://mungfali.net/dashboard/docs) · [Pricing](https://mungfali.net/pricing) · [OpenAPI spec](./openapi.yaml) · [Postman collection](./postman/Mungfali-Image-Search-API.postman_collection.json)

---

## Product overview

**Mungfali Image Search API** helps you add high-quality image search to apps, dashboards, and automation pipelines without expensive legacy providers. One `GET /v1/search` call returns up to **150** structured image objects — **no pagination required**.

Trusted by **5,000+ developers**. Average response time **~243 ms**. **250 free searches/month** — no credit card required.

---

## Why Developers Choose Mungfali

- **Up to 150 images per request** — fewer round trips than Bing (10) or many Google CSE setups.
- **Simple REST + JSON** — one endpoint, predictable schema (`name`, `value[]`, rich metadata per image).
- **Transparent pricing** — from **$0** to **$99/mo**; ~**$1.90 per 1,000** requests on paid tiers vs ~$15 for Bing.
- **Built-in SafeSearch** — family-friendly results via `safeSearch`.
- **Transparent PNG filtering** — optional removal of images with transparent backgrounds.
- **Domain include/exclude** — curate sources for brand-safe or licensed workflows.
- **Global edge delivery** — low latency worldwide.
- **Official SDKs** — PHP, Python, Node.js, and Go clients with a single `search()` method.

---

## Google Images API Alternative

Google does not offer a public “Google Images API” for general image search. Teams usually wire up **Google Custom Search JSON API** (limited image results, CSE setup, separate billing) or scrape via aggregators like **SerpAPI**.

**Mungfali** is built specifically for **image search**:

| Capability | Mungfali | Google Custom Search API |
|------------|----------|---------------------------|
| Purpose-built image search | Yes | Partial (CSE configuration) |
| Results per request | Up to 150 | Typically 10 per call |
| Setup complexity | API key only | Search engine ID + API key |
| Starting price | $0 (250 req/mo) | Pay-as-you-go + CSE limits |
| Pagination required | **No** | Yes |

If you need a **Google Images API alternative** that is fast to integrate and priced for startups, Mungfali is designed for that workflow.

---

## Features

- Up to **150** image results per search in a **single call** (no pagination)
- **SafeSearch** (`off`, `moderate`, `strict`)
- **Transparent background filtering**
- **Domain allow/block lists**
- **Structured JSON**: `imageUrl`, `hostUrl`, dimensions, `FamilyFriendly`, `isTransparent`, `accentColor`
- **Bearer token** authentication
- **Rate limits** by plan with clear error codes
- **Postman collection** and **OpenAPI 3** spec included in this repo

---

## Quick start

1. **Sign up** at [mungfali.net/signup](https://mungfali.net/signup) (250 free searches/month).
2. Copy your API key from the [dashboard](https://mungfali.net/dashboard).
3. Call the search endpoint (example below), then parse the JSON `value` array and render `imageUrl` in your UI.

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=electric car" \
  --data-urlencode "count=150" \
  -H "Authorization: Bearer YOUR_API_KEY"
```

---

## Example request and response

### Request

```http
GET /v1/search?q=luxury+car&count=150 HTTP/1.1
Host: mungfali.net
Authorization: Bearer mng_your_api_key_here
Accept: application/json
```

### Response `200 OK`

```json
{
  "name": "luxury car",
  "value": [
    {
      "name": "Luxury Car Photos, Download The BEST Free Luxury Car Stock Photos & HD ...",
      "FamilyFriendly": true,
      "imageUrl": "https://images.pexels.com/photos/3729464/pexels-photo-3729464.jpeg?cs=srgb&dl=pexels-mikebirdy-3729464.jpg&fm=jpg",
      "hostUrl": "https://www.pexels.com/search/luxury%20car/",
      "contentSize": "1781235 B",
      "width": 5456,
      "height": 3632,
      "isTransparent": false,
      "accentColor": "C21C09"
    },
    {
      "name": "Black luxury sedan on coastal highway at sunset",
      "FamilyFriendly": true,
      "imageUrl": "https://images.unsplash.com/photo-1503376780353-7e6692767b70?auto=format&fit=crop&w=1200&q=80",
      "hostUrl": "https://unsplash.com/s/photos/luxury-car",
      "contentSize": "412800 B",
      "width": 1200,
      "height": 800,
      "isTransparent": false,
      "accentColor": "1A1A2E"
    }
  ]
}
```

---

## Pricing and free tier

| Plan | Monthly price | Requests / month | Best for |
|------|---------------|------------------|----------|
| **Free** | $0 | 250 | Prototypes, evaluation |
| **Starter** | $19 | 10,000 | Early products |
| **Growth** | $49 | 50,000 | Growing apps (most popular) |
| **Pro** | $99 | 150,000 | Production workloads |

- **Annual billing:** save **20%** at checkout.
- **Overages:** billed per plan terms on the [pricing page](https://mungfali.net/pricing).
- **Unused searches** do not roll over.

[View full pricing →](https://mungfali.net/pricing) · [Create free account →](https://mungfali.net/signup)

---

## Comparison with other image search APIs

| Feature | **Mungfali** | Google Custom Search API | Bing Image Search API | SerpAPI |
|---------|--------------|--------------------------|----------------------|---------|
| Primary use | Image search | Web + image via CSE | Image search | Multi-engine SERP |
| Typical results / call | **150** | ~10 | 10–50 (tiered) | Varies by engine |
| Pagination required | **No** | Yes | Often | Yes |
| SafeSearch | Built-in | Limited | Yes | Varies |
| Transparent PNG filter | Yes | No | No | No |
| Starting price | **$0** (250/mo) | Pay-as-you-go + CSE | Paid tiers | Paid tiers |
| Cost per 1,000 requests (indicative) | **~$1.90** | Varies | **~$15** | **~$10+** |
| Setup | API key | API key + Search Engine ID | Azure resource | API key |

---

## SDK usage

Lightweight clients live under [`sdks/`](./sdks/). Each exposes `search(query, options?)`.

### PHP

```php
<?php
require 'sdks/php/MungfaliClient.php';

$client = new MungfaliClient(getenv('MUNGfALI_API_KEY'));
$result = $client->search('mountain landscape', ['count' => 150]);
foreach ($result['value'] as $image) {
    echo $image['imageUrl'] . PHP_EOL;
}
```

### Python

```python
from mungfali_client import MungfaliClient

client = MungfaliClient(api_key="mng_your_api_key")
data = client.search("mountain landscape", count=150)
for image in data["value"]:
    print(image["imageUrl"])
```

### Node.js

```javascript
const { MungfaliClient } = require('./sdks/nodejs/mungfali-client');

const client = new MungfaliClient(process.env.MUNGfALI_API_KEY);
const data = await client.search('mountain landscape', { count: 150 });
data.value.forEach((img) => console.log(img.imageUrl));
```

### Go

```go
client := mungfali.NewClient(os.Getenv("MUNGfALI_API_KEY"))
data, err := client.Search(ctx, "mountain landscape", mungfali.SearchOptions{Count: 150})
```

See [examples/](./examples/) for full cURL, PHP, Python, Node.js, and Go samples.

---

## Common Use Cases

- **Content and media apps** — hero images, galleries, and editorial tooling
- **E-commerce** — product inspiration, mood boards, and merchandising aids
- **Marketing automation** — campaign asset discovery with domain filters
- **Internal dashboards** — research panels for design and SEO teams
- **AI / ML pipelines** — collect labeled image URLs and metadata at scale
- **Education platforms** — SafeSearch-enabled image pickers for students

---

## Documentation in this repository

| Topic | Link |
|-------|------|
| Authentication | [docs/authentication.md](./docs/authentication.md) |
| Search endpoint | [docs/search.md](./docs/search.md) |
| Error codes | [docs/error-codes.md](./docs/error-codes.md) |
| Rate limits | [docs/rate-limits.md](./docs/rate-limits.md) |
| FAQ | [docs/faq.md](./docs/faq.md) |
| OpenAPI 3 | [openapi.yaml](./openapi.yaml) |
| Examples | [examples/](./examples/) |

Live docs: [mungfali.net/dashboard/docs](https://mungfali.net/dashboard/docs)

---

## FAQ

### Is there a free tier?

Yes — **250 searches per month** on the Free plan. [Sign up](https://mungfali.net/signup) without a credit card.

### How do I authenticate?

Send `Authorization: Bearer YOUR_API_KEY`. See [Authentication](./docs/authentication.md). The live API also accepts `X-API-Key` for the same key.

### What is the best Google Images API alternative for startups?

Mungfali offers a single image-focused endpoint, up to 150 results per call, and transparent pricing — without configuring a Custom Search Engine.

### Does unused quota roll over?

No. Quota resets each billing period. You can upgrade mid-cycle for more capacity.

### What format are responses?

JSON with a query `name` and a `value` array of image objects. See [Search](./docs/search.md).

### How fast is the API?

Average response time is under **300 ms** for typical queries.

---

## Contributing

We welcome issues and pull requests. See [CONTRIBUTING.md](./CONTRIBUTING.md) and [CODE_OF_CONDUCT.md](./CODE_OF_CONDUCT.md).

## Security

Report vulnerabilities privately — see [SECURITY.md](./SECURITY.md).

## License

[MIT License](./LICENSE) © Mungfali

---

**Links:** [Website](https://mungfali.net) · [Sign up](https://mungfali.net/signup) · [Pricing](https://mungfali.net/pricing) · [Docs](https://mungfali.net/dashboard/docs) · [Contact](mailto:hello@mungfali.com)
