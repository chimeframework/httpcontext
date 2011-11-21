package httpcontext

import (
	_ "errors"
	"net/http"
)

type Request struct {
	*http.Request
	attributes map[string]string
}

func NewRequest(request *http.Request) *Request {
	this := &Request{}
	this.Request = request
	this.attributes = make(map[string]string)
	return this
}

func (this *Request) GetAttributes() map[string]string {
	return this.attributes
}

func (this *Request) HasAttribute(name string) bool {
	_, ok := this.attributes[name]
	return ok
}

func (this *Request) AddAttributes(attributes map[string]string) {
	for key, value := range attributes {
		this.attributes[key] = value
	}
}

func (this *Request) GetAttribute(name string) (attribute string, ok bool) {
	attribute, ok = this.attributes[name]
	return
}
