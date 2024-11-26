# auxquest
A CLI tool for DM'ing Dungeons and Dragons written in Go. While I still
largely run a pen & paper campaign, I always keep a laptop up and running for
looking up info and keeping notes

Currently a work in progress.

## Commands

> **help** lists all available commands

> **--inspect or -i <monster name or id>a**
> Get basic information about a monster from the D&D 5e
> API.

> **--fight or -f <monster-1> <monster-2>
> Simulator a turn-based fight between two monsters.
> Slightly skewed as it does not factor movement or flying
> into the battles, so extremely agile or flying monsters
> will have less of an edge than they normal would.
