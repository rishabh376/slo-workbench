# Setup - slo-workbench

Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.

1. **Infra first**: `cd terraform && terraform init && terraform apply`
2. `kubectl apply -k k8s/ -n slo-workbench`
3. `kubectl port-forward svc/slo-workbench 8080:80 -n slo-workbench`
4. `curl localhost:8080/healthz`

Env vars: `PORT=8080`, `LOG_LEVEL=info`

Dashboards and alerts are built around signal quality, not noise, with clear paging criteria.
