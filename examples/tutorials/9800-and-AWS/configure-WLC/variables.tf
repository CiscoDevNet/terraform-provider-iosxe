# device variables
variable "host_url" {
  description = "Device host path starting with 'https://'"
  type        = string
  sensitive   = true
}

variable "host_ip" {
  description = "Device host IP address"
  type        = string
  sensitive   = true
}

variable "insecure" {
  description = "Device insecure mode boolean"
  type        = string
  sensitive   = true
}

variable "device_username" {
  description = "Device username"
  type        = string
  sensitive   = true
}

variable "device_password" {
  description = "Device password"
  type        = string
  sensitive   = true
}