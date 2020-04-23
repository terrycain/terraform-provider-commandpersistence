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