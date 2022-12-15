package main

import (
	"context"
	"fmt"
	"go.temporal.io/sdk/client"
	"hello-world-temporal/app"
	"log"
)

func main() {
	// Create the client object just once per process
	c, err := client.Dial(client.Options{})
	if err != nil {
		log.Fatalln("unable to create Temporal client", err)
	}
	defer c.Close()

	options := client.StartWorkflowOptions{
		ID:        "greeting-workflow",
		TaskQueue: app.GreetingTaskQueue,
	}

	// Start the workflow
	name := "John"
	we, err := c.ExecuteWorkflow(context.Background(), options, app.GreetingWorkflow, name)
	if err != nil {
		log.Fatalln("unable to complete Workflow result", err)
	}

	var greeting string
	err = we.Get(context.Background(), &greeting)
	if err != nil {
		log.Fatalln("unable to get Workflow result", err)
	}

	printResults(greeting, we.GetID(), we.GetRunID())
}

func printResults(greeting string, workflowID, runID string) {
	fmt.Printf("\nWorkflowID: %s RunID: %s\n", workflowID, runID)
	fmt.Printf("\n%s\n\n", greeting)
}
