# Skrive
Secure and sleek dosage logging for the terminal.

## Usage
`skrive [path to doses.dat]`

Skrive selects a path for the doses file in the following order:

1. The path provided as the first argument, as shown above
2. The path given by the environment variable `SKRIVE_DOSES_PATH`, if set
3. The path to a file called `doses.dat` in your home directory
    - `~/doses.dat` on Unix-like platforms
    - `C:\Users\<user>\doses.dat` on Windows

Skrive will attempt to create the file if it does not exist.
