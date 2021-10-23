// Code generated by "packer-sdc mapstructure-to-hcl2"; DO NOT EDIT.

package builder

import (
	"github.com/hashicorp/hcl/v2/hcldec"
	"github.com/solo-io/packer-plugin-arm-image/pkg/image/utils"
	"github.com/zclconf/go-cty/cty"
)

// FlatConfig is an auto-generated flat version of Config.
// Where the contents of a field with a `mapstructure:,squash` tag are bubbled up.
type FlatConfig struct {
	PackerBuildName        *string               `mapstructure:"packer_build_name" cty:"packer_build_name" hcl:"packer_build_name"`
	PackerBuilderType      *string               `mapstructure:"packer_builder_type" cty:"packer_builder_type" hcl:"packer_builder_type"`
	PackerCoreVersion      *string               `mapstructure:"packer_core_version" cty:"packer_core_version" hcl:"packer_core_version"`
	PackerDebug            *bool                 `mapstructure:"packer_debug" cty:"packer_debug" hcl:"packer_debug"`
	PackerForce            *bool                 `mapstructure:"packer_force" cty:"packer_force" hcl:"packer_force"`
	PackerOnError          *string               `mapstructure:"packer_on_error" cty:"packer_on_error" hcl:"packer_on_error"`
	PackerUserVars         map[string]string     `mapstructure:"packer_user_variables" cty:"packer_user_variables" hcl:"packer_user_variables"`
	PackerSensitiveVars    []string              `mapstructure:"packer_sensitive_variables" cty:"packer_sensitive_variables" hcl:"packer_sensitive_variables"`
	ISOChecksum            *string               `mapstructure:"iso_checksum" required:"true" cty:"iso_checksum" hcl:"iso_checksum"`
	RawSingleISOUrl        *string               `mapstructure:"iso_url" required:"true" cty:"iso_url" hcl:"iso_url"`
	ISOUrls                []string              `mapstructure:"iso_urls" cty:"iso_urls" hcl:"iso_urls"`
	TargetPath             *string               `mapstructure:"iso_target_path" cty:"iso_target_path" hcl:"iso_target_path"`
	TargetExtension        *string               `mapstructure:"iso_target_extension" cty:"iso_target_extension" hcl:"iso_target_extension"`
	CommandWrapper         *string               `mapstructure:"command_wrapper" cty:"command_wrapper" hcl:"command_wrapper"`
	OutputDir              *string               `mapstructure:"output_directory" cty:"output_directory" hcl:"output_directory"`
	OutputFile             *string               `mapstructure:"output_filename" cty:"output_filename" hcl:"output_filename"`
	ImageType              *utils.KnownImageType `mapstructure:"image_type" cty:"image_type" hcl:"image_type"`
	ImageMounts            []string              `mapstructure:"image_mounts" cty:"image_mounts" hcl:"image_mounts"`
	MountPath              *string               `mapstructure:"mount_path" cty:"mount_path" hcl:"mount_path"`
	ChrootMounts           [][]string            `mapstructure:"chroot_mounts" cty:"chroot_mounts" hcl:"chroot_mounts"`
	AdditionalChrootMounts [][]string            `mapstructure:"additional_chroot_mounts" cty:"additional_chroot_mounts" hcl:"additional_chroot_mounts"`
	ResolvConf             *ResolvConfBehavior   `mapstructure:"resolv-conf" cty:"resolv-conf" hcl:"resolv-conf"`
	LastPartitionExtraSize *uint64               `mapstructure:"last_partition_extra_size" cty:"last_partition_extra_size" hcl:"last_partition_extra_size"`
	TargetImageSize        *uint64               `mapstructure:"target_image_size" cty:"target_image_size" hcl:"target_image_size"`
	QemuBinary             *string               `mapstructure:"qemu_binary" cty:"qemu_binary" hcl:"qemu_binary"`
	DisableEmbedded        *bool                 `mapstructure:"disable_embedded" cty:"disable_embedded" hcl:"disable_embedded"`
	QemuArgs               []string              `mapstructure:"qemu_args" cty:"qemu_args" hcl:"qemu_args"`
}

