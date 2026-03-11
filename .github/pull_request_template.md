<!--- Please ensure that the WIP label is not being applied if ready for review -->
<!--- If the WIP label is applied to your PR, no one will look at it -->
<!--- Please feel free to ask for help -->

## ✔️ Type of Change

- [ ] 🚀 New feature (new resource/data source, new attributes to existing resource)
- [ ] 🐛 Bug fix (corrects an issue in resource behavior, RESTCONF/NETCONF handling)
- [ ] ⚠️ Breaking change (modifies existing resource schema or removes attributes)
- [ ] 📄 Documentation update (resource docs, examples, registry docs)
- [ ] 🛠️ Refactoring / tech debt (internal improvements, no user-facing changes)
- [ ] ⚙️ Other (please describe):

## 🔗 Related Issue(s)

<!--- Link the relevant issue(s). Use "Fixes #123" to auto-close on merge. -->

## 📝 Proposed Changes

<!--- Describe what this PR does and why. Include YANG paths if adding new attributes. -->

---

## 🧪 Testing Evidence

> **⚠️ PRs without testing evidence will not be reviewed** (unless this is a docs-only or chore-only change — check N/A below).

- [ ] N/A — this PR does not touch definitions, generated code, or provider logic (e.g., docs-only, CI, chore)

### 1. Code Generation

<!--- If modifying `gen/definitions/*.yaml`, confirm code generation runs cleanly -->

```bash
# Paste `make gen NAME=<resource>` output
```

- [ ] `make gen` completes without errors
- [ ] Generated code matches YANG model expectations

### 2. Acceptance Tests

<!--- Provider uses Go acceptance tests against a real IOS-XE device -->

```bash
# Paste `make testacc TESTARGS="-run TestAcc<Resource>"` output or relevant snippet
```

- [ ] Acceptance tests pass against a real device
- [ ] Device type and IOS-XE version: _______

---

## 🔗 Cross-Repo Links

<!--- Provider changes often trigger downstream module + schema updates. -->

- Module PR: <!-- e.g., netascode/terraform-iosxe-nac-iosxe#123 -->
- Schema PR: <!-- e.g., netascode/nac-iosxe#789 -->

## ✅ Checklist

### Pre-Submission

- [ ] I have merged/rebased from `main` and resolved all merge conflicts
- [ ] `make gen` runs cleanly for affected resources
- [ ] Acceptance tests pass against a real IOS-XE device
- [ ] `go build ./...` compiles without errors

### Code Quality

- [ ] Definition YAML follows existing patterns in `gen/definitions/`
- [ ] YANG paths are correct and validated against YangSuite
- [ ] Examples in `examples/resources/` are updated for new/changed attributes
- [ ] Resource documentation in `docs/` is regenerated and accurate

### Labels & Assignment

- [ ] I have assigned at least one label that aligns with `release.yaml` categories: `feature`, `enhancement`, `bug`, `fix`, `breaking-change`, `documentation`, `refactor`, `tech-debt`, `chore`, or `dependencies`
- [ ] Label(s) match the linked GitHub issue label(s) where applicable
- [ ] PR is assigned to myself
- [ ] One or more reviewers are assigned
