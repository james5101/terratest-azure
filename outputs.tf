output "storage_account_name" {
  description = "id of the security group provisioned"
  value       = "${azurerm_storage_account.terratest.name}"
}

output "resource_group_name" {
  description = "id of the security group provisioned"
  value       = "${azurerm_resource_group.terratest.name}"
}