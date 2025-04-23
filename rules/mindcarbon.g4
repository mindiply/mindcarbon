grammar mindcarbon;

program: statement*;

statement: assignment
          | computationDef
          | expr SM
          ;

computationDef: COMPUTATION name=ID LPAREN (paramDef (COMMA paramDef)*)? RPAREN LCURLY assignment* expr RCURLY;
paramDef: name=ID COLUMN type=ID (LPAREN unit=ID RPAREN)?;

assignment: ID EQUALS expr SM;
expr: LPAREN expr RPAREN # grouping
          | expr '^'<assoc=right> expr # exponentiation
          | expr DIV expr # division
          | expr MUL expr # multiplication
          | expr ADD expr # addition
          | expr MIN expr # subtraction
          | fnName=ID LPAREN (expr (',' expr)*)? RPAREN # functionCall
          | ID # identifier
          | MIN? FLOAT # floatConstant
          | MIN? INT # intConstant
          | QUOTED_STRING # stringConstant
          ;

COMPUTATION: [Cc][Oo][Mm][Pp][Uu][Tt][Aa][Tt][Ii][Oo][Nn];


ID: [a-zA-Z][a-zA-Z0-9_]*;
QUOTED_STRING: '"' (ESC_SEQ | ~["\\])* '"';
ESC_SEQ: '\\' (['"\\/bfnrt] | UNICODE);
UNICODE: 'u' HEX HEX HEX HEX;
HEX: [0-9a-fA-F];
SINGLE_QUOTE: '\'';

INT: [1-9] DIGIT*;
FLOAT: INT.INT;

fragment DIGIT: [0-9]+;

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
