// Code generated by go-swagger; DO NOT EDIT.

package op

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new op API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for op API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
AddServiceAmount add service amount API
*/
func (a *Client) AddServiceAmount(params *AddServiceAmountParams) (*AddServiceAmountNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewAddServiceAmountParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "addServiceAmount",
		Method:             "POST",
		PathPattern:        "/add-service-amount",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &AddServiceAmountReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*AddServiceAmountNoContent), nil

}

/*
DelStation del station API
*/
func (a *Client) DelStation(params *DelStationParams) (*DelStationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewDelStationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "delStation",
		Method:             "POST",
		PathPattern:        "/del-station",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &DelStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*DelStationNoContent), nil

}

/*
GetPing get ping API
*/
func (a *Client) GetPing(params *GetPingParams) (*GetPingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewGetPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "getPing",
		Method:             "GET",
		PathPattern:        "/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &GetPingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*GetPingOK), nil

}

/*
Info info API
*/
func (a *Client) Info(params *InfoParams) (*InfoOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewInfoParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "info",
		Method:             "GET",
		PathPattern:        "/info",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &InfoReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*InfoOK), nil

}

/*
Kasse kasse API
*/
func (a *Client) Kasse(params *KasseParams) (*KasseOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewKasseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "kasse",
		Method:             "POST",
		PathPattern:        "/kasse",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &KasseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*KasseOK), nil

}

/*
Load load API
*/
func (a *Client) Load(params *LoadParams) (*LoadOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "load",
		Method:             "POST",
		PathPattern:        "/load",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoadReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoadOK), nil

}

/*
LoadMoney load money API
*/
func (a *Client) LoadMoney(params *LoadMoneyParams) (*LoadMoneyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadMoneyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loadMoney",
		Method:             "POST",
		PathPattern:        "/load-money",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoadMoneyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoadMoneyOK), nil

}

/*
LoadRelay load relay API
*/
func (a *Client) LoadRelay(params *LoadRelayParams) (*LoadRelayOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewLoadRelayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "loadRelay",
		Method:             "POST",
		PathPattern:        "/load-relay",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &LoadRelayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*LoadRelayOK), nil

}

/*
OpenStation open station API
*/
func (a *Client) OpenStation(params *OpenStationParams) (*OpenStationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewOpenStationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "openStation",
		Method:             "POST",
		PathPattern:        "/open-station",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &OpenStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*OpenStationNoContent), nil

}

/*
Ping ping API
*/
func (a *Client) Ping(params *PingParams) (*PingOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewPingParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "ping",
		Method:             "POST",
		PathPattern:        "/ping",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &PingReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*PingOK), nil

}

/*
ProgramRelays program relays API
*/
func (a *Client) ProgramRelays(params *ProgramRelaysParams) (*ProgramRelaysOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProgramRelaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "programRelays",
		Method:             "POST",
		PathPattern:        "/program-relays",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProgramRelaysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProgramRelaysOK), nil

}

/*
Programs programs API
*/
func (a *Client) Programs(params *ProgramsParams) (*ProgramsOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewProgramsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "programs",
		Method:             "POST",
		PathPattern:        "/programs",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &ProgramsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ProgramsOK), nil

}

/*
Save save API
*/
func (a *Client) Save(params *SaveParams) (*SaveNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "save",
		Method:             "POST",
		PathPattern:        "/save",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SaveNoContent), nil

}

/*
SaveCollection save collection API
*/
func (a *Client) SaveCollection(params *SaveCollectionParams) (*SaveCollectionNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saveCollection",
		Method:             "POST",
		PathPattern:        "/save-collection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SaveCollectionNoContent), nil

}

/*
SaveIfNotExists save if not exists API
*/
func (a *Client) SaveIfNotExists(params *SaveIfNotExistsParams) (*SaveIfNotExistsNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveIfNotExistsParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saveIfNotExists",
		Method:             "POST",
		PathPattern:        "/save-if-not-exists",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveIfNotExistsReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SaveIfNotExistsNoContent), nil

}

