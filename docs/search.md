# Search endpoint

## `GET /v1/search`

Returns image results for a text query in a **single response** (up to 150 images). There is **no pagination**.

**Base URL:** `https://mungfali.net`

## Query parameters

| Parameter | Required | Description |
|-----------|----------|-------------|
| `q` | **Yes** | Search terms (URL-encoded). |
| `count` | No | Maximum images to return, 1–150. Default `150`. |
| `safeSearch` | No | `off`, `moderate`, or `strict`. Default `moderate`. |
| `filterTransparent` | No | When `true`, exclude transparent PNGs. Default `false`. |

## Example request

```http
GET /v1/search?q=electric+car&count=150 HTTP/1.1
Host: mungfali.net
Authorization: Bearer mng_your_api_key
Accept: application/json
```

## Example response

```json
{
  "name": "electric car",
  "value": [
    {
      "name": "Modern electric vehicle charging at station",
      "FamilyFriendly": true,
      "imageUrl": "https://images.pexels.com/photos/9799997/pexels-photo-9799997.jpeg?auto=compress&cs=tinysrgb&w=1200",
      "hostUrl": "https://www.pexels.com/search/electric%20car/",
      "contentSize": "892144 B",
      "width": 4000,
      "height": 2667,
      "isTransparent": false,
      "accentColor": "2D5A27"
    }
  ]
}
```

## Response fields

### Root

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Normalized query |
| `value` | array | Image result objects (up to 150) |

### Image object (`value[]`)

| Field | Type | Description |
|-------|------|-------------|
| `name` | string | Title or description |
| `FamilyFriendly` | boolean | SafeSearch classification |
| `imageUrl` | string | Direct image URL |
| `hostUrl` | string | Source page URL |
| `contentSize` | string | File size label |
| `width` | integer | Width in pixels |
| `height` | integer | Height in pixels |
| `isTransparent` | boolean | Transparent background flag |
| `accentColor` | string | Dominant color (hex, no `#`) |

## OpenAPI

Full schema: [openapi.yaml](../openapi.yaml)

## Related

- [Authentication](./authentication.md)
- [Error codes](./error-codes.md)
