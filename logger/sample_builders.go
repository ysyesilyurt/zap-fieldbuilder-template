package logger

import "go.uber.org/zap"

/* Sample builder implementations for zap-fieldbuilder-template - ysyesilyurt 2021 */

const (
	LogFieldAccountId   = "account-id"
	LogFieldIssuer      = "issuer"
	LogFieldRequestId   = "request-id"
	LogFieldMethod      = "method"
	LogFieldPath        = "path"
	LogFieldIP          = "ip"
	LogFieldUrl         = "url"
	LogTimestamp        = "timestamp"
	LogFieldStatus      = "status"
	LogFieldDurationMs  = "duration-ms"
	LogFieldIsPanic     = "is-panic"
	LogFieldRequestDump = "request-dump"
	LogFieldHeaders     = "headers"
	LogFieldStackTrace  = "stack"
	LogFieldType        = "type"
	LogFieldMessageType = "message-type"
	LogFieldEventId     = "event-id"
	LogFieldEventType   = "event-type"
	LogFieldEventTopic  = "event-topic"
)

func (zfb *ZapFieldBuilder) Error(err error) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.Error(err))
}

func (zfb *ZapFieldBuilder) AccountId(accountId int) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.Int(LogFieldAccountId, accountId))
}

func (zfb *ZapFieldBuilder) RequestId(reqId string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldRequestId, reqId))
}

func (zfb *ZapFieldBuilder) Method(method string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldMethod, method))
}

func (zfb *ZapFieldBuilder) Path(path string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldPath, path))
}

func (zfb *ZapFieldBuilder) Url(url string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldUrl, url))
}

func (zfb *ZapFieldBuilder) Ip(ip string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldIP, ip))
}

func (zfb *ZapFieldBuilder) Timestamp(timestamp string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogTimestamp, timestamp))
}

func (zfb *ZapFieldBuilder) Status(status int) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.Int(LogFieldStatus, status))
}

func (zfb *ZapFieldBuilder) DurationMs(durationMs int64) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.Int64(LogFieldDurationMs, durationMs))
}

func (zfb *ZapFieldBuilder) Issuer(iss string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldIssuer, iss))
}

func (zfb *ZapFieldBuilder) IsPanic(isPanic bool) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.Bool(LogFieldIsPanic, isPanic))
}

func (zfb *ZapFieldBuilder) RequestDump(rd string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldRequestDump, rd))
}

func (zfb *ZapFieldBuilder) Headers(h []string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.Strings(LogFieldHeaders, h))
}

func (zfb *ZapFieldBuilder) Stacktrace(st []string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.Strings(LogFieldStackTrace, st))
}

func (zfb *ZapFieldBuilder) Type(t string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldType, t))
}

func (zfb *ZapFieldBuilder) MessageType(mt string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldMessageType, mt))
}

func (zfb *ZapFieldBuilder) EventId(eventId int) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.Int(LogFieldEventId, eventId))
}

func (zfb *ZapFieldBuilder) EventType(eventType string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldEventType, eventType))
}

func (zfb *ZapFieldBuilder) EventTopic(topic string) *ZapFieldBuilder {
	return zfb.addOrReplaceField(zap.String(LogFieldEventTopic, topic))
}
