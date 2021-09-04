package render

import (
	"fmt"
	"net/http"
	"os"

	"helm.sh/helm/v3/pkg/chart"
	"helm.sh/helm/v3/pkg/engine"
)

const templateFileName = "templates/template.yml"

func HandleRenderTemplate(res http.ResponseWriter, req *http.Request) {
	environment := os.Getenv("ENVIRONMENT")
	if environment == "dev" {
		res.Header().Set("Access-Control-Allow-Origin", "*")
	}

	renderedTemplate, err := renderTemplate(res, req)
	if err != nil {
		res.WriteHeader(http.StatusUnprocessableEntity)
		res.Write([]byte(err.Error()))
		return
	}
	res.Write([]byte(renderedTemplate))
}

func renderTemplate(res http.ResponseWriter, req *http.Request) (string, error) {
	rawChart, err := getRawChart(req)
	if err != nil {
		return "", fmt.Errorf("cannot get raw chart: %v", err)
	}

	values, err := getValues(rawChart)
	if err != nil {
		return "", fmt.Errorf("cannot get values from raw chart: %v", err)
	}

	userChart, err := getUserChart(rawChart)
	if err != nil {
		return "", fmt.Errorf("cannot get user defined chart: %v", err)
	}

	renderedTemplate, err := getRenderedTemplate(&userChart, values)
	if err != nil {
		return "", fmt.Errorf("cannot render template: %v", err)
	}

	return renderedTemplate, nil
}

func getRenderedTemplate(chart *chart.Chart, values map[string]interface{}) (string, error) {
	renderedTemplate, err := engine.Render(chart, values)
	if err != nil {
		return "", fmt.Errorf("cannot render template using engine: %v", err)
	}
	templateName := fmt.Sprintf("%s/%s", chart.Metadata.Name, templateFileName)

	return renderedTemplate[templateName], nil
}
