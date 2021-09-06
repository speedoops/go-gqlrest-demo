package i18n

var LangKey = contextKey("lang")

type contextKey string

func (c contextKey) String() string {
	return "trace/tracespec context key " + string(c)
}