/*
SaveMoney save money API
*/
func (a *Client) SaveMoney(params *SaveMoneyParams) (*SaveMoneyNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveMoneyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saveMoney",
		Method:             "POST",
		PathPattern:        "/save-money",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveMoneyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SaveMoneyNoContent), nil

}

/*
SaveRelay save relay API
*/
func (a *Client) SaveRelay(params *SaveRelayParams) (*SaveRelayNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSaveRelayParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "saveRelay",
		Method:             "POST",
		PathPattern:        "/save-relay",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SaveRelayReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SaveRelayNoContent), nil

}

/*
SetKasse set kasse API
*/
func (a *Client) SetKasse(params *SetKasseParams) (*SetKasseNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetKasseParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setKasse",
		Method:             "POST",
		PathPattern:        "/set-kasse",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetKasseReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetKasseNoContent), nil

}

/*
SetProgramName set program name API
*/
func (a *Client) SetProgramName(params *SetProgramNameParams) (*SetProgramNameNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetProgramNameParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setProgramName",
		Method:             "POST",
		PathPattern:        "/set-program-name",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetProgramNameReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetProgramNameNoContent), nil

}

/*
SetProgramRelays set program relays API
*/
func (a *Client) SetProgramRelays(params *SetProgramRelaysParams) (*SetProgramRelaysNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetProgramRelaysParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setProgramRelays",
		Method:             "POST",
		PathPattern:        "/set-program-relays",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetProgramRelaysReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetProgramRelaysNoContent), nil

}

/*
SetStation set station API
*/
func (a *Client) SetStation(params *SetStationParams) (*SetStationNoContent, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewSetStationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "setStation",
		Method:             "POST",
		PathPattern:        "/set-station",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &SetStationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*SetStationNoContent), nil

}

/*
StationByHash station by hash API
*/
func (a *Client) StationByHash(params *StationByHashParams) (*StationByHashOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStationByHashParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stationByHash",
		Method:             "POST",
		PathPattern:        "/station-by-hash",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StationByHashReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StationByHashOK), nil

}

/*
StationReportCurrentMoney station report current money API
*/
func (a *Client) StationReportCurrentMoney(params *StationReportCurrentMoneyParams) (*StationReportCurrentMoneyOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStationReportCurrentMoneyParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stationReportCurrentMoney",
		Method:             "POST",
		PathPattern:        "/station-report-current-money",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StationReportCurrentMoneyReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StationReportCurrentMoneyOK), nil

}

/*
StationReportDates station report dates API
*/
func (a *Client) StationReportDates(params *StationReportDatesParams) (*StationReportDatesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStationReportDatesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stationReportDates",
		Method:             "POST",
		PathPattern:        "/station-report-dates",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StationReportDatesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StationReportDatesOK), nil

}

/*
StationsVariables stations variables API
*/
func (a *Client) StationsVariables(params *StationsVariablesParams) (*StationsVariablesOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStationsVariablesParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "stationsVariables",
		Method:             "POST",
		PathPattern:        "/stations-variables",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StationsVariablesReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StationsVariablesOK), nil

}

/*
Status status API
*/
func (a *Client) Status(params *StatusParams) (*StatusOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "status",
		Method:             "GET",
		PathPattern:        "/status",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StatusReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusOK), nil

}

/*
StatusCollection status collection API
*/
func (a *Client) StatusCollection(params *StatusCollectionParams) (*StatusCollectionOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewStatusCollectionParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "statusCollection",
		Method:             "GET",
		PathPattern:        "/status-collection",
		ProducesMediaTypes: []string{"application/json"},
		ConsumesMediaTypes: []string{"application/json"},
		Schemes:            []string{"http"},
		Params:             params,
		Reader:             &StatusCollectionReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*StatusCollectionOK), nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}