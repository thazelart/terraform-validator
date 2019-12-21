---
id: recursivity
title: Recursivity
sidebar_label: Recursivity
---
## layers
Each time you define a new layer in your `.terraform-validator.yaml`, this is added to terraform-validator configuration for the current directory and its sub-directories.

for example, here is my root directory configuration file:
```yaml
# .terraform-validator.yaml
layers:
  cust1:
    files:
      default:
        authorized_blocks:
          - variable
          - output
          - provider
          - terraform
          - module
          - data
          - locals
    ensure_providers_version: true
    ensure_terraform_version: true
    block_pattern_name: "^[a-z0-9-]*$"
```

Starting from the root directory and for all its sub-directories, `cust1` layer will be available.

If a sub-directory define another `cust1`, it will replace this configuration for this sub-directory and its own sub-directories.

## current_layer
The default `current_layer` is by default the one choosen in the parent directory.

## tips
During on journey through your stacks and terraform modules, terraform-validator won't run test in folders that do not contain any `.tf` files BUT will read the configurations.                    
That way, you can easily configure a large number of folders at once. Let's take this example:
```bash
|-- modules
|   |-- .terraform-validator.yaml
|   `-- mod1
|       |-- README.md
|       |-- main.tf
|       `-- provider.tf
|   `-- mod2
|       |-- README.md
|       |-- main.tf
|       `-- provider.tf
```
The `.terraform-validator.yaml` file inside the `modules` folder define the configuration for both `mod1` and `mod2` !
