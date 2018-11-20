Go Title Generator (WIP)
========================

Using a markov-chain approach to generate news headlines. The chain can be trained by textfiles contains lines of headlines.
 

```
$ ./gotitlegenerator --help
Usage of ./gotitlegenerator:
  -input string
        the path to load the titles from (default "model.txt")
  -ngram int
        ngram size to generate the chain from (default 1)
  -words int
        number of words to create the title from (default 6)```

