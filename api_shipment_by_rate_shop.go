/*
Shipping APIs

The Shipping APIs include a variety of operations that allow users to manage and track their shipping requests.

API version: 1.0.0
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package sp360shipping

import (
	"bytes"
	"context"
	"io"
	"net/http"
	"net/url"
)


// ShipmentByRateShopAPIService ShipmentByRateShopAPI service
type ShipmentByRateShopAPIService service

type ApiCancelShipmentByIdV2Request struct {
	ctx context.Context
	ApiService *ShipmentByRateShopAPIService
	shipmentReprintCancelV2 *ShipmentReprintCancelV2
	xPBDeveloperPartnerId *string
	xPBLocationId *string
	xPBTransactionId *string
}

func (r ApiCancelShipmentByIdV2Request) ShipmentReprintCancelV2(shipmentReprintCancelV2 ShipmentReprintCancelV2) ApiCancelShipmentByIdV2Request {
	r.shipmentReprintCancelV2 = &shipmentReprintCancelV2
	return r
}

// This is the Developer Partner ID assigned by PB, which is used in API for communication with strategic business partners.&lt;br /&gt; When the developer is the only business partner, this field is not required.
func (r ApiCancelShipmentByIdV2Request) XPBDeveloperPartnerId(xPBDeveloperPartnerId string) ApiCancelShipmentByIdV2Request {
	r.xPBDeveloperPartnerId = &xPBDeveloperPartnerId
	return r
}

// This is the Location ID assigned as per the Developer Partner&#39;s location.
func (r ApiCancelShipmentByIdV2Request) XPBLocationId(xPBLocationId string) ApiCancelShipmentByIdV2Request {
	r.xPBLocationId = &xPBLocationId
	return r
}

// A unique Transaction ID provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system.
func (r ApiCancelShipmentByIdV2Request) XPBTransactionId(xPBTransactionId string) ApiCancelShipmentByIdV2Request {
	r.xPBTransactionId = &xPBTransactionId
	return r
}

func (r ApiCancelShipmentByIdV2Request) Execute() (*CancelShipmentV2, *http.Response, error) {
	return r.ApiService.CancelShipmentByIdV2Execute(r)
}

/*
CancelShipmentByIdV2 Cancel Shipment

The operation cancel/void shipment.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCancelShipmentByIdV2Request
*/
func (a *ShipmentByRateShopAPIService) CancelShipmentByIdV2(ctx context.Context) ApiCancelShipmentByIdV2Request {
	return ApiCancelShipmentByIdV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return CancelShipmentV2
func (a *ShipmentByRateShopAPIService) CancelShipmentByIdV2Execute(r ApiCancelShipmentByIdV2Request) (*CancelShipmentV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *CancelShipmentV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentByRateShopAPIService.CancelShipmentByIdV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/shipments/cancel"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shipmentReprintCancelV2 == nil {
		return localVarReturnValue, nil, reportError("shipmentReprintCancelV2 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPBDeveloperPartnerId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-Developer-Partner-Id", r.xPBDeveloperPartnerId, "")
	}
	if r.xPBLocationId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-LocationId", r.xPBLocationId, "")
	}
	if r.xPBTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-TransactionId", r.xPBTransactionId, "")
	}
	// body params
	localVarPostBody = r.shipmentReprintCancelV2
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v []ErrorsInner
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v []NotFoundErrorsInner
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiCreateShipmentV2Request struct {
	ctx context.Context
	ApiService *ShipmentByRateShopAPIService
	createShipmentV2Request *CreateShipmentV2Request
	xPBDeveloperPartnerId *string
	xPBLocationId *string
	xPBTransactionId *string
}

func (r ApiCreateShipmentV2Request) CreateShipmentV2Request(createShipmentV2Request CreateShipmentV2Request) ApiCreateShipmentV2Request {
	r.createShipmentV2Request = &createShipmentV2Request
	return r
}

// The Developer Partner ID assigned by PB, which is used in API for the communication with the strategic business partners. &lt;br /&gt; When the developer is the only business partner, this field is not required.
func (r ApiCreateShipmentV2Request) XPBDeveloperPartnerId(xPBDeveloperPartnerId string) ApiCreateShipmentV2Request {
	r.xPBDeveloperPartnerId = &xPBDeveloperPartnerId
	return r
}

