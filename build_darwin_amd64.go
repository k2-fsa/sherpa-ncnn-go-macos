//go:build darwin && amd64 && !ios

package sherpa_onnx

// #cgo LDFLAGS: -L ${SRCDIR}/lib/x86_64-apple-darwin -lsherpa-ncnn-c-api -lsherpa-ncnn-core -lkaldi-native-fbank-core -lncnn -Wl,-rpath,${SRCDIR}/lib/x86_64-apple-darwin
import "C"
