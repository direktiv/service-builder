// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"context"
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PostHandlerFunc turns a function with the right signature into a post handler
type PostHandlerFunc func(PostParams) middleware.Responder

// Handle executing the request and returning a response
func (fn PostHandlerFunc) Handle(params PostParams) middleware.Responder {
	return fn(params)
}

// PostHandler interface for that can handle valid post params
type PostHandler interface {
	Handle(PostParams) middleware.Responder
}

// NewPost creates a new http.Handler for the post operation
func NewPost(ctx *middleware.Context, handler PostHandler) *Post {
	return &Post{Context: ctx, Handler: handler}
}

/* Post swagger:route POST / post

Post post API

*/
type Post struct {
	Context *middleware.Context
	Handler PostHandler
}

func (o *Post) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		*r = *rCtx
	}
	var Params = NewPostParams()
	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request
	o.Context.Respond(rw, r, route.Produces, route, res)

}

// PostBody post body
//
// swagger:model PostBody
type PostBody struct {

	// Amazon access key ID
	// Example: BBQJTV4BBQJTNWBBQJY5
	// Required: true
	AccessKey *string `json:"access-key"`

	// commands
	// Example: ["ec2 create-security-group --group-name secgroup1 --description mySecGroup","ec2 authorize-security-group-ingress --group-name secgroup1 --cidr 0.0.0.0/0 --protocol tcp --port 443"]
	Commands []string `json:"commands"`

	// Region in which to execute command
	// Example: us-east-1
	Region *string `json:"region,omitempty"`

	// Amazon secret access key
	// Example: ezaEezaEs0q5WezaEs0q5ezaEs0q5aEs0ezaEs0q5
	// Required: true
	SecretKey *string `json:"secret-key"`
}

// Validate validates this post body
func (o *PostBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateAccessKey(formats); err != nil {
		res = append(res, err)
	}

	if err := o.validateSecretKey(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *PostBody) validateAccessKey(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"access-key", "body", o.AccessKey); err != nil {
		return err
	}

	return nil
}

func (o *PostBody) validateSecretKey(formats strfmt.Registry) error {

	if err := validate.Required("body"+"."+"secret-key", "body", o.SecretKey); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this post body based on context it is used
func (o *PostBody) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *PostBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *PostBody) UnmarshalBinary(b []byte) error {
	var res PostBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}
