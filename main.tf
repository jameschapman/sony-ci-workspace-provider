

terraform {
  required_providers {
    ci = {
      source = "umusic/sony/ci"
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
  name = "submissions"
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