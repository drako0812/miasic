grammar MIASIC;

/*
 * Parser Rules
 */

program: exit_statement EOF;
exit_statement: KW_EXIT WS+ int_lit WS* SEMICOLON WS*;
int_lit: DECINT | HEXINT | OCTINT | BININT;

/*
 * Lexer Rules
 */
fragment DECDIGIT1: [1-9];
fragment DECDIGIT: [0-9];
fragment HEXSPEC: '0' ( 'x' | 'X');
fragment HEXDIGIT: [0-9a-fA-F];
fragment OCTSPEC: '0' ( 'o' | 'O');
fragment OCTDIGIT: [0-7];
fragment BINSPEC: '0' ( 'b' | 'B');
fragment BINDIGIT: [01];

KW_EXIT: 'EXIT';
SEMICOLON: ';';
IDENT: [_a-zA-Z][_a-zA-Z0-9]*;
WS: (WHITESPACE | NEWLINE);
WHITESPACE: (' ' | '\t');
NEWLINE: ('\r'? '\n' | '\r')+;
DECINT: ('0' | (DECDIGIT1 DECDIGIT*));
HEXINT: HEXSPEC HEXDIGIT+;
OCTINT: OCTSPEC OCTDIGIT+;
BININT: BINSPEC BINDIGIT+;