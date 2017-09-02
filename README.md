# Confirm

A tiny but useful utility for confirming actions, typically
used in user-executable scripts or makefiles.

## Installation


### Install with Go Get

You can install Go from the [official Go website](https://golang.org/dl/).
Confirm requires Go version 1.6+. Once you have Go installed, simply run:

```
go get -u github.com/albrow/confirm
```

You probably also want to add __$GOPATH/bin__ to __PATH__ so
that you can run confirm directly. To do that permanently, you
can add this line to __~/.bashrc__ or __~/.bash_profile__:

```bash
export PATH=$GOPATH/bin
```

Alternatively, you can run confirm by specifying the full
path: __$GOPATH/bin/confirm__.

### Install Binaries Directly

If you don't want to install Go, you can simply download pre-compiled binaries
for Darwin (Mac OS) and Linux from the
[Releases page](https://github.com/albrow/confirm/releases).

## Example Usage

You can run confirm directly from the command-line, but it wouldn't
be particularly useful that way. Usually, you'll want to use confirm
in user-executable scripts or makefiles.

For example, you might use confirm in a deployment script:

__deploy.sh__:
```bash
# Confirm will return a non-zero status code if it doesn't
# receive confirmation. Calling set -e ensures the script
# will stop running if that happens.
set -e

# The first thing we do is use confirm to make sure the user
# really meant to deploy.
confirm 'Are you sure you want to deploy?'

# Then you can keep running the deployment script.
echo 'Deploying application to production...'

# Do more deployment related stuff...
```

When you run __deploy.sh__, confirm will print out the message
you provided. Here's what it would look like:

```bash
> ./deploy.sh
Are you sure you want to deploy?
> yes
Deploying application to production...
```

It will wait for you to type in 'yes' or 'y' (case-insensitive)
before continuing. You can type 'no' or 'n' to cancel. Using
the key combination `Ctrl+C` will also cancel.

## Advanced Usage

You can supply flags to change the default behavior.

For example, you can use the `--confirm-with` and `--case-sensitive`
flags to wait for the user to type in a specific value exactly before
continuing:

```bash
confirm \
--case-sensitive \
--confirm-with='my_test_database' \
'Are you sure you want to delete the database? Please enter the name of the database to confirm. (my_test_database)' && \
dropdb my_test_database
```

For more information and documentation about all the flags
and options that confirm supports, run `confirm --help`.
