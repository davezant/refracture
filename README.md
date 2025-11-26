# Refracture

**Version:** 0.2.0
**Author:** Davezant
**License:** Unlicense

---

## Overview

**Refracture** is a command-line tool designed to help Godot developers clean up messy projects. It automatically refactors and organizes your Godot project files, placing them in a clean, structured directory, while also handling special folders and performing string replacements as needed.

Perfect for lazy or busy developers who want a neat, export-ready project structure without manually sorting files.

---

## Features

* Refactors Godot project files into a clean structure.
* Automatically copies and organizes project files.
* Optionally includes special folders such as `addons`.
* Performs string replacement in project files to adapt paths after restructuring.
* Exports the cleaned project to a designated folder.
* Works recursively through project directories.

---

## Installation

1. Clone the repository:

```bash
git clone https://github.com/yourusername/refracture.git
```

2. Build the executable:

```bash
cd refracture
go build -o refracture ./cmd
```

3. Add `refracture` to your system PATH or run it directly from the build folder.

---

## Usage

```bash
refracture --src <path-to-your-godot-project> --out <path-to-output-folder> [--addons]
```

### Flags

| Flag       | Shorthand | Description                                                   |
| ---------- | --------- | ------------------------------------------------------------- |
| `--src`    | `-s`      | Path to the Godot project to refactor (required)              |
| `--out`    | `-o`      | Path where the refactored project will be exported (required) |
| `--addons` | `-a`      | Include `addons` folder in the refactored project (optional)  |

### Example

```bash
refracture --src ./MyGodotGame --out ./CleanedGame --addons
```

---

## How It Works

1. **Check Project:** Verifies that a `project.godot` file exists in the source path.
2. **Scan Files:** Recursively collects all project files.
3. **Designate & Organize:** Creates a structured folder layout.
4. **Copy Files:** Copies files into their designated locations.
5. **Copy Special Folders:** Optionally copies `addons` or other special folders.
6. **Replace Strings:** Updates file references to match the new folder structure.
7. **Finish:** Outputs a fully refactored project in the specified destination folder.

---

## Contribution

Contributions are welcome! Please fork the repository and open a pull request with your improvements.

---

## License

This is free and unencumbered software released into the public domain under [The Unlicense](https://unlicense.org/).

You can copy, modify, publish, use, compile, sell, or distribute this software, either in source code form or as a compiled binary, for any purpose, commercial or non-commercial, and by any means.
