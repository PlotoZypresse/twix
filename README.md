# twix

A command-line duplicate image finder written in Go. twix uses both cryptographic hashing and perceptual hashing to find exact duplicates and visually similar images in your directories.

## Installation

```bash
go install github.com/PlotoZypresse/twix@latest
```

Make sure your `GOPATH` is in your `PATH` environment variable to use the installed binary.

## Usage

```bash
twix <folder-path>
```

Replace `<folder-path>` with the directory you want to scan for duplicate images.

## Features

- Exact Duplicate Detection: Uses SHA256 hashing to find identical image files
- Similar Image Detection: Uses perceptual hashing to find visually similar images
- Recursive Scanning: Automatically scans all subdirectories
- Multi-format Support: Works with JPEG and PNG images

## How It Works

twix employs two methods to detect duplicates:

1. SHA256 Hashing: Creates unique fingerprints of image files to find exact duplicates
2. Perceptual Hashing: Analyzes image content to find visually similar images, even if they're different files

## Example Output

```bash
$ twix photos
Duplicate images found at: /photos/vacation/beach1.jpg and /photos/backup/beach-copy.jpg
Duplicate images found at: /photos/party/image.png and /photos/edited/party-edited.png
Finding duplicates took: 1.234s
```

## Supported Formats

- JPEG (.jpg, .jpeg)
- PNG (.png)

## Building from Source

If you prefer to build manually:

```bash
git clone https://github.com/your-username/twix
cd twix
go build -o twix main.go
```

## License

MIT License