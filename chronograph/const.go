package chronograph

const (
	Default         = "1970-01-01T00:00:00.000Z"
	Layout          = "2006-01-02T15:04:05.000Z"
	timestampSample = "YYYY-MM-DDTHH:MM:SS.NNNZ"
)

const (
	errorCity      = "City (%s) is incorrect. contains invalid substring. cannot hold non-alpha (exc. underscore, hyphen)"
	errorTimestamp = "Timestamp (%s) is incorrect. does not match supported time layout (" + Default + ")"
)

const (
	promptTimestamp = "Please enter a valid timestamp (" + timestampSample + ")"
)
