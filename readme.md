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

## Thingy API Server

When run, by default, the Thingy API Server adds 34 `Thingy` items to the in-memory struct used as a database.

### Retrieve a list of `Thingy` from the DB (`GET`)

The results only displays `thingiesPerPage` results (10 by default), starting at the first element in the DB

```console
$ curl -s localhost:8080/api/v1/thingy
    "id": "01HS33PEN2AFGQ3G0BH7920HGJ",
    "name": "Thingy-0"
  },
  {
    "id": "01HS33PEN3X4SKWGEJ1JQR0HCY",
    "name": "Thingy-1"
  },
  ...
```

A user may provide the *offset* parameter to specify which is the first element retrieved from the DB:

```console
$ curl -s localhost:8080/api/v1/thingy?offset=20
[
  {
    "id": "01HS33PEN3X4SKWGEJ2PZ8XRNK",
    "name": "Thingy-20"
  },
  {
    "id": "01HS33PEN3X4SKWGEJ2R6KJ0N2",
    "name": "Thingy-21"
  },
  {
    "id": "01HS33PEN3X4SKWGEJ2VJBGC14",
    "name": "Thingy-22"
  },
  ...
```

### Get a specifc `Thingy`

At some point I decided to standarize the JSON returned by the API in two main fields: `data` and `error`... But it seems that, for some reason, I did not implemented it for returning a *list* of *Thingies*...

```console
$ curl localhost:8080/api/v1/thingy/id/01HSVGEMYMHNYQEW2TMXM44M90
{
  "data": {
    "id": "01HSVGEMYMHNYQEW2TMXM44M90",
    "name": "Thingy-6"
  },
  "error": null
}
```

