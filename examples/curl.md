# cURL examples

Set your API key:

```bash
export MUNGfALI_API_KEY="mng_your_api_key_here"
```

## Basic search

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=electric car" \
  -H "Authorization: Bearer $MUNGfALI_API_KEY" \
  -H "Accept: application/json"
```

## Paginated search (page 1, 20 per page)

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=mountain landscape" \
  --data-urlencode "page=1" \
  --data-urlencode "per_page=20" \
  -H "Authorization: Bearer $MUNGfALI_API_KEY"
```

## SafeSearch strict + filter transparent PNGs

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=office workspace" \
  --data-urlencode "safeSearch=strict" \
  --data-urlencode "filterTransparent=true" \
  -H "Authorization: Bearer $MUNGfALI_API_KEY"
```

## Pretty-print with jq

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=luxury car" \
  --data-urlencode "per_page=3" \
  -H "Authorization: Bearer $MUNGfALI_API_KEY" | jq '.value[].imageUrl'
```

Expected shape (truncated):

```json
{
  "name": "luxury car",
  "value": [ { "imageUrl": "https://...", "hostUrl": "https://..." } ],
  "page": 1,
  "per_page": 3,
  "total": 150
}
```
