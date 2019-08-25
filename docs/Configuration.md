# Customized configuration

## Main idea
terraform-validator provides some customizations via the `.terraform-validator.yaml` file.
The defaults fit for most projects.

your configuration file can contains two parameters: `classes` and `current_folder_class`.


## `classes`
Classes is a map of class. By default, classes contains one class named `default` :
```
classes:
  default:
    files:
      main.tf:
        mandatory: true
        authorized_blocks:
      variables.tf:
        mandatory: true,
        authorized_blocks:
          - variable
      outputs.tf:
        mandatory: true
        authorized_blocks:
          - output
      providers.tf:
        mandatory: true
        authorized_blocks:
          - provider
      backend.tf:
        mandatory: true
        authorized_blocks:
          - terraform
      default:
        mandatory: false
        authorized_blocks:
          - resource
          - module
          - data
          - locals
    ensure_terraform_version: false
    ensure_providers_version: false
    block_pattern_name: "^[a-z0-9_]*$"
```
### What is a class ?
A class is a set of parameter that define the complete configuration of terraform-validator.

#### files
files is a map of filename that contains what rules define each files.
* `mandatory`: set if the file is mandatory.                            
  *Type: `boolean`*                          
  *Default: `false`*
* `authorized_blocks`: is a list of authorized terraform block types.                          
  *Type: `List`*                    
  *Default: `<empty>`*                    
  available terraform block types:
    * variable
    * output
    * provider
    * terraform
    * resource
    * module
    * data
    * locals

If a match does not match exactly one of the files, the configuration will be taken from `default`.

```
# .terraform-validator.yaml
...
      # main.tf is mandatory with no block inside
      main.tf:
        mandatory: true
        authorized_blocks:
      # variables.tf is not mandatory with only variable blocks inside
      variables.tf:
        mandatory: true,
        authorized_blocks:
          - variable
      # other files will match default config
      default:
        mandatory: false
        authorized_blocks:
          - resource
          - module
          - data
          - locals
...
```

#### ensure_terraform_version
*Type: `boolean`*                 
*Default: `false`*                  
Configure terraform-validator in order to ensure (or not) if the terraform version has been set.

#### ensure_providers_version
*Type: `boolean`*                 
*Default: `false`*    
Configure terraform-validator in order to ensure (or not) if the providers versions has been set.

#### block_pattern_name
*Type: `string`*                 
*Default: `^[a-z0-9_]*$`*    
Configure the pattern that should match each terraform resources.


## `current_folder_class`
*Type: `string`*
*Default: `default`*                       
The `current_folder_class` permit you to select the configuration class you want to use in the current folder.

```
# .terraform-validator.yaml
current_folder_class: my_config
```


## Recursivity and configuration
### `classes`
Each time you define a new class in your `.terraform-validator.yaml`, this is added to terraform-validator configuration for the current directory and its sub-directories.

for example, here is my root directory configuration file:
```
# .terraform-validator.yaml
classes:
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

Starting from the root directory and for all its sub-directories, `cust1` class will be available.

If a sub-directory define another `cust1`, it will replace this configuration for this sub-directory and its own sub-directories.

### `current_folder_class`
The default `current_folder_class` is by default the one choosen in the parent directory.

### tips
During on journey through your stacks and terraform modules, terraform-validator won't run test in folders that do not contain any `.tf` files BUT will read the configurations.                    
That way, you can easily configure a large number of folders at once. Let's take this example:
```
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
