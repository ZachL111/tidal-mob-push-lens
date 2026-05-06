# Review Journal

The cases below are the review handles I would use before changing the implementation.

The local checks classify each case as `ship`, `watch`, or `hold`. That gives the project a small review vocabulary that matches its mobile workflows focus without claiming live deployment or external usage.

## Cases

- `baseline`: `form pressure`, score 191, lane `ship`
- `stress`: `sync drift`, score 143, lane `ship`
- `edge`: `local state`, score 229, lane `ship`
- `recovery`: `conflict cost`, score 176, lane `ship`
- `stale`: `form pressure`, score 240, lane `ship`

## Note

A future change should add new cases before it changes the scoring rule.
