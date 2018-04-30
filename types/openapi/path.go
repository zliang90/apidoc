// Copyright 2018 by caixw, All rights reserved.
// Use of this source code is governed by a MIT
// license that can be found in the LICENSE file.

package openapi

// PathItem 每一条路径的详细描述信息
type PathItem struct {
	Ref         string       `json:"ref,omitempty" yaml:"ref,omitempty"`
	Summary     string       `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description string       `json:"description,omitempty" yaml:"description,omitempty"`
	Get         *Operation   `json:"get,omitempty" yaml:"get,omitempty"`
	Put         *Operation   `json:"put,omitempty" yaml:"put,omitempty"`
	Post        *Operation   `json:"post,omitempty" yaml:"post,omitempty"`
	Delete      *Operation   `json:"delete,omitempty" yaml:"delete,omitempty"`
	Options     *Operation   `json:"options,omitempty" yaml:"options,omitempty"`
	Head        *Operation   `json:"head,omitempty" yaml:"head,omitempty"`
	Patch       *Operation   `json:"patch,omitempty" yaml:"patch,omitempty"`
	Trace       *Operation   `json:"trace,omitempty" yaml:"trace,omitempty"`
	Servers     []*Server    `json:"servers,omitempty" yaml:"servers,omitempty"`
	Parameters  []*Parameter `json:"parameters,omitempty" yaml:"parameters,omitempty"`
}

// Operation 描述对某一个资源的操作具体操作
type Operation struct {
	Tags         []string               `json:"tags,omitempty" yaml:"tags,omitempty"`
	Summary      string                 `json:"summary,omitempty" yaml:"summary,omitempty"`
	Description  string                 `json:"description,omitempty" yaml:"description,omitempty"`
	ExternalDocs *ExternalDocumentation `json:"externalDocs,omitempty" yaml:"externalDocs,omitempty"`
	OperationID  string                 `json:"operationId,omitempty" yaml:"operationId,omitempty" `
	Parameters   []*Parameter           `json:"parameters,omitempty" yaml:"parameters,omitempty"`
	RequestBody  *RequestBody           `json:"requestBody,omitempty" yaml:"requestBody,omitempty"`
	Responses    map[string]*Response   `json:"responses" yaml:"responses"`
	Callbacks    map[string]*Callback   `json:"callbacks,omitempty" yaml:"callbacks,omitempty"`
	Deprecated   bool                   `json:"deprecated,omitempty" yaml:"deprecated,omitempty"`
	Security     []*SecurityRequirement `json:"security,omitempty" yaml:"security,omitempty"`
	Servers      []*Server              `json:"servers,omitempty" yaml:"servers,omitempty"`
}

// RequestBody 请求内容
type RequestBody struct {
	Description string                `json:"description,omitempty" yaml:"description,omitempty"`
	Content     map[string]*MediaType `json:"content" yaml:"content"`
	Required    bool                  `json:"required,omitempty" yaml:"required,omitempty" `
}

// MediaType 媒体类型
type MediaType struct {
	Schema   *Schema              `json:"schema,omitempty" yaml:"schema,omitempty"`
	Example  ExampleValue         `json:"example,omitempty" yaml:"example,omitempty"`
	Examples map[string]*Example  `json:"examples,omitempty" yaml:"examples,omitempty"`
	Encoding map[string]*Encoding `json:"encoding,omitempty" yaml:"encoding,omitempty"`
}

// Encoding 定义编码
type Encoding struct {
	Style
	ContentType string             `json:"contentType,omitempty" yaml:"contentType,omitempty"`
	Headers     map[string]*Header `json:"headers,omitempty" yaml:"headers,omitempty"`
}

// Callback Object
//
// NOTE: 暂时未用到
type Callback PathItem

// Response 每个 API 的返回信息
type Response struct {
	Description string                `json:"description" yaml:"description"`
	Headers     map[string]*Header    `json:"headers,omitempty" yaml:"headers,omitempty"`
	Content     map[string]*MediaType `json:"content,omitempty" yaml:"content,omitempty"`
	Links       map[string]*Link      `json:"links,omitempty" yaml:"links,omitempty"`
}

// Sanitize 数据检测
func (req *RequestBody) Sanitize() *Error {
	if len(req.Content) == 0 {
		return newError("content", "不能为空")
	}

	for key, mt := range req.Content {
		if err := mt.Sanitize(); err != nil {
			err.Field = "content[" + key + "]." + err.Field
			return err
		}
	}

	return nil
}

// Sanitize 数据检测
func (resp *Response) Sanitize() *Error {
	if resp.Description == "" {
		return newError("description", "不能为空")
	}

	for key, header := range resp.Headers {
		if err := header.Sanitize(); err != nil {
			err.Field = "headers[" + key + "]." + err.Field
			return err
		}
	}

	for key, mt := range resp.Content {
		if err := mt.Sanitize(); err != nil {
			err.Field = "content[" + key + "]." + err.Field
			return err
		}
	}

	for key, link := range resp.Links {
		if err := link.Sanitize(); err != nil {
			err.Field = "links[" + key + "]." + err.Field
			return err
		}
	}

	return nil
}

// Sanitize 数据检测
func (mt *MediaType) Sanitize() *Error {
	if mt.Schema != nil {
		if err := mt.Sanitize(); err != nil {
			err.Field = "schema." + err.Field
			return err
		}
	}

	for key, en := range mt.Encoding {
		if err := en.Sanitize(); err != nil {
			err.Field = "encoding[" + key + "]." + err.Field
			return err
		}
	}
	return nil
}

// Sanitize 数据检测
func (en *Encoding) Sanitize() *Error {
	if err := en.Style.Sanitize(); err != nil {
		return err
	}

	for key, header := range en.Headers {
		if err := header.Sanitize(); err != nil {
			err.Field = "headers[" + key + "]." + err.Field
			return err
		}
	}

	return nil
}