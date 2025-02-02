package models

import "go.mongodb.org/mongo-driver/bson/primitive"

// InsolvencyResourceDao contains the meta-data for the insolvency resource in Mongo
type InsolvencyResourceDao struct {
	ID            primitive.ObjectID         `bson:"_id"`
	TransactionID string                     `bson:"transaction_id"`
	Etag          string                     `bson:"etag"`
	Kind          string                     `bson:"kind"`
	Data          InsolvencyResourceDaoData  `bson:"data"`
	Links         InsolvencyResourceLinksDao `bson:"links"`
}

// InsolvencyResourceDaoData contains the data for the insolvency resource in Mongo
type InsolvencyResourceDaoData struct {
	CompanyNumber string                    `bson:"company_number"`
	CaseType      string                    `bson:"case_type"`
	CompanyName   string                    `bson:"company_name"`
	Practitioners []PractitionerResourceDao `bson:"practitioners"`
}

// InsolvencyResourceLinksDao contains the links for the insolvency resource
type InsolvencyResourceLinksDao struct {
	Self             string `bson:"self"`
	Transaction      string `bson:"transaction"`
	ValidationStatus string `bson:"validation_status"`
}

// PractitionerResourceDao contains the data for for the practitioner resource in Mongo
type PractitionerResourceDao struct {
	ID              string                       `bson:"id"`
	IPCode          string                       `bson:"ip_code"`
	FirstName       string                       `bson:"first_name"`
	LastName        string                       `bson:"last_name"`
	TelephoneNumber string                       `bson:"telephone_number"`
	Email           string                       `bson:"email"`
	Address         AddressResourceDao           `bson:"address"`
	Role            string                       `bson:"role"`
	Links           PractitionerResourceLinksDao `bson:"links"`
}

// AddressResourceDao contains the data for any addresses in Mongo
type AddressResourceDao struct {
	AddressLine1 string `bson:"address_line_1"`
	AddressLine2 string `bson:"address_line_2"`
	Country      string `bson:"country"`
	Locality     string `bson:"locality"`
	Region       string `bson:"region"`
	PostalCode   string `bson:"postal_code"`
}

type PractitionerResourceLinksDao struct {
	Self string `bson:"self"`
}
