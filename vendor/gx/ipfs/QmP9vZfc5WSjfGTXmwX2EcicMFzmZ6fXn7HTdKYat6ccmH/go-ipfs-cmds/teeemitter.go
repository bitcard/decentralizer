package cmds

import (
	"gx/ipfs/QmQp2a2Hhb7F6eK2A5hN8f9aJy4mtkEikL9Zj4cgB7d1dD/go-ipfs-cmdkit"
)

// NewTeeEmitter creates a new ResponseEmitter.
// Writing to it will write to both the passed ResponseEmitters.
func NewTeeEmitter(re1, re2 ResponseEmitter) ResponseEmitter {
	return &teeEmitter{re1, re2}
}

type teeEmitter struct {
	ResponseEmitter

	re ResponseEmitter
}

func (re *teeEmitter) Close() error {
	err1 := re.ResponseEmitter.Close()
	err2 := re.re.Close()

	if err1 != nil {
		return err1
	}

	// XXX we drop the second error if both fail
	return err2
}

func (re *teeEmitter) Emit(v interface{}) error {
	err1 := re.ResponseEmitter.Emit(v)
	err2 := re.re.Emit(v)

	if err1 != nil {
		return err1
	}

	// XXX we drop the second error if both fail
	return err2
}

func (re *teeEmitter) SetLength(l uint64) {
	re.ResponseEmitter.SetLength(l)
	re.re.SetLength(l)
}

func (re *teeEmitter) SetError(err interface{}, code cmdkit.ErrorType) {
	re.ResponseEmitter.SetError(err, code)
	re.re.SetError(err, code)
}

type TeeError struct {
	err1, err2 error
}

func (err TeeError) BothNil() bool {
	return err.err1 == nil && err.err2 == nil
}

func (err TeeError) Error() string {
	if err.err1 != nil && err.err2 != nil {
		return "1: " + err.err1.Error() + "\n2: " + err.err2.Error()
	} else if err.err1 != nil {
		return "1: " + err.err1.Error()
	} else if err.err2 != nil {
		return "2: " + err.err2.Error()
	} else {
		return ""
	}
}
