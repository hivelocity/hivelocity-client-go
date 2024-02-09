/*
 * Hivelocity API
 *
 * Interact with Hivelocity
 *
 * API version: 2.0
 * Generated by: Swagger Codegen (https://github.com/swagger-api/swagger-codegen.git)
 */

package swagger

type DeviceInitialCreds struct {
	// Date password will expire from API.
	PasswordReturnsUntil int32 `json:"passwordReturnsUntil,omitempty"`
	// Password for initial ssh access. Will expire from api 7 days after initial provision.
	Password string `json:"password,omitempty"`
	// Port for initial ssh access.
	Port int32 `json:"port,omitempty"`
	// Link to encrypted locker containing password for initial ssh access. Locker contents be expired from api 7 days after initial provision.
	LockerUrl string `json:"lockerUrl,omitempty"`
	// User for initial ssh access.
	User string `json:"user,omitempty"`
}
