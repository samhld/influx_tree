import "influxdata/influxdb/schema"
schema.tagValues(bucket: "%s", tag: "%s")
|> yield(name: "%s")