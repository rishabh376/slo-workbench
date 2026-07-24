# Operations Runbook - slo-workbench

Observability Teams coordinate dedupe dashboard alerts for OpenTelemetry spans, so log spend is controlled without losing incident context.

**On-call? Read this first.**

## Deploy
```
kubectl apply -k k8s/ -n slo-workbench
kubectl rollout status deploy/slo-workbench -n slo-workbench
```

## Rollback (when things go red)
```
./scripts/rollback.sh
# usually < 30sec
```

## Scale up NOW
```
kubectl scale deploy/slo-workbench --replicas=10 -n slo-workbench
```

## Logs
```
kubectl logs -l app=slo-workbench -n slo-workbench -f --tail=200
```

## Crashloop?
```
kubectl describe pod -l app=slo-workbench -n slo-workbench
kubectl logs <pod> --previous -n slo-workbench
```
90% of the time: missing env var in ConfigMap.

## Metrics
`curl localhost:8080/metrics` or Grafana dashboard in `monitoring/`

## DR / Failover
See `docs/DISASTER_RECOVERY.md` - RTO 15m

We actually tested this. 11 min.
