package types

import "time"

type RessourceLinkResp struct {
	Link    string    `json:"link"`
	Expires time.Time `json:"expires"`
}

type RessourceLinkChunksResp struct {
	Type string `json:"type"`
	Data struct {
		Success   bool          `json:"success"`
		ChunkInfo ChunkInfoType `json:"chunk_info"`
	} `json:"data"`
}

type ChunkInfoType struct {
	ChunkSize       int      `json:"chunk_size"`
	NumChunks       int      `json:"num_chunks"`
	Rows            int      `json:"rows"`
	BaseDownloadURL string   `json:"base_download_url"`
	ChunkFileNames  []string `json:"chunk_file_names"`
}

type ChunkReq struct {
	Rows int `json:"rows"`
}
