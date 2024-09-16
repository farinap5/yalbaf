---
title: Lexer
date: 2024-09-15
refs: 
tags:
---
O lexer a linha particionada em tokens, tal que

```
SELECT X FROM A
```

Será

```
Token {SELECT, OPERAÇÃO SELECT}
Token {X, DEFINE COLUNA}
Token {FROM, ESCOLHE TABELA}
Token {A, TABELA}
```

----

Problema

A expressão pode eventualmente surgir da seguinte forma:

```
SELECT/**/X/**/FROM/**/X
```

Então, não é possível quebrar a expressão pelos espaços. É necessário olhar o limite do token com base na definição.

O token pode acabar de algumas formas na expressão:

```
TOKEN;
TOKENx20
TOKEN/*
TOKEN#
TOKEN\n
TOKEN--
```