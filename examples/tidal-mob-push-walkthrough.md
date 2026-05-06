# Tidal Mob Push Lens Walkthrough

This note is the quickest way to read the extra review model in `tidal-mob-push-lens`.

| Case | Focus | Score | Lane |
| --- | --- | ---: | --- |
| baseline | form pressure | 191 | ship |
| stress | sync drift | 143 | ship |
| edge | local state | 229 | ship |
| recovery | conflict cost | 176 | ship |
| stale | form pressure | 240 | ship |

Start with `stale` and `stress`. They create the widest contrast in this repository's fixture set, which makes them better review anchors than the middle cases.

The useful comparison is `form pressure` against `sync drift`, not the raw score alone.
