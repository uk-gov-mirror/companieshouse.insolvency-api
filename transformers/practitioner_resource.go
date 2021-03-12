package transformers

import "github.com/companieshouse/insolvency-api/models"

// PractitionerResourceRequestToDB will take the input request from the REST call and transform it to a dao ready for
// insertion into the database
func PractitionerResourceRequestToDB(req *models.PractitionerRequest) *models.PractitionerResourceDao {

	dao := &models.PractitionerResourceDao{
		FirstName: req.FirstName,
		LastName:  req.LastName,
		Address: models.AddressResourceDao{
			AddressLine1: req.Address.AddressLine1,
			AddressLine2: req.Address.AddressLine2,
			Country:      req.Address.Country,
			Locality:     req.Address.Locality,
			Region:       req.Address.Region,
			PostalCode:   req.Address.PostalCode,
		},
		Role: req.Role,
	}

	return dao
}

// PractitionerResourceDaoToCreatedResponse will transform an practitioner resource dao that has successfully been created into
// a http response entity
func PractitionerResourceDaoToCreatedResponse(model *models.PractitionerResourceDao) *models.CreatedPractitionerResource {
	return &models.CreatedPractitionerResource{
		FirstName: model.FirstName,
		LastName:  model.LastName,
		Address: models.CreatedAddressResource{
			AddressLine1: model.Address.AddressLine1,
			AddressLine2: model.Address.AddressLine2,
			Country:      model.Address.Country,
			Locality:     model.Address.Locality,
			Region:       model.Address.Region,
			PostalCode:   model.Address.PostalCode,
		},
		Role: model.Role,
	}
}
