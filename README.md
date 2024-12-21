# auxquest
A CLI tool for DM'ing Dungeons and Dragons written in Go. While I still
largely run a pen & paper campaign, I always keep a laptop up and running for
looking up info and keeping notes

Currently a work in progress.

![Screenshot of auxquest in action](https://cdn.discordapp.com/attachments/1169489864597716993/1316158251854987374/auxquest11.png?ex=675a0764&is=6758b5e4&hm=30f1b33dfbed106a09f9249e6440d7db86df01be598fd330d98297cdc9175429&)

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

Want to contribute to auxquest at all?

### Clone the repo:
```bash
git clone https://github.com/jake-abed/auxquest
cd auxquest
```

### Build it:
```bash
go build
```

### Submit a pull request!

If you want to contribute anything at all, please fork the repository
and open up a pull request to the 'main' branch.
