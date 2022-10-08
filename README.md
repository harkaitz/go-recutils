# GO-RECUTILS

## Reason for the fork.

This is a fork of `http://www.git.cypherpunks.ru/?p=gorecfile.git;a=summary` repository.

I had issues downloading the code from `www.git.cypherpunks.ru` safely (2022-10-XX):

1.- I was getting an invalid certificate (2022).
2.- The connection is slow and closes (even ignoring all security practices).

The package name is `recfile`.

## Updates from www.git.cypherpunks.ru:

- 2022-10-08 : Initial fork.

## Original README:

GNU recutils'es recfile parser/writer on pure Go.
recfiles are human-editable, plaintext databases. This library allows
you to read records and their fields from it. Look for cmd/gorecsel as
an example usage.

* iterate through the records and their fields
* ignore comments
* support continuation lines (\$) and multilines (^+)

It is free software: see the file COPYING for copying conditions.

## License and authors:

- License. GPL-3
- Author: Sergey Matveev <stargrave@stargrave.org>
