package main

import (
	"embed"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"net/http/httputil"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

//go:embed templates/*
var f embed.FS

type tool struct {
	Name     string
	Path     string
	Callback func(c *gin.Context)
}

func serve(addr string) error {

	gin.SetMode(gin.ReleaseMode)

	r := gin.New()

	r.Use(
		gin.LoggerWithFormatter(jsonLogFormatter),
		gin.Recovery(),
	)

	templ := template.Must(template.New("").ParseFS(f, "templates/*.tmpl"))
	r.SetHTMLTemplate(templ)

	tools := []tool{
		{
			Name: "Headers",
			Path: "/headers",
			Callback: func(c *gin.Context) {
				requestDump, err := httputil.DumpRequest(c.Request, true)
				if err != nil {
					fmt.Println(err)
				}
				c.Set("headers", c.Request.Header)

				c.String(200, string(requestDump))
			},
		},
		{
			Name: "Slow",
			Path: "/slow",
			Callback: func(c *gin.Context) {
				delay, err := strconv.Atoi(c.DefaultQuery("delay", "10"))
				if err != nil {
					c.String(500, "Error: delay must be a integer")
					return
				}
				time.Sleep(time.Duration(delay) * time.Second)
				c.String(200, fmt.Sprintf("Response of slow page! (%d seconds). Use ?delay=10 for adjustments.", delay))
			},
		},
		{
			Name: "Picture",
			Path: "/picture",
			Callback: func(c *gin.Context) {
				url := fmt.Sprintf(
					"https://picsum.photos/%s/%s",
					c.DefaultQuery("width", "200"),
					c.DefaultQuery("height", "300"),
				)
				response, err := http.Get(url)
				if err != nil {
					c.String(500, "Error: "+err.Error())
					return
				}
				defer response.Body.Close()
				b, err := io.ReadAll(response.Body)
				if err != nil {
					c.String(500, "Error: "+err.Error())
					return
				}

				c.Data(200, response.Header.Get("Content-Type"), b)
			},
		},
	}

	r.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl", gin.H{
			"items": tools,
		})
	})
	for _, t := range tools {
		r.GET(t.Path, t.Callback)
		r.GET(t.Path+"/*subpath", t.Callback)
	}

	jsonBytes, _ := json.Marshal(struct {
		TimeStamp string `json:"time_stamp"`
		Message   string `json:"message"`
	}{
		TimeStamp: time.Now().Format(time.RFC3339),
		Message:   "Server started at " + addr,
	})
	fmt.Println(string(jsonBytes))

	return r.Run(addr)
}

func jsonLogFormatter(params gin.LogFormatterParams) string {
	data := struct {
		TimeStamp  string      `json:"time_stamp"`
		StatusCode int         `json:"status_code"`
		Duration   string      `json:"duration"`
		ClientIP   string      `json:"client_ip"`
		Method     string      `json:"method"`
		Path       string      `json:"path"`
		Headers    interface{} `json:"headers,omitempty"`
	}{
		TimeStamp:  params.TimeStamp.Format(time.RFC3339),
		StatusCode: params.StatusCode,
		Duration:   params.Latency.String(),
		ClientIP:   params.ClientIP,
		Method:     params.Method,
		Path:       params.Path,
	}

	if headers, ok := params.Keys["headers"]; ok {
		data.Headers = headers
	}

	jsonBytes, _ := json.Marshal(data)
	return string(jsonBytes) + "\n"
}
