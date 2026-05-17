"""Lightweight Mungfali Image Search API client."""

from __future__ import annotations

import json
import urllib.error
import urllib.parse
import urllib.request
from typing import Any


class MungfaliClient:
    BASE_URL = "https://mungfali.net/v1/search"

    def __init__(self, api_key: str, timeout: int = 30) -> None:
        self.api_key = api_key
        self.timeout = timeout

    def search(
        self,
        query: str,
        *,
        page: int | None = None,
        per_page: int | None = None,
        safe_search: str | None = None,
        filter_transparent: bool | None = None,
    ) -> dict[str, Any]:
        params: dict[str, str] = {"q": query}
        if page is not None:
            params["page"] = str(page)
        if per_page is not None:
            params["per_page"] = str(per_page)
        if safe_search is not None:
            params["safeSearch"] = safe_search
        if filter_transparent is not None:
            params["filterTransparent"] = "true" if filter_transparent else "false"

        url = f"{self.BASE_URL}?{urllib.parse.urlencode(params)}"
        request = urllib.request.Request(
            url,
            headers={
                "Authorization": f"Bearer {self.api_key}",
                "Accept": "application/json",
            },
            method="GET",
        )

        try:
            with urllib.request.urlopen(request, timeout=self.timeout) as response:
                body = response.read().decode("utf-8")
        except urllib.error.HTTPError as exc:
            detail = exc.read().decode("utf-8", errors="replace")
            raise RuntimeError(f"API error {exc.code}: {detail}") from exc

        return json.loads(body)
