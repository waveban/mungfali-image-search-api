# Pagination

Mungfali can return **up to 150 images** for a single query. You may either:

1. **Single request** — omit `page` and `per_page` (or set `per_page=150`) to receive the full result set in one response.
2. **Paginated requests** — use `page` and `per_page` to fetch smaller chunks for UI carousels or memory-constrained clients.

## Parameters

| Parameter | Default | Range | Description |
|-----------|---------|-------|-------------|
| `page` | `1` | ≥ 1 | Current page (1-based). |
| `per_page` | `150` | 1–150 | Number of images per page. |

## Examples

### Page 1 (20 results)

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=mountain" \
  --data-urlencode "page=1" \
  --data-urlencode "per_page=20" \
  -H "Authorization: Bearer $MUNGfALI_API_KEY"
```

### Page 2

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=mountain" \
  --data-urlencode "page=2" \
  --data-urlencode "per_page=20" \
  -H "Authorization: Bearer $MUNGfALI_API_KEY"
```

## Response metadata

Paginated responses include:

```json
{
  "page": 2,
  "per_page": 20,
  "total": 150
}
```

Use `total` and `per_page` to compute whether another page exists:

```
has_more = (page * per_page) < total
```

## When to skip pagination

For batch jobs, ML datasets, or backends that can handle a larger payload, a **single call** with `per_page=150` minimizes latency and quota usage versus multiple paginated calls for the same query.

## Related

- [Search endpoint](./search.md)
- [Rate limits](./rate-limits.md)
