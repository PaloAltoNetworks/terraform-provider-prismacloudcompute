package cnnf

import (
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/collection"
	"github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute/internal/api/common"
)

type Rule struct {
	Disabled     bool               `json:"disabled,omitempty"`
	Dst          int                `json:"dst,omitempty"` // ID of each dst firewall entity. 20 bits are used. Max legal value: 2^20-1
	Effect       string             `json:"effect,omitempty"`
	ID           int                `json:"id,omitempty"`
	Modified     string             `json:"modified,omitempty"`
	Name         string             `json:"name,omitempty"`
	Notes        string             `json:"notes,omitempty"`
	Owner        string             `json:"owner,omitempty"`
	Ports        []common.PortRange `json:"ports,omitempty"`
	PreviousName string             `json:"previousName,omitempty"`
	Src          int                `json:"src,omitempty"`
}

type AllowAllConnections struct {
	Inbound  []string `json:"inbound,omitempty"`
	Outbound []string `json:"outbound,omitempty"`
}

type Subnet struct {
	CIDR string `json:"cidr,omitempty"`
	Name string `json:"name,omitempty"`
}

type Entities struct {
	Id          int                     `json:"_id,omitempty"`
	AllowAll    AllowAllConnections     `json:"allowAll,omitempty"`
	Collections []collection.Collection `json:"collection,omitempty"`
	Domains     []string                `json:"domains,omitempty"`
	Name        string                  `json:"name,omitempty"`
	Subnets     []Subnet                `json:"subnets,omitempty"`
	Type        string                  `json:"type,omitempty"`
}
