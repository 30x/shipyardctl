package cmd

import (
  "os"
  "fmt"
)


// RequireOrgName used to short circuit commands
// requiring the Apigee org name if it is not present
func RequireOrgName() error {
	if orgName == "" {
		if orgName = os.Getenv("APIGEE_ORG"); orgName == "" {
			return fmt.Errorf("Missing required flag '--org', or place in environment as APIGEE_ORG.")
		}
	}

	return nil
}

// RequireEnvName used to short circuit commands
// requiring the Apigee env name if it is not present
func RequireEnvName() error {
	if envName == "" {
		if envName = os.Getenv("APIGEE_ENV"); envName == "" {
			return fmt.Errorf("Missing required flag '--env', or place in environment as APIGEE_ENV.")
		}
	}

	return nil
}

// RequireAppName used to short circuit commands
// requiring the app name if it is not present
func RequireAppName() error {
	if appName == "" {
		return fmt.Errorf("Missing required flag '--name'.")
	}

	return nil
}

// RequireAppPath used to short circuit commands
// requiring the app path if it is not present
func RequireAppPath() error {
	if appPath == "" {
		return fmt.Errorf("Missing required flag '--path'.")
	}

	return nil
}

// RequireBundleName used to short circuit commands
// requiring the bundle name be provided via the name flag
func RequireBundleName() error {
	if bundleName == "" {
		return fmt.Errorf("Missing required flag '--name'.")
	}

	return nil
}

// RequireDirectory used to short circuit commands
// requiring a directory, if it is not present
func RequireDirectory() error {
	if directory == "" {
		return fmt.Errorf("Missing required flag '--directory'.")
	}

	return nil
}

// RequirePTSURL used to short circuit commands
// requiring a PTS URL, if it is not present
func RequirePTSURL() error {
	if ptsUrl == "" {
		return fmt.Errorf("Missing required flag '--pts-url'.")
	}

	return nil
}

// RequireHostnames used to short circuit commands
// requiring hostnames, if it is not present
func RequireHostnames() error {
	if hostnames == "" {
		return fmt.Errorf("Missing required flag '--hostnames'.")
	}

	return nil
}

// RequireZipPath used to short circuit commands
// requiring the path to a bundle zip, if it is not present
func RequireZipPath() error {
	if bundlePath == "" {
		return fmt.Errorf("Missing required flag '--zip-path'.")
	}

	return nil
}

// MakeBuildPath make build service path with given orgName
func MakeBuildPath() {
	basePath = fmt.Sprintf("/imagespaces/%s/images", orgName)
}