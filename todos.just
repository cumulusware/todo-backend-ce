# -*- Justfile -*-

local_url := "http://localhost:8080"

# List the available todos subcommands.
@default:
  just --list --unsorted --justfile todos.just

# Test the OPTIONS for todos.
options url=local_url:
  curl -I -X OPTIONS {{url}}/api/todos/

# Delete all todos.
deleteall url=local_url:
  curl -i -X DELETE {{url}}/api/todos/
