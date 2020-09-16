package redis

import (
	"testing"
)

func TestRConnect_Set(t *testing.T) {
	r, _ := DefaultClient("localhost:6379")
	defer r.Close()
	type args struct {
		key   string
		value string
	}
	tests := []struct {
		name    string
		r       RConnect
		args    args
		wantErr bool
	}{
		{
			name: "Set",
			args: args{
				key:   "MyTest",
				value: "My values",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RConnect{}
			if err := r.Set(tt.args.key, tt.args.value); (err != nil) != tt.wantErr {
				t.Errorf("RConnect.Set() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRConnect_Get(t *testing.T) {
	r, _ := DefaultClient("localhost:6379")
	defer r.Close()
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		r       RConnect
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "Get",
			args: args{
				key: "MyTest",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RConnect{}
			got, err := r.Get(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RConnect.Get() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RConnect.Get() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRConnect_Del(t *testing.T) {
	r, _ := DefaultClient("localhost:6379")
	defer r.Close()
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		r       RConnect
		args    args
		wantErr bool
	}{
		{
			name: "DELETE",
			args: args{
				key: "MyTest",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RConnect{}
			if err := r.Del(tt.args.key); (err != nil) != tt.wantErr {
				t.Errorf("RConnect.Delete() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRConnect_Expire(t *testing.T) {
	r, _ := DefaultClient("localhost:6379")
	defer r.Close()
	type args struct {
		key string
		ttl int
	}
	tests := []struct {
		name    string
		r       RConnect
		args    args
		wantErr bool
	}{
		{
			name: "Expire",
			args: args{
				key: "MyTest",
				ttl: 25,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RConnect{}
			if err := r.Expire(tt.args.key, tt.args.ttl); (err != nil) != tt.wantErr {
				t.Errorf("RConnect.Expire() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestRConnect_TTL(t *testing.T) {
	r, _ := DefaultClient("localhost:6379")
	defer r.Close()
	type args struct {
		key string
	}
	tests := []struct {
		name    string
		r       RConnect
		args    args
		want    int
		wantErr bool
	}{
		{
			name: "ttl",
			args: args{
				key: "MyTest",
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r := RConnect{}
			got, err := r.TTL(tt.args.key)
			if (err != nil) != tt.wantErr {
				t.Errorf("RConnect.TTL() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("RConnect.TTL() = %v, want %v", got, tt.want)
			}
		})
	}
}
