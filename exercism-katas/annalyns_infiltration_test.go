package annalyn

import "testing"

func TestCanFastAttack(t *testing.T) {
	t.Run("the knight is asleep", func(t *testing.T) {
		got := CanFastAttack(false)
		want := true

		assertError(t, got, want)
	})

	t.Run("the knight is awake", func(t *testing.T) {
		got := CanFastAttack(true)
		want := false

		assertError(t, got, want)
	})
}

func TestCanSpy(t *testing.T) {
	t.Run("someone is awake", func(t *testing.T) {
		got := CanSpy(false, true, false)
		want := true

		assertError(t, got, want)
	})

	t.Run("nobody is awake", func(t *testing.T) {
		got := CanSpy(false, false, false)
		want := false

		assertError(t, got, want)
	})

	t.Run("everyone is awake", func(t *testing.T) {
		got := CanSpy(true, true, true)
		want := true

		assertError(t, got, want)
	})
}

func TestCanSignal(t *testing.T) {
	t.Run("prisoner is awake, archer is sleeping", func(t *testing.T) {
		got := CanSignalPrisoner(false, true)
		want := true

		assertError(t, got, want)
	})

	t.Run("archer is awake, prisoner is sleeping", func(t *testing.T) {
		got := CanSignalPrisoner(true, false)
		want := false

		assertError(t, got, want)
	})

	t.Run("both are awake", func(t *testing.T) {
		got := CanSignalPrisoner(true, true)
		want := false

		assertError(t, got, want)
	})

	t.Run("both are asleep", func(t *testing.T) {
		got := CanSignalPrisoner(false, false)
		want := false

		assertError(t, got, want)
	})
}

func TestCanFreePrisoner(t *testing.T) {
	t.Run("pet dog absent, archer and knight sleeping, prisoner awake", func(t *testing.T) {
		got := CanFreePrisoner(false, false, true, false)
		want := true

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("knight and archer awake, prisoner awake, no dog", func(t *testing.T) {
		got := CanFreePrisoner(true, true, true, false)
		want := false

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("knight and archer awake, prisoner awake, pet dog included", func(t *testing.T) {
		got := CanFreePrisoner(true, true, true, true)
		want := false

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("knight is sleeping, archer awake, prisoner awake, pet dog included", func(t *testing.T) {
		got := CanFreePrisoner(false, true, true, true)
		want := false

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("knight is awake, archer sleeping, prisoner awake, pet dog included", func(t *testing.T) {
		got := CanFreePrisoner(true, false, true, true)
		want := true

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})

	t.Run("all characters are sleeping, dog included", func(t *testing.T) {
		got := CanFreePrisoner(false, false, false, true)
		want := true

		if got != want {
			t.Errorf("expected %v but got %v", want, got)
		}
	})
}

func assertError(t testing.TB, got, want bool) {
	t.Helper()
	if got != want {
		t.Errorf("expected %v but got %v", want, got)
	}
}
