# Architecture - slo-workbench

> Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.

**Pain:** Log costs grow until teams disable useful debug data.  
**Success metric:** Reduce duplicate observability alerts by 55% across shared services.

## Decision log

We chose K8s + go because the service needs to be deployed, observed, and rolled back under real production conditions. The architecture reflects the operational problem behind Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.: we need something that ships quickly, is observable, and can be recovered without hand-waving.

We also prioritized The observability stack itself is hardened so monitoring remains available during incidents. and Dashboards and alerts are built around signal quality, not noise, with clear paging criteria. so the architecture is built around real operational signal, not just feature scope.

See `docs/architecture/ADR-001-runtime-choice.md`

## Components

- go service with health, readiness, and domain-specific endpoints for Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.
- Prometheus metrics and structured logging for operators
- Docker image with non-root runtime and secure defaults
- K8s Deployment + HPA 2-10, 70% CPU
- GitOps (FluxCD) and rollback scripts
- The operational stack implied by the project: Prometheus, Grafana, Loki, Tempo, OpenTelemetry, Helm

## Trade-offs

| Choice | Why | Cost |
|---|---|---|
| K8s over serverless | team expertise, portability, predictable ops | more YAML |
| Self-hosted observability | lower cost than SaaS alternatives | we run it |
| RollingUpdate first | simpler for v0.1, canary later | faster shipping |

## Data flow

```
Client -> Ingress -> Service -> Pod
Pod -> /metrics -> Prometheus -> Alertmanager -> Slack
Domain actions -> API -> logs/metrics -> operators
```

## Scaling

HPA 2-10 pods. Stateless. 50m/64Mi request.

## Why not...

- **Serverless?** Cold start would hurt the main user flow and violate the reliability target.
- **Managed SaaS only?** The project must remain cost-effective and under operational control.
