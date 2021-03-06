package GoGym

import (
	"net/http"
	"net/url"
	"reflect"
)

const (
	GETMethod     = "GET"
	POSTMethod    = "POST"
	PUTMethod     = "PUT"
	PATCHMethod   = "PATCH"
	DELETEMethod  = "DELETE"
	OPTIONSMethod = "OPTIONS"
)

const (
	ServiceRequest = "Request"
)

// Request service
type Request struct {
	app *Gym // Service Container

	Method     string
	Header     http.Header
	Query      url.Values
	Form       url.Values
	RouteParam RouteParam
}

// Prepare is a method prepares the Request service
func (r *Request) Prepare(g *Gym) {
	r.InjectServiceContainer(g)
	r.RouteParam.variables = make(map[string]string)
}

// WhoIsYourBoss is a method sets the service container into the Request
func (r *Request) InjectServiceContainer(g *Gym) {
	r.app = g
}

// CallYourBoss is a method gets the service container
func (r *Request) GetServiceContainer() *Gym {
	return r.app
}

func (r *Request) CallMethod(method string, param []interface{}) []reflect.Value {
	return nil
}

// Accept is a method gets the http request and parse it
func (r *Request) accept(request *http.Request) {
	request.ParseForm()
	r.Method = request.Method
	r.Query = request.Form
	r.Form = request.PostForm
	r.Header = request.Header
}

// BindPathVar is a method binding values to related variable in the uri
func (r *Request) bindPathVar(tokens []Token) {
	for _, v := range tokens {
		if v.IsParam {
			//r.PathVar[v.Name] = v.Value
			r.RouteParam.Set(v.Name, v.Value)
		}
	}
}

// PathVar stores all variables of defined path
type RouteParam struct {
	variables map[string]string
}

// Get is a PathVar getter
func (r *RouteParam) Get(varName string) string {
	value, isSet := r.variables[varName]
	if !isSet {
		return ""
	}
	return value
}

// Set is a PathVar setter
func (r *RouteParam) Set(varName, value string) {
	r.variables[varName] = value
}

// All is a method for getting all path variables
func (r *RouteParam) All() map[string]string {
	return r.variables
}