// FlatMapstructure returns a new FlatConfig.
// FlatConfig is an auto-generated flat version of Config.
// Where the contents a fields with a `mapstructure:,squash` tag are bubbled up.
func (*Config) FlatMapstructure() interface{ HCL2Spec() map[string]hcldec.Spec } {
	return new(FlatConfig)
}

// HCL2Spec returns the hcl spec of a Config.
// This spec is used by HCL to read the fields of Config.
// The decoded values from this spec will then be applied to a FlatConfig.
func (*FlatConfig) HCL2Spec() map[string]hcldec.Spec {
	s := map[string]hcldec.Spec{
		"packer_build_name":          &hcldec.AttrSpec{Name: "packer_build_name", Type: cty.String, Required: false},
		"packer_builder_type":        &hcldec.AttrSpec{Name: "packer_builder_type", Type: cty.String, Required: false},
		"packer_core_version":        &hcldec.AttrSpec{Name: "packer_core_version", Type: cty.String, Required: false},
		"packer_debug":               &hcldec.AttrSpec{Name: "packer_debug", Type: cty.Bool, Required: false},
		"packer_force":               &hcldec.AttrSpec{Name: "packer_force", Type: cty.Bool, Required: false},
		"packer_on_error":            &hcldec.AttrSpec{Name: "packer_on_error", Type: cty.String, Required: false},
		"packer_user_variables":      &hcldec.AttrSpec{Name: "packer_user_variables", Type: cty.Map(cty.String), Required: false},
		"packer_sensitive_variables": &hcldec.AttrSpec{Name: "packer_sensitive_variables", Type: cty.List(cty.String), Required: false},
		"iso_checksum":               &hcldec.AttrSpec{Name: "iso_checksum", Type: cty.String, Required: false},
		"iso_url":                    &hcldec.AttrSpec{Name: "iso_url", Type: cty.String, Required: false},
		"iso_urls":                   &hcldec.AttrSpec{Name: "iso_urls", Type: cty.List(cty.String), Required: false},
		"iso_target_path":            &hcldec.AttrSpec{Name: "iso_target_path", Type: cty.String, Required: false},
		"iso_target_extension":       &hcldec.AttrSpec{Name: "iso_target_extension", Type: cty.String, Required: false},
		"command_wrapper":            &hcldec.AttrSpec{Name: "command_wrapper", Type: cty.String, Required: false},
		"output_directory":           &hcldec.AttrSpec{Name: "output_directory", Type: cty.String, Required: false},
		"output_filename":            &hcldec.AttrSpec{Name: "output_filename", Type: cty.String, Required: false},
		"image_type":                 &hcldec.AttrSpec{Name: "image_type", Type: cty.String, Required: false},
		"image_mounts":               &hcldec.AttrSpec{Name: "image_mounts", Type: cty.List(cty.String), Required: false},
		"mount_path":                 &hcldec.AttrSpec{Name: "mount_path", Type: cty.String, Required: false},
		"chroot_mounts":              &hcldec.AttrSpec{Name: "chroot_mounts", Type: cty.List(cty.List(cty.String)), Required: false},
		"additional_chroot_mounts":   &hcldec.AttrSpec{Name: "additional_chroot_mounts", Type: cty.List(cty.List(cty.String)), Required: false},
		"resolv-conf":                &hcldec.AttrSpec{Name: "resolv-conf", Type: cty.String, Required: false},
		"last_partition_extra_size":  &hcldec.AttrSpec{Name: "last_partition_extra_size", Type: cty.Number, Required: false},
		"target_image_size":          &hcldec.AttrSpec{Name: "target_image_size", Type: cty.Number, Required: false},
		"qemu_binary":                &hcldec.AttrSpec{Name: "qemu_binary", Type: cty.String, Required: false},
		"disable_embedded":           &hcldec.AttrSpec{Name: "disable_embedded", Type: cty.Bool, Required: false},
		"qemu_args":                  &hcldec.AttrSpec{Name: "qemu_args", Type: cty.List(cty.String), Required: false},
	}
	return s
}