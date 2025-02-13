package main

const goIndexTemplate = `package services

import(
	{{ range $service := .services }}"github.com/micro/services/clients/go/{{ $service.Name}}"
{{ end }}
)

func NewClient(token string) *Client {
	return &Client{
		token: token,
		{{ range $service := .services }}
		{{ title $service.Name }}Service: {{ $service.Name }}.New{{ title $service.Name}}Service(token),{{end}}
	}
}

type Client struct {
	token string
{{ range $service := .services }}
	{{ title $service.Name }}Service *{{ $service.Name }}.{{ title $service.Name }}Service{{end}}
}
`

const goServiceTemplate = `{{ $service := .service }}package {{ $service.Name }}

import(
	"github.com/micro/micro-go/client"
)

func New{{ title $service.Name }}Service(token string) *{{ title $service.Name }}Service {
	return &{{ title $service.Name }}Service{
		client: client.NewClient(&client.Options{
			Token: token,
		}),
	}
}

type {{ title $service.Name }}Service struct {
	client *client.Client
}

{{ range $key, $req := $service.Spec.Components.RequestBodies }}
{{ $endpointName := requestTypeToEndpointName $key}}{{ if endpointComment $endpointName $service.Spec.Components.Schemas }}{{ endpointComment $endpointName $service.Spec.Components.Schemas }}{{ end }}func (t *{{ title $service.Name }}Service) {{ $endpointName }}(request *{{ requestType $key }}) (*{{ requestTypeToResponseType $key }}, error) {
	rsp := &{{ requestTypeToResponseType $key }}{}
	return rsp, t.client.Call("{{ $service.Name }}", "{{ requestTypeToEndpointPath $key}}", request, rsp)
}
{{ end }}


{{ range $typeName, $schema := $service.Spec.Components.Schemas }}
type {{ title $typeName }} struct {{ "{" }}
{{ recursiveTypeDefinition "go" $service.Name $typeName $service.Spec.Components.Schemas }}{{ "}" }}
{{end}}
`

const goExampleTemplate = `{{ $service := .service }}package example

import(
	"fmt"
	"os"

	"github.com/micro/services/clients/go/{{ $service.Name}}"
)

{{ if endpointComment .endpoint $service.Spec.Components.Schemas }}{{ endpointComment .endpoint $service.Spec.Components.Schemas }}{{ end }}func {{ .funcName }}() {
	{{ $service.Name }}Service := {{ $service.Name }}.New{{ title $service.Name }}Service(os.Getenv("MICRO_API_TOKEN"))
	rsp, err := {{ $service.Name }}Service.{{ title .endpoint }}(&{{ $service.Name }}.{{ title .endpoint }}Request{
		{{ goExampleRequest $service.Name .endpoint $service.Spec.Components.Schemas .example.Request }}
	})
	fmt.Println(rsp, err)
}
`

const curlExampleTemplate = `{{ $service := .service }}curl "http://localhost:8080/{{ $service.Name }}/{{ title .endpoint }}" \
-H "Content-Type: application/json" \
-H "Authorization: Bearer $MICRO_API_TOKEN" \
-d '{{ tsExampleRequest $service.Name .endpoint $service.Spec.Components.Schemas .example.Request }}'`
