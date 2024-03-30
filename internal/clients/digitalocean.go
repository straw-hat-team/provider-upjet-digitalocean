/*
Copyright 2021 Upbound Inc.
*/

package clients

import (
	"context"
	"encoding/json"

	"github.com/crossplane/crossplane-runtime/pkg/resource"
	"github.com/pkg/errors"
	"k8s.io/apimachinery/pkg/types"
	"sigs.k8s.io/controller-runtime/pkg/client"

	"github.com/crossplane/upjet/pkg/terraform"

	"github.com/crossplane-contrib/provider-digitalocean/apis/v1beta1"
)

const (
	// error messages
	errNoProviderConfig     = "no providerConfigRef provided"
	errGetProviderConfig    = "cannot get referenced ProviderConfig"
	errTrackUsage           = "cannot track ProviderConfig usage"
	errExtractCredentials   = "cannot extract credentials"
	errUnmarshalCredentials = "cannot unmarshal digitalocean credentials as JSON"
)

// TerraformSetupBuilder builds Terraform a terraform.SetupFn function which
// returns Terraform provider setup configuration
func TerraformSetupBuilder(version, providerSource, providerVersion string) terraform.SetupFn {
	return func(ctx context.Context, client client.Client, mg resource.Managed) (terraform.Setup, error) {
		ps := terraform.Setup{
			Version: version,
			Requirement: terraform.ProviderRequirement{
				Source:  providerSource,
				Version: providerVersion,
			},
		}

		configRef := mg.GetProviderConfigReference()
		if configRef == nil {
			return ps, errors.New(errNoProviderConfig)
		}
		pc := &v1beta1.ProviderConfig{}
		if err := client.Get(ctx, types.NamespacedName{Name: configRef.Name}, pc); err != nil {
			return ps, errors.Wrap(err, errGetProviderConfig)
		}

		t := resource.NewProviderConfigUsageTracker(client, &v1beta1.ProviderConfigUsage{})
		if err := t.Track(ctx, mg); err != nil {
			return ps, errors.Wrap(err, errTrackUsage)
		}

		data, err := resource.CommonCredentialExtractor(ctx, pc.Spec.Credentials.Source, client, pc.Spec.Credentials.CommonCredentialSelectors)
		if err != nil {
			return ps, errors.Wrap(err, errExtractCredentials)
		}

		if err := json.Unmarshal(data, &ps.Configuration); err != nil {
			return ps, errors.Wrap(err, errUnmarshalCredentials)
		}

		//creds := map[string]any{}
		//if err := json.Unmarshal(data, &creds); err != nil {
		//	return ps, errors.Wrap(err, errUnmarshalCredentials)
		//}
		//// Set credentials in Terraform provider configuration.
		//ps.Configuration = map[string]any{
		//	"token":               creds["token"],
		//	"spaces_access_id":    creds["spaces_access_id"],
		//	"spaces_secret_key":   creds["spaces_secret_key"],
		//	"api_endpoint":        creds["api_endpoint"],
		//	"spaces_endpoint":     creds["spaces_endpoint"],
		//	"requests_per_second": creds["requests_per_second"],
		//	"http_retry_max":      creds["http_retry_max"],
		//	"http_retry_wait_min": creds["http_retry_wait_min"],
		//	"http_retry_wait_max": creds["http_retry_wait_max"],
		//}

		return ps, nil
	}
}