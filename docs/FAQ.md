# FAQ - slo-workbench

**Q: What problem does this actually solve?**
A: Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.

Pain point: Log costs grow until teams disable useful debug data.

**Q: Free to run?**
A: Yep. GitHub Free - $0/mo. See `terraform/INFRA_REQUIREMENTS.md`

**Q: Production ready?**
A: Yes - SLOs, security hardening, DR runbook, Helm, Terraform, monitoring.

**Q: Stack?**
A: Prometheus, Grafana, Loki, Tempo, OpenTelemetry, Helm

**Q: What infra do I need?**
A: `terraform/INFRA_REQUIREMENTS.md` has the full breakdown - CPU/RAM/storage/network/cost.
