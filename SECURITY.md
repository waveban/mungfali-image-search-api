# Security Policy

## Supported versions

| Version | Supported          |
| ------- | -------------------- |
| 1.0.x   | :white_check_mark: |

This repository documents the Mungfali Image Search API. Security fixes for the **hosted API** are applied on the server side; this repo receives documentation and client updates as needed.

## Reporting a vulnerability

**Please do not report security vulnerabilities through public GitHub issues.**

Instead, report them privately:

1. Email **hello@mungfali.com** with subject line `Security: Mungfali Image Search API`.
2. Include a description of the issue, steps to reproduce, and impact assessment.
3. If applicable, include proof-of-concept code or request/response samples **without** embedding live API keys.

We will acknowledge receipt within **3 business days** and aim to provide an initial assessment within **7 business days**.

## What to include

- Affected endpoint(s) and parameters
- Whether authentication is required
- Your contact information for follow-up

## Safe disclosure

We ask that you allow reasonable time for remediation before public disclosure. We will credit reporters who wish to be acknowledged, unless you prefer to remain anonymous.

## API key safety

- Never commit API keys to this repository.
- Rotate keys immediately if exposed in logs, screenshots, or version control.
- Use environment variables or secret managers in production.

Thank you for helping keep Mungfali and our users secure.
