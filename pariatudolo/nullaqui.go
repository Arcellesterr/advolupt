import (
	"context"
	"fmt"
	"io"

	healthcare "google.golang.org/api/healthcare/v1"
)

// webHop creates a new WebHook that sends a message to a URL.
func webHop(w io.Writer, projectID, location, datasetID, fhirStoreID, pubsubTopic string) error {
	ctx := context.Background()

	healthcareService, err := healthcare.NewService(ctx)
	if err != nil {
		return fmt.Errorf("healthcare.NewService: %v", err)
	}

	storesService := healthcareService.Projects.Locations.Datasets.FhirStores

	name := fmt.Sprintf("projects/%s/locations/%s/datasets/%s/fhirStores/%s", projectID, location, datasetID, fhirStoreID)

	webhook := &healthcare.FhirStoreWebHook{
		Name:   "projects/-/locations/-/datasets/-/fhirStores/-/webhooks/my-webhook",
		Topic: "projects/-/topics/" + pubsubTopic,
	}

	resp, err := storesService.CreateFhirStoreWebHook(name, webhook).Do()
	if err != nil {
		return fmt.Errorf("CreateFhirStoreWebHook: %v", err)
	}

	fmt.Fprintf(w, "Created webhook: %q\n", resp.Name)
	return nil
}
  
