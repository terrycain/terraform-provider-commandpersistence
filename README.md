Terraform Provider
==================

- Website: https://www.terraform.io
- [![Gitter chat](https://badges.gitter.im/hashicorp-terraform/Lobby.png)](https://gitter.im/hashicorp-terraform/Lobby)
- Mailing list: [Google Groups](http://groups.google.com/group/terraform-tool)

<img src="https://cdn.rawgit.com/hashicorp/terraform-website/master/content/source/assets/images/logo-hashicorp.svg" width="600px">

Requirements
------------

-	[Terraform](https://www.terraform.io/downloads.html) 0.10.x
-	[Go](https://golang.org/doc/install) 1.16 (to build the provider plugin)

Usage
---------------------

Near identical to the `external` provider except the command is run once on create, and never again.

```
# For example, restrict template version in 0.1.x
resource "commandpersistence_cmd" "example" {
  program = ["python3", "${path.root}/example.py"]

  query = {
    # arbitrary map of string -> string, passed
    # to the external program as the data query.
    id = "abc123"
  }
}
```

Building The Provider
---------------------

Clone repository to: `$GOPATH/src/github.com/terrycain/terraform-provider-commandpersistence`

```sh
$ mkdir -p $GOPATH/src/github.com/terrycain; cd $GOPATH/src/github.com/terrycain
$ git clone git@github.com:terrycain/terraform-provider-commandpersistence
```

Enter the provider directory and build the provider

```sh
$ cd $GOPATH/src/github.com/terrycain/terraform-provider-commandpersistence
$ make build
```

Using the provider
----------------------
## Fill in for each provider

Developing the Provider
---------------------------

If you wish to work on the provider, you'll first need [Go](http://www.golang.org) installed on your machine (version 1.11+ is *required*). You'll also need to correctly setup a [GOPATH](http://golang.org/doc/code.html#GOPATH), as well as adding `$GOPATH/bin` to your `$PATH`.

To compile the provider, run `make build`. This will build the provider and put the provider binary in the `$GOPATH/bin` directory.

```sh
$ make build
...
$ $GOPATH/bin/terraform-provider-commandpersistence
...
```

In order to test the provider, you can simply run `make test`.

```sh
$ make test
```

In order to run the full suite of Acceptance tests, run `make testacc`.

*Note:* Acceptance tests create real resources, and often cost money to run.

```sh
$ make testacc
```

Making a release
----------------

Create a tag like v1.0.0 and the GitHub Actions CI will do the rest.
