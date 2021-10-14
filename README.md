# zap-fieldbuilder-template
A Handy Generic Logging `FieldBuilder` Template around [uber-go/zap](https://github.com/uber-go/zap). You can check [this tiny blog](https://medium.com/@ysyesilyurt/logging-field-builder-in-go-76064c663d8f) for further information 🙂

### Who is this template for?
Anybody who likes to code in Go, also uses Structured Logging with [uber-go/zap](https://github.com/uber-go/zap) and (most importantly) loves _cleaner_ and _neater_ code.

### What does it do?
Instead of having lots of logging lines repeatedly along the codebase such as the following

```
logger.Debug("Firing Event",
   zap.Int("event-id", id),
   zap.String("event-type", "TakeOff"),
   zap.String("issuer", user.getName()),
   zap.Int("account-id", user.getAccountId()),
   zap.Time("timestamp", time.Now().Format(time.RFC3339)))
```

... having those fields built flexibly and reusable like this

```
fields := logger.FieldBuilder().
   EventId(id).
   EventType("TakeOff").
   Issuer(user.getName()).
   AccountId(user.getAccountId()).
   Timestamp(time.Now().Format(time.RFC3339))
   
/* ... */

if /* ... */ {
	log.Debug("Firing Event...", fields.AnotherField(fieldPayload).Build()...)
} else {
	log.Debug("Firing Event...", fields.Build()...)
}

```

## Usage
1. Clone this template repository
```
git clone https://github.com/ysyesilyurt/zap-fieldbuilder-template.git
```
2. Using the `logger/field_builder.go` adopt this FieldBuilder template to your codebase. 

3. Simply add all your `Builder` methods that you are going to use in your logging statement as `zap.Field` when needed as needed. You can use the follwing `ZapFieldBuilder` method body pattern.
```
	func (zfb *ZapFieldBuilder) Status(status int) *ZapFieldBuilder {
		return zfb.addOrReplaceField(zap.Int(LogFieldStatus, status))
	}
```
4. Initialize your FieldBuilder using the constructor `FieldBuilder()` and just start building your `zap.Field` s with your builder methods.

5. Now you can log them to anywhere using `Build()` method which would yield you the underlying `[]zap.Field`.