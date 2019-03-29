package importer

import (
	"io/ioutil"
	"net/url"
	"os"

	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"

	"k8s.io/apimachinery/pkg/api/resource"

	"github.com/pkg/errors"

	"kubevirt.io/containerized-data-importer/pkg/image"
)

type fakeInfoOpRetVal struct {
	imgInfo *image.ImgInfo
	e       error
}

const TestImagesDir = "../../tests/images"

const (
	SmallActualSize  = 1024
	SmallVirtualSize = 1024
)

var (
	fakeSmallImageInfo = image.ImgInfo{Format: "", BackingFile: "", VirtualSize: SmallVirtualSize, ActualSize: SmallActualSize}
	fakeZeroImageInfo  = image.ImgInfo{Format: "", BackingFile: "", VirtualSize: 0, ActualSize: 0}
	fakeInfoRet        = fakeInfoOpRetVal{imgInfo: &fakeSmallImageInfo, e: nil}
)

type fakeQEMUOperations struct {
	e2             error
	e3             error
	ret4           fakeInfoOpRetVal
	e5             error
	e6             error
	resizeQuantity *resource.Quantity
}

type MockDataProvider struct {
	infoResponse     ProcessingPhase
	transferResponse ProcessingPhase
	processResponse  ProcessingPhase
	url              *url.URL
	transferPath     string
	calledPhases     []ProcessingPhase
}

// Info() is called to get initial information about the data
func (m *MockDataProvider) Info() (ProcessingPhase, error) {
	m.calledPhases = append(m.calledPhases, Info)
	if m.infoResponse == Error {
		return Error, errors.New("Info errored")
	}
	return m.infoResponse, nil
}

// Transfer() is called to transfer the data from the source to a temporary location.
func (m *MockDataProvider) Transfer(path string) (ProcessingPhase, error) {
	m.calledPhases = append(m.calledPhases, Transfer)
	m.transferPath = path
	if m.transferResponse == Error {
		return Error, errors.New("Transfer errored")
	}
	return m.transferResponse, nil
}

// Process() is called to do any special processing before giving the url to the data back to the processor
func (m *MockDataProvider) Process() (ProcessingPhase, error) {
	m.calledPhases = append(m.calledPhases, Process)
	if m.processResponse == Error {
		return Error, errors.New("Process errored")
	}
	return m.processResponse, nil
}

// Geturl returns the url that the data processor can use when converting the data.
func (m *MockDataProvider) GetURL() *url.URL {
	return m.url
}

// Close closes any readers or other open resources.
func (m *MockDataProvider) Close() error {
	return nil
}

