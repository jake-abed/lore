# Lore
A CLI tool for game masters running D&D, Pathfinder, and other tabletop game
campaigns. All written in Go. While I still largely run a pen & paper campaign,
I always keep a laptop up and running for looking up info and keeping notes.

Currently a work in progress, but still functional!

![Screenshot of Lore in action](./static/lore.png)

## Commands

> **help** lists all available commands

> **monsters**
> - **-i <monster name or id>** | Get basic information about a monster from the D&D 5e API.
> - **-f <monster-1> <monster-2>** | Simulate a turn-based fight between two monsters.
> - - Slightly skewed as it does not factor movement or flying into the battles, so extremely agile or flying monsters will have less of an edge than they normal would.

> **npcs** View, search, edit, and create custom NPCs for your campaign (WIP - not finished)
> - **-c** | Create a new NPC!
> - **-v <name>** | View an NPC with the provided name (exact match).
> - **-s <name>** | Lists all NPCs that partially match the provided name.
> - **-e <name>** | Edit an NPC with the provided name (exact match).

## Contributing

Want to contribute to Lore at all?

### Clone the repo:
```bash
git clone https://github.com/jake-abed/lore
cd lore
```

### Build it:
```bash
go build
```

### Submit a pull request!

If you want to contribute anything at all, please fork the repository
and open up a pull request to the 'main' branch.
