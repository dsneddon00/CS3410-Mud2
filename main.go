package main

import (
  "database/sql"
  "fmt"
  _ "github.com/mattn/go-sqlite3"
  "log"
  "os"
  "bufio"
  "strings"
)

var allCommands = make(map[string]func(string))
var zones = make(map[int]*Zone)
var rooms = make(map[int]*Room)
var directions = make(map[string]int)
var player = Player{}
var db *sql.DB

type Zone struct {
    ID    int
    Name  string
    Rooms []*Room
}

type Room struct {
    ID          int
    Zone        *Zone
    Name        string
    Description string
    Exits       [6]Exit
}

type Exit struct {
    To          *Room
    Description string
    Direction string
}

type Player struct {
  Room *Room
}

func commandLoop() error {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		line := scanner.Text()
		doCommand(line)
	}
	if err := scanner.Err(); err != nil {
		return fmt.Errorf("in main command loop: %v", err)
	}

	return nil
}

func addCommand(cmd string, f func(string)) {
	for i := range cmd {
		if i == 0 {
			continue
		}
		prefix := cmd[:i]
		allCommands[prefix] = f
	}
	allCommands[cmd] = f
}

func initCommands() {
	addCommand("smile", cmdSmile)
	addCommand("south", cmdSouth)
  addCommand("north", cmdNorth)
  addCommand("east", cmdEast)
  addCommand("west", cmdWest)
  addCommand("tiphat", cmdTipHat)
  addCommand("look", cmdLook)
  addCommand("up", cmdUp)
  addCommand("down", cmdDown)
  addCommand("recall", cmdRecall)
}

func doCommand(cmd string) error {
  words := strings.Fields(cmd)
	if len(words) == 0 {
		return nil
	} else if len(words) == 2 {
		if words[0] == "look" || words[0] == "loo" || words[0] == "lo" || words[0] == "l" {
      // ToLower for userproof because users are ducking idiots
			if f, exists := allCommands[strings.ToLower(words[0])]; exists {
				f(words[1])
				return nil
			}
		}
	}

	if f, exists := allCommands[strings.ToLower(words[0])]; exists {
		f("")
	} else {
		fmt.Printf("Huh?\n")
	}
	return nil
}

// directional commands

func cmdNorth(s string) {
  if len(player.Room.Exits[0].Description) == 0 {
    fmt.Printf("You can't go north from here, you're like an unatractive security guard giving me a patdown at an airport. \n")
  } else {
    // .To is a pointer to another room, man I hate go
    player.Room = player.Room.Exits[0].To
    fmt.Println("You proceed north.")
    fmt.Println(player.Room.Description)
    fmt.Printf("\n")

  }
}

func cmdSouth(s string) {
  if len(player.Room.Exits[2].Description) == 0 {
    fmt.Printf("You can't go south from here, I bet I could write a book about how much you don't know... \n")
  } else {
    // .To is a pointer to another room, man I hate go
    player.Room = player.Room.Exits[2].To
    fmt.Println("You proceed south.")
    fmt.Println(player.Room.Description)
    fmt.Printf("\n")

  }
}

func cmdEast(s string) {
  if len(player.Room.Exits[1].Description) == 0 {
    fmt.Printf("You can't go east from here, your wife can't cheat on you because she set her standards so low marrying you... \n")
  } else {
    // .To is a pointer to another room, man I hate go
    player.Room = player.Room.Exits[1].To
    fmt.Println("You proceed east.")
    fmt.Println(player.Room.Description)
    fmt.Printf("\n")

  }
}

func cmdWest(s string) {
  if len(player.Room.Exits[3].Description) == 0 {
    fmt.Printf("You can't go west from here, the middle finger was invented because of you... \n")
  } else {
    // .To is a pointer to another room, man I hate go
    player.Room = player.Room.Exits[3].To
    fmt.Println("You proceed west.")
    fmt.Println(player.Room.Description)
    fmt.Printf("\n")

  }
}

func cmdUp(s string) {
  if len(player.Room.Exits[4].Description) == 0 {
    fmt.Printf("You can't go up from here, I'll never forget the first time we met... But I'll keep trying, \n")
  } else {
    // .To is a pointer to another room, man I hate go
    player.Room = player.Room.Exits[4].To
    fmt.Println("You proceed up.")
    fmt.Println(player.Room.Description)
    fmt.Printf("\n")

  }
}

func cmdDown(s string) {
  if len(player.Room.Exits[5].Description) == 0 {
    fmt.Printf("You can't go up from here, idiot... \n")
  } else {
    // .To is a pointer to another room, man I hate go
    player.Room = player.Room.Exits[5].To
    fmt.Println("You proceed down.")
    fmt.Println(player.Room.Description)
    fmt.Printf("\n")

  }
}

