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

## Search with count (up to 150)

```bash
curl -s -G "https://mungfali.net/v1/search" \
  --data-urlencode "q=mountain landscape" \
  --data-urlencode "count=150" \
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
  -H "Authorization: Bearer $MUNGfALI_API_KEY" | jq '.value[].imageUrl'
```

Expected shape (truncated):

```json
{
  "name": "luxury car",
  "value": [ { "imageUrl": "https://...", "hostUrl": "https://..." } ]
}
```
