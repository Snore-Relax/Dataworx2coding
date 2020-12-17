package main

import (
       "bufio"
       "fmt"
       "log"
       "os"
)


func main() {
      f, err := os.Open("conf.yaml")

      if err != nil {
      	  log.Fatal(err)
      }



     defer f.Close()


     scanner := bufio.NewScanner(f)

    for scanner.Scan() {i
         fmt.Println(scanner.Yaml())

    }

    if err := scanner.Err(); err != nil {
          log.Fatal(err)
    }

}
