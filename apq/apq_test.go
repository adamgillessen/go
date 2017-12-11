package apq

import "testing"

func TestEnqueueDequeue(t *testing.T) {
	pq := New()
	pq.Enqueue("enqueued first", 1)
	pq.Enqueue("enqueued second", 0)

	if l := pq.Len(); l != 2 {
		t.Fatalf("Underlying heap has wrong length, got %d want 2", l)
	}

	if s := pq.Dequeue().(string); s != "enqueued second" {
		t.Fatalf("Dequeue() = %q, want 'enqueued second'", s)
	}

	if s := pq.Dequeue().(string); s != "enqueued first" {
		t.Fatalf("Dequeue() = %q, want 'enqueued first'", s)
	}

}

func TestUpdatePriority(t *testing.T) {
	pq := New()
	pq.Enqueue("a", 0)
	pq.Enqueue("b", 2)
	pq.Enqueue("c", 4)
	pq.Enqueue("d", 6)

	if err := pq.UpdatePriority("d", 1); err != nil {
		t.Fatalf("UpdatePriority('d', 1) = %v, want <nil>", err)
	}

	if err := pq.UpdatePriority("c", 3); err != nil {
		t.Fatalf("UpdatePriority('c', 3) = %v, want <nil>", err)
	}

	if err := pq.UpdatePriority("b", 5); err != nil {
		t.Fatalf("UpdatePriority('b', 5) = %v, want <nil>", err)
	}

	if err := pq.UpdatePriority("a", 7); err != nil {
		t.Fatalf("UpdatePriority('a', 7) = %v, want <nil>", err)
	}

	for i, x := range []string{"d", "c", "b", "a"} {
		if s := pq.Dequeue().(string); s != x {
			t.Errorf("Dequeue() = %q, want %q (the %dth item)", s, x, i)
		}
	}
}
