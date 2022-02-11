# CDK

Cloud Development Kit for Terraform ([CDKTF](https://github.com/hashicorp/terraform-cdk)) allows you to use familiar programming languages to define cloud infrastructure and provision it through HashiCorp Terraform. This gives you access to the entire Terraform ecosystem without learning HashiCorp Configuration Language (HCL) and lets you leverage the power of your existing toolchain for testing, dependency management, etc.

## Setup

This project uses [asdf](http://asdf-vm.com/) to install the dependencies, see [.tool-versions](.tools-versions) for the list.

```sh
asdf install
```

Alternatively ensure that the dependencies are installed on you host.

## Modules and providers

cdktf generates code for the providers that are used, therefore you need to run the following command every time the cdktf.json is updated. 

```sh
make get
```

## Build

Run make to the install the binary.

```sh
make install
```

## Run

```sh
Usage:
  make <target> e.g. make synthesize

Targets:
  synthesize            Synthesize Terraform resources to cdktf.out/
  diff                  Perform a diff (terraform plan) for the given stack
  deploy                Deploy the given stack
  destroy               Destroy the given stack
  get                   Generates provider-specific bindings
  ```

  The default stack name is cdk, to override this:

  ```sh
  make STACK=something-else synthesize
  ```