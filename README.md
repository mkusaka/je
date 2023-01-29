# je
print json escaped string from stdin or argument

## example

json escape from first argument
```
❯ go run main.go "\"\"\n"   
\"\"\\n
```

json escape string from stdin
```
❯ go run main.go <<EOF
asdf"
asdf
asdf
EOF
asdf\"\nasdf\nasdf\n
```

## install

install via go install
```bash
go install github.com/mkusaka/je@latest
```

install via script
```bash
bash <(curl https://raw.githubusercontent.com/mkusaka/je/main/scripts/install.bash)
```

## usage
## if you install from go
from argument
```bash
je <input>
```

from stdin
```bash
echo <input> | je
```

## if you install from bash
from argument
```bash
./je <input>
```

from stdin
```bash
echo <input> | ./je
```
