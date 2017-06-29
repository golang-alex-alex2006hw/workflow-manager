package server

import (
	"bytes"
	"context"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"strconv"

	"github.com/Clever/workflow-manager/gen-go/models"
	"github.com/go-errors/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/gorilla/mux"
	"gopkg.in/Clever/kayvee-go.v6/logger"
)

var _ = strconv.ParseInt
var _ = strfmt.Default
var _ = swag.ConvertInt32
var _ = errors.New
var _ = mux.Vars
var _ = bytes.Compare
var _ = ioutil.ReadAll

var formats = strfmt.Default
var _ = formats

// convertBase64 takes in a string and returns a strfmt.Base64 if the input
// is valid base64 and an error otherwise.
func convertBase64(input string) (strfmt.Base64, error) {
	temp, err := formats.Parse("byte", input)
	if err != nil {
		return strfmt.Base64{}, err
	}
	return *temp.(*strfmt.Base64), nil
}

// convertDateTime takes in a string and returns a strfmt.DateTime if the input
// is a valid DateTime and an error otherwise.
func convertDateTime(input string) (strfmt.DateTime, error) {
	temp, err := formats.Parse("date-time", input)
	if err != nil {
		return strfmt.DateTime{}, err
	}
	return *temp.(*strfmt.DateTime), nil
}

// convertDate takes in a string and returns a strfmt.Date if the input
// is a valid Date and an error otherwise.
func convertDate(input string) (strfmt.Date, error) {
	temp, err := formats.Parse("date", input)
	if err != nil {
		return strfmt.Date{}, err
	}
	return *temp.(*strfmt.Date), nil
}

func jsonMarshalNoError(i interface{}) string {
	bytes, err := json.MarshalIndent(i, "", "\t")
	if err != nil {
		// This should never happen
		return ""
	}
	return string(bytes)
}

// statusCodeForHealthCheck returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForHealthCheck(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	default:
		return -1
	}
}

func (h handler) HealthCheckHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	err := h.HealthCheck(ctx)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForHealthCheck(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(""))

}

// newHealthCheckInput takes in an http.Request an returns the input struct.
func newHealthCheckInput(r *http.Request) (*models.HealthCheckInput, error) {
	var input models.HealthCheckInput

	var err error
	_ = err

	return &input, nil
}

// statusCodeForGetJobsForWorkflow returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetJobsForWorkflow(obj interface{}) int {

	switch obj.(type) {

	case *[]models.Job:
		return 200

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.NotFound:
		return 404

	case []models.Job:
		return 200

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.NotFound:
		return 404

	default:
		return -1
	}
}

func (h handler) GetJobsForWorkflowHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newGetJobsForWorkflowInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.GetJobsForWorkflow(ctx, input)

	// Success types that return an array should never return nil so let's make this easier
	// for consumers by converting nil arrays to empty arrays
	if resp == nil {
		resp = []models.Job{}
	}

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForGetJobsForWorkflow(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForGetJobsForWorkflow(resp))
	w.Write(respBytes)

}

// newGetJobsForWorkflowInput takes in an http.Request an returns the input struct.
func newGetJobsForWorkflowInput(r *http.Request) (*models.GetJobsForWorkflowInput, error) {
	var input models.GetJobsForWorkflowInput

	var err error
	_ = err

	workflowNameStrs := r.URL.Query()["workflowName"]
	if len(workflowNameStrs) == 0 {
		return nil, errors.New("parameter must be specified")
	}

	if len(workflowNameStrs) > 0 {
		var workflowNameTmp string
		workflowNameStr := workflowNameStrs[0]
		workflowNameTmp, err = workflowNameStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.WorkflowName = workflowNameTmp
	}

	return &input, nil
}

// statusCodeForStartJobForWorkflow returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForStartJobForWorkflow(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.Job:
		return 200

	case *models.NotFound:
		return 404

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.Job:
		return 200

	case models.NotFound:
		return 404

	default:
		return -1
	}
}

func (h handler) StartJobForWorkflowHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newStartJobForWorkflowInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate(nil)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.StartJobForWorkflow(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForStartJobForWorkflow(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForStartJobForWorkflow(resp))
	w.Write(respBytes)

}

// newStartJobForWorkflowInput takes in an http.Request an returns the input struct.
func newStartJobForWorkflowInput(r *http.Request) (*models.JobInput, error) {
	var input models.JobInput

	var err error
	_ = err

	data, err := ioutil.ReadAll(r.Body)

	if len(data) > 0 {
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&input); err != nil {
			return nil, err
		}
	}

	return &input, nil
}

