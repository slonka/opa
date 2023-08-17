//go:build no_oci

package download

import (
	"context"
	"github.com/open-policy-agent/opa/bundle"
	"github.com/open-policy-agent/opa/plugins/rest"
)


// NewOCI returns a new Downloader that can be started.
func NewOCI(config Config, client rest.Client, path, storePath string) *OCIDownloader {
	panic("built without OCI support")
}

// WithCallback registers a function f to be called when download updates occur.
func (d *OCIDownloader) WithCallback(f func(context.Context, Update)) *OCIDownloader {
	panic("built without OCI support")
}

// WithLogAttrs sets an optional set of key/value pair attributes to include in
// log messages emitted by the downloader.
func (d *OCIDownloader) WithLogAttrs(attrs map[string]interface{}) *OCIDownloader {
	panic("built without OCI support")
}

// WithBundleVerificationConfig sets the key configuration used to verify a signed bundle
func (d *OCIDownloader) WithBundleVerificationConfig(config *bundle.VerificationConfig) *OCIDownloader {
	panic("built without OCI support")
}

// WithSizeLimitBytes sets the file size limit for bundles read by this downloader.
func (d *OCIDownloader) WithSizeLimitBytes(n int64) *OCIDownloader {
	panic("built without OCI support")
}

// WithBundlePersistence specifies if the downloaded bundle will eventually be persisted to disk.
func (d *OCIDownloader) WithBundlePersistence(persist bool) *OCIDownloader {
	panic("built without OCI support")
}

// ClearCache is deprecated. Use SetCache instead.
func (d *OCIDownloader) ClearCache() {
	panic("built without OCI support")
}

// SetCache sets the etag value to the SHA of the loaded bundle
func (d *OCIDownloader) SetCache(etag string) {
	panic("built without OCI support")
}

// Trigger can be used to control when the downloader attempts to download
// a new bundle in manual triggering mode.
func (d *OCIDownloader) Trigger(ctx context.Context) error {
	panic("built without OCI support")
}

// Start tells the Downloader to begin downloading bundles.
func (d *OCIDownloader) Start(ctx context.Context) {
	panic("built without OCI support")
}

// Stop tells the Downloader to stop downloading bundles.
func (d *OCIDownloader) Stop(context.Context) {
	panic("built without OCI support")
}
