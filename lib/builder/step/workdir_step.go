//  Copyright (c) 2018 Uber Technologies, Inc.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

package step

import (
	"fmt"
	"path/filepath"

	"github.com/uber/makisu/lib/context"
	"github.com/uber/makisu/lib/docker/image"
)

// WorkdirStep implements BuildStep and execute WORKDIR directive
type WorkdirStep struct {
	*baseStep

	workingDir string
}

// NewWorkdirStep returns a BuildStep from give build step.
func NewWorkdirStep(args string, workingDir string, commit bool) BuildStep {
	return &WorkdirStep{
		baseStep:   newBaseStep(Workdir, args, commit),
		workingDir: workingDir,
	}
}

// GenerateConfig generates a new image config base on config from previous step.
func (s *WorkdirStep) GenerateConfig(ctx *context.BuildContext, imageConfig *image.Config) (*image.Config, error) {
	config, err := image.NewImageConfigFromCopy(imageConfig)
	if err != nil {
		return nil, fmt.Errorf("copy image config: %s", err)
	}
	if filepath.IsAbs(s.workingDir) {
		config.Config.WorkingDir = ctx.RootDir
	}
	config.Config.WorkingDir = filepath.Join(config.Config.WorkingDir, s.workingDir)

	return config, nil
}
