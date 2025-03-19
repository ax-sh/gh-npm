//go:build mage

package main

// todo
//// Runs git tag with git cliff
//func Tag() error {
//	a, _ := sh.Output("git", "tag")
//	fmt.Printf("Tagging %s\n", a)
//	if err := sh.Run("echo", "mod", "download"); err != nil {
//		return err
//	}
//	return sh.Run("echo", "install", "./...")
//}
//
//// Release creates a new version tag and GitHub release with changelog
//func Tag() error {
//	// Get the bumped version using git-cliff
//	versionCmd := sh.Command("git", "cliff", "--bumped-version")
//	versionBytes, err := versionCmd.Output()
//	if err != nil {
//		return fmt.Errorf("failed to get bumped version: %w", err)
//	}
//	version := string(versionBytes)
//	log.Printf("Release version: %s", version)
//
//	//// Create a git tag with the version
//	//tagCmd := exec.Command("git", "tag", version)
//	//if err := tagCmd.Run(); err != nil {
//	//	return fmt.Errorf("failed to create git tag: %w", err)
//	//}
//	//
//	//// Push the tag to origin
//	//pushCmd := exec.Command("git", "push", "origin", version)
//	//if err := pushCmd.Run(); err != nil {
//	//	return fmt.Errorf("failed to push tag: %w", err)
//	//}
//	//
//	//// Generate the changelog
//	//cliffCmd := exec.Command("git", "cliff", "-l", "-s", "all")
//	//changelog, err := cliffCmd.Output()
//	//if err != nil {
//	//	return fmt.Errorf("failed to generate changelog: %w", err)
//	//}
//	//
//	//// Create a temporary file for the changelog
//	//tempFile, err := os.CreateTemp("", "changelog-*.md")
//	//if err != nil {
//	//	return fmt.Errorf("failed to create temp file: %w", err)
//	//}
//	//defer os.Remove(tempFile.Name())
//	//
//	//// Write the changelog to the temporary file
//	//if _, err := tempFile.Write(changelog); err != nil {
//	//	return fmt.Errorf("failed to write changelog to temp file: %w", err)
//	//}
//	//if err := tempFile.Close(); err != nil {
//	//	return fmt.Errorf("failed to close temp file: %w", err)
//	//}
//	//
//	//// Create a GitHub release with the changelog
//	//releaseCmd := exec.Command("gh", "release", "create", version, "--notes-file", tempFile.Name())
//	//if err := releaseCmd.Run(); err != nil {
//	//	return fmt.Errorf("failed to create GitHub release: %w", err)
//	//}
//	//
//	//fmt.Printf("Successfully created and pushed release %s\n", version)
//	//return nil
//}
