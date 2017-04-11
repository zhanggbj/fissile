package model

import (
	"io/ioutil"
	"os"
	"path/filepath"
	"sort"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestLoadRoleManifestOK(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	release, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/tor-good.yml")
	rolesManifest, err := LoadRoleManifest(roleManifestPath, []*Release{release})
	assert.NoError(err)
	assert.NotNil(rolesManifest)

	assert.Equal(roleManifestPath, rolesManifest.manifestFilePath)
	assert.Len(rolesManifest.Roles, 2)

	myrole := rolesManifest.Roles[0]
	assert.Equal([]string{
		"myrole.sh",
		"/script/with/absolute/path.sh",
	}, myrole.Scripts)

	foorole := rolesManifest.Roles[1]
	torjob := foorole.Jobs[0]
	assert.Equal("tor", torjob.Name)
	assert.NotNil(torjob.Release)
	assert.Equal("tor", torjob.Release.Name)
}

func TestGetScriptPaths(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	release, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/tor-good.yml")
	rolesManifest, err := LoadRoleManifest(roleManifestPath, []*Release{release})
	assert.NoError(err)
	assert.NotNil(rolesManifest)

	fullScripts := rolesManifest.Roles[0].GetScriptPaths()
	assert.Len(fullScripts, 3)
	for _, leafName := range []string{"environ.sh", "myrole.sh", "post_config_script.sh"} {
		assert.Equal(filepath.Join(workDir, "../test-assets/role-manifests", leafName), fullScripts[leafName])
	}
}

func TestLoadRoleManifestNotOKBadJobName(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	release, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/tor-bad.yml")
	_, err = LoadRoleManifest(roleManifestPath, []*Release{release})
	assert.NotNil(err)
	assert.Contains(err.Error(), "Cannot find job foo in release")
}

func TestLoadDuplicateReleases(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	release, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/tor-good.yml")
	_, err = LoadRoleManifest(roleManifestPath, []*Release{release, release})

	assert.NotNil(err)
	assert.Contains(err.Error(), "release tor has been loaded more than once")
}

func TestLoadRoleManifestMultipleReleasesOK(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	ntpReleasePath := filepath.Join(workDir, "../test-assets/ntp-release")
	ntpReleasePathBoshCache := filepath.Join(ntpReleasePath, "bosh-cache")
	ntpRelease, err := NewDevRelease(ntpReleasePath, "", "", ntpReleasePathBoshCache)
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	torRelease, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/multiple-good.yml")
	rolesManifest, err := LoadRoleManifest(roleManifestPath, []*Release{ntpRelease, torRelease})
	assert.NoError(err)
	assert.NotNil(rolesManifest)

	assert.Equal(roleManifestPath, rolesManifest.manifestFilePath)
	assert.Len(rolesManifest.Roles, 2)

	myrole := rolesManifest.Roles[0]
	assert.Len(myrole.Scripts, 1)
	assert.Equal("myrole.sh", myrole.Scripts[0])

	foorole := rolesManifest.Roles[1]
	torjob := foorole.Jobs[0]
	assert.Equal("tor", torjob.Name)
	assert.NotNil(torjob.Release)
	assert.Equal("tor", torjob.Release.Name)
}

func TestLoadRoleManifestMultipleReleasesNotOk(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	ntpReleasePath := filepath.Join(workDir, "../test-assets/ntp-release")
	ntpReleasePathBoshCache := filepath.Join(ntpReleasePath, "bosh-cache")
	ntpRelease, err := NewDevRelease(ntpReleasePath, "", "", ntpReleasePathBoshCache)
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	torRelease, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/multiple-bad.yml")
	_, err = LoadRoleManifest(roleManifestPath, []*Release{ntpRelease, torRelease})

	assert.NotNil(err)
	assert.Contains(err.Error(), "release foo has not been loaded and is referenced by job ntpd in role foorole")
}

func TestNonBoshRolesAreIgnoredOK(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	release, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/non-bosh-roles.yml")
	rolesManifest, err := LoadRoleManifest(roleManifestPath, []*Release{release})
	assert.NoError(err)
	assert.NotNil(rolesManifest)

	assert.Equal(roleManifestPath, rolesManifest.manifestFilePath)
	assert.Len(rolesManifest.Roles, 2)
}

func TestRolesSort(t *testing.T) {
	assert := assert.New(t)

	roles := Roles{
		{Name: "aaa"},
		{Name: "bbb"},
	}
	sort.Sort(roles)
	assert.Equal(roles[0].Name, "aaa")
	assert.Equal(roles[1].Name, "bbb")

	roles = Roles{
		{Name: "ddd"},
		{Name: "ccc"},
	}
	sort.Sort(roles)
	assert.Equal(roles[0].Name, "ccc")
	assert.Equal(roles[1].Name, "ddd")
}

func TestGetScriptSignatures(t *testing.T) {
	assert := assert.New(t)

	refRole := &Role{
		Name: "bbb",
		Jobs: Jobs{
			{
				SHA1: "Role 2 Job 1",
				Packages: Packages{
					{Name: "aaa", SHA1: "Role 2 Job 1 Package 1"},
					{Name: "bbb", SHA1: "Role 2 Job 1 Package 2"},
				},
			},
			{
				SHA1: "Role 2 Job 2",
				Packages: Packages{
					{Name: "ccc", SHA1: "Role 2 Job 2 Package 1"},
				},
			},
		},
	}

	firstHash, _ := refRole.GetScriptSignatures()

	workDir, err := ioutil.TempDir("", "fissile-test-")
	assert.NoError(err)
	defer os.RemoveAll(workDir)
	releasePath := filepath.Join(workDir, "role.yml")

	scriptName := "script.sh"
	scriptPath := filepath.Join(workDir, scriptName)
	err = ioutil.WriteFile(scriptPath, []byte("true\n"), 0644)
	assert.NoError(err)

	differentPatch := &Role{
		Name:    refRole.Name,
		Jobs:    Jobs{refRole.Jobs[0], refRole.Jobs[1]},
		Scripts: []string{scriptName},
		rolesManifest: &RoleManifest{
			manifestFilePath: releasePath,
		},
	}

	differentPatchHash, _ := differentPatch.GetScriptSignatures()
	assert.NotEqual(firstHash, differentPatchHash, "role hash should be dependent on patch string")

	err = ioutil.WriteFile(scriptPath, []byte("false\n"), 0644)
	assert.NoError(err)

	differentPatchFileHash, _ := differentPatch.GetScriptSignatures()
	assert.NotEqual(differentPatchFileHash, differentPatchHash, "role manifest hash should be dependent on patch contents")
}

func TestGetTemplateSignatures(t *testing.T) {
	assert := assert.New(t)

	differentTemplate1 := &Role{
		Name: "aaa",
		Jobs: Jobs{},
		Configuration: &Configuration{
			Templates: map[string]string{"foo": "bar"},
		},
	}

	differentTemplate2 := &Role{
		Name: "aaa",
		Jobs: Jobs{},
		Configuration: &Configuration{
			Templates: map[string]string{"bat": "baz"},
		},
	}

	differentTemplateHash1, _ := differentTemplate1.GetTemplateSignatures()
	differentTemplateHash2, _ := differentTemplate2.GetTemplateSignatures()
	assert.NotEqual(differentTemplateHash1, differentTemplateHash2, "template hash should be dependent on template contents")
}

func TestGetRoleManifestDevPackageVersion(t *testing.T) {
	assert := assert.New(t)

	refRole := &Role{
		Name: "bbb",
		Jobs: Jobs{
			{
				SHA1: "Role 2 Job 1",
				Packages: Packages{
					{Name: "aaa", SHA1: "Role 2 Job 1 Package 1"},
					{Name: "bbb", SHA1: "Role 2 Job 1 Package 2"},
				},
			},
			{
				SHA1: "Role 2 Job 2",
				Packages: Packages{
					{Name: "ccc", SHA1: "Role 2 Job 2 Package 1"},
				},
			},
		},
	}
	wrongJobOrder := &Role{
		Name: refRole.Name,
		Jobs: Jobs{refRole.Jobs[1], refRole.Jobs[0]},
	}
	altRole := &Role{
		Name: "aaa",
		Jobs: Jobs{
			{
				SHA1: "Role 1 Job 1",
				Packages: Packages{
					{Name: "zzz", SHA1: "Role 1	 Job 1 Package 1"},
				},
			},
		},
	}

	firstManifest := &RoleManifest{Roles: Roles{refRole, altRole}}
	firstHash, _ := firstManifest.GetRoleManifestDevPackageVersion(firstManifest.Roles, "")
	secondManifest := &RoleManifest{Roles: Roles{altRole, refRole}}
	secondHash, _ := secondManifest.GetRoleManifestDevPackageVersion(secondManifest.Roles, "")
	assert.Equal(firstHash, secondHash, "role manifest hash should be independent of role order")
	jobOrderManifest := &RoleManifest{Roles: Roles{wrongJobOrder, altRole}}
	jobOrderHash, _ := jobOrderManifest.GetRoleManifestDevPackageVersion(jobOrderManifest.Roles, "")
	assert.NotEqual(firstHash, jobOrderHash, "role manifest hash should be dependent on job order")
	differentExtraHash, _ := firstManifest.GetRoleManifestDevPackageVersion(firstManifest.Roles, "some string")
	assert.NotEqual(firstHash, differentExtraHash, "role manifest hash should be dependent on extra string")
}

func TestValidateRoleRun(t *testing.T) {
	newRun := func(memory int, cpu int, name string, externalPort string, internalPort string, protocol string) *RoleRun {
		return &RoleRun{
			VirtualCPUs: cpu,
			Memory:      memory,
			ExposedPorts: []*RoleRunExposedPort{{Name: name, External: externalPort,
				Internal: internalPort, Protocol: protocol,
			}},
		}
	}

	var (
		validRun      = newRun(10, 2, "test", "1", "2", "UDP")
		wrongProtocol = newRun(10, 2, "test", "1", "2", "AA")
		wrongPorts    = newRun(10, 2, "test", "0", "-1", "UDP")
		wrongParse    = newRun(10, 2, "test", "0", "qq", "UDP")
		negativeField = newRun(-10, 2, "test", "1", "2", "UDP")
	)

	tests := map[string]struct {
		run            *RoleRun
		isValid        bool
		expectedErrors string
	}{
		"nil": {nil, false,
			"Role 'nil': run: Required value"},
		"valid": {validRun, true, ``},
		"wrong protocol": {wrongProtocol, false,
			"Role 'wrong protocol': run.exposed-ports[test].protocol: Unsupported value: \"AA\": supported values: TCP, UDP"},
		"wrong ports": {wrongPorts, false,
			"Role 'wrong ports': run.exposed-ports[test].external: Invalid value: 0: must be between 1 and 65535, inclusive\n" +
				"Role 'wrong ports': run.exposed-ports[test].internal: Invalid value: -1: must be between 1 and 65535, inclusive"},
		"wrong parse": {wrongParse, false,
			"Role 'wrong parse': run.exposed-ports[test].external: Invalid value: 0: must be between 1 and 65535, inclusive\n" +
				"Role 'wrong parse': run.exposed-ports[test].internal: Invalid value: \"qq\": invalid syntax"},
		"negative field": {negativeField, false,
			`Role 'negative field': run.memory: Invalid value: -10: must be greater than or equal to 0`},
	}

	for name, tc := range tests {
		errs := validateRoleRun(tc.run, name)
		if tc.isValid && len(errs) > 0 {
			t.Errorf("%v: unexpected error: %v", name, errs)
		}
		if !tc.isValid && len(errs) == 0 {
			t.Errorf("%v: unexpected non-error", name)
		}
		if !tc.isValid && len(errs) > 0 {
			assert.Equal(t, tc.expectedErrors, errs.Errors())
		}
	}
}

func TestLoadRoleManifestVariablesSortedBAD(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	release, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/variables-badly-sorted.yml")
	rolesManifest, err := LoadRoleManifest(roleManifestPath, []*Release{release})
	assert.Equal(err.Error(),
		`configuration.variables: Invalid value: "FOO": Does not sort before 'BAR'
configuration.variables: Invalid value: "PELERINUL": Does not sort before 'ALPHA'`)
	assert.Nil(rolesManifest)
}

func TestLoadRoleManifestVariablesNotUsed(t *testing.T) {
	assert := assert.New(t)

	workDir, err := os.Getwd()
	assert.NoError(err)

	torReleasePath := filepath.Join(workDir, "../test-assets/tor-boshrelease")
	torReleasePathBoshCache := filepath.Join(torReleasePath, "bosh-cache")
	release, err := NewDevRelease(torReleasePath, "", "", torReleasePathBoshCache)
	assert.NoError(err)

	roleManifestPath := filepath.Join(workDir, "../test-assets/role-manifests/variables-without-usage.yml")
	rolesManifest, err := LoadRoleManifest(roleManifestPath, []*Release{release})
	assert.Equal(err.Error(),
		`configuration.variables: Not found: "No templates using 'SOME_VAR'"`)
	assert.Nil(rolesManifest)
}
