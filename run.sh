set -xe

go build && ./trie


# enable to print trie using graphiz svg and png
# dot -Tsvg out.dot > output.svg
dot -Tpng out.dot > output.png