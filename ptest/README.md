# BNF

```bnf
<expression> ::= ( <term> | ! <term> ) [ <combine_opts> <expression> ]*
<combine_opts> ::= ("-a" | "-o")
<term> ::= <factor> [ <between_opts> <factor> ] | <prefix_opts> <factor> 
<prefix_opts> ::= "-" [bcdefghknprstuwxzLOGS]
<between_opts> ::= ("=" | "==" | "!=" | ">" | ">=" | "<" | "<=" | "-nt" | "-ot" | "-ef" | "-eq" | "-ne" | "-lt" | "-le" | "-gt" | "-ge" | "-starts" | "-ends" | "-contains") 
<factor> ::= <string> | <number> | "(" expression ")"
<number> ::= [0-9]+
<string> ::= [a-zA-Z_$][-a-zA-Z_0-9$]+
```
