# deptest4glide

A test of glide (https://github.com/Masterminds/glide) adding transitive dependencies.

Glide appears to include more dependencies into glide.lock and the vendor/ directory than needed for a build.
The test program in this directory only needs golang.org/x/net (and its subpackage golang.org/x/net/atom) to be built and run - as shown in the log of 'go install -v -x' included in the go_install_log.txt file.  `go list` shows the following dependencies:
```
$ go list -e -f '{{join .Deps "\n"}}' | xargs go list -f '{{if not .Standard}}{{.ImportPath}}{{end}}'
golang.org/x/net/html
golang.org/x/net/html/atom
```
However, `glide update` also includes golang.org/x/crypto and golang.org/x/text - see the generated glide.lock file.  The question is why it does this (most likely due to the full scan of golang/org/x/net for dependencies) and whether it would be possible and desirable to change its handling of dependencies (or provide an option) so that only the dependencies that you need for your project would be included. glide.lock.want is what I would expect the generated glide.lock file for this project to be.

It would be great to use glide.yaml and glide.lock files as documentation: the list of dependencies required by the project to build - all the dependencies you need but only those that are required, no more and no less. 'Overincluding' confuses the situation, make glide.lock an unreliable source of information and makes users ask why packages like golang.org/x/crypto are being pulled when they are not needed for the project.

Related discussion on a glide issue: https://github.com/Masterminds/glide/issues/166
