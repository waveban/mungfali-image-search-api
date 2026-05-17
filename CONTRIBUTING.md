# Contributing to Mungfali Image Search API

Thank you for helping improve this repository. This project contains **documentation, SDKs, and examples** for the [Mungfali Image Search API](https://mungfali.net). It does not host the production API server.

## How to contribute

1. **Fork** the repository and create a branch from `main`.
2. **Make focused changes** — one topic per pull request (docs fix, new example, SDK improvement).
3. **Match existing style** — clear headings, working code samples, consistent base URL `https://mungfali.net`.
4. **Test examples** where possible against a valid API key (do not commit keys).
5. **Open a pull request** with a clear description and link any related issue.

## Reporting issues

- Use the [bug report template](.github/ISSUE_TEMPLATE/bug_report.yml) for incorrect docs, broken examples, or SDK bugs.
- Use the [feature request template](.github/ISSUE_TEMPLATE/feature_request.yml) for new languages, endpoints, or documentation sections.
- For **API outages or billing**, contact [hello@mungfali.com](mailto:hello@mungfali.com) or use your [dashboard](https://mungfali.net/dashboard).

## Security issues

Do **not** open public issues for vulnerabilities. See [SECURITY.md](./SECURITY.md).

## Code of conduct

By participating, you agree to abide by our [Code of Conduct](./CODE_OF_CONDUCT.md).

## Development tips

- Run markdown lint locally: `npx markdownlint-cli2 "**/*.md"`
- Keep OpenAPI, README, and `docs/` in sync when changing parameters or response shapes.
- Use realistic JSON in examples; avoid placeholder domains that do not resolve.

We appreciate your contributions.
