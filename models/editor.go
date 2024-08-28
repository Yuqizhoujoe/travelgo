package models

type BlockId string

type BlockToolData map[string]interface{}

type BlockTuneData map[string]interface{}

type EditorBlockData struct {
	ID    *BlockId      `json:"id,ompitemty"`
	Type  string        `json:"type"`
	Data  BlockToolData `json:"data"`
	Tunes BlockTuneData `json:"tunes,omitempty"`
}

type EditorData struct {
	Version string            `json:"version,omitempty"`
	Time    int64             `json:"time,omitempty"`
	Blocks  []EditorBlockData `json:"blocks"`
}
