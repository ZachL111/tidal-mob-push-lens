# tidal-mob-push-lens

`tidal-mob-push-lens` is a compact Go repository for mobile workflows, centered on this goal: Create a Go reference implementation for push workflows, centered on event replay, fixture event logs, and golden state snapshots.

## Use Case

This is intentionally local and self-contained so it can be inspected without credentials, services, or seeded history.

## Tidal Mob Push Lens Review Notes

`stale` and `stress` are the cases worth reading first. They show the optimistic and cautious ends of the fixture.

## Highlights

- `fixtures/domain_review.csv` adds cases for form pressure and sync drift.
- `metadata/domain-review.json` records the same cases in structured form.
- `config/review-profile.json` captures the read order and the two review questions.
- `examples/tidal-mob-push-walkthrough.md` walks through the case spread.
- The Go code includes a review path for `form pressure` and `sync drift`.
- `docs/field-notes.md` explains the strongest and weakest cases.

## Code Layout

The repository has two validation layers: the original compact policy fixture and the domain review fixture. They are separate so one can change without hiding failures in the other.

The Go addition stays small enough to inspect in one sitting.

## Run The Check

```powershell
powershell -NoProfile -ExecutionPolicy Bypass -File scripts/verify.ps1
```

## Regression Path

The same command runs the local verification path. The highest-scoring domain case is `stale` at 240, which lands in `ship`. The most cautious case is `stress` at 143, which lands in `ship`.

## Future Work

The fixture set is small enough to audit by hand. The next useful expansion is malformed input coverage, not extra surface area.
