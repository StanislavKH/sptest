package speedtest

import (
	"reflect"
	"testing"
)

// All speedcheck results different, fake data to be sure what we have no errors
var (
	TestResponse *Result = &Result{Download: 200, Upload: 200}
)

func TestSpeedTest_StartBenchmark(t *testing.T) {
	tests := []struct {
		name    string
		st      *SpeedTest
		want    *Result
		wantErr bool
	}{
		{
			name:    "no error test",
			st:      &SpeedTest{},
			want:    TestResponse,
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			st := &SpeedTest{}
			got, err := st.StartBenchmark()
			if (err != nil) != tt.wantErr {
				t.Errorf("SpeedTest.StartBenchmark() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(TestResponse, tt.want) {
				t.Errorf("SpeedTest.StartBenchmark() = %v, want %v", got, tt.want)
			}
		})
	}
}