// This is the Location ID assigned as per the Developer Partner&#39;s location.
func (r ApiCreateShipmentV2Request) XPBLocationId(xPBLocationId string) ApiCreateShipmentV2Request {
	r.xPBLocationId = &xPBLocationId
	return r
}

// A unique Transaction ID provided by the partner, which is used to enable debugging and linking between the client&#39;s transaction and the system.
func (r ApiCreateShipmentV2Request) XPBTransactionId(xPBTransactionId string) ApiCreateShipmentV2Request {
	r.xPBTransactionId = &xPBTransactionId
	return r
}

func (r ApiCreateShipmentV2Request) Execute() (*DomesticShipmentResponseV2, *http.Response, error) {
	return r.ApiService.CreateShipmentV2Execute(r)
}

/*
CreateShipmentV2 Create Shipment

The operation creates a new Shipment or generate a Shipment Label. Here, Shipment refers to process where an item is packed and shipped from one point (source) to the other (destination) using some carrier service.
 The shipment can be done for within domestic places or for overseas/International places. <br> To create a domestic shipment, the operation requires 'To' and 'From' locations or addresses respectively within the same country,  carrier services, and associated special services. <br> While for the International shipment, the operation requires international address(es) for delivery, which is 'To' address must be the international country location(s) (and not the same country mentioned in 'From' location or address), supported international carrier services, associated special service(s), and customs information.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiCreateShipmentV2Request
*/
func (a *ShipmentByRateShopAPIService) CreateShipmentV2(ctx context.Context) ApiCreateShipmentV2Request {
	return ApiCreateShipmentV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return DomesticShipmentResponseV2
func (a *ShipmentByRateShopAPIService) CreateShipmentV2Execute(r ApiCreateShipmentV2Request) (*DomesticShipmentResponseV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *DomesticShipmentResponseV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentByRateShopAPIService.CreateShipmentV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/shipments"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.createShipmentV2Request == nil {
		return localVarReturnValue, nil, reportError("createShipmentV2Request is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPBDeveloperPartnerId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-Developer-Partner-Id", r.xPBDeveloperPartnerId, "")
	}
	if r.xPBLocationId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-LocationId", r.xPBLocationId, "")
	}
	if r.xPBTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-TransactionId", r.xPBTransactionId, "")
	}
	// body params
	localVarPostBody = r.createShipmentV2Request
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v []ErrorsInner
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}

type ApiReprintShipmentByIdV2Request struct {
	ctx context.Context
	ApiService *ShipmentByRateShopAPIService
	shipmentReprintCancelV2 *ShipmentReprintCancelV2
	xPBDeveloperPartnerId *string
	xPBLocationId *string
	xPBTransactionId *string
}

func (r ApiReprintShipmentByIdV2Request) ShipmentReprintCancelV2(shipmentReprintCancelV2 ShipmentReprintCancelV2) ApiReprintShipmentByIdV2Request {
	r.shipmentReprintCancelV2 = &shipmentReprintCancelV2
	return r
}

// This is the Developer Partner ID assigned by PB, which is used in API for communication with strategic business partners. When the developer is the only business partner, this field is not required.
func (r ApiReprintShipmentByIdV2Request) XPBDeveloperPartnerId(xPBDeveloperPartnerId string) ApiReprintShipmentByIdV2Request {
	r.xPBDeveloperPartnerId = &xPBDeveloperPartnerId
	return r
}

// This is the Location ID assigned as per the Developer Partner&#39;s location.
func (r ApiReprintShipmentByIdV2Request) XPBLocationId(xPBLocationId string) ApiReprintShipmentByIdV2Request {
	r.xPBLocationId = &xPBLocationId
	return r
}

// A unique transaction Id provided by the partner which is used to enable debugging and linking between the client&#39;s transaction and the system.
func (r ApiReprintShipmentByIdV2Request) XPBTransactionId(xPBTransactionId string) ApiReprintShipmentByIdV2Request {
	r.xPBTransactionId = &xPBTransactionId
	return r
}

func (r ApiReprintShipmentByIdV2Request) Execute() (*ReprintShipmentV2, *http.Response, error) {
	return r.ApiService.ReprintShipmentByIdV2Execute(r)
}

