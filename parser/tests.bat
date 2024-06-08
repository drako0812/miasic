@echo off

setlocal enabledelayedexpansion

echo "Rebuilding Grammar..."
antlr4 MIASIC.g4

echo "Building JAVA parser..."
javac MIASIC*.java

echo "Running Tests..."

set "tests_dir=TESTS"
set "grun_command=grun MIASIC program"

set "total_files=0"
for %%f in ("%tests_dir%\*.miasic") do (
    set /a "total_files+=1"
)

set "processed_files=0"
for %%f in ("%tests_dir%\*.miasic") do (
    set /a "processed_files+=1"
    set "file=%%f"
    echo Testing file: !file!; !processed_files! / !total_files!
    %grun_command% !file!
)

echo "Done!"

endlocal
