package ErrorHandler

import "log"

func Handle(error error) {
	if error != nil {
		log.Fatal(error)
	}
}
