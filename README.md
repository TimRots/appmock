# appmock
---

CLI mocking ala busybox

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

## test the mock
```
ln -s ~/bin/appmock foo
./foo && echo $?
foo
bar
1
```

## how to build
```
# create bin directory in homedir and add it to $PATH
mkdir -p ~/bin
export PATH="$HOME/bin:$PATH"

# clone repo
git clone https://github.com/TimRots/appmock.git && cd appmock

# get dependencies
go get -t -v -d -u ./...

# runnning make will build the binary and place it in ~/bin
make
```
