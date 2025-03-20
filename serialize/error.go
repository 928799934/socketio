package serialize

import (
	erro "github.com/928799934/socketio/internal/errors"
)

const (
	ErrUnsupportedUseRead erro.String = "Serialize() method unsupported, use the Read() method instead"
	ErrUnsupported        erro.State  = "method: unsupported"
)
