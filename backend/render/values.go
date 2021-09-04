package render

import (
	"fmt"

	"helm.sh/helm/v3/pkg/chartutil"
	"sigs.k8s.io/yaml"
)

type TemplateVariables struct {
	Values       map[string]interface{}
	Release      map[string]interface{}
	Capabilities map[string]interface{}
}

func getValues(rawChart RawChart) (map[string]interface{}, error) {
	var templateVariables TemplateVariables

	err := yaml.Unmarshal([]byte(rawChart.Values), &templateVariables.Values)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("cannot load values: %v", err)
	}

	err = yaml.Unmarshal([]byte(rawChart.Release), &templateVariables.Release)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("cannot load release: %v", err)
	}

	err = yaml.Unmarshal([]byte(rawChart.Capabilities), &templateVariables.Capabilities)
	if err != nil {
		return map[string]interface{}{}, fmt.Errorf("cannot load capabilities: %v", err)
	}

	setCapabilitiesAPIVersionSet(&templateVariables)

	values := map[string]interface{}{
		"Values":       templateVariables.Values,
		"Release":      templateVariables.Release,
		"Capabilities": templateVariables.Capabilities,
	}

	return values, nil
}

func setCapabilitiesAPIVersionSet(cap *TemplateVariables) {
	capabilities := cap.Capabilities
	if len(capabilities) == 0 {
		return
	}
	if capabilities["APIVersions"] == nil {
		return
	}
	apiVersions := capabilities["APIVersions"].([]interface{})
	apiVersionSet := chartutil.VersionSet{}

	for _, apiVersion := range apiVersions {
		apiVersion := apiVersion.(string)
		apiVersionSet = append(apiVersionSet, apiVersion)
	}

	cap.Capabilities["APIVersions"] = apiVersionSet
}
