<a name="readme-top"></a>

[![Contributors][contributors-shield]][contributors-url]
[![Forks][forks-shield]][forks-url]
[![Stargazers][stars-shield]][stars-url]
[![Issues][issues-shield]][issues-url]
[![MIT License][license-shield]][license-url]
[![LinkedIn][linkedin-shield]][linkedin-url]

# GopherDB

Welcome to GopherDB! GopherDB is a flexible, distributed database platform that supports multiple concurrent workspaces, each uniquely identified. Tailor each workspace with options for Cache, Persistent, or hybrid configurations.

### Cache

Customize your cache with a number of eviction algorithms. GopherDB currently supports **LRU** (Least Recently Used).

Coming Soon - **LFU** (Least Frequently Used), **MRU** (Most Recently Used), **ARC** (Adaptive Replacement Cache), **TLRU** (Time-aware Least Recently Used)

### Disk

Persistent storage will include SQL, NoSQL, TimeSeries, and Vector databases, making it a comprehensive solution for managing diverse data needs in one unified platform.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

## Getting Started

Welcome to the GopherDB project repository! This document provides essential information for developers working on or contributing to the GopherDB application.

### Makefile

The make file currently only supports MacOS / Unix systems. This is used to build and install the gopherdb CLI globally on your system.

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Entry Points

_The following is informational. See_ **Running Commands** _for actual test runs._

#### Server

The entry point for the GopherDB server is located in the `server.go` file within the `server` directory. This file initializes the server with necessary configurations and starts listening for incoming requests.

```plaintext
/server/server.go
```

#### CLI

The entry point for the GopherDB CLI is located in the `root.go` file within the `cli/cmd` directory. This file is the root commands and configurations and is where all incoming requests will pass.

```plaintext
/cli/cmd/root.go
```

New commands should be their own go files under `cli/cmd` and the command var should be added to `root.go`'s init() func.

`root.go`:

    func init() {
        rootCmd.AddCommand(
            versionCmd,
            timezoneCmd,
            startCmd,
            // New Command Here
        )
        ...
    }

<p align="right">(<a href="#readme-top">back to top</a>)</p>

### Running Commands

Try running the gopherdb cli to start the server.

From root, open your terminal and run the following:

```plaintext
cmd cli && go run main.go help
```

#### CLI

The entry point for the GopherDB CLI is located in the `root.go` file within the `cli/cmd` directory. This file is the root commands and configurations and is where all incoming requests will pass.

```plaintext
/cli/cmd/root.go
```

<p align="right">(<a href="#readme-top">back to top</a>)</p>

<!-- MARKDOWN LINKS & IMAGES -->
<!-- https://www.markdownguide.org/basic-syntax/#reference-style-links -->

[contributors-shield]: https://img.shields.io/github/contributors/bwlee13/gopherdb.svg?style=for-the-badge
[contributors-url]: https://github.com/bwlee13/gopherdb/graphs/contributors
[forks-shield]: https://img.shields.io/github/forks/bwlee13/gopherdb.svg?style=for-the-badge
[forks-url]: https://github.com/bwlee13/gopherdb/network/members
[stars-shield]: https://img.shields.io/github/stars/bwlee13/gopherdb.svg?style=for-the-badge
[stars-url]: https://github.com/bwlee13/gopherdb/stargazers
[issues-shield]: https://img.shields.io/github/issues/bwlee13/gopherdb.svg?style=for-the-badge
[issues-url]: https://github.com/bwlee13/gopherdb/issues
[license-shield]: https://img.shields.io/github/license/bwlee13/gopherdb.svg?style=for-the-badge
[license-url]: https://github.com/bwlee13/gopherdb/blob/master/LICENSE
[linkedin-shield]: https://img.shields.io/badge/-LinkedIn-black.svg?style=for-the-badge&logo=linkedin&colorB=555
[linkedin-url]: https://www.linkedin.com/in/brandon-lee-68944885/
