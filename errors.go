package pigpen_cipher

import "errors"

// ErrPlaintextMustIsLetters 要加密的明文必须都是英文字母
var ErrPlaintextMustIsLetters = errors.New("plaintext must is letters")
