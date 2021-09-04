package render

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	"helm.sh/helm/v3/pkg/chart"
	"sigs.k8s.io/yaml"
)

type RawChart struct {
	Template     string `json:"template"`
	Values       string `json:"values"`
	Metadata     string `json:"chart"`
	Release      string `json:"release"`
	Capabilities string `json:"capabilities"`
	Helpers      string `json:"helpers"`
}

func getRawChart(req *http.Request) (RawChart, error) {
	body, err := ioutil.ReadAll(req.Body)
	if err != nil {
		return RawChart{}, fmt.Errorf("cannot read request body: %v", err)
	}

	var rawChart RawChart

	err = json.Unmarshal(body, &rawChart)
	if err != nil {
		return RawChart{}, fmt.Errorf("cannot parse request body json: %v", err)
	}

	return rawChart, nil
}

func getUserChart(rawChart RawChart) (chart.Chart, error) {
	templateFile := chart.File{Name: templateFileName, Data: []byte(rawChart.Template)}
	helpersFile := chart.File{Name: "templates/_helpers.tpl", Data: []byte(rawChart.Helpers)}
	var metadata chart.Metadata

	err := yaml.Unmarshal([]byte(rawChart.Metadata), &metadata)
	if err != nil {
		return chart.Chart{}, fmt.Errorf("cannot load chart metadata: %v", err)
	}

	userChart := chart.Chart{Templates: []*chart.File{&templateFile, &helpersFile}, Metadata: &metadata}

	return userChart, nil
}
