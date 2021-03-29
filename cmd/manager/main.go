// Copyright 2019 The Kubernetes Authors.
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

package main

import (
	"flag"
	"log"
	"net/http"

	_ "net/http/pprof"

	"github.com/spf13/pflag"

	"k8s.io/klog"

	"github.com/open-cluster-management/multicloud-operators-subscription/cmd/manager/exec"
)

func main() {
	exec.ProcessFlags()

	klog.InitFlags(nil)

	pflag.CommandLine.AddGoFlagSet(flag.CommandLine)
	pflag.Parse()

	defer klog.Flush()

	pflag.Parse()

	// we need a webserver to get the pprof webserver
	go func() {
		log.Println(http.ListenAndServe("localhost:6060", nil))
	}()

	exec.RunManager()
}
