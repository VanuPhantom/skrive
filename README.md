# Skrive
Secure and sleek dosage logging for the terminal.

## Usage
```
skrive [-f path to doses.dat]
skrive log [-f path to doses.dat] [<quantity> <substance> <route> [time-spec]]
```

The log command's time argument can be formatted using either `dhm` formatting or as a number of minutes.  
For example, 90 minutes can be formatted as follows;
- `1h30m`
- `1h30`
- `90`

Skrive selects a path for the doses file in the following order:

1. The path provided with the `-f` flag, as shown above
2. The path given by the environment variable `SKRIVE_DOSES_PATH`, if set
3. The path to a file called `doses.dat` in your home directory
    - `~/doses.dat` on Unix-like platforms
    - `C:\Users\<user>\doses.dat` on Windows

Skrive will attempt to create the file if it does not exist.
