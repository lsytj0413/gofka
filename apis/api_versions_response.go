package apis

// ApiVersionsResponse is the response for ApiVersionsRequest
type ApiVersionsResponse struct {
	// ErrorCode is the top-level error code.
	ErrorCode int16

	// ApiKeys is the list of api supported by the broker.
	// Type: ARRAY
	// Type: COMPACT_ARRAY, Version >= 3
	ApiKeys []ApiKey

	// ThrottleTimeMs is the duration in milliseconds for request was throttled due to
	//    a quota violation.
	// Version: >= 1
	ThrottleTimeMs int32

	// TAG_BUFFER: COMPACT_ARRAY
	// Version: >= 3
}

// ApiKey is the api & versions
type ApiKey struct {
	// ApiKey is the Key of each RequestApi
	ApiKey int16

	// MinVersion is the broker support min version.
	MinVersion int16

	// MaxVersion is the broker support max version.
	MaxVersion int16

	// TAG_BUFFER: COMPACT_ARRAY
	// Version: >= 3
}
