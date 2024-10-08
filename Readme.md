# About the Credential utility

<!-- TODO: add a section on how to run tests. -->
<!-- You can have three h2 titles: Install, Run, and Test. Then, you can narrow them down
by explaining in a more specific way. -->

This CLI utility exposes the following commands which can be used to generate credentials like password(s)
This CLI utility exposes the following commands which can be used to generate credentials like password(s).
It was build with love using https://github.com/spf13/cobra-cli

## How to install or create binaries

Simply run

```shell
go install ./cmd/credential
make sure that you have feteched all tags from the git repo.

```shell
git fetch --all --tags
```

To install, simply run

```shell
make install
```
this will create a binary file named `credential` in your $GOPATH.
You may add $GOPATH to your $PATH to run the `credential` utility

To build binaries for distribution, simply run

```shell
make binary
```
This will create binaries for Darwin-linux, Linux-amd64, Linux-arm64 in `./out/bin`

## How to run tests

To run the tests from the code, simply run

```shell
make test
```
## How to run this utility

To generate a password run the following command

```shell
~/go/bin/credential password
```

or if you have the `credential` in your $PATH

```shell
credential password
```
### Available options for generating passwords

<!-- TODO: expand this section and put some run examples. -->

- Length of the password(s) can be specified using `length=8` forexample
- To dump the password(s) to console set `output=0`
- To dump the password(s) to file set `output=1` and set filepath using `file=/your/filepath`
- To control the number of passwords use `count=20` for example
- If you want to include special characters in your passwords(recommended) use the `includeSpecialCharacters true` flag
- For some reason if you need to generate URL-safe/Base64Encoded passwords (Never ever send passwords in URLs)
use `urlSafe=true`

Some examples to run the utility( Assuming you have put `credential` in your $PATH)

- To generate a password with default options
    ```shell
    credential password
    ```

    where the default options are:
    - length=7
    - output=terminal
    - includeSpecialCharacters=true
    - urlSafe=False

- To generate a password with specified length and output to default file `./passwords.txt`
    ```shell
    credential password --length=40 --output=1 
    ```

- To generate 10 passwords with specified length and output to specific file `./myPassword.txt`
    ```shell
    credential password --length=40 --output=1 --count=10 --file=./myPassword.txt
    ```

- To generate 10 passwords with specified length and output to console
    ```shell
    credential password --length=40 --count=10
    ```

- To generate 10 URL safe passwords with specified length and output to console
    ```shell
    credential password --length=40 --count=10 --urlSafe=true
    ```
    
## How to run/test it without installing

Simply replace `~/go/bin/credential` with `go run cmd/credential/main.go` and you can try everything described above
