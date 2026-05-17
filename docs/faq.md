# FAQ

## General

### What is the Mungfali Image Search API?

A REST API that returns up to **150** image results per query as structured JSON — built for developers who need affordable image search without complex search-engine setup.

### Is this a Google Images API?

Google does not offer a public general-purpose Google Images API. Mungfali is a **Google Images API alternative** aimed at teams that would otherwise use Custom Search, Bing, or SerpAPI.

### How do I get an API key?

[Sign up free](https://mungfali.net/signup) — 250 searches/month, no credit card. Your key appears in the [dashboard](https://mungfali.net/dashboard).

## Authentication & requests

### Bearer token or X-API-Key?

Both work. Documentation standardizes on `Authorization: Bearer mng_...`. See [Authentication](./authentication.md).

### What is the search endpoint?

`GET https://mungfali.net/v1/search` with required query parameter `q`.

### Do I need pagination?

No. You can retrieve up to 150 images in one call. Use `page` and `per_page` when you want smaller pages. See [Pagination](./pagination.md).

## Billing & limits

### Is there a free tier?

Yes — **250 requests/month** on the Free plan.

### Do unused searches roll over?

No. Quota resets each billing cycle.

### How much does paid usage cost?

Paid plans start at **$19/month** for 10,000 requests. Effective cost is about **$1.90 per 1,000** requests on Growth-tier economics vs ~$15 for Bing Image Search. See [pricing](https://mungfali.net/pricing).

## SEO & product questions

### What keywords describe this API?

image search API, REST image search, JSON image API, Bing alternative, SerpAPI alternative, affordable image search, developer image API.

### What fields are in each image result?

`imageUrl`, `hostUrl`, `width`, `height`, `FamilyFriendly`, `isTransparent`, `accentColor`, and more — see [Search](./search.md).

### How fast are responses?

Typical responses complete in under **300 ms**.

## Support

- **Documentation:** [mungfali.net/dashboard/docs](https://mungfali.net/dashboard/docs)
- **Email:** [hello@mungfali.com](mailto:hello@mungfali.com)
- **Security:** [SECURITY.md](../SECURITY.md)
