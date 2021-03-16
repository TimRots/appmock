# appmock
---

CLI mocking ala busybox.

## create a config for app foo
```
mkdir -p ~/.appmock
cat << EOMOCK > ~/.appmock/foo.yml
---
flags:
stdout: |-
  foo
stderr: |-
  bar
exit-status: 1
EOMOCK
```

## test a mock
```
ln -sf ~/bin/appmock foo
./foo && echo $?
foo
bar
1
```
