---
layout: "commandpersistence"
page_title: "Template: commandpersistence_cmd"
sidebar_current: "docs-template-commandpersistence-cmd"
description: |-
  CommandPersistence Command resource.
---

# commandpersistence_cmd

Runs a command and stores its output in the resource. The command is ran 
only once during resource creating time. Works near identically to the
`external` provider except its ran only once.

## Example Usage

E.g.

```hcl
resource "commandpersistence_cmd" "example" {
  program = ["python3", "${path.root}/example.py"]

  query = {
    # arbitrary map from strings to strings, passed
    # to the external program as the data query.
    id = "abc123"
  }
}

output "example_output" {
    value = commandpersistence_cmd.example.result
}
```

## Argument Reference

The following arguments are supported:

* `program` - (Required) A list of strings, whose first element is the program
  to run and whose subsequent elements are optional command line arguments
  to the program. Terraform does not execute the program through a shell, so
  it is not necessary to escape shell metacharacters nor add quotes around
  arguments containing spaces.

* `working_dir` - (Optional) Working directory of the program.
  If not supplied, the program will run in the current directory.

* `query` - (Optional) A map of string values to pass to the external program
  as the query arguments. If not supplied, the program will receive an empty
  object as its input.

## Attributes Reference

The following attributes are exported:

* `result` - A map of string values returned from the external program.

## Import

The output of a command can be imported. The import ID must point to a file containing valid JSON where its only string -> string, no nested objects etc...

```
$ terraform import commandpersistence_cmd.example /some/file.json
```