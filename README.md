# twix

A command-line duplicate image finder written in Go. twix uses both cryptographic hashing and perceptual hashing to find exact duplicates and visually similar images in your directories.

## Installation

```bash
go install github.com/PlotoZypresse/twix@latest
```

Make sure your `GOPATH` is in your `PATH` environment variable to use the installed binary.

## Usage

```bash
twix <folder-path> [mode]
```

### Arguments

- `<folder-path>` - Directory to scan for duplicate images
- `[mode]` - Optional operation mode (default: `-h`)

### Operation Modes

- `-h` - Hash mode only (default): Uses SHA256 hashing to find exact duplicates
- `-p` - Perceptual hash mode only: Uses perceptual hashing to find visually similar images
- `-hp` - Both modes (currently displays "TODO" message)

## Features

- **Exact Duplicate Detection**: Uses SHA256 hashing to find identical image files
- **Similar Image Detection**: Uses perceptual hashing to find visually similar images
- **Multiple Operation Modes**: Choose between hash comparison, perceptual comparison, or both
- **Recursive Scanning**: Automatically scans all subdirectories
- **Multi-format Support**: Works with JPEG and PNG images
- **Performance Timing**: Shows execution time for duplicate detection

## How It Works

twix employs two methods to detect duplicates:

1. **SHA256 Hashing**: Creates unique fingerprints of image files to find exact duplicates - perfect for identical files
2. **Perceptual Hashing**: Analyzes image content to find visually similar images, even if they're different files or have minor modifications

## Example Output

```bash
$ twix photos -h
Duplicate images found at: /photos/vacation/beach1.jpg and /photos/backup/beach-copy.jpg
Finding duplicates took: 1.234s

$ twix photos -p
Duplicate images found at: /photos/party/image.png and /photos/edited/party-edited.png
Finding duplicates took: 3.456s
```

## Supported Formats

- JPEG (.jpg, .jpeg)
- PNG (.png)

## Building from Source

If you prefer to build manually:

```bash
git clone https://github.com/PlotoZypresse/twix
cd twix
go build twix.go
```

## License

MIT License

---

### Notes on Operation Modes

- **Hash Mode (-h)**: Fastest method, only finds 100% identical files
- **Perceptual Hash Mode (-p)**: Slower but finds visually similar images with minor differences
- **Combined Mode (-hp)**: Planned feature that will use both methods for comprehensive detection

Choose the mode based on your needs: use hash mode for quick exact duplicate finding, or perceptual mode when you suspect images have been resized, recompressed, or slightly modified.