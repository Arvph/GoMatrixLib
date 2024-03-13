.PHONY: test test_coverage clean


test:
	go test ./matrix/ -v

test_coverage: 
	@mkdir ./report
	@go test ./... -coverprofile=coverage.out
	@go tool cover -html=coverage.out -o coverage.html
	@mv ./coverage.out ./report/
	@mv ./coverage.html ./report/
	@rm -fr coverage.out
	@rm -fr coverage.html

clean:
	@rm -fr *.out
	@rm -fr *.html
	@rm -fr report