---
title: Parser BNF
date: 2024-09-26
refs: 
tags:
---
```ebnf
<sttmSeq> ::= <sttm> ( ";" <sttm> )* ";"?
<sttm> ::= <selectSttm>
<selectSttm> ::= "SELECT" <resultColumn> ( "," <resultColumn> )* ( "FROM" <tableName> ( "WHERE" <expr> )? "GROUP BY" <expr> )?
<resultColumn> ::= "*" | <tableName> "." "*" | <expr> ( "AS" <columnAlias> )?
<tableName> ::= [a-z]+
<columnAlias> ::= [a-z]+
```
