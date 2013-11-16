package iface

import (
    "github.com/dpx-infinity/imaged/message"
)

type Type interface {
    Push(msg message.Envelope)
    Register(ch chan<- message.Envelope, matcher func(message.Envelope) bool)
}
