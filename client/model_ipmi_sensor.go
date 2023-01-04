/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type IpmiSensor struct {
	SensorId string  `json:"sensorId,omitempty"`
	Group    string  `json:"group,omitempty"`
	Units    string  `json:"units,omitempty"`
	Status   bool    `json:"status,omitempty"`
	Name     string  `json:"name,omitempty"`
	Reading  float32 `json:"reading,omitempty"`
}
