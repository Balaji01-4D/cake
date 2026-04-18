# Plans

## Version Goals

### **v1 — Core Experience (MVP)**
Focus: Make it immediately useful and better than existing tools in daily workflow.

**Key Features:**
- Parse and display targets from the current directory's `Makefile` (supports `Makefile`, `makefile`, `GNUmakefile`)
- Real-time fuzzy search (type to filter targets instantly)
- Always-visible **split layout**:
  - Left: Scrollable list of targets with fuzzy matching
  - Right: Live preview of the recipe (exact commands) for the currently selected/hovered target
- Run selected target with original `make` (full compatibility)
- Live output view when running `make` (with option to go back)
- Clean, modern styling with Lip Gloss
- Keyboard-first + basic mouse support
- `q` to quit, sensible defaults, zero configuration

**Goal for v1:**  
A fast, focused tool that solves the most common pain point — “What does this target actually run?” — with the best hover preview experience among Makefile TUIs.

### **v2 — Recursive & Power User Features**
Focus: Make `cake` shine in real-world monorepos and complex projects.

**Planned Features:**
- Automatic detection of Makefiles in subdirectories (1–2 levels deep by default)
- Support for running targets from subfolders using `make -C <subdir> <target>`
- Option to select a specific Makefile with `-f` flag support
- Grouped or prefixed targets (e.g., `api/build`, `frontend/test`)
- Dry-run mode (`make -n`)
- Copy recipe to clipboard (`c` key)
- Help screen with all keybindings
- Config file support (optional themes, default directory depth, etc.)
- Better error handling and graceful fallback when no Makefile is found
- Performance improvements for large Makefiles

**Stretch goals for v2:**
- Simple tree navigation for projects with many sub-Makefiles
- Variable inspection (show expanded values)
- Integration with `just` or other task runners (future-proofing)
