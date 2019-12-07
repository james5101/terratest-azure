resource "azurerm_resource_group" "terratest" {
  name     = "terratest-example-rg"
  location = "eastus"
}

resource "azurerm_storage_account" "terratest" {
  name                     = "teratestexample1"
  resource_group_name      = "${azurerm_resource_group.terratest.name}"
  location                 = "${azurerm_resource_group.terratest.location}"
  account_tier             = "Standard"
  account_replication_type = "LRS"

  tags = {
    purpose = "terratest-example"
  }
}