// emote commands

func cmdSmile(s string) {
	fmt.Printf("You smile happily.\n")
}

func cmdTipHat(s string) {
  fmt.Printf("You tip your hat politely.\n")
}

// action commands

func cmdLook(s string) {
  // finds player room name and description, uses that to find the exits corresponding to each room
  if len(s) == 0 {
		fmt.Println(player.Room.Name, "\n", player.Room.Description, "\n")
		for i, exit := range player.Room.Exits  {
			if len(exit.Description) > i-i {
				fmt.Println(exit.Direction, exit.Description, "\n")
			}
		}
  } else if s == "north" || s == "nort" || s == "nor" || s == "no" || s == "n" {
		var desc = player.Room.Exits[0].Description
		if len(desc) > 0 {
			fmt.Println(player.Room.Exits[0].Description)
		} else {
			fmt.Printf("There is no where to go in this direction...\n")
		}

	} else if s == "east" || s == "eas" || s == "ea" || s == "e" {
		var desc = player.Room.Exits[1].Description
		if len(desc) > 0 {
			fmt.Println(player.Room.Exits[1].Description)
		} else {
			fmt.Printf("There is no where to go in this direction...\n")
		}
	} else if s == "south" || s == "sout" || s == "sou" || s == "so" || s == "s" {
		var desc = player.Room.Exits[2].Description
		if len(desc) > 0 {
			fmt.Println(player.Room.Exits[2].Description)
		} else {
			fmt.Printf("There is no where to go in this direction...\n")
		}
	} else if s == "west" || s == "wes" || s == "we" || s == "w" {
		var desc = player.Room.Exits[3].Description
		if len(desc) > 0 {
			fmt.Println(player.Room.Exits[3].Description)
		} else {
			fmt.Printf("There is no where to go in this direction...\n")
		}
	} else if s == "up" || s == "u" {
		var desc = player.Room.Exits[4].Description
		if len(desc) > 0 {
			fmt.Println(player.Room.Exits[4].Description)
		} else {
			fmt.Printf("There is no where to go in this direction...\n")
		}
	} else if s == "down" || s == "dow" || s == "do" || s == "d" {
		var desc = player.Room.Exits[5].Description
		if len(desc) > 0 {
			fmt.Println(player.Room.Exits[5].Description)
		} else {
			fmt.Printf("There is no where to go in this direction...\n")
		}
	}
}

// Mud Part 2

// Objectives 1 and 2

func readZones(stmt *sql.Stmt) error {
  // new iteration of readZones

  // allocate the Query
  rows, err := stmt.Query()
  // error checker
  if err != nil {
    log.Fatalf("zone query: %v", err)
  }
  defer rows.Close()
  // reads each of the zones
  for rows.Next() {
    var id int
    var name string
    var rooms []*Room
    // the actual scan itself + an error checker
    if err := rows.Scan(&id, &name); err != nil {
      log.Fatalf("reading zones: %v", err)
    }
    // creates the zone
    var zone = Zone{id, name, rooms}
    zones[id] = &zone
  }
  // final error checker
  if err := rows.Err(); err != nil {
    log.Fatal(err)
  }
  return nil

  /*
  // creating the path to the database
  path := "world.db"

  options :=
    "?" + "_busy_timeout=10000" +
      "&" + "_foreign_keys=ON" +
      "&" + "_journey_mode=WAL" +
      "&" + "_synchronous=NORMAL"

  // launching the command to open the file
  db, err := sql.Open("sqlite3", path + options)
  // if we have trouble opening the database we launch an error
  if err != nil {
    log.Fatalf("opening database: %v", err)
  }
  defer db.Close()

  // Launches a query
  rows, err := db.Query(`SELECT * FROM zones`)
  // if we have trouble with the query we launch an error
  if err != nil {
    log.Fatalf("zone query: %v", err)
  }
  defer rows.Close()

  // creates a map (kinda like a dictionary) of zones
  var zones = make(map[int]*Zone)
  // loops and appends to zone
  for rows.Next() {
    var id int
    var name string
    var rooms []*Room
    if err := rows.Scan(&id, &name); err != nil {
      log.Fatalf("reading zones: %v", err)
    }
    // creates the Zone object, pushing the ID, name, and room into a zone object
    var zone = Zone{id, name, rooms}
    zones[id] = &zone
    fmt.Printf("id::%d name: %s\n", id, name)
  }
  */
}

// objective 4

