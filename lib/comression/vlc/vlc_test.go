package vlc

import (
	"testing"

	"archiver/lib/comression/vlc/table/shannon_fano"
)

//func Test_encodeBin(t *testing.T) {
//	tests := []struct {
//		name string
//		str  string
//		want string
//	}{
//		{
//			name: "base test",
//			str:  "!ted",
//			want: "001000100110100101",
//		},
//	}
//	for _, tt := range tests {
//		t.Run(tt.name, func(t *testing.T) {
//			if got := encodeBin(tt.str); got != tt.want {
//				t.Errorf("encodeBin() = %v, want %v", got, tt.want)
//			}
//		})
//	}
//}

func Test_encodeContent(t *testing.T) {
	tests := []struct {
		name    string
		text    string
		wantErr bool
	}{
		{
			name: "base test",
			text: "abbbcc",
		},
		{
			name: "text with whitespace",
			text: "abbb cc",
		},
		{
			name: "text",
			text: "code",
		},
		{
			name: "base test",
			text: "Although the details of our Code of Conduct have been adjusted over time our goals have not We want the Go community to be as inclusive welcoming helpful and respectful as possible If you want to use or discuss Go we want you to feel welcome here The community is large enough that instead of assuming everyone knows what is expected of them our Code of Conduct serves as an agreement setting explicit expectations for our behavior in both online and offline interactions If we dont live up to the agreement people can point that out and we can correct our behavior",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tblGen := shannon_fano.NewGenerator()
			encoderDecoder := New(tblGen)
			encodedData := encoderDecoder.Encode(tt.text)

			data, err := encoderDecoder.Decode(encodedData)
			if (err != nil) != tt.wantErr {
				t.Errorf("Decode() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if data != tt.text {
				t.Errorf("Decode() got = %v, want %v", data, tt.text)
			}
		})
	}
}
