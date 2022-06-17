import fileinput
import re

# This script goes through the provided file, and replaces any " #<number>",
# with the valid mark down formatted link to it. e.g.
# " [#number](https://github.com/MonikaCat/provenance/issues/<number>)
# Note that if the number is for a PR, github will auto-redirect you when you click the link.
# It is safe to run the script multiple times in succession.
#
# Example:
#
# $ python ./scripts/linkify.py CHANGELOG.md
for line in fileinput.input(inplace=1):
    line = re.sub(r"\sPR([0-9]+)", r" [PR \1](https://github.com/MonikaCat/provenance/pull/\1)", line.rstrip())
    line = re.sub(r"\s#([0-9]+)", r" [#\1](https://github.com/MonikaCat/provenance/issues/\1)", line)
    print(line)
