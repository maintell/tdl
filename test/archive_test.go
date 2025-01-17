package test

import (
	"os"
	"path/filepath"

	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

var _ = Describe("Test tdl archive", func() {
	AfterEach(func() {
		// remove zip files
		files, err := filepath.Glob("*.zip")
		Expect(err).To(Succeed())
		for _, file := range files {
			Expect(os.Remove(file)).To(Succeed())
		}
	})

	When("backup", func() {
		It("should success", func() {
			exec(cmd, []string{"backup"}, true)

			files, err := filepath.Glob("tdl-backup-*.zip")
			Expect(err).To(Succeed())
			Expect(len(files)).To(Equal(1))

			_, err = os.Stat(files[0])
			Expect(err).To(Succeed())
		})

		It("should success with custom file name", func() {
			exec(cmd, []string{"backup", "-d", "custom.zip"}, true)

			_, err := os.Stat("custom.zip")
			Expect(err).To(Succeed())
		})
	})

	When("recover", func() {
		It("should success", func() {
			exec(cmd, []string{"backup", "-d", "custom.zip"}, true)

			exec(cmd, []string{"recover", "-f", "custom.zip"}, true)
		})

		It("should fail if do not specify file name", func() {
			exec(cmd, []string{"recover"}, false)
		})

		It("should fail with invalid file name", func() {
			exec(cmd, []string{"recover", "-f", "foo.zip"}, false)
		})
	})
})
