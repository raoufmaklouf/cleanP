# cleanP
 Tool for bug hunters when you have a big list of endpoints this tool helps you to extract unique parameters clean and without duplication of values with preserving the original values


# install 
`go get -u github.com/raoufmaklouf/cleanP`


# usage
` cleanP file.txt`

# example

```
> cat EndPointFile.txt
http://www.testme.com/path1/aaaa?q=test1
http://www.testme.com/path2/aaab?q=test1
http://www.testme.com/path1/aaaa?q=test2
http://www.testme.com/path2/aaac?q=test2
http://www.testme.com/path1/aaab?q=test3
http://www.testme.com/path2/aaaa?q=test3
http://www.testme.com/path1/aaac?q=test1&id=123
http://www.testme.com/path2/aaac?q=test1&id=456
http://www.testme.com/path1/aaac?q=test1&id=123
http://www.testme.com/path2/aaaa?q=test1&id=456
```
```
> cleanP EndPointFile.txt
http://www.testme.com/path1/aaaa?q=test1
http://www.testme.com/path2/aaab?q=test1
http://www.testme.com/path1/aaac?q=test1&id=123
http://www.testme.com/path2/aaac?q=test1&id=456
```
