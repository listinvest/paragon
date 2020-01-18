// Code generated by entc, DO NOT EDIT.

package ent

import (
	"context"
	"log"
	"time"

	"github.com/facebookincubator/ent/dialect/sql"
)

// dsn for the database. In order to run the tests locally, run the following command:
//
//	 ENT_INTEGRATION_ENDPOINT="root:pass@tcp(localhost:3306)/test?parseTime=True" go test -v
//
var dsn string

func ExampleCredential() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the credential's edges.

	// create credential vertex with its edges.
	c := client.Credential.
		Create().
		SetPrincipal("string").
		SetSecret("string").
		SetType("string").
		SetFails(1).
		SaveX(ctx)
	log.Println("credential created:", c)

	// query edges.

	// Output:
}
func ExampleJob() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the job's edges.
	t0 := client.Task.
		Create().
		SetQueueTime(time.Now()).
		SetLastChangedTime(time.Now()).
		SetClaimTime(time.Now()).
		SetExecStartTime(time.Now()).
		SetExecStopTime(time.Now()).
		SetContent("string").
		SetOutput("string").
		SetError("string").
		SetSessionID("string").
		SaveX(ctx)
	log.Println("task created:", t0)
	t1 := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t1)
	j3 := client.Job.
		Create().
		SetName("string").
		SetCreationTime(time.Now()).
		SetContent("string").
		SaveX(ctx)
	log.Println("job created:", j3)

	// create job vertex with its edges.
	j := client.Job.
		Create().
		SetName("string").
		SetCreationTime(time.Now()).
		SetContent("string").
		AddTasks(t0).
		AddTags(t1).
		SetNext(j3).
		SaveX(ctx)
	log.Println("job created:", j)

	// query edges.
	t0, err = j.QueryTasks().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tasks: %v", err)
	}
	log.Println("tasks found:", t0)

	t1, err = j.QueryTags().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tags: %v", err)
	}
	log.Println("tags found:", t1)

	j3, err = j.QueryNext().First(ctx)
	if err != nil {
		log.Fatalf("failed querying next: %v", err)
	}
	log.Println("next found:", j3)

	// Output:
}
func ExampleTag() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the tag's edges.

	// create tag vertex with its edges.
	t := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t)

	// query edges.

	// Output:
}
func ExampleTarget() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the target's edges.
	t0 := client.Task.
		Create().
		SetQueueTime(time.Now()).
		SetLastChangedTime(time.Now()).
		SetClaimTime(time.Now()).
		SetExecStartTime(time.Now()).
		SetExecStopTime(time.Now()).
		SetContent("string").
		SetOutput("string").
		SetError("string").
		SetSessionID("string").
		SaveX(ctx)
	log.Println("task created:", t0)
	t1 := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t1)
	c2 := client.Credential.
		Create().
		SetPrincipal("string").
		SetSecret("string").
		SetType("string").
		SetFails(1).
		SaveX(ctx)
	log.Println("credential created:", c2)

	// create target vertex with its edges.
	t := client.Target.
		Create().
		SetName("string").
		SetPrimaryIP("string").
		SetMachineUUID("string").
		SetPublicIP("string").
		SetPrimaryMAC("string").
		SetHostname("string").
		SetLastSeen(time.Now()).
		AddTasks(t0).
		AddTags(t1).
		AddCredentials(c2).
		SaveX(ctx)
	log.Println("target created:", t)

	// query edges.
	t0, err = t.QueryTasks().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tasks: %v", err)
	}
	log.Println("tasks found:", t0)

	t1, err = t.QueryTags().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tags: %v", err)
	}
	log.Println("tags found:", t1)

	c2, err = t.QueryCredentials().First(ctx)
	if err != nil {
		log.Fatalf("failed querying credentials: %v", err)
	}
	log.Println("credentials found:", c2)

	// Output:
}
func ExampleTask() {
	if dsn == "" {
		return
	}
	ctx := context.Background()
	drv, err := sql.Open("mysql", dsn)
	if err != nil {
		log.Fatalf("failed creating database client: %v", err)
	}
	defer drv.Close()
	client := NewClient(Driver(drv))
	// creating vertices for the task's edges.
	t0 := client.Tag.
		Create().
		SetName("string").
		SaveX(ctx)
	log.Println("tag created:", t0)

	// create task vertex with its edges.
	t := client.Task.
		Create().
		SetQueueTime(time.Now()).
		SetLastChangedTime(time.Now()).
		SetClaimTime(time.Now()).
		SetExecStartTime(time.Now()).
		SetExecStopTime(time.Now()).
		SetContent("string").
		SetOutput("string").
		SetError("string").
		SetSessionID("string").
		AddTags(t0).
		SaveX(ctx)
	log.Println("task created:", t)

	// query edges.
	t0, err = t.QueryTags().First(ctx)
	if err != nil {
		log.Fatalf("failed querying tags: %v", err)
	}
	log.Println("tags found:", t0)

	// Output:
}
