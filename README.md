# slo-workbench

> Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.

**Domain:** monitoring | **Complexity:** intermediate | **Stack:** Prometheus, Grafana, Loki, Tempo, OpenTelemetry, Helm

| | |
|---|---|
| Use case | Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context. |
| Pain point | Log costs grow until teams disable useful debug data. |
| Success metric | Reduce duplicate observability alerts by 55% across shared services. |
| Cost | $0/mo on GitHub Free |

## Why this exists

Log costs grow until teams disable useful debug data.. Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.

> This stack makes logs, metrics, and traces useful together so teams can spot issues before customers do. For this implementation, the service is designed to remove the operational friction behind Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context..

## What this product does

This project ships a production-oriented service that exposes health, readiness, observability, and domain-specific endpoints for Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.. The implementation is intended to be deployed behind Kubernetes with secure defaults, CI validation, and rollback-ready manifests.

## Quick start

```bash
docker build -t slo-workbench:local .
docker run -p 8080:8080 slo-workbench:local
curl http://localhost:8080/healthz
```

## Architecture

```
GitHub CI -> Container Registry -> Kubernetes -> Prometheus -> Operators / SRE
```

See `docs/ARCHITECTURE.md`

## Infrastructure

Full spec: `terraform/INFRA_REQUIREMENTS.md`

TL;DR: K8s cluster, 2-10 pods HPA, $0 on GitHub Free

## Docs

- `docs/ARCHITECTURE.md` - components, tradeoffs
- `docs/DESIGN_DISCUSSION.md` - how we decided to build this
- `docs/DEPLOYMENT.md` - GitOps, rollback
- `docs/SECURITY.md` - threat model
- `docs/TROUBLESHOOTING.md` - 3am runbook
- `terraform/INFRA_REQUIREMENTS.md` - **infra: CPU/RAM/storage/cost**

Interview prep: `interview/`
