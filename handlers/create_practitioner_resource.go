package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/companieshouse/chs.go/log"
	"github.com/companieshouse/insolvency-api/constants"
	"github.com/companieshouse/insolvency-api/dao"
	"github.com/companieshouse/insolvency-api/models"
	"github.com/companieshouse/insolvency-api/transformers"
	"github.com/companieshouse/insolvency-api/utils"
	"github.com/gorilla/mux"
)

// HandleCreatePractitionerResource creates a new practitioner resource
func HandleCreatePractitionerResource(svc dao.Service) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, req *http.Request) {
		log.InfoR(req, "start POST request for practitioner resource")

		vars := mux.Vars(req)
		transactionID, err := utils.GetTransactionIDFromVars(vars)
		if err != nil {
			log.ErrorR(req, err)
			m := models.NewMessageResponse("transaction id is not in the url path")
			utils.WriteJSONWithStatus(w, req, m, http.StatusBadRequest)
			return
		}

		// Decode incoming request and check if it is valid
		var request models.PractitionerRequest
		err = json.NewDecoder(req.Body).Decode(&request)
		if err != nil {
			log.ErrorR(req, fmt.Errorf("invalid request"))
			m := models.NewMessageResponse(fmt.Sprintf("failed to read request body for transaction %s", transactionID))
			utils.WriteJSONWithStatus(w, req, m, http.StatusBadRequest)
			return
		}

		// Check if practitioner role supplied is valid
		if ok := constants.IsInRoleList(request.Role); !ok {
			log.ErrorR(req, fmt.Errorf("invalid practitioner role"))
			m := models.NewMessageResponse(fmt.Sprintf("the practitioner supplied is not valid %s", request.Role))
			utils.WriteJSONWithStatus(w, req, m, http.StatusBadRequest)
			return
		}

		// Update insolvency resource in mongo with practitioner data
		model := transformers.PractitionerResourceRequestToDB(&request)

		err = svc.CreatePractitionerResource(model, transactionID)
		if err != nil {
			log.ErrorR(req, fmt.Errorf("failed to create practitioner resource in database"))
			m := models.NewMessageResponse(fmt.Sprintf("there was a problem handling your request for transaction %s", transactionID))
			utils.WriteJSONWithStatus(w, req, m, http.StatusInternalServerError)
			return
		}

		utils.WriteJSONWithStatus(w, req, transformers.PractitionerResourceDaoToCreatedResponse(model), http.StatusOK)
	})
}
