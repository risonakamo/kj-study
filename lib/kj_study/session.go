// funcs implementing user session system

package kj_study

import (
	"errors"
	jisho_ws "kj-study/lib/jisho/word_sentence"
	"kj-study/lib/utils"
)

// possible states a word sentence pair can be in
type WordSentenceStatus string
const (
    WordSentenceStatus_normal WordSentenceStatus="normal"
    WordSentenceStatus_activeGreen WordSentenceStatus="active-green"
    WordSentenceStatus_activeRed WordSentenceStatus="active-red"
)

// top level kj study session data
type KjStudySession struct {
    WordSentences []WordSentencePair `json:"wordSentences"`

    // data file being used for this session
    Datafile string `json:"datafile"`
}

// status of a word sentence pair. includes the pair information and user selections
// of the pair
type WordSentencePair struct {
    Word string  `json:"word"`
    Sentence string `json:"sentence"`
    Status WordSentenceStatus `json:"status"`
}

// generate a new session from a word sentence file.
// sentences are shuffled
func GenerateNewSession(
    wordSentenceFilesDir string,
    wordSentenceFilename string,
    sentencesPerWordMin int,
    sentencesPerWordMax int,
) KjStudySession {
    var wordSentences jisho_ws.WordSentenceDict=jisho_ws.ReadSingleSplitDict(
        wordSentenceFilesDir,
        wordSentenceFilename,
    )

    var wordSentenceSubset jisho_ws.WordSentenceDict=jisho_ws.GetSentenceSubset(
        wordSentences,
        sentencesPerWordMin,
        sentencesPerWordMax,
    )

    var sentencePairs []WordSentencePair=wordSentenceDictToPairs(wordSentenceSubset)
    utils.ShuffleArray[WordSentencePair](sentencePairs)

    return KjStudySession{
        WordSentences: sentencePairs,
        Datafile: wordSentenceFilename,
    }

}

// convert word sentence dict to pairs array. should probably shuffle the array
func wordSentenceDictToPairs(sentencesDict jisho_ws.WordSentenceDict) []WordSentencePair {
    var pairs []WordSentencePair

    var word string
    var sentences []string
    for word,sentences = range sentencesDict {
        var sentence string
        for _,sentence = range sentences {
            pairs=append(pairs,WordSentencePair{
                Word: word,
                Sentence: sentence,
                Status: WordSentenceStatus_normal,
            })
        }
    }

    return pairs
}

// mutate session. find a target word/sentence, and set state to specified state
func SetPairState(
    session *KjStudySession,

    word string,
    sentence string,
    newState WordSentenceStatus,
) error {
    var pairI int
    for pairI = range session.WordSentences {
        if (session.WordSentences[pairI].Word==word &&
        session.WordSentences[pairI].Sentence==sentence) {
            session.WordSentences[pairI].Status=newState
            return nil
        }
    }

    return errors.New("missing pair error")
}