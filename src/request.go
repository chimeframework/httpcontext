package httpcontext

type Request struct {
    attributes map[string]string
}

func NewRequest() *Request {
    this := &Request{}
    this.attributes = make(map[string]string)
    return this
}

func (this *Request) GetAttributes() map[string]string{
        return this.attributes
}
