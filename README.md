# Skrive
Secure and sleek dosage logging for the terminal.

## Usage
`skrive [path to doses.dat]`

Skrive looks for a doses file in this order:

1. At the path of the program argument, as shown above
2. At the path given by environment variable `SKRIVE_DOSES_PATH`, if set
3. At `~/doses.dat`

Skrive will attempt to create the file if it does not exist.
