# defer-judgement

> `defer judgement()` — my Go learning repo. The code is a work in progress; the lessons are the point.

A personal log of me learning Go through algorithms, scripts, and whatever else seems worth poking at. The name is intentional: don't judge the code yet, judge whether I'm learning something.

## What's actually here

- [`algorithms/`](./algorithms) — data structures & algorithms practice, one folder per problem. Each has its own README explaining the approach, complexity, and what I learned.

That's it for now. Other folders (e.g. `scripts/`, `cmd/`) will appear when there's something real to put in them, not before.

## Running things

Each algorithm folder is self-contained and testable:

```bash
# run all tests
go test ./algorithms/...

# run a specific problem
go test ./algorithms/<problem-name>/
```

## Layout convention

Every algorithm lives in its own folder under `algorithms/`, named in `kebab-case`:

```
algorithms/
└── two-sum/
    ├── README.md        ← problem, approach, complexity, lessons
    ├── solution.go
    └── solution_test.go
```

The per-folder README is the important part — it's what makes this navigable later. See [`docs/algorithm-template.md`](./docs/algorithm-template.md) for the template I copy each time.

## Index

The full catalogue of solutions, grouped by topic, lives in [`algorithms/README.md`](./algorithms/README.md).

## Learning notes

- **Per-problem lessons** live inside each algorithm's own README.
- **Cross-cutting notes** — Go idioms I keep relearning, slice/goroutine gotchas, testing patterns I'm trying — live in the [Wiki](../../wiki).
- **Session journal** — three dated lines about what clicked — lives in [`LOG.md`](./LOG.md).

## Status

Personal, unstable, not for production, not accepting contributions. If something here is useful to you, take it. If something here is wrong, I probably haven't caught it yet — open an issue if you feel like it. 🔥

## License

[MIT](./LICENSE).