var _ = Describe("Data Processor", func() {
	It("should call the right phases based on the responses from the provider, Transfer should pass the scratch data dir as a path", func() {
		mdp := &MockDataProvider{
			infoResponse:     Transfer,
			transferResponse: Process,
			processResponse:  Complete,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "1G")
		err := dp.ProcessData()
		Expect(err).ToNot(HaveOccurred())
		Expect(3).To(Equal(len(mdp.calledPhases)))
		Expect(Info).To(Equal(mdp.calledPhases[0]))
		Expect(Transfer).To(Equal(mdp.calledPhases[1]))
		Expect(Process).To(Equal(mdp.calledPhases[2]))
		Expect("scratchDataDir").To(Equal(mdp.transferPath))
	})

	It("should call the right phases based on the responses from the provider, TransferTarget should pass the data dir as a path", func() {
		mdp := &MockDataProvider{
			infoResponse:     TransferTarget,
			transferResponse: Process,
			processResponse:  Complete,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "1G")
		err := dp.ProcessData()
		Expect(err).ToNot(HaveOccurred())
		Expect(3).To(Equal(len(mdp.calledPhases)))
		Expect(Info).To(Equal(mdp.calledPhases[0]))
		Expect(Transfer).To(Equal(mdp.calledPhases[1]))
		Expect(Process).To(Equal(mdp.calledPhases[2]))
		Expect("dataDir").To(Equal(mdp.transferPath))
	})

	It("should error on Transfer phase", func() {
		mdp := &MockDataProvider{
			infoResponse:     TransferTarget,
			transferResponse: Error,
			processResponse:  Complete,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "1G")
		err := dp.ProcessData()
		Expect(err).To(HaveOccurred())
		Expect(2).To(Equal(len(mdp.calledPhases)))
		Expect(Info).To(Equal(mdp.calledPhases[0]))
		Expect(Transfer).To(Equal(mdp.calledPhases[1]))
	})

	It("should error on Unknown phase", func() {
		mdp := &MockDataProvider{
			infoResponse: ProcessingPhase("invalidphase"),
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "1G")
		err := dp.ProcessData()
		Expect(err).To(HaveOccurred())
		Expect(1).To(Equal(len(mdp.calledPhases)))
		Expect(Info).To(Equal(mdp.calledPhases[0]))
	})

	It("should call Convert after Process phase", func() {
		tmpDir, err := ioutil.TempDir("", "scratch")
		Expect(err).ToNot(HaveOccurred())
		defer os.RemoveAll(tmpDir)

		url, err := url.Parse("http://fakeurl-notreal.fake")
		Expect(err).ToNot(HaveOccurred())
		mdp := &MockDataProvider{
			infoResponse:     TransferTarget,
			transferResponse: Process,
			processResponse:  Convert,
			url:              url,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", tmpDir, "1G")
		qemuOperations := NewFakeQEMUOperations(nil, nil, fakeInfoRet, nil, nil, resource.NewScaledQuantity(int64(1500), 0))
		replaceQEMUOperations(qemuOperations, func() {
			err = dp.ProcessData()
			Expect(err).ToNot(HaveOccurred())
			Expect(3).To(Equal(len(mdp.calledPhases)))
			Expect(Info).To(Equal(mdp.calledPhases[0]))
			Expect(Transfer).To(Equal(mdp.calledPhases[1]))
			Expect(Process).To(Equal(mdp.calledPhases[2]))
			Expect("dataDir").To(Equal(mdp.transferPath))
		})
	})
})

var _ = Describe("Convert", func() {
	It("Should successfully convert and return resize", func() {
		url, err := url.Parse("http://fakeurl-notreal.fake")
		Expect(err).ToNot(HaveOccurred())
		mdp := &MockDataProvider{
			url: url,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "1G")
		qemuOperations := NewFakeQEMUOperations(nil, nil, fakeInfoOpRetVal{&fakeZeroImageInfo, errors.New("Scratch space required, and none found ")}, nil, nil, nil)
		replaceQEMUOperations(qemuOperations, func() {
			nextPhase, err := dp.convert(mdp.GetURL())
			Expect(err).ToNot(HaveOccurred())
			Expect(Resize).To(Equal(nextPhase))
		})
	})

	It("Should fail when validation fails and return Error", func() {
		url, err := url.Parse("http://fakeurl-notreal.fake")
		Expect(err).ToNot(HaveOccurred())
		mdp := &MockDataProvider{
			url: url,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "1G")
		qemuOperations := NewFakeQEMUOperations(nil, nil, fakeInfoOpRetVal{&fakeZeroImageInfo, errors.New("Scratch space required, and none found ")}, errors.New("Validation failure"), nil, nil)
		replaceQEMUOperations(qemuOperations, func() {
			nextPhase, err := dp.convert(mdp.GetURL())
			Expect(err).To(HaveOccurred())
			Expect(Error).To(Equal(nextPhase))
		})
	})

	It("Should fail when conversion fails and return Error", func() {
		url, err := url.Parse("http://fakeurl-notreal.fake")
		Expect(err).ToNot(HaveOccurred())
		mdp := &MockDataProvider{
			url: url,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "1G")
		qemuOperations := NewFakeQEMUOperations(errors.New("Conversion failure"), nil, fakeInfoOpRetVal{&fakeZeroImageInfo, errors.New("Scratch space required, and none found ")}, nil, nil, nil)
		replaceQEMUOperations(qemuOperations, func() {
			nextPhase, err := dp.convert(mdp.GetURL())
			Expect(err).To(HaveOccurred())
			Expect(Error).To(Equal(nextPhase))
		})
	})
})

var _ = Describe("Resize", func() {
	It("Should not resize and return complete, when requestedSize is blank", func() {
		url, err := url.Parse("http://fakeurl-notreal.fake")
		Expect(err).ToNot(HaveOccurred())
		mdp := &MockDataProvider{
			url: url,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "")
		nextPhase, err := dp.resize()
		Expect(err).ToNot(HaveOccurred())
		Expect(Complete).To(Equal(nextPhase))
	})

	It("Should not resize and return complete, when requestedSize is valid, but datadir doesn't exist (block device)", func() {
		url, err := url.Parse("http://fakeurl-notreal.fake")
		Expect(err).ToNot(HaveOccurred())
		mdp := &MockDataProvider{
			url: url,
		}
		dp := NewDataProcessor(mdp, "dest", "dataDir", "scratchDataDir", "1G")
		nextPhase, err := dp.resize()
		Expect(err).ToNot(HaveOccurred())
		Expect(Complete).To(Equal(nextPhase))
	})

	It("Should resize and return complete, when requestedSize is valid, and datadir exists", func() {
		tmpDir, err := ioutil.TempDir("", "data")
		Expect(err).ToNot(HaveOccurred())
		url, err := url.Parse("http://fakeurl-notreal.fake")
		Expect(err).ToNot(HaveOccurred())
		mdp := &MockDataProvider{
			url: url,
		}
		dp := NewDataProcessor(mdp, "dest", tmpDir, "scratchDataDir", "1G")
		qemuOperations := NewFakeQEMUOperations(nil, nil, fakeInfoOpRetVal{&fakeZeroImageInfo, nil}, nil, nil, nil)
		replaceQEMUOperations(qemuOperations, func() {
			nextPhase, err := dp.resize()
			Expect(err).ToNot(HaveOccurred())
			Expect(Complete).To(Equal(nextPhase))
		})
	})

	It("Should not resize and return error, when ResizeImage fails", func() {
		tmpDir, err := ioutil.TempDir("", "data")
		Expect(err).ToNot(HaveOccurred())
		url, err := url.Parse("http://fakeurl-notreal.fake")
		Expect(err).ToNot(HaveOccurred())
		mdp := &MockDataProvider{
			url: url,
		}
		dp := NewDataProcessor(mdp, "dest", tmpDir, "scratchDataDir", "1G")
		qemuOperations := NewQEMUAllErrors()
		replaceQEMUOperations(qemuOperations, func() {
			nextPhase, err := dp.resize()
			Expect(err).To(HaveOccurred())
			Expect(Error).To(Equal(nextPhase))
		})
	})
})

var _ = Describe("ResizeImage", func() {
	//fakeInfoRet has info.VirtualSize=1024
	table.DescribeTable("calling ResizeImage", func(qemuOperations image.QEMUOperations, imageSize string, totalSpace int64, wantErr bool) {
		replaceQEMUOperations(qemuOperations, func() {
			err := ResizeImage("dest", imageSize, totalSpace)
			if !wantErr {
				Expect(err).ToNot(HaveOccurred())
			} else {
				Expect(err).To(HaveOccurred())
			}
		})
	},
		table.Entry("successfully resize to imageSize when imageSize > info.VirtualSize and < totalSize", NewFakeQEMUOperations(nil, nil, fakeInfoRet, nil, nil, resource.NewScaledQuantity(int64(1500), 0)), "1500", int64(2048), false),
		table.Entry("successfully resize to totalSize when imageSize > info.VirtualSize and > totalSize", NewFakeQEMUOperations(nil, nil, fakeInfoRet, nil, nil, resource.NewScaledQuantity(int64(2048), 0)), "2500", int64(2048), false),
		table.Entry("successfully do nothing when imageSize = info.VirtualSize and > totalSize", NewFakeQEMUOperations(nil, nil, fakeInfoRet, nil, nil, resource.NewScaledQuantity(int64(1024), 0)), "1024", int64(1024), false),
		table.Entry("fail to resize to with blank imageSize", NewFakeQEMUOperations(nil, nil, fakeInfoRet, nil, nil, resource.NewScaledQuantity(int64(2048), 0)), "", int64(2048), true),
		table.Entry("fail to resize to with blank imageSize", NewQEMUAllErrors(), "", int64(2048), true),
	)
})

func replaceQEMUOperations(replacement image.QEMUOperations, f func()) {
	orig := qemuOperations
	if replacement != nil {
		qemuOperations = replacement
		defer func() { qemuOperations = orig }()
	}
	f()
}

func NewFakeQEMUOperations(e2, e3 error, ret4 fakeInfoOpRetVal, e5 error, e6 error, targetResize *resource.Quantity) image.QEMUOperations {
	return &fakeQEMUOperations{e2, e3, ret4, e5, e6, targetResize}
}

func (o *fakeQEMUOperations) ConvertToRawStream(*url.URL, string) error {
	return o.e2
}

func (o *fakeQEMUOperations) Validate(*url.URL, int64) error {
	return o.e5
}

func (o *fakeQEMUOperations) Resize(dest string, size resource.Quantity) error {
	if o.resizeQuantity != nil {
		Expect(o.resizeQuantity.Cmp(size)).To(Equal(0))
	}
	return o.e3
}

func (o *fakeQEMUOperations) Info(url *url.URL) (*image.ImgInfo, error) {
	return o.ret4.imgInfo, o.ret4.e
}

func (o *fakeQEMUOperations) CreateBlankImage(dest string, size resource.Quantity) error {
	return o.e6
}

func NewQEMUAllErrors() image.QEMUOperations {
	err := errors.New("qemu should not be called from this test override with replaceQEMUOperations")
	return NewFakeQEMUOperations(err, err, fakeInfoOpRetVal{nil, err}, err, err, nil)
}
