{{$svrType := .ServiceType}}
{{$svrName := .ServiceName}}
type {{.ServiceType}}HTTPServer interface {
{{- range .MethodSets}}
	{{.Name}}(context.Context, *{{.Request}}) (*{{.Reply}}, error)
{{- end}}
}

type {{$svrType}} struct{
	server {{.ServiceType}}HTTPServer
	router gee.IRoutes
	middlewares []gee.Handler
}

func (r *{{$svrType}}) RegisterService(){
    {{- range .Methods}}
    r.router.{{.Method}}("{{.Path}}", append(r.middlewares, r.{{.Name}}{{if not (eq .Num 0)}}{{.Num}}{{end}})...)
    {{- end}}
}

func (r *{{$svrType}}) Validate(in interface{}) error {
	if v, ok := in.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func Register{{.ServiceType}}HTTPServer(r gee.IRoutes, srv {{.ServiceType}}HTTPServer, middlewares ...gee.Handler) {
    s := &{{$svrType}}{
        server: srv,
        router: r,
        middlewares: middlewares,
    }
    s.RegisterService()
}

{{range .Methods}}
func (r *{{$svrType}}) {{.Name}}{{if not (eq .Num 0)}}{{.Num}}{{end}}(c *gee.Context) error {
    var in {{.Request}}
    {{- if .HasBody}}
    {{- if not (eq .Request "emptypb.Empty")}}
    if err := c.ShouldBindJSON(&in{{.Body}}); err != nil {
        return errors.BadRequest("InvalidParameter", err.Error())
    }
    {{- end}}
    {{- if not (eq .Body "")}}
    if err := c.ShouldBindQuery(&in); err != nil {
        return errors.BadRequest("InvalidParameter", err.Error())
    }
    {{- end}}
    {{- else}}
    if err := c.ShouldBindQuery(&in{{.Body}}); err != nil {
        return errors.BadRequest("InvalidParameter", err.Error())
    }
    {{- end}}
    {{- if .HasVars}}
    if err := c.ShouldBindUri(&in); err != nil {
        return errors.BadRequest("InvalidParameter", err.Error())
    }
    {{- end}}

    if err := r.Validate(&in); err != nil {
        return errors.BadRequest("InvalidParameter", err.Error())
    }

    out, err := r.server.{{.Name}}(c.Request.Context(), &in)
    if err != nil {
        return err
    }
    return c.JSON(out)
}
{{end}}


type {{.ServiceType}}HTTPClient interface {
{{- range .MethodSets}}
	{{.Name}}(ctx context.Context, req *{{.Request}}) (rsp *{{.Reply}}, err error)
{{- end}}
}

type {{.ServiceType}}HTTPClientImpl struct{
	cc       *gee.Client
}

func (c *{{$svrType}}HTTPClientImpl) Validate(in interface{}) error {
	if v, ok := in.(interface{ Validate() error }); ok {
		if err := v.Validate(); err != nil {
			return err
		}
	}
	return nil
}

func New{{.ServiceType}}HTTPClient (httpClient *http.Client, endpoint string, options ...gee.ClientOption) {{.ServiceType}}HTTPClient {
	client := gee.NewClient(httpClient, endpoint, options...)

	return &{{.ServiceType}}HTTPClientImpl{
		cc: client,
	}
}

{{range .MethodSets}}
func (c *{{$svrType}}HTTPClientImpl) {{.Name}}(ctx context.Context, in *{{.Request}}) (*{{.Reply}}, error) {
	var out {{.Reply}}
	path := "{{.Path}}"

    if err := c.Validate(&in); err != nil {
        return nil, errors.BadRequest("InvalidParameter", err.Error())
    }

	err := c.cc.Invoke(ctx, "{{.Method}}", path, in{{.Body}}, &out{{.ResponseBody}})
	if err != nil {
		return nil, err
	}
	return &out, nil
}
{{end}}
