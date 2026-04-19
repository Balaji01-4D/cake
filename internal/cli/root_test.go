package cli

import "testing"

func TestDefaultMakeCommandPrefix(t *testing.T) {
	tests := []struct {
		name     string
		makefile string
		want     string
	}{
		{
			name:     "default Makefile",
			makefile: "Makefile",
			want:     "make",
		},
		{
			name:     "custom makefile",
			makefile: "temp.mk",
			want:     `make -f "temp.mk"`,
		},
		{
			name:     "custom makefile with spaces",
			makefile: "build files/temp file.mk",
			want:     `make -f "build files/temp file.mk"`,
		},
		{
			name:     "empty makefile path",
			makefile: "",
			want:     "make",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := defaultMakeCommandPrefix(tt.makefile)
			if got != tt.want {
				t.Fatalf("defaultMakeCommandPrefix(%q) = %q, want %q", tt.makefile, got, tt.want)
			}
		})
	}
}
