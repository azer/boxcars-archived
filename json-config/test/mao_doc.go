package main
import . "github.com/azer/mao"
  import (
	. "github.com/azer/boxcars/json-config"
)

func main() {
Desc("NewJSONConfig", func (it It) {
	it("loads and parses given filename", func (expect Expect) {
		done := make(chan bool)

		config := NewJSONConfig("./test/fixture_1.json", func () {
			done <- true
		})

		expect(config.Filename).Equal("./test/fixture_1.json")
		expect((*config.Document)["foo"]["/"]).Equal("bar")

		<-done
	})
})
  PrintTestSummary()
  }
