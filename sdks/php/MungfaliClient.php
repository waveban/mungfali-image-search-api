<?php

declare(strict_types=1);

/**
 * Lightweight Mungfali Image Search API client.
 */
final class MungfaliClient
{
    private const BASE_URL = 'https://mungfali.net/v1/search';

    public function __construct(
        private readonly string $apiKey,
        private readonly int $timeoutSeconds = 30,
    ) {
    }

    /**
     * @param array{page?: int, per_page?: int, safeSearch?: string, filterTransparent?: bool} $options
     * @return array<string, mixed>
     */
    public function search(string $query, array $options = []): array
    {
        $params = ['q' => $query];

        if (isset($options['page'])) {
            $params['page'] = (int) $options['page'];
        }
        if (isset($options['per_page'])) {
            $params['per_page'] = (int) $options['per_page'];
        }
        if (isset($options['safeSearch'])) {
            $params['safeSearch'] = (string) $options['safeSearch'];
        }
        if (isset($options['filterTransparent'])) {
            $params['filterTransparent'] = $options['filterTransparent'] ? 'true' : 'false';
        }

        $url = self::BASE_URL . '?' . http_build_query($params);

        $ch = curl_init($url);
        curl_setopt_array($ch, [
            CURLOPT_RETURNTRANSFER => true,
            CURLOPT_TIMEOUT => $this->timeoutSeconds,
            CURLOPT_HTTPHEADER => [
                'Authorization: Bearer ' . $this->apiKey,
                'Accept: application/json',
            ],
        ]);

        $body = curl_exec($ch);
        $status = (int) curl_getinfo($ch, CURLINFO_HTTP_CODE);
        $error = curl_error($ch);
        curl_close($ch);

        if ($body === false) {
            throw new RuntimeException('Request failed: ' . $error);
        }

        if ($status < 200 || $status >= 300) {
            throw new RuntimeException(sprintf('API error %d: %s', $status, $body));
        }

        return json_decode($body, true, 512, JSON_THROW_ON_ERROR);
    }
}
