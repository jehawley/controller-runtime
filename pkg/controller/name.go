/*
Copyright 2020 The Kubernetes Authors.

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

package controller

import (
	"fmt"
	"runtime/debug"
	"sync"

	"k8s.io/apimachinery/pkg/util/sets"
)

var nameLock sync.Mutex
var usedNames sets.Set[string]

func checkName(name string) error {
	nameLock.Lock()
	defer nameLock.Unlock()
	if usedNames == nil {
		usedNames = sets.Set[string]{}
	}

	fmt.Sprintf("\n <3 name: %s", name)
	debug.PrintStack()
	if usedNames.Has(name) {
		return fmt.Errorf("controller with name %s already exists. Controller names must be unique to avoid multiple controllers reporting to the same metric", name)
	}

	usedNames.Insert(name)

	return nil
}
