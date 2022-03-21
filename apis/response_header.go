package apis

// ResponseHeader is the header of response message, all message will begin with the header
type ResponseHeader struct {
	// CorrelationId is the id the each response, this is same as the associated request.
	CorrelationId int32

	// TAG_BUFFER: COMPACT_ARRAY
	// Version: >= 1
	// More Information: https://cwiki.apache.org/confluence/display/KAFKA/KIP-482%3A+The+Kafka+Protocol+should+Support+Optional+Tagged+Fields
	// Each Tag format is:
	//   + length: UNSIGNED_VARINT
	//   + field.1.tag: UNSIGNED_VARINT
	//   + field.1.length: UNSIGNED_VARINT
	//   + field.1.data:
}
