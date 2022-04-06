variable "do_token" {
  type      = string
  sensitive = true
}

variable "repo" {
  type    = string
  default = "https://github.com/Arriven/db1000n"
}

variable "instance_count" {
  type    = number
  default = 1
}

# https://docs.digitalocean.com/reference/api/api-reference/#operation/list_instance_sizes
variable "instance_size_slug" {
  type    = string
  default = "professional-xs"
}

# https://docs.digitalocean.com/reference/api/api-reference/#operation/list_all_regions
variable "region" {
  type    = string
  default = "nyc1"
}
