# Python example

Requires Python 3.9+.

## Using the SDK

```python
import os
import sys

sys.path.insert(0, os.path.join(os.path.dirname(__file__), "..", "sdks", "python"))
from mungfali_client import MungfaliClient

client = MungfaliClient(api_key=os.environ["MUNGfALI_API_KEY"])

data = client.search(
    "electric car",
    page=1,
    per_page=10,
    safe_search="moderate",
)

for image in data["value"]:
    print(image["imageUrl"])
```

## Using requests (no SDK)

```python
import os
import requests

API_KEY = os.environ["MUNGfALI_API_KEY"]

response = requests.get(
    "https://mungfali.net/v1/search",
    params={"q": "electric car", "page": 1, "per_page": 10},
    headers={"Authorization": f"Bearer {API_KEY}", "Accept": "application/json"},
    timeout=30,
)
response.raise_for_status()

data = response.json()
print(data["name"], len(data["value"]))
for image in data["value"]:
    print(image["imageUrl"])
```
