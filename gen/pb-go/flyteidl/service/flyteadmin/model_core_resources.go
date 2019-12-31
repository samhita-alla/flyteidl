/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// A customizable interface to convey resources requested for a container. This can be interpretted differently for different container engines.
type CoreResources struct {
	// The desired set of resources requested. ResourceNames must be unique within the list.
	Requests []ResourcesResourceEntry `json:"requests,omitempty"`
	// Defines a set of bounds (e.g. min/max) within which the task can reliably run. ResourceNames must be unique within the list.
	Limits []ResourcesResourceEntry `json:"limits,omitempty"`
}
