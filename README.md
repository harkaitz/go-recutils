# GO-RECUTILS

GNU recutils'es recfile parser/writer on pure Go.
recfiles are human-editable, plaintext databases. This library allows
you to read records and their fields from it. Look for cmd/gorecsel as
an example usage.

* iterate through the records and their fields
* ignore comments
* support continuation lines (\$) and multilines (^+)

It is free software: see the file COPYING for copying conditions.

## Reasons for the fork.

This repository is a forked from [cypherpunks.ru](http://www.git.cypherpunks.ru/?p=gorecfile.git;a=summary) in 2022-10-08.

I had issues downloading the code safely (2022-10-XX):

- I was getting an invalid certificate (2022).
- The connection was slow and closed unexpectedly.

I will keep the package name the same `recfile`. This way you only need
to change the "import" statement to use this fork.

## Updates from cypherpunks.ru:

- 2022-10-08 : Initial fork.

## License and authors:

- License. GPL-3
- Author: Sergey Matveev <stargrave@stargrave.org>
