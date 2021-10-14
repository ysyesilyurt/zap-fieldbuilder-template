package logger

import "go.uber.org/zap"

/* Generic FieldBuilder Template for uber-go/zap - ysyesilyurt 2021 */

type ZapFieldBuilder struct {
	fields                    []zap.Field
	existingFieldIndexMapping map[string]int
}

/* FieldBuilder is the generic constructor for ZapFieldBuilder

	ZapFieldBuilder allows you to chain fields with builder methods and as soon as you call Build, it yields
	the underlying []zap.Field according to the order in builder chain.

	If same builder method called multiple times on a ZapFieldBuilder variable, order does not change (order in
	the first call to the field will be used) and underlying zap.Field is going to replaced every time.

	You can methods to ZapFieldBuilder with ptr receiver as builders (for sample builders check sample_builders.go)
	In builder method definitions below call with your zap.Field would be sufficient. A sample builder:

		func (zfb *ZapFieldBuilder) Status(status int) *ZapFieldBuilder {
			return zfb.addOrReplaceField(zap.Int(LogFieldStatus, status))
		}
*/
func FieldBuilder() *ZapFieldBuilder {
	return &ZapFieldBuilder{
		fields:                    make([]zap.Field, 0),
		existingFieldIndexMapping: make(map[string]int),
	}
}

// Appends the zap.Field
func (zfb *ZapFieldBuilder) appendField(field zap.Field) {
	zfb.fields = append(zfb.fields, field)
	zfb.existingFieldIndexMapping[field.Key] = len(zfb.fields) - 1
}

// Replaces an already existing zap.Field
func (zfb *ZapFieldBuilder) replaceField(field zap.Field) {
	zfb.fields[zfb.existingFieldIndexMapping[field.Key]] = field
}

// Adds the zap.Field if it hasn't been added before or else replaces the existing one (preserving the order)
func (zfb *ZapFieldBuilder) addOrReplaceField(field zap.Field) *ZapFieldBuilder {
	if _, exists := zfb.existingFieldIndexMapping[field.Key]; exists {
		zfb.replaceField(field)
	} else {
		zfb.appendField(field)
	}
	return zfb
}

// Builds (i.e. Yields) the resulting []zap.Field (Use directly inside your zap log line)
func (zfb *ZapFieldBuilder) Build() []zap.Field {
	return zfb.fields
}

/* Implement your custom Builders below this line */