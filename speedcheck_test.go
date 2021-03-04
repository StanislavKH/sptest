package speedcheck

import (
	"reflect"
	"testing"

	"github.com/StanislavKH/sptest/fast"
	"github.com/StanislavKH/sptest/speedtest"
)

// All speedcheck results different, fake data to be sure what we have no errors
var (
	TestResponse *Response = &Response{Download: 200, Upload: 200}
)

func TestSpCheck_GetNetworkSpeed(t *testing.T) {
	type fields struct {
		stest *speedtest.SpeedTest
		ftest *fast.SpeedTest
	}
	type args struct {
		op *Options
	}
	tests := []struct {
		name    string
		fields  fields
		args    args
		want    *Response
		wantErr bool
	}{
		{
			name:    "Option test case Netflix Fast.",
			args:    args{op: &Options{CheckType: TypeFast}},
			want:    TestResponse,
			wantErr: false,
		},
		{
			name:    "Option test case Speedtest.",
			args:    args{op: &Options{CheckType: TypeSpeed}},
			want:    TestResponse,
			wantErr: false,
		},
		{
			name:    "Option empty",
			args:    args{op: &Options{}},
			want:    TestResponse,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SpCheck{
				stest: tt.fields.stest,
				ftest: tt.fields.ftest,
			}
			got, err := s.GetNetworkSpeed(tt.args.op)
			if (err != nil) != tt.wantErr {
				t.Errorf("SpCheck.GetNetworkSpeed() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(TestResponse, tt.want) {
				t.Errorf("SpCheck.GetNetworkSpeed() = %v, want %v", got, tt.want)
			}
		})
	}
}

func BenchmarkFast(b *testing.B) {
	b.ReportAllocs()
	sp := New()
	res, err := sp.GetNetworkSpeed(&Options{CheckType: TypeFast})
	_ = res
	if err != nil {
		b.Error(err)
	}
}

func BenchmarkSpeedtest(b *testing.B) {
	b.ReportAllocs()
	sp := New()
	res, err := sp.GetNetworkSpeed(&Options{CheckType: TypeSpeed})
	_ = res
	if err != nil {
		b.Error(err)
	}
}
