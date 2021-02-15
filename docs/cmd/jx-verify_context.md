## jx-verify context

Verifies the current kubernetes context matches a given name

***Aliases**: ctx*

### Usage

```
jx-verify context
```

### Synopsis

Verifies the current kubernetes context matches a given name

### Examples

  # populate the pods don't have missing images
  jx verify context -c "gke_$PROJECT_ID-bdd_$REGION_$CLUSTER_NAME"%!(EXTRA string=jx-verify)

### Options

```
  -c, --context string   The kubernetes context to match against
  -h, --help             help for context
```

### SEE ALSO

* [jx-verify](jx-verify.md)	 - commands for verifying Jenkins X environments

###### Auto generated by spf13/cobra on 15-Feb-2021