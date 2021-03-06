/*
 * DC/OS
 *
 * DC/OS API
 *
 * API version: 1.0.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package dcos

type CosmosPackageSearchDetails struct {
	Name           string                      `json:"name"`
	CurrentVersion string                      `json:"currentVersion"`
	Description    string                      `json:"description"`
	Framework      bool                        `json:"framework"`
	Selected       bool                        `json:"selected,omitempty"`
	Images         CosmosPackageResourceImages `json:"images,omitempty"`
	Tags           []string                    `json:"tags"`
	Versions       map[string]interface{}      `json:"versions"`
}
