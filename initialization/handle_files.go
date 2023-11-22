package initialization

import (
	"sort"

	"github.com/spacemeshos/post/config"
	"github.com/spacemeshos/post/shared"
	"go.uber.org/zap"
)

func prepareFileIndexes(cfg config.Config, opts config.InitOpts, logger *zap.Logger) ([]int, error) {
	layout, err := deriveFilesLayout(cfg, opts)
	if err != nil {
		return nil, err
	}
	maxFileIndex := layout.LastFileIdx
	minFileIndex := layout.FirstFileIdx
	from := minFileIndex
	to := maxFileIndex
	if (opts.FromFileIdx != 0) {
		from = opts.FromFileIdx
	}
	if (opts.ToFileIdx != nil) {
		to = *opts.ToFileIdx
	}
	fileIndexMap := make(map[int]bool, layout.LastFileIdx-layout.FirstFileIdx+1)
	for i := 0; i <= maxFileIndex; i++ {
		if (i >= from && i <= to) {
			fileIndexMap[i] = true
		}
	}
	files, err := GetFiles(opts.DataDir, shared.IsInitFile)
	for _, file := range files {
		name := file.Name()
		fileIndex, err := shared.ParseFileIndex(name)
		if err != nil && name != MetadataFileName {
			// TODO(mafa): revert back to warning, see https://github.com/spacemeshos/go-spacemesh/issues/4789
			logger.Debug("found unrecognized file", zap.String("fileName", name))
			continue
		}
		if fileIndex > maxFileIndex || fileIndex < minFileIndex {
			continue
		}
		if fileIndex == layout.LastFileIdx {
			if shared.NumLabels(uint64(file.Size()), config.BitsPerLabel) < uint64(layout.LastFileNumLabels) {
				continue
			} else {
				delete(fileIndexMap, fileIndex)
				continue
			}
		}
		if shared.NumLabels(uint64(file.Size()), config.BitsPerLabel) < uint64(layout.FileNumLabels) {
			continue
		}
		delete(fileIndexMap, fileIndex)
	}
	keys := make([]int, 0, len(fileIndexMap))
	for k, _ := range fileIndexMap {
		keys = append(keys, k)
	}
	sort.Ints(keys)
	return keys, nil
}

func (init *Initializer) IsAlreadyInitialization() bool {
	indexes, _ := prepareFileIndexes(init.cfg, init.opts, init.logger)
	if len(indexes) > 0 {
		return false
	}
	return true
}