// statusCodeForCancelJob returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForCancelJob(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.NotFound:
		return 404

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.NotFound:
		return 404

	default:
		return -1
	}
}

func (h handler) CancelJobHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newCancelJobInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = h.CancelJob(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForCancelJob(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	w.WriteHeader(200)
	w.Write([]byte(""))

}

// newCancelJobInput takes in an http.Request an returns the input struct.
func newCancelJobInput(r *http.Request) (*models.CancelJobInput, error) {
	var input models.CancelJobInput

	var err error
	_ = err

	jobIdStr := mux.Vars(r)["jobId"]
	if len(jobIdStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	jobIdStrs := []string{jobIdStr}

	if len(jobIdStrs) > 0 {
		var jobIdTmp string
		jobIdStr := jobIdStrs[0]
		jobIdTmp, err = jobIdStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.JobId = jobIdTmp
	}

	data, err := ioutil.ReadAll(r.Body)
	if len(data) == 0 {
		return nil, errors.New("parameter must be specified")
	}

	if len(data) > 0 {
		input.Reason = &models.CancelReason{}
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(input.Reason); err != nil {
			return nil, err
		}
	}

	return &input, nil
}

// statusCodeForGetJob returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetJob(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.Job:
		return 200

	case *models.NotFound:
		return 404

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.Job:
		return 200

	case models.NotFound:
		return 404

	default:
		return -1
	}
}

func (h handler) GetJobHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	jobId, err := newGetJobInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = models.ValidateGetJobInput(jobId)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.GetJob(ctx, jobId)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForGetJob(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForGetJob(resp))
	w.Write(respBytes)

}

// newGetJobInput takes in an http.Request an returns the jobId parameter
// that it contains. It returns an error if the request doesn't contain the parameter.
func newGetJobInput(r *http.Request) (string, error) {
	jobId := mux.Vars(r)["jobId"]
	if len(jobId) == 0 {
		return "", errors.New("Parameter jobId must be specified")
	}
	return jobId, nil
}

// statusCodeForPostStateResource returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForPostStateResource(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.StateResource:
		return 201

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.StateResource:
		return 201

	default:
		return -1
	}
}

func (h handler) PostStateResourceHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newPostStateResourceInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate(nil)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.PostStateResource(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForPostStateResource(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForPostStateResource(resp))
	w.Write(respBytes)

}

// newPostStateResourceInput takes in an http.Request an returns the input struct.
func newPostStateResourceInput(r *http.Request) (*models.NewStateResource, error) {
	var input models.NewStateResource

	var err error
	_ = err

	data, err := ioutil.ReadAll(r.Body)

	if len(data) > 0 {
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&input); err != nil {
			return nil, err
		}
	}

	return &input, nil
}

// statusCodeForGetStateResource returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetStateResource(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.NotFound:
		return 404

	case *models.StateResource:
		return 200

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.NotFound:
		return 404

	case models.StateResource:
		return 200

	default:
		return -1
	}
}

func (h handler) GetStateResourceHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newGetStateResourceInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.GetStateResource(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForGetStateResource(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForGetStateResource(resp))
	w.Write(respBytes)

}

// newGetStateResourceInput takes in an http.Request an returns the input struct.
func newGetStateResourceInput(r *http.Request) (*models.GetStateResourceInput, error) {
	var input models.GetStateResourceInput

	var err error
	_ = err

	namespaceStr := mux.Vars(r)["namespace"]
	if len(namespaceStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	namespaceStrs := []string{namespaceStr}

	if len(namespaceStrs) > 0 {
		var namespaceTmp string
		namespaceStr := namespaceStrs[0]
		namespaceTmp, err = namespaceStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Namespace = namespaceTmp
	}

	nameStr := mux.Vars(r)["name"]
	if len(nameStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	nameStrs := []string{nameStr}

	if len(nameStrs) > 0 {
		var nameTmp string
		nameStr := nameStrs[0]
		nameTmp, err = nameStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Name = nameTmp
	}

	return &input, nil
}

// statusCodeForPutStateResource returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForPutStateResource(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.StateResource:
		return 201

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.StateResource:
		return 201

	default:
		return -1
	}
}

