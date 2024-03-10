# Observability

## Components
- Logging: Gcloud logging
- Metrics (Visual graph + Alert): Prometheus + Grafana
- Tracing: Jaeger

## Alerts patterns

## Alert rules
- increase() vs rate()
  - increase(): total count of occurrences over a time window: 
    - `increase(http_requests_total{status=~"5.."}[5m]) > 100`
  - rate(): # of occurrences / second (over a time window, per 1 server)
    - `rate(http_requests_total[5m]) > 0.5`
  - sum(rate()): # of occurrences / second (over a time window, for all servers)
    - `sum(rate(http_requests_total[5m])) > 0.5`