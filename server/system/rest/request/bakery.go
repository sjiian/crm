package request

// This file is auto-generated.
//
// Changes to this file may cause incorrect behavior and will be lost if
// the code is regenerated.
//
// Definitions file that controls how this file is generated:
//

import (
	"encoding/json"
	"fmt"
	"github.com/cortezaproject/corteza/server/pkg/payload"
	"github.com/go-chi/chi/v5"
	"io"
	"mime/multipart"
	"net/http"
	"strings"
)

// dummy vars to prevent
// unused imports complain
var (
	_ = chi.URLParam
	_ = multipart.ErrMessageTooLarge
	_ = payload.ParseUint64s
	_ = strings.ToLower
	_ = io.EOF
	_ = fmt.Errorf
	_ = json.NewEncoder
)

type (
	// Internal API interface
	BakeryBake struct {
		// Quantity POST parameter
		//
		// Quantity
		Quantity int

		// Code POST parameter
		//
		// Code
		Code string
	}
)

// NewBakeryBake request
func NewBakeryBake() *BakeryBake {
	return &BakeryBake{}
}

// Auditable returns all auditable/loggable parameters
func (r BakeryBake) Auditable() map[string]interface{} {
	return map[string]interface{}{
		"quantity": r.Quantity,
		"code":     r.Code,
	}
}

// Auditable returns all auditable/loggable parameters
func (r BakeryBake) GetQuantity() int {
	return r.Quantity
}

// Auditable returns all auditable/loggable parameters
func (r BakeryBake) GetCode() string {
	return r.Code
}

// Fill processes request and fills internal variables
func (r *BakeryBake) Fill(req *http.Request) (err error) {

	if strings.HasPrefix(strings.ToLower(req.Header.Get("content-type")), "application/json") {
		err = json.NewDecoder(req.Body).Decode(r)

		switch {
		case err == io.EOF:
			err = nil
		case err != nil:
			return fmt.Errorf("error parsing http request body: %w", err)
		}
	}

	{
		// Caching 32MB to memory, the rest to disk
		if err = req.ParseMultipartForm(32 << 20); err != nil && err != http.ErrNotMultipart {
			return err
		} else if err == nil {
			// Multipart params

			if val, ok := req.MultipartForm.Value["quantity"]; ok && len(val) > 0 {
				r.Quantity, err = payload.ParseInt(val[0]), nil
				if err != nil {
					return err
				}
			}

			if val, ok := req.MultipartForm.Value["code"]; ok && len(val) > 0 {
				r.Code, err = val[0], nil
				if err != nil {
					return err
				}
			}
		}
	}

	{
		if err = req.ParseForm(); err != nil {
			return err
		}

		// POST params

		if val, ok := req.Form["quantity"]; ok && len(val) > 0 {
			r.Quantity, err = payload.ParseInt(val[0]), nil
			if err != nil {
				return err
			}
		}

		if val, ok := req.Form["code"]; ok && len(val) > 0 {
			r.Code, err = val[0], nil
			if err != nil {
				return err
			}
		}
	}

	return err
}
