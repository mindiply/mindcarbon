grammar mindcarbon;

program: statement*;

statement: expr SM
          | computationDef
          | assignment
          ;

computationDef: COMPUTATION name=ID LPAREN (paramDef (COMMA paramDef)*)? RPAREN block;
paramDef: name=ID COLUMN type=ID (LPAREN unit=ID RPAREN)?;

block: LCURLY bstat* expr RCURLY;
bstat: assignment | expr SM;

assignment: ID EQUALS expr SM;
expr: fnName=expr LPAREN (prmName=ID COLUMN expr (',' prmName=ID COLUMN expr)*)? RPAREN # functionCall
          | <assoc=right> expr EXP expr # expOp
          | expr op=(MUL|DIV) expr # mulOrDivOp
          | expr op=(ADD|MIN) expr # addOrMinOp
          | ID # idResolution
          | MIN? (FLOAT | INT) # numberConstant
          | QUOTED_STRING # stringConstant
          | LPAREN expr RPAREN # grouping
          ;

COMPUTATION: [Cc][Oo][Mm][Pp][Uu][Tt][Aa][Tt][Ii][Oo][Nn];


ID: [a-zA-Z][a-zA-Z0-9_]*;
QUOTED_STRING: '"' (ESC_SEQ | ~["\\])* '"';
ESC_SEQ: '\\' (['"\\/bfnrt] | UNICODE);
UNICODE: 'u' HEX HEX HEX HEX;
SINGLE_QUOTE: '\'';

INT: '0' | ([1-9] DIGIT*);
FLOAT: DIGIT+ '.' DIGIT+;

fragment DIGIT: [0-9];
fragment HEX: [0-9a-fA-F];

DIV: '/';
MUL: '*';
ADD: '+';
MIN: '-';
EXP: '^';
EQUALS: '=';

COLUMN: ':';
COMMA: ',';
LPAREN: '(';
RPAREN: ')';
LCURLY: '{';
RCURLY: '}';
SM: ';';

COMMENT: '//' ~[\r\n]* -> skip;
WS: [ \t\r\n]+ -> skip;
