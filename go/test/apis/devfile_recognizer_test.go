package recognizer

/*******************************************************************************
 * Copyright (c) 2022 Red Hat, Inc.
 * Distributed under license by Red Hat, Inc. All rights reserved.
 * This program is made available under the terms of the
 * Eclipse Public License v2.0 which accompanies this distribution,
 * and is available at http://www.eclipse.org/legal/epl-v20.html
 *
 * Contributors:
 * Red Hat, Inc.
 ******************************************************************************/
import (
	"testing"

	"github.com/redhat-developer/alizer/go/pkg/apis/language"
	"github.com/redhat-developer/alizer/go/pkg/apis/recognizer"
)

func TestDetectQuarkusDevfile(t *testing.T) {
	detectDevFile(t, "quarkus", "java-quarkus")
}

func TestDetectMicronautDevfile(t *testing.T) {
	detectDevFile(t, "micronaut", "java-maven")
}

func TestDetectNodeJSDevfile(t *testing.T) {
	detectDevFile(t, "nodejs-ex", "nodejs")
}

func TestDetectDjangoDevfile(t *testing.T) {
	detectDevFile(t, "django", "python-django")
}

func TestDetectDjangoDevfileUsingLanguages(t *testing.T) {
	languages := []language.Language{
		{
			Name: "Python",
			Aliases: []string{
				"python3",
			},
			UsageInPercentage: 88.23,
			Frameworks: []string{
				"Django",
			},
			Tools:          []string{},
			CanBeComponent: false,
		},
		{
			Name: "Shell",
			Aliases: []string{
				"sh",
			},
			UsageInPercentage: 11.77,
			Frameworks:        []string{},
			Tools:             []string{},
			CanBeComponent:    false,
		},
	}
	detectDevFileUsingLanguages(t, "", languages, "python-django")
}

func TestDetectQuarkusDevfileUsingLanguages(t *testing.T) {
	detectDevFileUsingLanguages(t, "quarkus", []language.Language{}, "java-quarkus")
}

func TestDetectMicronautDevfileUsingLanguages(t *testing.T) {
	detectDevFileUsingLanguages(t, "micronaut", []language.Language{}, "java-maven")
}

func TestDetectNodeJSDevfileUsingLanguages(t *testing.T) {
	detectDevFileUsingLanguages(t, "nodejs-ex", []language.Language{}, "nodejs")
}

func TestDetectGoDevfile(t *testing.T) {
	detectDevFile(t, "golang-gin-app", "go")
}

func detectDevFile(t *testing.T, projectName string, devFileName string) {
	detectDevFileFunc := func(devFileTypes []recognizer.DevFileType) (recognizer.DevFileType, error) {
		testingProjectPath := GetTestProjectPath(projectName)
		return recognizer.SelectDevFileFromTypes(testingProjectPath, devFileTypes)
	}
	detectDevFileInner(t, devFileName, detectDevFileFunc)
}

func detectDevFileUsingLanguages(t *testing.T, projectName string, languages []language.Language, devFileName string) {
	if projectName != "" {
		testingProjectPath := GetTestProjectPath(projectName)
		var err error
		languages, err = recognizer.Analyze(testingProjectPath)
		if err != nil {
			t.Error(err)
		}
	}
	detectDevFileFunc := func(devFileTypes []recognizer.DevFileType) (recognizer.DevFileType, error) {
		return recognizer.SelectDevFileUsingLanguagesFromTypes(languages, devFileTypes)
	}
	detectDevFileInner(t, devFileName, detectDevFileFunc)
}

func detectDevFileInner(t *testing.T, devFileName string, detectFuncInner func([]recognizer.DevFileType) (recognizer.DevFileType, error)) {
	devFileTypes := getDevFileTypes()
	devFileType, err := detectFuncInner(devFileTypes)
	if err != nil {
		t.Error(err)
	}

	if devFileType.Name != devFileName {
		t.Error("Expected value " + devFileName + " but it was" + devFileType.Name)
	}
}

func getDevFileTypes() []recognizer.DevFileType {
	return []recognizer.DevFileType{
		{
			Name:        "java",
			Language:    "java",
			ProjectType: "java",
			Tags:        make([]string, 0),
		},
		{
			Name:        "java-quarkus",
			Language:    "java",
			ProjectType: "quarkus",
			Tags: []string{
				"Java",
				"Quarkus",
			},
		},
		{
			Name:        "java-maven",
			Language:    "java",
			ProjectType: "java",
			Tags: []string{
				"Java",
				"Maven",
			},
		},
		{
			Name:        "java-spring",
			Language:    "java",
			ProjectType: "spring",
			Tags: []string{
				"Java",
				"Spring",
			},
		},
		{
			Name:        "java-vertx",
			Language:    "java",
			ProjectType: "vertx",
			Tags: []string{
				"Java",
				"Vert.x",
			},
		},
		{
			Name:        "java-wildfly",
			Language:    "java",
			ProjectType: "wildfly",
			Tags: []string{
				"Java",
				"Wildfly",
			},
		},
		{
			Name:        "nodejs",
			Language:    "nodejs",
			ProjectType: "nodejs",
			Tags: []string{
				"NodeJS",
				"Express",
			},
		},
		{
			Name:        "python-django",
			Language:    "python",
			ProjectType: "django",
			Tags: []string{
				"Python",
				"pip",
			},
		},
		{
			Name:        "go",
			Language:    "go",
			ProjectType: "go",
			Tags: []string{
				"go",
			},
		},
	}
}
