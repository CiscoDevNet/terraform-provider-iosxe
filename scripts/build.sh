XC_ARCH=${XC_ARCH:-"386 amd64 arm"}
XC_OS=${XC_OS:-linux darwin windows freebsd openbsd}
XC_EXCLUDE_OSARCH='!darwin/arm !darwin/386'

gox \
-os="${XC_OS}" \
-arch="${XC_ARCH}" \
-osarch="${XC_EXCLUDE_OSARCH}" \
-output "pkg/{{.OS}}_{{.Arch}}/${PWD##*/}"