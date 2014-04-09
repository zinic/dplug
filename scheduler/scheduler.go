package scheduler

import "github.com/zinic/dplug/model"
import "github.com/zinic/dplug/provider"

type Layout struct {
    Host       *model.Host
    Containers *[]model.Container
}

type DeploymentTracker interface {
    GetLayout() Layout
}

func FindCandidateHost(creds *provider.Credentials, image *model.Image, cloud provider.CloudProvider) *model.Host {
    for _, host := range cloud.List(creds) {
        if host.Memory.Available >= image.MemoryRequired {
            return &host
        }
    }

    return nil
}
