package kj_study

import "kj-study/lib/utils"

// write session to target file
func WriteSession(filename string,session *KjStudySession) {
    var e error=utils.WriteYaml(filename,session)

    if e!=nil {
        panic(e)
    }
}

// get kj study session from file. if did not exist, returns empty session
func GetSession(filename string) KjStudySession {
    var result KjStudySession
    var e error
    result,e=utils.ReadYaml[KjStudySession](filename)

    if e!=nil {
        return result
    }

    return result
}