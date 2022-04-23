# cleanP
 Tool for bug hunters when you have a big list of endpoints this tool helps you to extract unique parameters clean and without duplication of values with preserving the original values


### install 
`go get -u github.com/raoufmaklouf/cleanP`


### usage
` cat file.txt | cleanP `

### example

```
> cat EndPointFile.txt
http://www.testme.com/prod/abc?size=S
http://www.testme.com/prod/xyz?size=M
http://www.testme.com/prod/aaa?size=L
http://www.testme.com/prod/fff?size=XL
http://www.testme.com/prod/rrr?size=XXL
http://www.testme.com/user/bob?q=show
http://www.testme.com/user/jack?q=remove&id=123
http://www.testme.com/user/nilson?q=add&id=456
http://www.testme.com/user/ramzi?q=add&id=123
http://www.testme.com/user/nina?q=add&id=456
```
```
> cat EndPointFile.txt | cleanP 
http://www.testme.com/prod/abc?size=S
http://www.testme.com/user/bob?q=show
http://www.testme.com/user/jack?q=remove&id=123
```
