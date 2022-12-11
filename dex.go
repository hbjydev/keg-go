package libkeg

import (
	"encoding/csv"
	"errors"
	"strconv"
	"strings"
	"time"
)

var (
  ErrNodeNotEnoughFields = errors.New("dex tsv entry does not have 3 parts")
  ErrNodeIdNotInt = errors.New("dex node id is not an int")
  ErrInvalidTimestamp = errors.New("timestamp must be formatted as 2006-01-02 15:04:05Z")
)

// Node represents a KEG node of any type.
type Node struct {
	// Id is the incrementing ID of the KEG node.
	Id int `yaml:"id" json:"id"`

	// Created is the time the KEG node was created.
	Created KegTimestamp `yaml:"created" json:"created"`

	// Title is the string title of the KEG node.
	Title string `yaml:"title" json:"title"`
}

// NodeFromTsvSlice converts a slice formatted as [id, created, title]
// info a Node, returning an error if the record is invalid.
//
//   node, err := libkeg.NodeFromTsvSlice([]string{"1", "2022-12-11 22:33:00Z", "Hello, world!"})
func NodeFromTsvSlice(data []string) (Node, error) {
  if len(data) != 3 {
    return Node{}, ErrNodeNotEnoughFields
  }

  id, err := strconv.Atoi(data[0]) // convert id to an int
  if err != nil {
    return Node{}, ErrNodeIdNotInt
  }

  cat, err := time.Parse(KegTimestampLayout, data[1])
  if err != nil {
    return Node{}, ErrInvalidTimestamp
  }

  node := Node{
    Id: id,
    Created: KegTimestamp{
      Time: cat,
    },
    Title: data[2],
  }

  return node, nil
}

// NodeMap is a map of nodes, keyed by the ID of the node in the KEG.
type NodeMap map[int]Node

// Dex is an index of nodes
type Dex struct {
  // Nodes is a slice containing all of the [Node] entries in the KEG.
	Nodes []Node
}

// DexFromNodesTsv converts a TSV represented in a string into a Dex
// struct, using the format "(id)\t(created)\t(title)".
//
//   dex, err := libkeg.DexFromNodesTsv("1\t2022-12-11 22:32:00Z\tHello, world!")
func DexFromNodesTsv(tsv string) (Dex, error) {
  // Create CSV reader from TSV string and set delimiter to \t.
  r := strings.NewReader(tsv)
  cr := csv.NewReader(r)
  cr.Comma = '\t'
  cr.LazyQuotes = true

  // Read all TSV entries
  records, err := cr.ReadAll()
  if err != nil {
    return Dex{}, err
  }

  // Loop through entries and convert to a KEG Node struct, returning an
  // error if one occurs.
  nodes := []Node{}
  for _, v := range records {
    node, err := NodeFromTsvSlice(v)
    if err != nil {
      return Dex{}, err
    }
    nodes = append(nodes, node)
  }

  dex := Dex{Nodes: nodes}
  return dex, nil
}
