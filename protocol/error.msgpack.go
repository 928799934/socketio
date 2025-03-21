package protocol

import erro "github.com/928799934/socketio/internal/errors"

const (
	ErrDecodeBase64Failed erro.StringF = "failed to decode msgpack base64 field:: %w"
	ErrDecodeFieldFailed  erro.StringF = "failed to decode msgpack field:: %w"
	ErrEncodeFieldFailed  erro.StringF = "failed to encode msgpack field:: %w"
)
