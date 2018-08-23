/*
 * Grafeas API
 *
 * An API to insert and retrieve annotations on cloud artifacts.
 *
 * API version: v1alpha1
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package grafeas

type ApiPackageManagerLocation struct {
	// The cpe_uri in [cpe format](https://cpe.mitre.org/specification/) denoting the package manager version distributing a package.
	CpeUri string `json:"cpe_uri,omitempty"`
	// The version installed at this location.
	Version *VulnerabilityTypeVersion `json:"version,omitempty"`
	// The path from which we gathered that this package/version is installed.
	Path string `json:"path,omitempty"`
}
