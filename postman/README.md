# Postman collection

## Import

1. Open [Postman](https://www.postman.com/downloads/).
2. Click **Import** → **Upload Files**.
3. Select `Mungfali-Image-Search-API.postman_collection.json`.

## Set your API key

1. In the sidebar, select the **Mungfali Image Search API** collection.
2. Open the **Variables** tab.
3. Set **Current value** for `api_key` to your key from [mungfali.net/dashboard](https://mungfali.net/dashboard) (starts with `mng_`).
4. Click **Save**.

Collection-level Bearer auth uses `{{api_key}}` for all requests under **Search**. Error examples override auth only where noted.

## Requests included

- **Search images (basic)** — `q` only
- **Search images (count=150)** — optional `count` parameter
- **Search images (SafeSearch strict)** — `safeSearch` + `filterTransparent`
- **Missing query (400)** / **Invalid API key (401)** — error samples