func readRooms(stmt *sql.Stmt) (map[int]*Room, error) {
  rows, err := stmt.Query()
  if(err != nil) {
    log.Fatalf("room query: %v", err)
  }
  defer rows.Close()

  var rooms = make(map[int]*Room)

  for rows.Next() {
    var id int
    var zone_id int
    var name string
    var description string
    var exits [6]Exit
    if err := rows.Scan(&id, &zone_id, &name, &description); err != nil {
      log.Fatalf("reading rooms: %v", err)
    }
    var room = Room{id, zones[zone_id], name, description, exits}
    rooms[id] = &room
    zones[zone_id].Rooms = append(zones[zone_id].Rooms, &room)
  }
  if err := rows.Err(); err != nil {
    log.Fatal(err)
  }
  return rooms, nil
}

func readExits(stmt *sql.Stmt) error {
  // open the query
  rows, err := stmt.Query()
  // error checker
  if(err != nil) {
    log.Fatalf("exit query: %v", err)
  }
  defer rows.Close()

  // loops through, runs through the scan
  for rows.Next() {
    var from_room_id int
    var to_room_id int
    var direction string
    var description string
    // scans the rows using pointers
    if err := rows.Scan(&from_room_id, &to_room_id, &direction, &description); err != nil {
      log.Fatalf("reading exits: %v", err)
    }
    // grabs the exit
    var exit = Exit{rooms[to_room_id], description, direction}
    // goes to the from room id, grabs the direction from the array of directions, then sets all that to exit
    // tldr seting the corresponding room to the corresponding exit
    rooms[from_room_id].Exits[directions[direction]] = exit
  }
  if err := rows.Err(); err != nil {
    log.Fatal(err)
  }
  player.Room = rooms[3001]
  fmt.Println(player.Room.Name, "\n", player.Room.Description)
  return nil
}

func main() {
  // allocating directions
  directions["n"] = 0
  directions["e"] = 1
  directions["s"] = 2
  directions["w"] = 3
  directions["u"] = 4
  directions["d"] = 5

  log.SetFlags(log.Ltime | log.Lshortfile)

  // put the world.db path in
  path := "world.db"

  // allocates the options
  options :=
    "?" + "_busy_timeout=10000" +
      "&" + "_foreign_keys=ON" +
      "&" + "_journey_mode=WAL" +
      "&" + "_synchronous=NORMAL"

  // launching the command to open the file
  db, err := sql.Open("sqlite3", path + options)
  // if we have trouble opening the database we launch an error
  if err != nil {
    log.Fatalf("opening database: %v", err)
  }
  defer db.Close()

  // Transaction Section
  tx, err := db.Begin()
  if(err != nil) {
    log.Fatalf("begin zone read transaction: %v", err)
  }
  stmt, err := tx.Prepare(`SELECT * FROM zones`)
  if(err != nil) {
    log.Fatalf("prepare zone read transaction: %v", err)
  }
  defer stmt.Close()

  err = readZones(stmt)
  // commits and rollbacks
  if(err != nil) {
    // if it fails (doesn't return nil) we rollback the transaction
    tx.Rollback()
  } else {
    // nil has been returned, we shall now commit
    tx.Commit()
  }

  tx, err = db.Begin()
  if(err != nil) {
    log.Fatalf("begin room read transaction: %v", err)
  }
  stmt, err = tx.Prepare(`SELECT * FROM rooms`)
  if err != nil {
    log.Fatalf("prepare room read transaction: %v", err)
  }

  defer stmt.Close()
  rooms, err = readRooms(stmt)
  if(err != nil) {
    tx.Rollback()
  } else {
    tx.Commit()
  }


  // Exit Read transaction
  tx, err = db.Begin()
  if(err != nil) {
    log.Fatalf("begin exit read transaction: %v", err)
  }
  stmt, err = tx.Prepare(`SELECT * FROM exits`)
  if err != nil {
    log.Fatalf("prepare exit read transaction: %v", err)
  }
  defer stmt.Close()
  err = readExits(stmt)
  if(err != nil) {
    tx.Rollback()
  } else {
    tx.Commit()
  }

  initCommands()
  if err := commandLoop(); err != nil {
    log.Fatalf("%v", err)
  }

}

func cmdRecall(s string) {
  player.Room = rooms[3001]
  fmt.Printf("You have returned to the Temple of Midgaaurd. \n")
}


/*
From all my viewers who were kind enough to understand I couldn't stream tonight because of this assignment, I thank you.
You all are a bunch of chads.
Oh and Curtis, if you're watching, you're a chad too.

if (assignmentTurnedIn = true)
{
      jexGrade == 100;
}
*/
