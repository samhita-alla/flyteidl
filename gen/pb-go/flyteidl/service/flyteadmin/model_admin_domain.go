/*
 * flyteidl/service/admin.proto
 *
 * No description provided (generated by Swagger Codegen https://github.com/swagger-api/swagger-codegen)
 *
 * API version: version not set
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package flyteadmin

// Namespace within a project commonly used to differentiate between different service instances. e.g. \"production\", \"development\", etc.
type AdminDomain struct {
	// Globally unique domain name.
	Id string `json:"id,omitempty"`
	// Display name.
	Name string `json:"name,omitempty"`
}
