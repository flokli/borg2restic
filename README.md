# borg2restic

A tool to help convert a borg repository to restic.

It assumes the following environment variables to be set:

 - `BORG_REPO` set to the old borg repository
 - `RESTIC_REPOSITORY` set to the restic repository

It is highly recommended to also set `BORG_PASSPHRASE` (or `BORG_PASSCOMMAND`),
and `RESTIC_PASSWORD[_FILE]` to allow non-interactive access to the two repositories.

It will mount the borg repository, collect a list of all snapshots, and will
invoke restic to "back up from the mountpoint".

It does this by changing the working directory of the restic process to the
location of the mounted borg contents (and an optional subpath in there).

Additionally, the path shown in `restic snapshots` can be overridden, so it
doesn't show artifacts from the conversion.
This needs the version of `restic` in `$PATH` to have
https://github.com/restic/restic/pull/3200 applied.

Timestamps are preserved (to the best of my knowledge). It's also possible to
explicitly set the hostname for the `restic backup` command.

Check the `--help` output for more help.
