package gateways

import "time"

type IResource interface {
	GetID() string
	GetName() string
	GetUrl() string
	GetCreated() time.Time
	GetUpdated() time.Time
}
