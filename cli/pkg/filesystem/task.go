/*
 Copyright 2021 The KubeSphere Authors.

 Licensed under the Apache License, Version 2.0 (the "License");
 you may not use this file except in compliance with the License.
 You may obtain a copy of the License at

     http://www.apache.org/licenses/LICENSE-2.0

 Unless required by applicable law or agreed to in writing, software
 distributed under the License is distributed on an "AS IS" BASIS,
 WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 See the License for the specific language governing permissions and
 limitations under the License.
*/

package filesystem

import (
	"fmt"

	"github.com/beclab/Olares/cli/pkg/core/action"
	"github.com/beclab/Olares/cli/pkg/core/connector"
	"github.com/pkg/errors"
)

type ChownFileAndDir struct {
	action.BaseAction
	Path string
}

func (c *ChownFileAndDir) Execute(runtime connector.Runtime) error {
	exist, err := runtime.GetRunner().FileExist(c.Path)
	if err != nil {
		return err
	}

	if exist {
		userId, err := runtime.GetRunner().Cmd("echo $(id -u)", false, false)
		if err != nil {
			return errors.Wrap(errors.WithStack(err), "get user id failed")
		}

		userGroupId, err := runtime.GetRunner().Cmd("echo $(id -g)", false, false)
		if err != nil {
			return errors.Wrap(errors.WithStack(err), "get user group id failed")
		}

		chownKubeConfig := fmt.Sprintf("chown -R %s:%s %s", userId, userGroupId, c.Path)
		if _, err := runtime.GetRunner().SudoCmd(chownKubeConfig, false, false); err != nil {
			return errors.Wrapf(errors.WithStack(err), "chown user %s failed", c.Path)
		}
	}
	return nil
}
