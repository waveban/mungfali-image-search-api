# PHP example

Requires PHP 8.0+ and the cURL extension (or use the included SDK).

## Using the SDK

```php
<?php
require __DIR__ . '/../sdks/php/MungfaliClient.php';

$apiKey = getenv('MUNGfALI_API_KEY') ?: 'mng_your_api_key_here';
$client = new MungfaliClient($apiKey);

$result = $client->search('electric car', [
    'count' => 150,
    'safeSearch' => 'moderate',
]);

foreach ($result['value'] as $image) {
    echo $image['imageUrl'] . PHP_EOL;
}
```

## Raw cURL (no SDK)

```php
<?php
$query = http_build_query([
    'q' => 'electric car',
    'count' => 150,
]);

$ch = curl_init('https://mungfali.net/v1/search?' . $query);
curl_setopt_array($ch, [
    CURLOPT_RETURNTRANSFER => true,
    CURLOPT_HTTPHEADER => [
        'Authorization: Bearer ' . getenv('MUNGfALI_API_KEY'),
        'Accept: application/json',
    ],
]);

$response = curl_exec($ch);
$status = curl_getinfo($ch, CURLINFO_HTTP_CODE);
curl_close($ch);

if ($status !== 200) {
    throw new RuntimeException('API error: ' . $response);
}

$data = json_decode($response, true, 512, JSON_THROW_ON_ERROR);
print_r($data['value'][0]['imageUrl'] ?? 'No results');
```
