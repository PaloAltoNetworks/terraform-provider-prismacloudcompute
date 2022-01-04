# How to contribute
First off, thanks for taking the time to contribute!
Following these guidelines helps keep the project maintainable, easy to contribute to, and more secure.

## Where to start
There are many ways to contribute.
You can fix a bug, improve the documentation, submit bug reports and feature requests, or take a first shot at a feature you need for yourself.

Pull requests are necessary for all contributions of code or documentation.

## New to open source?
If you're **new to open source** and not sure what a pull request is, welcome!
We're glad to have you!
All of us once had a contribution to make and didn't know where to start.

[Learn how to make a pull request](https://github.com/PaloAltoNetworks/.github/blob/master/Learn-GitHub.md#learn-how-to-make-a-pull-request)

## Fixing a typo or other small issue
Many fixes require little effort or review, such as:

- Typos, white space and formatting changes
- Comment clean up
- Change logging messages or debugging output

These small changes can be made directly in GitHub if you like.

Click the pencil icon in GitHub above the file to edit the file directly in GitHub.
This will automatically create a fork and pull request with the change.
See: [Make a small change with a Pull Request](https://www.freecodecamp.org/news/how-to-make-your-first-pull-request-on-github/)

## Bug fixes and features
For something that is bigger than a one or two line fix, go through the process of making a fork and pull request yourself:

1. Create your own fork of the code
2. Clone the fork locally
3. Make the changes in your local clone
4. Push the changes from local to your fork
5. Create a pull request to pull the changes from your fork back into the upstream repository.
The pull request should be a single commit.

Please use a clear commit message. We'll review every PR and might offer feedback or request changes before
merging.


# Contributing
If you're interested in developing the provider, see below for a basic setup guide.

## Prerequisites
- Go 1.17+
- Terraform 1.0.0+

## Building the provider
0. Set `$GOPATH` if not already set.
    ```bash
    export GOPATH=$(go env GOPATH)
    ```
1. Fetch the repository and navigate to its directory.
    ```bash
    git clone https://github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute.git "$GOPATH"/src/github.com/PaloAltoNetworks/terraform-provider-prismacloudcompute && cd $_
    ```
2. Make your changes.
3. Build and install the provider with your changes.
This also moves the compiled binary to the appropriate location.
    ```bash
    # macOS-specific OS_ARCH; adjust as necessary
    make install OS_ARCH=darwin_amd64 VERSION=0.0.0-testing
    ```
4. Point your terraform file to this local plugin.
    ```terraform
    terraform {
      required_providers {
        prismacloudcompute = {
          source  = "paloaltonetworks.com/prismacloud/prismacloudcompute"
          version = "0.0.0-testing"
        }
      }
    }
    ```

## Developing the provider
See Makefile for available `make` targets.

## Testing the provider
Until a test suite is built, you must test your changes to the provider manually.

## Submitting your changes
Squash all of your changes to a single commit, and submit a pull request.
