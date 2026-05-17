# Error codes

Errors are returned as JSON with an `error` object containing `code` and `message`.

## HTTP status summary

| Status | Code | Meaning |
|--------|------|---------|
| 400 | `invalid_request` | Missing `q`, invalid parameter type, or out-of-range value |
| 401 | `unauthorized` | Missing or invalid API key |
| 403 | `forbidden` | Key valid but not permitted for this action |
| 429 | `rate_limit_exceeded` | Too many requests or monthly quota exhausted |
| 500 | `internal_error` | Server error — retry with backoff |
| 503 | `service_unavailable` | Temporary outage — retry with backoff |

## Examples

### 400 — missing query

```json
{
  "error": {
    "code": "invalid_request",
    "message": "Query parameter \"q\" is required."
  }
}
```

### 401 — invalid key

```json
{
  "error": {
    "code": "unauthorized",
    "message": "Invalid or missing Bearer token."
  }
}
```

### 429 — quota exceeded

```json
{
  "error": {
    "code": "rate_limit_exceeded",
    "message": "Monthly request quota exceeded. Upgrade at https://mungfali.net/pricing"
  }
}
```

## Client handling tips

- **401:** Verify the key in the dashboard; check for stray whitespace in headers.
- **429:** Back off exponentially; upgrade plan or wait for quota reset.
- **5xx:** Retry up to 3 times with jitter; log `error.code` for support.

## Related

- [Rate limits](./rate-limits.md)
- [Authentication](./authentication.md)
