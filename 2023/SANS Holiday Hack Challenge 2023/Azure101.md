Cheat Sheet Links:

1. General: https://learn.microsoft.com/en-us/cli/azure/reference-index?view=azure-cli-latest
2. Group: https://learn.microsoft.com/en-us/cli/azure/group?view=azure-cli-latest
3. App: https://learn.microsoft.com/en-us/cli/azure/functionapp?view=azure-cli-latest
4. VM: https://learn.microsoft.com/en-us/cli/azure/vm?view=azure-cli-latest

Commands in the challege:

1. `az account show | less`
2. `az group list | less`
3. `az functionapp --resource-group northpole-rg1 | less`
4. `az vm list --resource-group northpole-rg2 --show-details | less`
5. `az vm run-command invoke --resource-group northpole-rg2 --name NP-VM1 --command-id RunShellScript --scripts "ls"`
