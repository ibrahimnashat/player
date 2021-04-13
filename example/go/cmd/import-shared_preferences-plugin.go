package main

// DO NOT EDIT, this file is generated by hover at compile-time for the shared_preferences plugin.

import (
	flutter "github.com/go-flutter-desktop/go-flutter"
	shared_preferences "github.com/go-flutter-desktop/plugins/shared_preferences"
)

func init() {
	// Only the init function can be tweaked by plugin maker.
	options = append(options, flutter.AddPlugin(&shared_preferences.SharedPreferencesPlugin{
		VendorName:      flutter.ProjectOrganizationName,
		ApplicationName: flutter.ProjectName,
	}))
}