# Latest SBOM Files

This directory contains the Software Bill of Materials (SBOM) for the most recent release of terraform-provider-iosxe.

## Files

- `sbom.spdx.json` - SBOM in SPDX 2.3 format (ISO/IEC 5962:2021)
- `sbom.cyclonedx.json` - SBOM in CycloneDX format (OWASP)

## Maintenance

These files are **automatically updated** after each release by the GitHub Actions workflow. The `update-sbom-in-repo` job:

1. Downloads SBOM files generated during the release
2. Copies them to this directory
3. Updates version metadata in this README
4. Commits and pushes changes to the main branch

Typically, these files are updated within 2-3 minutes of a release being published.

### Manual Update (If Automation Fails)

If the automated update fails, you can manually update the files:

```bash
# Set the new version
VERSION=v0.14.0  # Replace with actual new version

# Download SBOM files from the release
curl -LO https://github.com/CiscoDevNet/terraform-provider-iosxe/releases/download/${VERSION}/terraform-provider-iosxe_${VERSION}_sbom.spdx.json
curl -LO https://github.com/CiscoDevNet/terraform-provider-iosxe/releases/download/${VERSION}/terraform-provider-iosxe_${VERSION}_sbom.cyclonedx.json

# Copy to sbom/latest/ directory
cp terraform-provider-iosxe_${VERSION}_sbom.spdx.json sbom/latest/sbom.spdx.json
cp terraform-provider-iosxe_${VERSION}_sbom.cyclonedx.json sbom/latest/sbom.cyclonedx.json

# Update this README with version information, then commit and push
git add sbom/latest/
git commit -m "chore: update SBOM to ${VERSION}"
git push
```

## Current Version

**Provider Version**: v0.16.0
**Generated**: 2026-02-24
**Syft Version**: v1.41.2

## Usage

For usage instructions and examples, see [SBOM.md](../../SBOM.md) in the repository root.

## Historical Versions

All version-specific SBOMs are stored in the repository at `sbom/{version}/`:
- Browse all versions: https://github.com/CiscoDevNet/terraform-provider-iosxe/tree/main/sbom
- Example: `sbom/v0.14.3/sbom.spdx.json`
