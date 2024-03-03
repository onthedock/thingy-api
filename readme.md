# Thingy API

This is a *dummy* API for building a test Terraform Provider.

## Install Terraform

From the [oficial documentation](https://developer.hashicorp.com/terraform/install), we install Terraform using the package manager of our Linux distribution:

```console
wget -O- https://apt.releases.hashicorp.com/gpg | sudo gpg --dearmor -o /usr/share/keyrings/hashicorp-archive-keyring.gpg
echo "deb [signed-by=/usr/share/keyrings/hashicorp-archive-keyring.gpg] https://apt.releases.hashicorp.com $(lsb_release -cs) main" | sudo tee /etc/apt/sources.list.d/hashicorp.list
sudo apt update && sudo apt install terraform
```

As we are using a *devcontainer*, we add the previous instructions to the `.devcontainer/post-create.sh` script.

Once Terraform has been installed, we can validate it by running:

```console
$ terraform version
Terraform v1.7.4
on linux_amd64
```
