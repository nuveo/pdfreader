package main

// Example program for pdfread.go

// The program takes a PDF file as argument and writes the MediaBoxes of the
// pages.

import (
  "os";
  "fmt";
  "pdfread";
)

func main() {
  pd := pdfread.Load(os.Args[1]);
  if pd != nil {
    pg := pd.Pages();
    for k := range pg {
      fmt.Printf("%s\n", pd.Att("/MediaBox", pg[k]))
    }
  }
}