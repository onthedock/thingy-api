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
{
  "data": [
    {
      "id": "01HSVPR87X4BDCEN5XYBF6RZCH",
      "name": "Thingy-0"
    },
    {
      "id": "01HSVPR87YERSPD27QBXMJESRR",
      "name": "Thingy-1"
    },
    ...
  ],
  "error": null
}
  ...
```

A user may provide the *offset* parameter to specify which is the first element retrieved from the DB:

```console
$ curl -s localhost:8080/api/v1/thingy?offset=25 | jq
{
  "data": [
    {
      "id": "01HSVPR87YERSPD27QDN59DSPX",
      "name": "Thingy-25"
    },
    {
      "id": "01HSVPR87YERSPD27QDQANYFVJ",
      "name": "Thingy-26"
    },
    ...
  ],
  "error": null
}
```

### Get a specifc `Thingy`

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

### Add a new `Thingy`

Using the `POST` method, we provide a `name` for the `Thingy` and the API will generate an ULID for it automatically:

```console
$ curl -X POST localhost:8080/api/v1/thingy/name/new-thingy-name
{
  "data":{
    "id":"01HSVHKYW5K88N10E3T4CENASX",
    "name":"new-thingy-name"
  },
  "err":null
}
```

### Create a new `Thingy` providing all the fields

We may be interested in adding a `Thingy` with a particular value for its `Id`, instead of relying on an automatically generated one.

For that, we can use the `PUT` method:

```console
$ curl -X PUT -d '{"Id": "01HSVHVTTVEHT0Y6GMCX1MF3XE", "name": "my-thingy"}' http://localhost:8080/api/v1/thingy | jq
{
  "data": "01HSVHVTTVEHT0Y6GMCX1MF3XE",
  "err": null
}
```

As we are providing the full specification of a `Thingy`, including its `Id`, the `PUT` operation is idempotent; we may overwrite the `Thingy` in the DB as many times as we want.

We can validate that it has been created retrieving it using:

```console
$ curl -s -X GET http://localhost:8080/api/v1/thingy/id/01HSVHVTTVEHT0Y6GMCX1MF3XE
{
  "data": {
    "id": "01HSVHVTTVEHT0Y6GMCX1MF3XE",
    "name": "my-thingy"
  },
  "error": null
}
```

### Remove a `Thingy` from the DB

To remove a `Thingy`, we specify its `Id`:

```console
 $ curl -X DELETE http://localhost:8080/api/v1/thingy/id/01HSVHVTTVEHT0Y6GMCX1MF3XE
{
  "data":"01HSVHVTTVEHT0Y6GMCX1MF3XE",
  "error":null
}
```

If we try to delete a non-existing `Thingy`, we get a *Not Found* error:

```console
$ curl -X DELETE http://localhost:8080/api/v1/thingy/id/01HSVHVTTVEHT0Y6GMCX1MF3XE | jq
{
  "data": null,
  "err": "thingy not found"
}
```
