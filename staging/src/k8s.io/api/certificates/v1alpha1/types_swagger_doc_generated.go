/*
Copyright The Kubernetes Authors.

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package v1alpha1

// This file contains a collection of methods that can be used from go-restful to
// generate Swagger API documentation for its models. Please read this PR for more
// information on the implementation: https://github.com/emicklei/go-restful/pull/215
//
// TODOs are ignored from the parser (e.g. TODO(andronat):... || TODO:...) if and only if
// they are on one line! For multiple line or blocks that you want to ignore use ---.
// Any context after a --- is ignored.
//
// Those methods can be generated by using hack/update-codegen.sh

// AUTO-GENERATED FUNCTIONS START HERE. DO NOT EDIT.
var map_ClusterTrustBundle = map[string]string{
	"":         "ClusterTrustBundle is a cluster-scoped container for X.509 trust anchors (root certificates).\n\nClusterTrustBundle objects are considered to be readable by any authenticated user in the cluster, because they can be mounted by pods using the `clusterTrustBundle` projection.  All service accounts have read access to ClusterTrustBundles by default.  Users who only have namespace-level access to a cluster can read ClusterTrustBundles by impersonating a serviceaccount that they have access to.\n\nIt can be optionally associated with a particular assigner, in which case it contains one valid set of trust anchors for that signer. Signers may have multiple associated ClusterTrustBundles; each is an independent set of trust anchors for that signer. Admission control is used to enforce that only users with permissions on the signer can create or modify the corresponding bundle.",
	"metadata": "metadata contains the object metadata.",
	"spec":     "spec contains the signer (if any) and trust anchors.",
}

func (ClusterTrustBundle) SwaggerDoc() map[string]string {
	return map_ClusterTrustBundle
}

var map_ClusterTrustBundleList = map[string]string{
	"":         "ClusterTrustBundleList is a collection of ClusterTrustBundle objects",
	"metadata": "metadata contains the list metadata.",
	"items":    "items is a collection of ClusterTrustBundle objects",
}

func (ClusterTrustBundleList) SwaggerDoc() map[string]string {
	return map_ClusterTrustBundleList
}

var map_ClusterTrustBundleSpec = map[string]string{
	"":            "ClusterTrustBundleSpec contains the signer and trust anchors.",
	"signerName":  "signerName indicates the associated signer, if any.\n\nIn order to create or update a ClusterTrustBundle that sets signerName, you must have the following cluster-scoped permission: group=certificates.k8s.io resource=signers resourceName=<the signer name> verb=attest.\n\nIf signerName is not empty, then the ClusterTrustBundle object must be named with the signer name as a prefix (translating slashes to colons). For example, for the signer name `example.com/foo`, valid ClusterTrustBundle object names include `example.com:foo:abc` and `example.com:foo:v1`.\n\nIf signerName is empty, then the ClusterTrustBundle object's name must not have such a prefix.\n\nList/watch requests for ClusterTrustBundles can filter on this field using a `spec.signerName=NAME` field selector.",
	"trustBundle": "trustBundle contains the individual X.509 trust anchors for this bundle, as PEM bundle of PEM-wrapped, DER-formatted X.509 certificates.\n\nThe data must consist only of PEM certificate blocks that parse as valid X.509 certificates.  Each certificate must include a basic constraints extension with the CA bit set.  The API server will reject objects that contain duplicate certificates, or that use PEM block headers.\n\nUsers of ClusterTrustBundles, including Kubelet, are free to reorder and deduplicate certificate blocks in this file according to their own logic, as well as to drop PEM block headers and inter-block data.",
}

func (ClusterTrustBundleSpec) SwaggerDoc() map[string]string {
	return map_ClusterTrustBundleSpec
}

var map_PodCertificateRequest = map[string]string{
	"":         "PodCertificateRequest encodes a pod requesting a certificate from a given signer.\n\nKubelets use this API to implement podCertificate projected volumes",
	"metadata": "metadata contains the object metadata.",
	"spec":     "spec contains the details about the certificate being requested.",
	"status":   "status contains the issued certificate, and a standard set of conditions.",
}

func (PodCertificateRequest) SwaggerDoc() map[string]string {
	return map_PodCertificateRequest
}

var map_PodCertificateRequestList = map[string]string{
	"":         "PodCertificateRequestList is a collection of PodCertificateRequest objects",
	"metadata": "metadata contains the list metadata.",
	"items":    "items is a collection of PodCertificateRequest objects",
}

func (PodCertificateRequestList) SwaggerDoc() map[string]string {
	return map_PodCertificateRequestList
}

var map_PodCertificateRequestSpec = map[string]string{
	"":                     "PodCertificateRequestSpec describes the certificate request.  All fields are immutable after creation.",
	"signerName":           "signerName indicates the requested signer.\n\nAll signer names beginning with `kubernetes.io` are reserved for use by the Kubernetes project.  There is currently one well-known signer documented by the Kubernetes project, `kubernetes.io/kube-apiserver-client-pod`, which will issue client certificates understood by kube-apiserver.  It is currently unimplemented.",
	"podName":              "podName is the name of the pod into which the certificate will be mounted.",
	"podUID":               "podUID is the UID of the pod into which the certificate will be mounted.",
	"serviceAccountName":   "serviceAccountName is the name of the service account the pod is running as.",
	"serviceAccountUID":    "serviceAccountUID is the UID of the service account the pod is running as.",
	"nodeName":             "nodeName is the name of the node the pod is assigned to.",
	"nodeUID":              "nodeUID is the UID of the node the pod is assigned to.",
	"maxExpirationSeconds": "maxExpirationSeconds is the maximum lifetime permitted for the certificate.\n\nIf omitted, kube-apiserver will set it to 86400(24 hours). kube-apiserver will reject values shorter than 3600 (1 hour).  The maximum allowable value is 7862400 (91 days).\n\nThe signer implementation is then free to issue a certificate with any lifetime *shorter* than MaxExpirationSeconds, but no shorter than 3600 seconds (1 hour).  This constraint is enforced by kube-apiserver. `kubernetes.io` signers will never issue certificates with a lifetime longer than 24 hours.",
	"pkixPublicKey":        "pkixPublicKey is the PKIX-serialized public key the signer will issue the certificate to.\n\nThe key must be one of RSA3072, RSA4096, ECDSAP256, ECDSAP384, ECDSAP521, or ED25519. Note that this list may be expanded in the future.\n\nSigner implementations do not need to support all key types supported by kube-apiserver and kubelet.  If a signer does not support the key type used for a given PodCertificateRequest, it must deny the request by setting a status.conditions entry with a type of \"Denied\" and a reason of \"UnsupportedKeyType\". It may also suggest a key type that it does support in the message field.",
	"proofOfPossession":    "proofOfPossession proves that the requesting kubelet holds the private key corresponding to pkixPublicKey.\n\nIt is contructed by signing the ASCII bytes of the pod's UID using `pkixPublicKey`.\n\nkube-apiserver validates the proof of possession during creation of the PodCertificateRequest.\n\nIf the key is an RSA key, then the signature is over the ASCII bytes of the pod UID, using RSASSA-PSS from RFC 8017 (as implemented by the golang function crypto/rsa.SignPSS with nil options).\n\nIf the key is an ECDSA key, then the signature is as described by [SEC 1, Version 2.0](https://www.secg.org/sec1-v2.pdf) (as implemented by the golang library function crypto/ecdsa.SignASN1)\n\nIf the key is an ED25519 key, the the signature is as described by the [ED25519 Specification](https://ed25519.cr.yp.to/) (as implemented by the golang library crypto/ed25519.Sign).",
}

func (PodCertificateRequestSpec) SwaggerDoc() map[string]string {
	return map_PodCertificateRequestSpec
}

var map_PodCertificateRequestStatus = map[string]string{
	"":                 "PodCertificateRequestStatus describes the status of the request, and holds the certificate data if the request is issued.",
	"conditions":       "conditions applied to the request.\n\nThe types \"Issued\", \"Denied\", and \"Failed\" have special handling.  At most one of these conditions may be present, and they must have status \"True\".\n\nIf the request is denied with `Reason=UnsupportedKeyType`, the signer may suggest a key type that will work in the message field.",
	"certificateChain": "certificateChain is populated with an issued certificate by the signer. This field is set via the /status subresource. Once populated, this field is immutable.\n\nIf the certificate signing request is denied, a condition of type \"Denied\" is added and this field remains empty. If the signer cannot issue the certificate, a condition of type \"Failed\" is added and this field remains empty.\n\nValidation requirements:\n 1. certificateChain must consist of one or more PEM-formatted certificates.\n 2. Each entry must be a valid PEM-wrapped, DER-encoded ASN.1 Certificate as\n    described in section 4 of RFC5280.\n\nIf more than one block is present, and the definition of the requested spec.signerName does not indicate otherwise, the first block is the issued certificate, and subsequent blocks should be treated as intermediate certificates and presented in TLS handshakes.  When projecting the chain into a pod volume, kubelet will drop any data in-between the PEM blocks, as well as any PEM block headers.",
	"notBefore":        "notBefore is the time at which the certificate becomes valid.  The value must be the same as the notBefore value in the leaf certificate in certificateChain.  This field is set via the /status subresource.  Once populated, it is immutable. The signer must set this field at the same time it sets certificateChain.",
	"beginRefreshAt":   "beginRefreshAt is the time at which the kubelet should begin trying to refresh the certificate.  This field is set via the /status subresource, and must be set at the same time as certificateChain.  Once populated, this field is immutable.\n\nThis field is only a hint.  Kubelet may start refreshing before or after this time if necessary.",
	"notAfter":         "notAfter is the time at which the certificate expires.  The value must be the same as the notAfter value in the leaf certificate in certificateChain.  This field is set via the /status subresource.  Once populated, it is immutable.  The signer must set this field at the same time it sets certificateChain.",
}

func (PodCertificateRequestStatus) SwaggerDoc() map[string]string {
	return map_PodCertificateRequestStatus
}

// AUTO-GENERATED FUNCTIONS END HERE
