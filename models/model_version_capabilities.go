package models

// VersionCapabilities - Additional capabilities the printer has. The object is expected to be extended in the future with more capabilities. The absence of a capability in the object, or the complete absence of the object means the printer doesn't support such capability (probably doesn't even know such capability might exist). 
type VersionCapabilities struct {

	// The printer supports uploading GCodes by the PUT method (as described in this document). It is capable of doing the PUT and HEAD to /api/v1/files/{storage}/{path} and it is capable of answering the /api/v1/storage endpoint.  In absence of this capability, client MAY opt to try the legacy \"octoprint\" POST to /api/files/{storage}. 
	UploadByPut bool `json:"upload-by-put,omitempty"`
}
