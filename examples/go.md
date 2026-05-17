# Go example

Requires Go 1.21+.

## Using the SDK

```go
package main

import (
	"context"
	"fmt"
	"os"

	"github.com/mungfali/mungfali-image-search-api/sdks/go/mungfali"
)

func main() {
	client := mungfali.NewClient(os.Getenv("MUNGfALI_API_KEY"))

	data, err := client.Search(context.Background(), "electric car", mungfali.SearchOptions{
		Page:    1,
		PerPage: 10,
	})
	if err != nil {
		panic(err)
	}

	for _, img := range data.Value {
		fmt.Println(img.ImageURL)
	}
}
```

## Standard library only

```go
package main

import (
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
)

func main() {
	apiKey := os.Getenv("MUNGfALI_API_KEY")
	u, _ := url.Parse("https://mungfali.net/v1/search")
	q := u.Query()
	q.Set("q", "electric car")
	q.Set("page", "1")
	q.Set("per_page", "10")
	u.RawQuery = q.Encode()

	req, _ := http.NewRequest(http.MethodGet, u.String(), nil)
	req.Header.Set("Authorization", "Bearer "+apiKey)
	req.Header.Set("Accept", "application/json")

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, _ := io.ReadAll(res.Body)
	if res.StatusCode != http.StatusOK {
		panic(fmt.Sprintf("status %d: %s", res.StatusCode, body))
	}

	var data struct {
		Name  string `json:"name"`
		Value []struct {
			ImageURL string `json:"imageUrl"`
		} `json:"value"`
	}
	_ = json.Unmarshal(body, &data)
	for _, img := range data.Value {
		fmt.Println(img.ImageURL)
	}
}
```
