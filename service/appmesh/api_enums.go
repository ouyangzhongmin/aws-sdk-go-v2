// Code generated by private/model/cli/gen-api/main.go. DO NOT EDIT.

package appmesh

type DurationUnit string

// Enum values for DurationUnit
const (
	DurationUnitMs DurationUnit = "ms"
	DurationUnitS  DurationUnit = "s"
)

func (enum DurationUnit) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum DurationUnit) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type EgressFilterType string

// Enum values for EgressFilterType
const (
	EgressFilterTypeAllowAll EgressFilterType = "ALLOW_ALL"
	EgressFilterTypeDropAll  EgressFilterType = "DROP_ALL"
)

func (enum EgressFilterType) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum EgressFilterType) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type GrpcRetryPolicyEvent string

// Enum values for GrpcRetryPolicyEvent
const (
	GrpcRetryPolicyEventCancelled         GrpcRetryPolicyEvent = "cancelled"
	GrpcRetryPolicyEventDeadlineExceeded  GrpcRetryPolicyEvent = "deadline-exceeded"
	GrpcRetryPolicyEventInternal          GrpcRetryPolicyEvent = "internal"
	GrpcRetryPolicyEventResourceExhausted GrpcRetryPolicyEvent = "resource-exhausted"
	GrpcRetryPolicyEventUnavailable       GrpcRetryPolicyEvent = "unavailable"
)

func (enum GrpcRetryPolicyEvent) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum GrpcRetryPolicyEvent) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HttpMethod string

// Enum values for HttpMethod
const (
	HttpMethodConnect HttpMethod = "CONNECT"
	HttpMethodDelete  HttpMethod = "DELETE"
	HttpMethodGet     HttpMethod = "GET"
	HttpMethodHead    HttpMethod = "HEAD"
	HttpMethodOptions HttpMethod = "OPTIONS"
	HttpMethodPatch   HttpMethod = "PATCH"
	HttpMethodPost    HttpMethod = "POST"
	HttpMethodPut     HttpMethod = "PUT"
	HttpMethodTrace   HttpMethod = "TRACE"
)

func (enum HttpMethod) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HttpMethod) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type HttpScheme string

// Enum values for HttpScheme
const (
	HttpSchemeHttp  HttpScheme = "http"
	HttpSchemeHttps HttpScheme = "https"
)

func (enum HttpScheme) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum HttpScheme) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type ListenerTlsMode string

// Enum values for ListenerTlsMode
const (
	ListenerTlsModeDisabled   ListenerTlsMode = "DISABLED"
	ListenerTlsModePermissive ListenerTlsMode = "PERMISSIVE"
	ListenerTlsModeStrict     ListenerTlsMode = "STRICT"
)

func (enum ListenerTlsMode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum ListenerTlsMode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type MeshStatusCode string

// Enum values for MeshStatusCode
const (
	MeshStatusCodeActive   MeshStatusCode = "ACTIVE"
	MeshStatusCodeDeleted  MeshStatusCode = "DELETED"
	MeshStatusCodeInactive MeshStatusCode = "INACTIVE"
)

func (enum MeshStatusCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum MeshStatusCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type PortProtocol string

// Enum values for PortProtocol
const (
	PortProtocolGrpc  PortProtocol = "grpc"
	PortProtocolHttp  PortProtocol = "http"
	PortProtocolHttp2 PortProtocol = "http2"
	PortProtocolTcp   PortProtocol = "tcp"
)

func (enum PortProtocol) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum PortProtocol) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type RouteStatusCode string

// Enum values for RouteStatusCode
const (
	RouteStatusCodeActive   RouteStatusCode = "ACTIVE"
	RouteStatusCodeDeleted  RouteStatusCode = "DELETED"
	RouteStatusCodeInactive RouteStatusCode = "INACTIVE"
)

func (enum RouteStatusCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum RouteStatusCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type TcpRetryPolicyEvent string

// Enum values for TcpRetryPolicyEvent
const (
	TcpRetryPolicyEventConnectionError TcpRetryPolicyEvent = "connection-error"
)

func (enum TcpRetryPolicyEvent) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum TcpRetryPolicyEvent) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type VirtualNodeStatusCode string

// Enum values for VirtualNodeStatusCode
const (
	VirtualNodeStatusCodeActive   VirtualNodeStatusCode = "ACTIVE"
	VirtualNodeStatusCodeDeleted  VirtualNodeStatusCode = "DELETED"
	VirtualNodeStatusCodeInactive VirtualNodeStatusCode = "INACTIVE"
)

func (enum VirtualNodeStatusCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum VirtualNodeStatusCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type VirtualRouterStatusCode string

// Enum values for VirtualRouterStatusCode
const (
	VirtualRouterStatusCodeActive   VirtualRouterStatusCode = "ACTIVE"
	VirtualRouterStatusCodeDeleted  VirtualRouterStatusCode = "DELETED"
	VirtualRouterStatusCodeInactive VirtualRouterStatusCode = "INACTIVE"
)

func (enum VirtualRouterStatusCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum VirtualRouterStatusCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}

type VirtualServiceStatusCode string

// Enum values for VirtualServiceStatusCode
const (
	VirtualServiceStatusCodeActive   VirtualServiceStatusCode = "ACTIVE"
	VirtualServiceStatusCodeDeleted  VirtualServiceStatusCode = "DELETED"
	VirtualServiceStatusCodeInactive VirtualServiceStatusCode = "INACTIVE"
)

func (enum VirtualServiceStatusCode) MarshalValue() (string, error) {
	return string(enum), nil
}

func (enum VirtualServiceStatusCode) MarshalValueBuf(b []byte) ([]byte, error) {
	b = b[0:0]
	return append(b, enum...), nil
}
