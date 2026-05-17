# Authentication

The Mungfali Image Search API uses **API keys** issued from your [dashboard](https://mungfali.net/dashboard) after [signup](https://mungfali.net/signup).

## Bearer token (recommended)

Send your key in the `Authorization` header:

```http
Authorization: Bearer mng_your_api_key_here
```

### cURL example

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=sunset" \
  -H "Authorization: Bearer $MUNGfALI_API_KEY"
```

## Alternative: X-API-Key header

The production API also accepts:

```http
X-API-Key: mng_your_api_key_here
```

Use whichever fits your HTTP client; both authenticate the same key.

## Key format

- Keys are prefixed with `mng_`.
- Store keys in environment variables or a secrets manager — never commit them to git.

## Unauthorized responses

Missing or invalid keys return **401** with:

```json
{
  "error": {
    "code": "unauthorized",
    "message": "Invalid or missing Bearer token."
  }
}
```

See [Error codes](./error-codes.md).

## Related

- [Search endpoint](./search.md)
- [Rate limits](./rate-limits.md)
- [FAQ](./faq.md)
