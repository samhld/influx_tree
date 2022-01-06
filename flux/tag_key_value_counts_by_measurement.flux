import "influxdata/influxdb/schema"
cardinalityByTag = (bucket, measurement) => schema.tagKeys(bucket: bucket, predicate: (r) => r._measurement == measurement)
	|> filter(fn: (r) => r._value != "_start" and r._value != "_stop")
	|> map(
		fn: (r) => ({
			tag: r._value,
			_value: (schema.tagValues(bucket: bucket, tag: r._value)
				|> count()
				|> findRecord(fn: (key) => true, idx: 0))._value,
		}),
	)
	|> group(columns: ["tag"])
	|> sum()
cardinalityByTag(bucket: "%s", measurement: "%s")
	|> group()