# Rate limits

Limits depend on your **subscription plan**. Each successful `GET /v1/search` counts as **one request** toward your monthly quota.

## Monthly quotas

| Plan | Requests / month |
|------|------------------|
| Free | 250 |
| Starter | 10,000 |
| Growth | 50,000 |
| Pro | 150,000 |

See current plans on [mungfali.net/pricing](https://mungfali.net/pricing).

## Burst behavior

Short bursts are allowed for normal application traffic. Sustained traffic above your plan may receive **429** responses with `rate_limit_exceeded`.

## Quota exhaustion

When your monthly quota is used:

- HTTP **429** is returned.
- Response includes guidance to upgrade.
- Unused searches **do not roll over** to the next period.

## Best practices

1. **Cache** popular queries server-side with TTL aligned to your product needs.
2. **Use `count`** when you need fewer than 150 images in a single response.
3. **Monitor** usage in the [dashboard](https://mungfali.net/dashboard).
4. **Upgrade** before launch traffic spikes.

## Related

- [Error codes](./error-codes.md)
- [FAQ](./faq.md)
- [Pricing](https://mungfali.net/pricing)
