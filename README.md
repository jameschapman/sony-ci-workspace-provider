<!-- PROJECT LOGO -->
<br />
<p align="center">
  <h3 align="center">Sony CI Terraform provider</h3>

  <p align="center">
    IMPORTANT: This is a very limited provider which currently only provides basic folder creation/deletion functionality.
  </p>
</p>

<!-- TABLE OF CONTENTS -->

## Table of Contents

- [About the Project](#about-the-project)
  - [Built With](#built-with)
- [Getting Started](#getting-started)
  - [Prerequisites](#prerequisites)
  - [Installation](#installation)
- [Usage](#usage)
- [Roadmap](#roadmap)
- [Contributing](#contributing)
- [License](#license)

<!-- ABOUT THE PROJECT -->

## About The Project

### Built With

- [Go](https://golang.org/)

<!-- GETTING STARTED -->

## Getting Started

To use this Provider follow these steps.

### Prerequisites

This is an example of how to list things you need to use the software and how to install them.

- Install [Terraform](https://www.terraform.io/downloads.html]) - this should be v0.13.x onwards
- Install [Go](https://golang.org/doc/install)

### Installation

1. Clone the repo

```sh
git clone https://github.com/jameschapman/sony-ci-provider.git
```

2. Build the provider. Note you can change the version number from 1.0.0 in the file name, however ensure you carry this version number through on all of the steps

```sh
// Windows
go build -o terraform-provider-ci_1.0.0.exe

// Unix
go build -o terraform-provider-ci_1.0.0
```

3. Copy Provider to plugin directory. Note, you should replace [host] with a host name of your choice, this could be your company name for example.

```sh
// Windows
move terraform-provider-ci_1.0.0.exe %APPDATA%\terraform.d\plugins\umusic\[host]\ci\1.0.0\windows_amd64\terraform-provider-ci_1.0.0.exe

// Windows
cp terraform-provider-ci_1.0.0 ~/.terraform.d/plugins/[host]/sony/ci/1.0.0/linux_amd64/terraform-provider-ci_1.0.0
```

<!-- USAGE EXAMPLES -->

## Usage

Now reference the Provider in your Terraform files. Ensure you replace [host] with your host name and enter your client_id, client_secret, user, and password details.

```terraform
terraform {
  required_providers {
    ci = {
      source = "[host]/sony/ci"
      version = "1.0.0"
    }
  }
}


provider "ci" {
	client_id      = ""
	client_secret  = ""
    user          = ""
	password      = ""
}

# Parent folder
resource "ci_folder" "parent" {
  name = "parent folder"
  workspace_id  = ""
  parent_id  = ""
}

# Sub folder 1
resource "ci_folder" "submissions" {
  name = "sub folder 1"
  workspace_id  = ""
  parent_id  = ci_folder.parent.id
}

# Sub folder 2
resource "ci_folder" "workorders" {
  name = "sub folder 2"
  workspace_id  = ""
  parent_id  = ci_folder.parent.id
}
```

Now run...

```sh
terraform init
terraform apply
```

<!-- ROADMAP -->

## Roadmap

Note that I have no plans to improve on this Provider unless I need to make a change, however you can still see the [open issues](https://github.com/jameschapman/sony-ci-provider/issues) for a list of known problems.

<!-- CONTRIBUTING -->

## Contributing

Feel free to contribute to the project.

1. Fork the Project
2. Create your Feature Branch (`git checkout -b feature/AmazingFeature`)
3. Commit your Changes (`git commit -m 'Add some AmazingFeature'`)
4. Push to the Branch (`git push origin feature/AmazingFeature`)
5. Open a Pull Request

<!-- LICENSE -->

## License

Distributed under the MIT License.
