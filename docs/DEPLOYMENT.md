# Deployment - slo-workbench

**Use case:** Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.

## Infra requirements
→ `terraform/INFRA_REQUIREMENTS.md` - CPU/RAM/storage/network/cost breakdown

## GitOps
Flux watches `main` and reconciles changes automatically.

## Rollback
```
./scripts/rollback.sh
kubectl rollout undo deploy/slo-workbench -n slo-workbench
```
Usually < 30 sec.

## Cost
GitHub Free - $0/mo

The observability stack itself is hardened so monitoring remains available during incidents.
