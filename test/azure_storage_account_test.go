package test

import(
	// "fmt"
	// "io/ioutil"
	"testing"
	// "context"
	// "github.com/Azure/azure-sdk-for-go/services/storage/mgmt/2019-06-01/storage"
	// "github.com/Azure/azure-sdk-for-go/services/resources/mgmt/2017-05-10/resources"
	// "github.com/gruntwork-io/terratest/modules/azure"
	"github.com/gruntwork-io/terratest/modules/terraform"
	"github.com/stretchr/testify/assert"
)

func TestStorageAccountTerratest(t *testing.T)  {
	t.Parallel()

	terraformOptions := &terraform.Options{
		TerraformDir: "../",
	}

	defer terraform.Destroy(t, terraformOptions)
	terraform.InitAndApply(t, terraformOptions)

	storageaccount := terraform.Output(t, terraformOptions, "storage_account_name")
	resourcegroup := terraform.Output(t, terraformOptions, "resource_group_name")
	
	expectedstorageaccount := "teratestexample1"
	expectedresourcegroup := "terratest-example-rg"

	assert.Equal(t, storageaccount, expectedstorageaccount)
	assert.Equal(t, resourcegroup, expectedresourcegroup)
	


}