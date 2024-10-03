---
title: Parser
date: 2024-09-26
refs: 
tags:
---
https://github.com/AlecKazakova/sqlite-bnf

https://github.com/antlr/grammars-v4/blob/master/sql/sqlite/SQLiteParser.g4


---


O parser deve ser capaz de iniciar a leitura a partir de qualquer ponto da arvore. Caso tendo a seguinte sequência

```
a" OR 1 = 1--
```

Ele deve ser capaz de ler essa sequência de tokens, detectar uma congruência e definir com base em um *trashold* que essa string é um comando.

---

Sistema de **trashold**:

```sql
SELECT a FROM b ;
1      2   3  4 5
```

No caso há uma sequência de 5 tokens válidos (como informado pelo parser). Nesse caso da para afirmar que há uma sequência congruente com a gramática.