package scheduler

import (
	"testing"
	. "github.com/bborbe/assert"
)

func TestImplements(t *testing.T) {
	scheduler, err := NewScheduler([]string{"a"})
	if err != nil {
		t.Fatal(err)
	}
	var i *Scheduler
	err = AssertThat(scheduler, Implements(i).Message("check type"))
	if err != nil {
		t.Fatal(err)
	}
}

func TestNewSchedulerThrowErrorWithoutNodes(t *testing.T) {
	// nil list
	{
		_, err := NewScheduler(nil)
		if err == nil {
			t.Fatal("error expected")
		}
	}
	// empty list
	{
		_, err := NewScheduler([]string{})
		if err == nil {
			t.Fatal("error expected")
		}

	}
}

func TestNext(t *testing.T) {
	scheduler, err := NewScheduler([]string{"a","b","c"})
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(scheduler.Next(),Is("a"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(scheduler.Next(),Is("b"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(scheduler.Next(),Is("c"))
	if err != nil {
		t.Fatal(err)
	}
	err = AssertThat(scheduler.Next(),Is("a"))
	if err != nil {
		t.Fatal(err)
	}
}
