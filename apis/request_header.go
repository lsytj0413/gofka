package apis

// RequestHeader is the header of request message, all message will begin with the header
type RequestHeader struct {
	// RequestApiKey is the ApiKey of each RequestApi, this is unique of every request.
	// EX: the ApiVersions Request's key is 18
	RequestApiKey int16

	// RequestApiVersion is the ApiVersion of current RequestApi.
	// EX: the ApiVersions Request have four verions, 0/1/23
	RequestApiVersion int16

	// CorrelationId is the id the each request, this should be unique of (client, broker).
	// If the client use non-blocking IO, which send requests even while awaiting responses,
	//  the CorrelationId can help to distinct which request it's response to.
	CorrelationId int32

	// ClientId can use to distinct with other client.
	// Version: >= 1
	// Type: NUALLABLE_STRING
	ClientId string

	// TAG_BUFFER: COMPACT_ARRAY
	// Version: >= 2
	// More Information: https://cwiki.apache.org/confluence/display/KAFKA/KIP-482%3A+The+Kafka+Protocol+should+Support+Optional+Tagged+Fields
	// Each Tag format is:
	//   + length: UNSIGNED_VARINT
	//   + field.1.tag: UNSIGNED_VARINT
	//   + field.1.length: UNSIGNED_VARINT
	//   + field.1.data:
}
