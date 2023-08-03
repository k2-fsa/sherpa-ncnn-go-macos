//go:build darwin && arm64 && !ios

package sherpa_ncnn

// #cgo LDFLAGS: -L ${SRCDIR}/lib/aarch64-apple-darwin -lsherpa-ncnn-c-api -lsherpa-ncnn-core -lkaldi-native-fbank-core -lncnn -Wl,-rpath,${SRCDIR}/lib/aarch64-apple-darwin
import "C"
