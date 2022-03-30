package main

import (
  "fmt"
  "math/rand"
  "time"
)

func main() {
  rand.Seed(time.Now().UnixNano())
  isHeistOn := true
  eludedGuards := rand.Intn(100)

  if eludedGuards >= 50 {
    fmt.Println("Looks like you've managed to make it past the guards. Good job, but remember, this is the first step.")
  } else {
    fmt.Println("Plan a better disguise next time?")
  }

  openedVault := rand.Intn(100)

   if isHeistOn && openedVault >= 70 {
    fmt.Println("Grab and GO!")
  } else if isHeistOn {
    isHeistOn = false
    fmt.Print("The vault cant't be opened.")
  }

  leftSafely := rand.Intn(5)

  if isHeistOn {
    switch leftSafely {
      case 0:
        isHeistOn = false
        fmt.Println("The heist was a bust! You got caught on the way out.")
      case 1:
        isHeistOn = false
        fmt.Println("Turns out the vault doesnt open from the inside...")
      case 2:
        isHeistOn = false
        fmt.Println("Jeepers its the cops!")
      case 3:
        isHeistOn = false
        fmt.Println("Headed for the big house...sigh")
      default:
        fmt.Println("Start the getaway car! We are out of here!")
    }
  }

  if isHeistOn {
    amtStolen := 10000 + rand.Intn(1000000)
    fmt.Println("You took", "$", amtStolen, "hey that's not too bad!")
  }
}
