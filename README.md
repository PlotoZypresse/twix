# twix

A command-line duplicate image finder written in Go. twix uses both cryptographic hashing and perceptual hashing to find exact duplicates and visually similar images in your directories.

## Installation

```bash
go install github.com/PlotoZypresse/twix@latest
```

Make sure your `GOPATH` is in your `PATH` environment variable to use the installed binary.

## Usage

```bash
twix [mode] [distance] <folder-path>
```

### Arguments

- `[mode]` - Optional operation mode (default: `-h`)
- `[distance]` - Optional for setting distance threshold during pHash comparison (default: `5`)
- `<folder-path>` - Directory to scan for duplicate images

### Flags

- `-h` - Hash mode only (default): Uses SHA256 hashing to find exact duplicates
- `-p` - Perceptual hash mode only: Uses perceptual hashing to find visually similar images
- `-hp` - Both modes (currently displays "TODO" message)
- `-distance=5` - Distance threshold for pHash comparison

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
$ twix -h photos 
Duplicate images found at: /photos/vacation/beach1.jpg and /photos/backup/beach-copy.jpg
Finding duplicates took: 1.234s

$ twix -p -distance=10 photos 
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

If the perceptual mode returns a lot of duplicates try lowering the value. If no or few duplicates are returned try a higher value. The optimal value depends on the images.

## Future
The pHash comparison currently runs in $O(n^2)$ time. For large amount of images this option can take quite some time. Will eb worked on in the future.
The combined mode of hash and pHash mode is not implemented yet.

## Misc
Always check images before deleting