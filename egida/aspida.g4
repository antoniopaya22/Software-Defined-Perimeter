grammar Aspida;

options
{
    language = Go;
}
/*

Aspida Full grammar

@author Antonio Paya Gonzalez
*/

/*

    =====================================  PARSER RULES  =====================================

*/

program             : main hosts tasks variables?;

// =================================================== Blocks
main                : MAIN_KW ':' '{' main_content '}';
hosts               : HOSTS_KW ':' STRING NS;
tasks               : TASKS_KW  ':' '{' tasks_content '}';
variables           : VARS_KW  ':' '{' vars_content '}';

// ======================= MAIN Block
main_content        : main_prop (main_prop)*
                    ;
main_prop           : name NS                   #nameMain
                    | connection NS             #connectionMain
                    | description NS            #descriptionMain
                    ;
name                : 'name' ':' value
                    | 'NAME' ':' value
                    ;
connection          : 'connection' ':' connection_type
                    | 'CONNECTION' ':' connection_type
                    ;
connection_type     : LOCAL_KW                  #connectionLocal
                    | SSH_KW                    #connectionSSH
                    ;
description         : 'description' ':' value
                    | 'DESCRIPTION' ':' value
                    ;

// ======================= TASKS Block
tasks_content       : tasks_prop (tasks_prop)*      #tContent
                    | ifStat (elifStat)* elseStat   #ifStatement
                    ;
tasks_prop          : sections                  #tSections
                    | points                    #tPoints
                    | controls                  #tControls
                    | exclusions                #tExclusions
                    | tags                      #tTags
                    ;
sections            : 'sections' ':' str_array NS
                    | 'SECTIONS' ':' str_array NS
                    ;
points              : 'points' ':' str_array NS
                    | 'POINTS' ':' str_array NS
                    ;
controls            : 'controls' ':' str_array NS
                    | 'CONTROLS' ':' str_array NS
                    ;
exclusions          : 'exclusions' ':' str_array NS
                    | 'EXCLUSIONS' ':' str_array NS
                    ;

tags                : 'tags' ':' str_array NS
                    | 'TAGS' ':' str_array NS
                    ;

// =================================================== Statements
ifStat              : IF comparison '{' tasks_content '}'
                    ;
elifStat            : ELIF comparison '{' tasks_content '}'
                    ;
elseStat            : ELSE '{' tasks_content '}'
                    ;
comparison          : value comp_op value*
                    ;

// ======================= VARS Block
vars_content        : vars_prop (vars_prop)*
                    ;
vars_prop           : STRING ':' value NS                   #varDef
                    | STRING ':' '{' vars_content '}' NS    #varObjDef
                    ;

// =================================================== Values

comp_op: '<'|'>'|'=='|'>='|'<='|'!=';

value
   : STRING     #StringVal
   | NUMBER     #NumberVal
   | 'true'     #TrueVal
   | 'false'    #FalseVal
   | 'null'     #NullVal
   | array      #ArrayVal
   ;

str_array
   : '[' STRING (',' STRING)* ']'
   | '[' ']'
   ;

array
   : '[' value (',' value)* ']'
   | '[' ']'
   ;

// =================================================== Skippables
COMMENT             :  '#' ~('\r' | '\n')* -> skip ;
WHITESPACE          : (' ' | '\t') -> skip ;
NEWLINE             : ('\r'? '\n' | '\r')+ -> skip;



/*

    =====================================  LEXER RULES  =====================================

*/

// =================================================== Fragments

fragment ESC
    : '\\' (["\\/bfnrt] | UNICODE)
    ;
fragment UNICODE
    : 'u' HEX HEX HEX HEX
    ;
fragment HEX
    : [0-9a-fA-F]
    ;
fragment SAFECODEPOINT
    : ~ ["\\\u0000-\u001F]
    ;
fragment INT
    : '0' | [1-9] [0-9]*
    ;
// no leading zeros
fragment EXP
   : [Ee] [+\-]? INT
   ;


STRING
    : '"' (ESC | SAFECODEPOINT)* '"'
    ;
NUMBER
    : '-'? INT ('.' [0-9] +)? EXP?
    ;


// Keywords
NS              : ';';      /*New Sentence*/
MAIN_KW         : 'MAIN';
HOSTS_KW        : 'HOST';
TASKS_KW        : 'TASKS';
VARS_KW         : 'VARS';
LOCAL_KW        : 'LOCAL' | 'local';
SSH_KW          : 'SSH' | 'ssh';
IF              : 'IF';
ELIF            : 'ELIF';
ELSE            : 'ELSE';
OR              : 'OR';
AND             : 'AND';
NOT             : 'NOT';
IS              : 'IS';