func (h handler) PutStateResourceHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newPutStateResourceInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.PutStateResource(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForPutStateResource(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForPutStateResource(resp))
	w.Write(respBytes)

}

// newPutStateResourceInput takes in an http.Request an returns the input struct.
func newPutStateResourceInput(r *http.Request) (*models.PutStateResourceInput, error) {
	var input models.PutStateResourceInput

	var err error
	_ = err

	namespaceStr := mux.Vars(r)["namespace"]
	if len(namespaceStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	namespaceStrs := []string{namespaceStr}

	if len(namespaceStrs) > 0 {
		var namespaceTmp string
		namespaceStr := namespaceStrs[0]
		namespaceTmp, err = namespaceStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Namespace = namespaceTmp
	}

	nameStr := mux.Vars(r)["name"]
	if len(nameStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	nameStrs := []string{nameStr}

	if len(nameStrs) > 0 {
		var nameTmp string
		nameStr := nameStrs[0]
		nameTmp, err = nameStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Name = nameTmp
	}

	data, err := ioutil.ReadAll(r.Body)

	if len(data) > 0 {
		input.NewStateResource = &models.NewStateResource{}
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(input.NewStateResource); err != nil {
			return nil, err
		}
	}

	return &input, nil
}

// statusCodeForGetWorkflows returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetWorkflows(obj interface{}) int {

	switch obj.(type) {

	case *[]models.Workflow:
		return 200

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case []models.Workflow:
		return 200

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	default:
		return -1
	}
}

func (h handler) GetWorkflowsHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	resp, err := h.GetWorkflows(ctx)

	// Success types that return an array should never return nil so let's make this easier
	// for consumers by converting nil arrays to empty arrays
	if resp == nil {
		resp = []models.Workflow{}
	}

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForGetWorkflows(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForGetWorkflows(resp))
	w.Write(respBytes)

}

// newGetWorkflowsInput takes in an http.Request an returns the input struct.
func newGetWorkflowsInput(r *http.Request) (*models.GetWorkflowsInput, error) {
	var input models.GetWorkflowsInput

	var err error
	_ = err

	return &input, nil
}

// statusCodeForNewWorkflow returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForNewWorkflow(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.Workflow:
		return 201

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.Workflow:
		return 201

	default:
		return -1
	}
}

func (h handler) NewWorkflowHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newNewWorkflowInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate(nil)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.NewWorkflow(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForNewWorkflow(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForNewWorkflow(resp))
	w.Write(respBytes)

}

// newNewWorkflowInput takes in an http.Request an returns the input struct.
func newNewWorkflowInput(r *http.Request) (*models.NewWorkflowRequest, error) {
	var input models.NewWorkflowRequest

	var err error
	_ = err

	data, err := ioutil.ReadAll(r.Body)

	if len(data) > 0 {
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(&input); err != nil {
			return nil, err
		}
	}

	return &input, nil
}

// statusCodeForGetWorkflowVersionsByName returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetWorkflowVersionsByName(obj interface{}) int {

	switch obj.(type) {

	case *[]models.Workflow:
		return 200

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.NotFound:
		return 404

	case []models.Workflow:
		return 200

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.NotFound:
		return 404

	default:
		return -1
	}
}

func (h handler) GetWorkflowVersionsByNameHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newGetWorkflowVersionsByNameInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.GetWorkflowVersionsByName(ctx, input)

	// Success types that return an array should never return nil so let's make this easier
	// for consumers by converting nil arrays to empty arrays
	if resp == nil {
		resp = []models.Workflow{}
	}

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForGetWorkflowVersionsByName(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForGetWorkflowVersionsByName(resp))
	w.Write(respBytes)

}

// newGetWorkflowVersionsByNameInput takes in an http.Request an returns the input struct.
func newGetWorkflowVersionsByNameInput(r *http.Request) (*models.GetWorkflowVersionsByNameInput, error) {
	var input models.GetWorkflowVersionsByNameInput

	var err error
	_ = err

	nameStr := mux.Vars(r)["name"]
	if len(nameStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	nameStrs := []string{nameStr}

	if len(nameStrs) > 0 {
		var nameTmp string
		nameStr := nameStrs[0]
		nameTmp, err = nameStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Name = nameTmp
	}

	latestStrs := r.URL.Query()["latest"]

	if len(latestStrs) == 0 {
		latestStrs = []string{"true"}
	}
	if len(latestStrs) > 0 {
		var latestTmp bool
		latestStr := latestStrs[0]
		latestTmp, err = strconv.ParseBool(latestStr)
		if err != nil {
			return nil, err
		}
		input.Latest = &latestTmp
	}

	return &input, nil
}

// statusCodeForUpdateWorkflow returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForUpdateWorkflow(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.NotFound:
		return 404

	case *models.Workflow:
		return 201

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.NotFound:
		return 404

	case models.Workflow:
		return 201

	default:
		return -1
	}
}

