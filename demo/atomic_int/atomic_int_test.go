package atomic_int

import "testing"

var a AtomicInt

func TestAtomicInt_Add(t *testing.T) {
	type args struct {
		delta int
	}
	tests := []struct {
		name string
		ai   *AtomicInt
		args args
	}{
	// TODO: Add test cases.
		{"lzx",&a,args{1}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tmp := a
			tt.ai.Add(tt.args.delta)
			if a != tmp+1 {
				t.Errorf("a = %v ,want %v",a,tmp+1)
			}
		})
	}
}

func TestAtomicInt_Load(t *testing.T) {
	tests := []struct {
		name string
		ai   *AtomicInt
		want int
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ai.Load(); got != tt.want {
				t.Errorf("AtomicInt.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtomicInt64_Add(t *testing.T) {
	type args struct {
		delta int64
	}
	tests := []struct {
		name string
		ai   *AtomicInt64
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ai.Add(tt.args.delta)
		})
	}
}

func TestAtomicInt64_Load(t *testing.T) {
	tests := []struct {
		name string
		ai   *AtomicInt64
		want int64
	}{
	// TODO: Add test cases.

	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ai.Load(); got != tt.want {
				t.Errorf("AtomicInt64.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtomicInt32_Add(t *testing.T) {
	type args struct {
		delta int32
	}
	tests := []struct {
		name string
		ai   *AtomicInt32
		args args
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.ai.Add(tt.args.delta)
		})
	}
}

func TestAtomicInt32_Load(t *testing.T) {
	tests := []struct {
		name string
		ai   *AtomicInt32
		want int32
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ai.Load(); got != tt.want {
				t.Errorf("AtomicInt32.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestAtomicBool_Load(t *testing.T) {
	tests := []struct {
		name string
		ab   *AtomicBool
		want bool
	}{
	// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := tt.ab.Load(); got != tt.want {
				t.Errorf("AtomicBool.Load() = %v, want %v", got, tt.want)
			}
		})
	}
}