/*
ReprintShipmentByIdV2 Reprint Shipment

The operation reprints Shipment by the shipmentId. It retrieves an existing shipping label to reprint. The API sends the shipmentID returned by the original created shipment request. Use this only if the shipping label in the Create Shipment response is missing or lost.

 @param ctx context.Context - for authentication, logging, cancellation, deadlines, tracing, etc. Passed from http.Request or context.Background().
 @return ApiReprintShipmentByIdV2Request
*/
func (a *ShipmentByRateShopAPIService) ReprintShipmentByIdV2(ctx context.Context) ApiReprintShipmentByIdV2Request {
	return ApiReprintShipmentByIdV2Request{
		ApiService: a,
		ctx: ctx,
	}
}

// Execute executes the request
//  @return ReprintShipmentV2
func (a *ShipmentByRateShopAPIService) ReprintShipmentByIdV2Execute(r ApiReprintShipmentByIdV2Request) (*ReprintShipmentV2, *http.Response, error) {
	var (
		localVarHTTPMethod   = http.MethodPost
		localVarPostBody     interface{}
		formFiles            []formFile
		localVarReturnValue  *ReprintShipmentV2
	)

	localBasePath, err := a.client.cfg.ServerURLWithContext(r.ctx, "ShipmentByRateShopAPIService.ReprintShipmentByIdV2")
	if err != nil {
		return localVarReturnValue, nil, &GenericOpenAPIError{error: err.Error()}
	}

	localVarPath := localBasePath + "/api/v2/shipments/reprint"

	localVarHeaderParams := make(map[string]string)
	localVarQueryParams := url.Values{}
	localVarFormParams := url.Values{}
	if r.shipmentReprintCancelV2 == nil {
		return localVarReturnValue, nil, reportError("shipmentReprintCancelV2 is required and must be specified")
	}

	// to determine the Content-Type header
	localVarHTTPContentTypes := []string{"application/json"}

	// set Content-Type header
	localVarHTTPContentType := selectHeaderContentType(localVarHTTPContentTypes)
	if localVarHTTPContentType != "" {
		localVarHeaderParams["Content-Type"] = localVarHTTPContentType
	}

	// to determine the Accept header
	localVarHTTPHeaderAccepts := []string{"application/json"}

	// set Accept header
	localVarHTTPHeaderAccept := selectHeaderAccept(localVarHTTPHeaderAccepts)
	if localVarHTTPHeaderAccept != "" {
		localVarHeaderParams["Accept"] = localVarHTTPHeaderAccept
	}
	if r.xPBDeveloperPartnerId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-Developer-Partner-Id", r.xPBDeveloperPartnerId, "")
	}
	if r.xPBLocationId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-LocationId", r.xPBLocationId, "")
	}
	if r.xPBTransactionId != nil {
		parameterAddToHeaderOrQuery(localVarHeaderParams, "X-PB-TransactionId", r.xPBTransactionId, "")
	}
	// body params
	localVarPostBody = r.shipmentReprintCancelV2
	req, err := a.client.prepareRequest(r.ctx, localVarPath, localVarHTTPMethod, localVarPostBody, localVarHeaderParams, localVarQueryParams, localVarFormParams, formFiles)
	if err != nil {
		return localVarReturnValue, nil, err
	}

	localVarHTTPResponse, err := a.client.callAPI(req)
	if err != nil || localVarHTTPResponse == nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	localVarBody, err := io.ReadAll(localVarHTTPResponse.Body)
	localVarHTTPResponse.Body.Close()
	localVarHTTPResponse.Body = io.NopCloser(bytes.NewBuffer(localVarBody))
	if err != nil {
		return localVarReturnValue, localVarHTTPResponse, err
	}

	if localVarHTTPResponse.StatusCode >= 300 {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: localVarHTTPResponse.Status,
		}
		if localVarHTTPResponse.StatusCode == 400 {
			var v []ErrorsInner
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 401 {
			var v ModelError
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		if localVarHTTPResponse.StatusCode == 404 {
			var v []NotFoundErrorsInner
			err = a.client.decode(&v, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
			if err != nil {
				newErr.error = err.Error()
				return localVarReturnValue, localVarHTTPResponse, newErr
			}
					newErr.error = formatErrorMessage(localVarHTTPResponse.Status, &v)
					newErr.model = v
			return localVarReturnValue, localVarHTTPResponse, newErr
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	err = a.client.decode(&localVarReturnValue, localVarBody, localVarHTTPResponse.Header.Get("Content-Type"))
	if err != nil {
		newErr := &GenericOpenAPIError{
			body:  localVarBody,
			error: err.Error(),
		}
		return localVarReturnValue, localVarHTTPResponse, newErr
	}

	return localVarReturnValue, localVarHTTPResponse, nil
}
