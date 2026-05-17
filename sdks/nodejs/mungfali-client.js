/**
 * Lightweight Mungfali Image Search API client.
 */
class MungfaliClient {
  static BASE_URL = 'https://mungfali.net/v1/search';

  /**
   * @param {string} apiKey
   * @param {{ timeoutMs?: number }} [options]
   */
  constructor(apiKey, options = {}) {
    this.apiKey = apiKey;
    this.timeoutMs = options.timeoutMs ?? 30_000;
  }

  /**
   * @param {string} query
   * @param {{ page?: number, per_page?: number, safeSearch?: string, filterTransparent?: boolean }} [opts]
   * @returns {Promise<Record<string, unknown>>}
   */
  async search(query, opts = {}) {
    const params = new URLSearchParams({ q: query });
    if (opts.page != null) params.set('page', String(opts.page));
    if (opts.per_page != null) params.set('per_page', String(opts.per_page));
    if (opts.safeSearch) params.set('safeSearch', opts.safeSearch);
    if (opts.filterTransparent != null) {
      params.set('filterTransparent', opts.filterTransparent ? 'true' : 'false');
    }

    const controller = new AbortController();
    const timer = setTimeout(() => controller.abort(), this.timeoutMs);

    try {
      const response = await fetch(`${MungfaliClient.BASE_URL}?${params}`, {
        headers: {
          Authorization: `Bearer ${this.apiKey}`,
          Accept: 'application/json',
        },
        signal: controller.signal,
      });

      const body = await response.text();
      if (!response.ok) {
        throw new Error(`API error ${response.status}: ${body}`);
      }

      return JSON.parse(body);
    } finally {
      clearTimeout(timer);
    }
  }
}

module.exports = { MungfaliClient };
