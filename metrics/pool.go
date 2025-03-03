package metrics

import (
	"log/slog"
	"strconv"
	"strings"

	"github.com/cloudbase/garm/auth"
	"github.com/prometheus/client_golang/prometheus"
)

// CollectPoolMetric collects the metrics for the pool objects
func (c *GarmCollector) CollectPoolMetric(ch chan<- prometheus.Metric, hostname string, controllerID string) {
	ctx := auth.GetAdminContext()

	pools, err := c.runner.ListAllPools(ctx)
	if err != nil {
		slog.With(slog.Any("error", err)).ErrorContext(ctx, "listing pools")
		return
	}

	type poolInfo struct {
		Name string
		Type string
	}

	poolNames := make(map[string]poolInfo)
	for _, pool := range pools {
		if pool.EnterpriseName != "" {
			poolNames[pool.ID] = poolInfo{
				Name: pool.EnterpriseName,
				Type: string(pool.PoolType()),
			}
		} else if pool.OrgName != "" {
			poolNames[pool.ID] = poolInfo{
				Name: pool.OrgName,
				Type: string(pool.PoolType()),
			}
		} else {
			poolNames[pool.ID] = poolInfo{
				Name: pool.RepoName,
				Type: string(pool.PoolType()),
			}
		}

		var poolTags []string
		for _, tag := range pool.Tags {
			poolTags = append(poolTags, tag.Name)
		}

		poolInfo, err := prometheus.NewConstMetric(
			c.poolInfo,
			prometheus.GaugeValue,
			1,
			pool.ID,                     // label: id
			pool.Image,                  // label: image
			pool.Flavor,                 // label: flavor
			pool.Prefix,                 // label: prefix
			string(pool.OSType),         // label: os_type
			string(pool.OSArch),         // label: os_arch
			strings.Join(poolTags, ","), // label: tags
			pool.ProviderName,           // label: provider
			poolNames[pool.ID].Name,     // label: pool_owner
			poolNames[pool.ID].Type,     // label: pool_type
		)
		if err != nil {
			slog.With(slog.Any("error", err)).ErrorContext(ctx, "cannot collect poolInfo metric")
			continue
		}
		ch <- poolInfo

		poolStatus, err := prometheus.NewConstMetric(
			c.poolStatus,
			prometheus.GaugeValue,
			bool2float64(pool.Enabled),
			pool.ID,                          // label: id
			strconv.FormatBool(pool.Enabled), // label: enabled
		)
		if err != nil {
			slog.With(slog.Any("error", err)).ErrorContext(ctx, "cannot collect poolStatus metric")
			continue
		}
		ch <- poolStatus

		poolMaxRunners, err := prometheus.NewConstMetric(
			c.poolMaxRunners,
			prometheus.GaugeValue,
			float64(pool.MaxRunners),
			pool.ID, // label: id
		)
		if err != nil {
			slog.With(slog.Any("error", err)).ErrorContext(ctx, "cannot collect poolMaxRunners metric")
			continue
		}
		ch <- poolMaxRunners

		poolMinIdleRunners, err := prometheus.NewConstMetric(
			c.poolMinIdleRunners,
			prometheus.GaugeValue,
			float64(pool.MinIdleRunners),
			pool.ID, // label: id
		)
		if err != nil {
			slog.With(slog.Any("error", err)).ErrorContext(ctx, "cannot collect poolMinIdleRunners metric")
			continue
		}
		ch <- poolMinIdleRunners

		poolBootstrapTimeout, err := prometheus.NewConstMetric(
			c.poolBootstrapTimeout,
			prometheus.GaugeValue,
			float64(pool.RunnerBootstrapTimeout),
			pool.ID, // label: id
		)
		if err != nil {
			slog.With(slog.Any("error", err)).ErrorContext(ctx, "cannot collect poolBootstrapTimeout metric")
			continue
		}
		ch <- poolBootstrapTimeout
	}
}
