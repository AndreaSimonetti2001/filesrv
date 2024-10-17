package try

import "filesrv/logger"

func Catch(e error) {
	if e != nil {
		logger.Error("%v", e)
	}
}
