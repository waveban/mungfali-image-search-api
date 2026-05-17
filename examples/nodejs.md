# Node.js example

Requires Node.js 18+ (native `fetch`).

## Using the SDK

```javascript
const path = require('path');
const { MungfaliClient } = require(path.join(__dirname, '../sdks/nodejs/mungfali-client'));

const client = new MungfaliClient(process.env.MUNGfALI_API_KEY);

async function main() {
  const data = await client.search('electric car', {
    page: 1,
    per_page: 10,
    safeSearch: 'moderate',
  });

  for (const image of data.value) {
    console.log(image.imageUrl);
  }
}

main().catch(console.error);
```

## Native fetch (no SDK)

```javascript
const apiKey = process.env.MUNGfALI_API_KEY;
const params = new URLSearchParams({
  q: 'electric car',
  page: '1',
  per_page: '10',
});

const response = await fetch(`https://mungfali.net/v1/search?${params}`, {
  headers: {
    Authorization: `Bearer ${apiKey}`,
    Accept: 'application/json',
  },
});

if (!response.ok) {
  throw new Error(`API error ${response.status}: ${await response.text()}`);
}

const data = await response.json();
console.log(data.name, data.value.length);
data.value.forEach((img) => console.log(img.imageUrl));
```
