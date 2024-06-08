@echo off
setlocal

set ANTLR_JAR=./antlr-4.13.1-complete.jar
set ANTLR_OPTIONS=-Dlanguage=Go -no-visitor -package parser

for %%f in (*.g4) do (
    java -Xmx500M -cp "%ANTLR_JAR%" org.antlr.v4.Tool %ANTLR_OPTIONS% %%f
)

endlocal
