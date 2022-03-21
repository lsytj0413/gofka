package apis

// ApiVersionsRequest is the request for list broker support ApiKey & ApiVersion
type ApiVersionsRequest struct {
	// ClientSoftwareName is the name of client.
	// Version: >= 3
	// Type: COMPACT_STRING
	ClientSoftwareName string

	// ClientSoftwareVersion is the version of client.
	// Version: >= 3
	// Type: COMPACT_STRING
	ClientSoftwareVersion string

	// TAG_BUFFER: COMPACT_ARRAY
	// Version: >= 3
}
