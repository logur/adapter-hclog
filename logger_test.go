package hclog

import (
	"bytes"
	"regexp"
	"strings"
	"testing"

	"github.com/hashicorp/go-hclog"
	"logur.dev/logur"
	"logur.dev/logur/conformance"
)

// nolint: gochecknoglobals
var logLineRegex = regexp.MustCompile(`.* \[(.*)\] {1,2}(.*): (.*)`)

func TestLogger(t *testing.T) {
	suite := conformance.TestSuite{
		LoggerFactory: func(level logur.Level) (logur.Logger, conformance.TestLogger) {
			var buf bytes.Buffer
			logger := hclog.New(&hclog.LoggerOptions{
				Level:  hclog.Level(level + 1),
				Output: &buf,
			})

			return New(logger), conformance.TestLoggerFunc(func() []logur.LogEvent {
				lines := strings.Split(strings.TrimSuffix(buf.String(), "\n"), "\n")

				events := make([]logur.LogEvent, len(lines))

				for key, line := range lines {
					match := logLineRegex.FindStringSubmatch(line)

					level, _ := logur.ParseLevel(strings.ToLower(match[1]))

					rawFields := strings.Fields(match[3])
					fields := make(map[string]interface{})

					for _, rawField := range rawFields {
						field := strings.SplitN(rawField, "=", 2)

						fields[field[0]] = field[1]
					}

					events[key] = logur.LogEvent{
						Line:   match[2],
						Level:  level,
						Fields: fields,
					}
				}

				return events
			})
		},
	}

	suite.Run(t)
}
