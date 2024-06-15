// provides functions to access a "daily word set"

package jisho_ws

// from a target sentence data dir, return a subset sentence data dict
func getDailySet(
    sentenceDataDir string,
) WordSentenceDict {
    var allSentences WordSentenceDict=ReadSentences(sentenceDataDir)


}