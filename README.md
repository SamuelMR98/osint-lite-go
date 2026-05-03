# OSINT Lite Go

A beginner-friendly OSINT CLI that checks public username presence across common platforms.

## Usage

```bash
go run main.go [flags] <username>
```

### Flags

- `-h`, `--help`: Show help message
- `-s`, `--social`: Check social media platforms only
- `-t`, `--tech`: Check tech platforms only
- `-j`, `--json`: Output results in JSON format
- `-jsave`, `--json-save`: Save results to a JSON file (e.g., results.json)

### Examples

```bash
# Check all platforms for the username "johndoe"
go run main.go johndoe

# Check only social media platforms for "johndoe"
go run main.go -s johndoe

# Check only tech platforms for "johndoe"
go run main.go -t johndoe

# Check all platforms and output results in JSON format
go run main.go -j johndoe
```