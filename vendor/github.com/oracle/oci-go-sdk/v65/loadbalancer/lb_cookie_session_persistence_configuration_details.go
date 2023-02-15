// Copyright (c) 2016, 2018, 2022, Oracle and/or its affiliates.  All rights reserved.
// This software is dual-licensed to you under the Universal Permissive License (UPL) 1.0 as shown at https://oss.oracle.com/licenses/upl or Apache License 2.0 as shown at http://www.apache.org/licenses/LICENSE-2.0. You may choose either license.
// Code generated. DO NOT EDIT.

// Load Balancing API
//
// API for the Load Balancing service. Use this API to manage load balancers, backend sets, and related items. For more
// information, see Overview of Load Balancing (https://docs.cloud.oracle.com/iaas/Content/Balance/Concepts/balanceoverview.htm).
//

package loadbalancer

import (
	"fmt"
	"github.com/oracle/oci-go-sdk/v65/common"
	"strings"
)

// LbCookieSessionPersistenceConfigurationDetails The configuration details for implementing load balancer cookie session persistence (LB cookie stickiness).
// Session persistence enables the Load Balancing service to direct all requests that originate from a single logical
// client to a single backend web server. For more information, see
// Session Persistence (https://docs.cloud.oracle.com/Content/Balance/Reference/sessionpersistence.htm).
// When you configure LB cookie stickiness, the load balancer inserts a cookie into the response. The parameters configured
// in the cookie enable session stickiness. This method is useful when you have applications and Web backend services
// that cannot generate their own cookies.
// Path route rules take precedence to determine the target backend server. The load balancer verifies that session stickiness
// is enabled for the backend server and that the cookie configuration (domain, path, and cookie hash) is valid for the
// target. The system ignores invalid cookies.
// To disable LB cookie stickiness on a running load balancer, use the
// UpdateBackendSet operation and specify `null` for the
// `LBCookieSessionPersistenceConfigurationDetails` object.
// Example: `LBCookieSessionPersistenceConfigurationDetails: null`
// **Note:** `SessionPersistenceConfigurationDetails` (application cookie stickiness) and `LBCookieSessionPersistenceConfigurationDetails`
// (LB cookie stickiness) are mutually exclusive. An error results if you try to enable both types of session persistence.
// **Warning:** Oracle recommends that you avoid using any confidential information when you supply string values using the API.
type LbCookieSessionPersistenceConfigurationDetails struct {

	// The name of the cookie inserted by the load balancer. If this field is not configured, the cookie name defaults
	// to "X-Oracle-BMC-LBS-Route".
	// Example: `example_cookie`
	// **Notes:**
	// *  Ensure that the cookie name used at the backend application servers is different from the cookie name used
	//    at the load balancer. To minimize the chance of name collision, Oracle recommends that you use a prefix
	//    such as "X-Oracle-OCI-" for this field.
	// *  If a backend server and the load balancer both insert cookies with the same name, the client or browser
	//    behavior can vary depending on the domain and path values associated with the cookie. If the name, domain,
	//    and path values of the `Set-cookie` generated by a backend server and the `Set-cookie` generated by the
	//    load balancer are all the same, the client or browser treats them as one cookie and returns only one of
	//    the cookie values in subsequent requests. If both `Set-cookie` names are the same, but the domain and path
	//    names are different, the client or browser treats them as two different cookies.
	CookieName *string `mandatory:"false" json:"cookieName"`

	// Whether the load balancer is prevented from directing traffic from a persistent session client to
	// a different backend server if the original server is unavailable. Defaults to false.
	// Example: `false`
	DisableFallback *bool `mandatory:"false" json:"disableFallback"`

	// The domain in which the cookie is valid. The `Set-cookie` header inserted by the load balancer contains a
	// domain attribute with the specified value.
	// This attribute has no default value. If you do not specify a value, the load balancer does not insert the domain
	// attribute into the `Set-cookie` header.
	// **Notes:**
	// *  RFC 6265 - HTTP State Management Mechanism (https://www.ietf.org/rfc/rfc6265.txt) describes client and
	//    browser behavior when the domain attribute is present or not present in the `Set-cookie` header.
	//    If the value of the `Domain` attribute is `example.com` in the `Set-cookie` header, the client includes
	//    the same cookie in the `Cookie` header when making HTTP requests to `example.com`, `www.example.com`, and
	//    `www.abc.example.com`. If the `Domain` attribute is not present, the client returns the cookie only for
	//    the domain to which the original request was made.
	// *  Ensure that this attribute specifies the correct domain value. If the `Domain` attribute in the `Set-cookie`
	//    header does not include the domain to which the original request was made, the client or browser might reject
	//    the cookie. As specified in RFC 6265, the client accepts a cookie with the `Domain` attribute value `example.com`
	//    or `www.example.com` sent from `www.example.com`. It does not accept a cookie with the `Domain` attribute
	//    `abc.example.com` or `www.abc.example.com` sent from `www.example.com`.
	// Example: `example.com`
	Domain *string `mandatory:"false" json:"domain"`

	// The path in which the cookie is valid. The `Set-cookie header` inserted by the load balancer contains a `Path`
	// attribute with the specified value.
	// Clients include the cookie in an HTTP request only if the path portion of the request-uri matches, or is a
	// subdirectory of, the cookie's `Path` attribute.
	// The default value is `/`.
	// Example: `/example`
	Path *string `mandatory:"false" json:"path"`

	// The amount of time the cookie remains valid. The `Set-cookie` header inserted by the load balancer contains
	// a `Max-Age` attribute with the specified value.
	// The specified value must be at least one second. There is no default value for this attribute. If you do not
	// specify a value, the load balancer does not include the `Max-Age` attribute in the `Set-cookie` header. In
	// most cases, the client or browser retains the cookie until the current session ends, as defined by the client.
	// Example: `3600`
	MaxAgeInSeconds *int `mandatory:"false" json:"maxAgeInSeconds"`

	// Whether the `Set-cookie` header should contain the `Secure` attribute. If `true`, the `Set-cookie` header
	// inserted by the load balancer contains the `Secure` attribute, which directs the client or browser to send the
	// cookie only using a secure protocol.
	// **Note:** If you set this field to `true`, you cannot associate the corresponding backend set with an HTTP
	// listener.
	// Example: `true`
	IsSecure *bool `mandatory:"false" json:"isSecure"`

	// Whether the `Set-cookie` header should contain the `HttpOnly` attribute. If `true`, the `Set-cookie` header
	// inserted by the load balancer contains the `HttpOnly` attribute, which limits the scope of the cookie to HTTP
	// requests. This attribute directs the client or browser to omit the cookie when providing access to cookies
	// through non-HTTP APIs. For example, it restricts the cookie from JavaScript channels.
	// Example: `true`
	IsHttpOnly *bool `mandatory:"false" json:"isHttpOnly"`
}

func (m LbCookieSessionPersistenceConfigurationDetails) String() string {
	return common.PointerString(m)
}

// ValidateEnumValue returns an error when providing an unsupported enum value
// This function is being called during constructing API request process
// Not recommended for calling this function directly
func (m LbCookieSessionPersistenceConfigurationDetails) ValidateEnumValue() (bool, error) {
	errMessage := []string{}

	if len(errMessage) > 0 {
		return true, fmt.Errorf(strings.Join(errMessage, "\n"))
	}
	return false, nil
}
