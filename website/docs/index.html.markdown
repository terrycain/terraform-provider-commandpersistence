---
layout: "commandpersistence"
page_title: "Provider: CommandPersistence"
sidebar_current: "docs-commandpersistence-index"
description: |-
  The CommandPersistence provider is used to store the output of a command in a resource.
---

# CommandPersistence Provider

The commandpersistence provider works just like the external provider but instead of a data source, a resource is declared.

Use the navigation to the left to read about the available data sources.

## Example Usage

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
