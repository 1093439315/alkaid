/*
 * Copyright 2020. The Alkaid Authors. All rights reserved.
 * Use of this source code is governed by a MIT-style
 * license that can be found in the LICENSE file.
 *
 * Alkaid is a BaaS service based on Hyperledger Fabric.
 *
 */

package types

import "github.com/yakumioto/alkaid/internal/utils/certificate"

const (
	SignCAType = "sign"
	TLSCAType  = "tls"

	OrdererOrgType = "orderer"
	PeerOrgType    = "peer"
)

// Organization in the network
type Organization struct {
	ID             int64    `json:"-"`
	OrganizationID string   `json:"organization_id,omitempty" binding:"required"`
	Name           string   `json:"name,omitempty" binding:"required"`
	NetworkID      []string `json:"network_id,omitempty"`
	Domain         string   `json:"domain,omitempty" binding:"required,fqdn"`
	Description    string   `json:"description,omitempty"`

	// Type value is orderer or peer
	Type string `json:"type,omitempty" binding:"required,oneof=orderer peer"`

	// The following fields are the fields that generate the certificate
	Country            string `json:"country,omitempty"`
	Province           string `json:"province,omitempty"`
	Locality           string `json:"locality,omitempty"`
	OrganizationalUnit string `json:"organizational_unit,omitempty"`
	StreetAddress      string `json:"street_address,omitempty"`
	PostalCode         string `json:"postal_code,omitempty"`

	// sign and tsl root ca
	SignCAPrivateKey  []byte `json:"-"`
	TLSCAPrivateKey   []byte `json:"-"`
	SignCACertificate []byte `json:"sign_ca_certificate,omitempty"`
	TLSCACertificate  []byte `json:"tlsca_certificate,omitempty"`

	CreateAt int64 `json:"create_at,omitempty"`
	UpdateAt int64 `json:"update_at,omitempty"`
}

// NewOrganization Default parameter
func NewOrganization() *Organization {
	return &Organization{
		Country:    "China",
		Province:   "Beijing",
		Locality:   "Beijing",
		PostalCode: "100000",
	}
}

func (o *Organization) GetCertificatePkixName() *certificate.PkixName {
	return &certificate.PkixName{
		OrgName:       o.OrganizationID,
		Country:       o.Country,
		Province:      o.Province,
		Locality:      o.Locality,
		OrgUnit:       o.OrganizationalUnit,
		StreetAddress: o.StreetAddress,
		PostalCode:    o.PostalCode,
	}
}
