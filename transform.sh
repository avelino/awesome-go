#!/bin/bash

## This bash script transforms the awesome-go README.md file such as to list all repositories
## in tables, as opposed to unordered list items, including columns for shield images displaying
## the number of stars, forks, issues and pull requests each github repository has.
##
## The script is run by specifying an input file and an output file.
##
## ex. transform README.md ALTREADME.md
##
## The above command will create a file called ALTREADME.md which is a transformed version of
##   the README.md file

cp $1 tmptransform
dos2unix tmptransform # convert newlines to *nix format to simplify regexes

echo "Escape HTML characters"
gawk -i inplace '{ print gensub(/&/, "\\&amp;", "g")}' tmptransform
gawk -i inplace '{ print gensub(/"/, "\\&quot;", "g")}' tmptransform

echo "Transform all list item links of github repositories"
gawk -i inplace '{ print gensub(/\* \[([^\)]+)\]\((https:\/\/github\.com\/([^\/]+\/[^\)]+))\)(.*)/, "<tr><td><a href=\"\\2\">\\1</a>\\4</td><td><img src=\"https://img.shields.io/github/stars/\\3\" alt=\"stars\"></td><td><img src=\"https://img.shields.io/github/forks/\\3\" alt=\"forks\"></td><td><img src=\"https://img.shields.io/github/issues/\\3\" alt=\"issues\"></td><td><img src=\"https://img.shields.io/github/issues-pr/\\3\" alt=\"pull requests\"></td></tr>", "g") }' tmptransform

echo "Transform remaining list item links (non github repositories)"
gawk -i inplace '{ print gensub(/\* \[([^\)]+)\]\((https[^\)]+)\)(.*)/, "<tr><td><a href=\"\\2\">\\1</a>\\3</td><td></td><td></td><td></td><td></td></tr>", "g") }' tmptransform

echo "Transform any other inline links"
gawk -i inplace '{ print gensub(/\[([^\)]+?)\]\((https[^\)]+)\)/, "<a href=\"\\2\">\\1</a>", "g")}' tmptransform

echo "Add newline before and after list items" # necessary for the 'paragraphing' in the next step to work"
gawk -i inplace '{ print gensub(/^(\*[^\*]*?)$/, "\n\\1\n", "g")}' tmptransform

echo "Remove duplicate empty lines" # possible artefacts of the previous operation"
cat -s tmptransform > tmpfile
mv tmpfile tmptransform

echo "Add table headers and closers" # around 'paragraphs' starting with <tr> and ending with </tr>
# I could for the life of me not figure out how to do this using gawk so I falled back to sed
sed -i '/./{H;$!d} ; x ; s/^\n\s*<tr>/\n<table style="width:100%">\n<tr><th style="width:60%">Library<\/th><th style="width:9%">Stars<\/th><th style="width:9%">Issues<\/th><th style="width:12%">Forks<\/th><th style="width:10%">PRs<\/th><\/tr>\n<tr>/ ; s/<\/tr>$/<\/tr>\n<\/table>/' tmptransform

echo "Create output file"
mv tmptransform $2

echo "Done"