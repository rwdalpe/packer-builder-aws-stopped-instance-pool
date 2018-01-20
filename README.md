# AWS Stopped Instance Pool Packer Plugin

A builder for packer which will create an AWS AMI using an existing EC2 instance
from a pool of stopped instances as its starting point instead of starting a new
EC2 instance from scratch. This will speed up the packing process, especially in
the case of Windows where there is an unavoidable wait time when starting a
new instance.

## Development

### Prerequisites

* Go 1.x installed
* `make` installed
* common shell utilities such as `find`, `rm`, and so on available on your `PATH`

### Building

This project manages its own `GOPATH` here in the project root for local
development. Consequently we've wrapped the normal go toolchain with a makefile
to help with that.

Simply replace `go` with `make` to handle most normal go development tasks.

```
make fmt # to format the project's go files
```

```
make get # to fetch dependencies into the project's GOPATH
```

```
make build # to build the project using the project's GOPATH
```

```
make install # to install the project into the project's GOPATH
```

```
make clean # to remove _only the project's_ source files from the project's GOPATH
```

```
make clean-go # to remove the entire project GOPATH
```

## License Information

Copyright (C) 2018 Robert Winslow Dalpe

This program is free software: you can redistribute it and/or modify
it under the terms of the GNU Affero General Public License as published by
the Free Software Foundation, either version 3 of the License, or
(at your option) any later version.

This program is distributed in the hope that it will be useful,
but WITHOUT ANY WARRANTY; without even the implied warranty of
MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE.  See the
GNU Affero General Public License for more details.

You should have received a copy of the GNU Affero General Public License
along with this program.  If not, see <http://www.gnu.org/licenses/>.

See [LICENSE](LICENSE) for more details.
