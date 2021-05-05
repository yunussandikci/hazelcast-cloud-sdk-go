package hazelcastcloud

import (
	"context"
	"encoding/json"
	"fmt"
	"github.com/yunussandikci/hazelcast-cloud-sdk-go/models"
	"github.com/stretchr/testify/assert"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
)

func TestInstanceTypeServiceOp_List(t *testing.T) {
	//given
	serveMux := http.NewServeMux()
	server := httptest.NewServer(serveMux)
	defer server.Close()

	serveMux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if m := http.MethodPost; m != r.Method {
			t.Errorf("Request method = %v, expected %v", r.Method, m)
		}
		var request GraphQLQuery
		json.NewDecoder(r.Body).Decode(&request)

		if strings.Contains(request.Query, "instanceTypes") {
			fmt.Fprint(w, `{"data":{"response":[{"name":"r5.4xlarge"},{"name":"r5.2xlarge"},{"name":"r5a.2xlarge"},{"name":"m5.large"},{"name":"m5.xlarge"},{"name":"r5a.xlarge"},{"name":"m5.4xlarge"},{"name":"r5.xlarge"},{"name":"r5a.4xlarge"},{"name":"r5.large"},{"name":"m5.2xlarge"}]}}`)
		} else {
			fmt.Fprint(w, `{"data":{"response":{"token":"token"}}}`)
		}

	})
	client, _, _ := NewFromCredentials("apiKey", "apiSecret", OptionEndpoint(server.URL))

	//when
	instanceTypes, _, _ := NewInstanceTypeService(client).List(context.TODO(), &models.InstanceTypeInput{})

	//then
	assert.Len(t, *instanceTypes, 11)
}

func ExampleInstanceTypeService_list() {
	client, _, _ := New()
	instanceTypes, _, _ := client.InstanceType.List(context.Background(), &models.InstanceTypeInput{CloudProvider: "aws"})
	fmt.Printf("Result: %#v", instanceTypes)
	//Output:Result: &[]models.InstanceType{models.InstanceType{Name:"r5.4xlarge", TotalMemory:127}, models.InstanceType{Name:"r5.2xlarge", TotalMemory:63}, models.InstanceType{Name:"r5a.2xlarge", TotalMemory:63}, models.InstanceType{Name:"m5.large", TotalMemory:7}, models.InstanceType{Name:"m5.xlarge", TotalMemory:15}, models.InstanceType{Name:"r5a.xlarge", TotalMemory:31}, models.InstanceType{Name:"m5.4xlarge", TotalMemory:63}, models.InstanceType{Name:"r5.xlarge", TotalMemory:31}, models.InstanceType{Name:"r5a.4xlarge", TotalMemory:127}, models.InstanceType{Name:"r5.large", TotalMemory:15}, models.InstanceType{Name:"m5.2xlarge", TotalMemory:31}}
}
