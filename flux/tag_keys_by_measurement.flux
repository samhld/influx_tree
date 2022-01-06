import "influxdata/influxdb/schema"
schema.tagKeys(bucket: "%s", predicate: (r) => r._measurement == "%s")
|> filter(fn: (r) => r._value != "_start" and r._value != "_stop")