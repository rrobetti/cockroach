// Code generated by gen/main.go. DO NOT EDIT.

package severity

import "github.com/cockroachdb/cockroach/pkg/util/log/logpb"

// UNKNOWN is populated into decoded log entries when the
// severity could not be determined.
const UNKNOWN = logpb.Severity_UNKNOWN

// INFO is used for informational messages, when no action
// is required as a result.
const INFO = logpb.Severity_INFO

// WARNING is used for situations which may require special handling,
// while normal operation is expected to resume automatically.
const WARNING = logpb.Severity_WARNING

// ERROR is used for situations that require special handling,
// when normal operation could not proceed as expected.
// Other operations can continue mostly unaffected.
const ERROR = logpb.Severity_ERROR

// FATAL is used for situations that require an immedate, hard
// server shutdown. A report is also sent to telemetry if telemetry
// is enabled.
const FATAL = logpb.Severity_FATAL

// NONE can be used in filters to specify that no messages
// should be emitted.
const NONE = logpb.Severity_NONE

// DEFAULT is the end sentinel. It is used during command-line
// handling to indicate that another value should be replaced instead
// (depending on which command is being run); see cli/flags.go for
// details.
const DEFAULT = logpb.Severity_DEFAULT