func (h handler) UpdateWorkflowHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newUpdateWorkflowInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.UpdateWorkflow(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForUpdateWorkflow(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForUpdateWorkflow(resp))
	w.Write(respBytes)

}

// newUpdateWorkflowInput takes in an http.Request an returns the input struct.
func newUpdateWorkflowInput(r *http.Request) (*models.UpdateWorkflowInput, error) {
	var input models.UpdateWorkflowInput

	var err error
	_ = err

	data, err := ioutil.ReadAll(r.Body)

	if len(data) > 0 {
		input.NewWorkflowRequest = &models.NewWorkflowRequest{}
		if err := json.NewDecoder(bytes.NewReader(data)).Decode(input.NewWorkflowRequest); err != nil {
			return nil, err
		}
	}

	nameStr := mux.Vars(r)["name"]
	if len(nameStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	nameStrs := []string{nameStr}

	if len(nameStrs) > 0 {
		var nameTmp string
		nameStr := nameStrs[0]
		nameTmp, err = nameStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Name = nameTmp
	}

	return &input, nil
}

// statusCodeForGetWorkflowByNameAndVersion returns the status code corresponding to the returned
// object. It returns -1 if the type doesn't correspond to anything.
func statusCodeForGetWorkflowByNameAndVersion(obj interface{}) int {

	switch obj.(type) {

	case *models.BadRequest:
		return 400

	case *models.InternalError:
		return 500

	case *models.NotFound:
		return 404

	case *models.Workflow:
		return 200

	case models.BadRequest:
		return 400

	case models.InternalError:
		return 500

	case models.NotFound:
		return 404

	case models.Workflow:
		return 200

	default:
		return -1
	}
}

func (h handler) GetWorkflowByNameAndVersionHandler(ctx context.Context, w http.ResponseWriter, r *http.Request) {

	input, err := newGetWorkflowByNameAndVersionInput(r)
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	err = input.Validate()

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.BadRequest{Message: err.Error()}), http.StatusBadRequest)
		return
	}

	resp, err := h.GetWorkflowByNameAndVersion(ctx, input)

	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		if btErr, ok := err.(*errors.Error); ok {
			logger.FromContext(ctx).AddContext("stacktrace", string(btErr.Stack()))
		}
		statusCode := statusCodeForGetWorkflowByNameAndVersion(err)
		if statusCode == -1 {
			err = models.InternalError{Message: err.Error()}
			statusCode = 500
		}
		http.Error(w, jsonMarshalNoError(err), statusCode)
		return
	}

	respBytes, err := json.MarshalIndent(resp, "", "\t")
	if err != nil {
		logger.FromContext(ctx).AddContext("error", err.Error())
		http.Error(w, jsonMarshalNoError(models.InternalError{Message: err.Error()}), http.StatusInternalServerError)
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(statusCodeForGetWorkflowByNameAndVersion(resp))
	w.Write(respBytes)

}

// newGetWorkflowByNameAndVersionInput takes in an http.Request an returns the input struct.
func newGetWorkflowByNameAndVersionInput(r *http.Request) (*models.GetWorkflowByNameAndVersionInput, error) {
	var input models.GetWorkflowByNameAndVersionInput

	var err error
	_ = err

	nameStr := mux.Vars(r)["name"]
	if len(nameStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	nameStrs := []string{nameStr}

	if len(nameStrs) > 0 {
		var nameTmp string
		nameStr := nameStrs[0]
		nameTmp, err = nameStr, error(nil)
		if err != nil {
			return nil, err
		}
		input.Name = nameTmp
	}

	versionStr := mux.Vars(r)["version"]
	if len(versionStr) == 0 {
		return nil, errors.New("parameter must be specified")
	}
	versionStrs := []string{versionStr}

	if len(versionStrs) > 0 {
		var versionTmp int64
		versionStr := versionStrs[0]
		versionTmp, err = swag.ConvertInt64(versionStr)
		if err != nil {
			return nil, err
		}
		input.Version = versionTmp
	}

	return &input, nil
}
