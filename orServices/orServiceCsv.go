package orServices

import (
	"fmt"
	"os"

	"github.com/MZDevinc/oneroster-edgems/models"

	// "github.com/jszwec/csvutil"
	"strings"

	"github.com/globalsign/mgo/bson"
	"github.com/gocarina/gocsv"
)

func ProcessFiles(dirPath string, orProcess models.ORProcess) error {

	// read the manifest csv file
	manifestPath := fmt.Sprintf("%s/manifest.csv", dirPath)
	manifestRows, err := ReadManifestCSV(manifestPath)
	if err != nil {
		fmt.Println(">> err ReadManifestCsv: ", err)
		return err
	}

	var orgDistrict []models.OROrg
	var districtIDs []bson.ObjectId

	// put the manifest data into   --- map[propertyName] = propertyValue---
	mainfestTable := make(map[string]string)
	for _, manifestRow := range manifestRows {
		switch manifestRow.PropertyName {

		// class sourceName
		case models.MANIFEST_PRO_SOURCE_SYSTEMNAME:
			mainfestTable[models.MANIFEST_PRO_SOURCE_SYSTEMNAME] = manifestRow.Value

		// Acadimic Sessions
		case models.MANIFEST_PRO_FILE_ACADEMICSESSIONS:
			mainfestTable[models.MANIFEST_PRO_FILE_ACADEMICSESSIONS] = manifestRow.Value

		// classes
		case models.MANIFEST_PRO_FILE_CLASSES:
			mainfestTable[models.MANIFEST_PRO_FILE_CLASSES] = manifestRow.Value

		// courses
		case models.MANIFEST_PRO_FILE_COURSES:
			mainfestTable[models.MANIFEST_PRO_FILE_COURSES] = manifestRow.Value

		// Enrollments
		case models.MANIFEST_PRO_FILE_ENROLLMENTS:
			mainfestTable[models.MANIFEST_PRO_FILE_ENROLLMENTS] = manifestRow.Value

		// Orgs
		case models.MANIFEST_PRO_FILE_ORGS:
			mainfestTable[models.MANIFEST_PRO_FILE_ORGS] = manifestRow.Value

		// Users
		case models.MANIFEST_PRO_FILE_USERS:
			mainfestTable[models.MANIFEST_PRO_FILE_USERS] = manifestRow.Value

		// Demographics ( we don't save it in Edgems, maybe we will need it later )
		case models.MANIFEST_PRO_FILE_DEMOGRAPHICS:
			mainfestTable[models.MANIFEST_PRO_FILE_DEMOGRAPHICS] = manifestRow.Value

			// we don't read the resources and results, we don't need it until now

			// case models.MANIFEST_PRO_FILE_RESULTS:
			// mainfestTable[models.MANIFEST_PRO_FILE_RESULTS] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_RESOURCES:
			// mainfestTable[models.MANIFEST_PRO_FILE_RESOURCES] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_LINEITEMS:
			// mainfestTable[models.MANIFEST_PRO_FILE_LINEITEMS] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_COURSERESOURCES:
			// mainfestTable[models.MANIFEST_PRO_FILE_COURSERESOURCES] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_CLASSRESOURCES:
			// mainfestTable[models.MANIFEST_PRO_FILE_CLASSRESOURCES] = manifestRow.Value
			// case models.MANIFEST_PRO_FILE_CATEGORIES:
			// mainfestTable[models.MANIFEST_PRO_FILE_CATEGORIES] = manifestRow.Value
		}

	}

	// the files should be readed in order
	//process Ditricts and schools
	if mainfestTable[models.MANIFEST_PRO_FILE_ORGS] != models.IMPORT_TYPE_ABSENT {
		doRollback := false
		if strings.Contains(strings.ToLower(mainfestTable[models.MANIFEST_PRO_SOURCE_SYSTEMNAME]), "classlink") {
			orgDistrict, districtIDs, doRollback, err = ProcessOrgsClassLinkCSV(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_ORGS])
		} else {
			orgDistrict, districtIDs, doRollback, err = ProcessOrgsCSV(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_ORGS])
		}

		if err != nil {
			if doRollback {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}
	//process Courses
	if mainfestTable[models.MANIFEST_PRO_FILE_COURSES] != models.IMPORT_TYPE_ABSENT {
		err = ProcessCoursesCSV(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_COURSES], districtIDs)
		if err != nil {
			if mainfestTable[models.MANIFEST_PRO_FILE_COURSES] != models.IMPORT_TYPE_BULK {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}

	//process Academic Session
	if mainfestTable[models.MANIFEST_PRO_FILE_ACADEMICSESSIONS] != models.IMPORT_TYPE_ABSENT {
		err = ProcessAcademicSessionsCSV(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_ACADEMICSESSIONS])
		if err != nil {
			if mainfestTable[models.MANIFEST_PRO_FILE_ACADEMICSESSIONS] != models.IMPORT_TYPE_BULK {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}
	//process Classes
	if mainfestTable[models.MANIFEST_PRO_FILE_CLASSES] != models.IMPORT_TYPE_ABSENT {
		err = ProcessClassesCSV(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_CLASSES], districtIDs)
		if err != nil {
			if mainfestTable[models.MANIFEST_PRO_FILE_CLASSES] != models.IMPORT_TYPE_BULK {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}

	//process Users
	if mainfestTable[models.MANIFEST_PRO_FILE_USERS] != models.IMPORT_TYPE_ABSENT {
		err = ProcessUsersCSV(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_USERS], districtIDs)
		if err != nil {
			if mainfestTable[models.MANIFEST_PRO_FILE_USERS] != models.IMPORT_TYPE_BULK {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}

	//process User Entrollments
	if mainfestTable[models.MANIFEST_PRO_FILE_ENROLLMENTS] != models.IMPORT_TYPE_ABSENT {
		err = ProcessEntrollmentCSV(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_ENROLLMENTS], districtIDs)
		if err != nil {
			if mainfestTable[models.MANIFEST_PRO_FILE_ENROLLMENTS] != models.IMPORT_TYPE_BULK {
				err = orProcess.RollBackOneRoster(orgDistrict)
			}
			return err
		}
	}

	//process Demographics  (we don't use it right now)
	// if mainfestTable[models.MANIFEST_PRO_FILE_DEMOGRAPHICS] != models.IMPORT_TYPE_ABSENT {
	// 	err = ProcessDemographics(dirPath, orProcess, mainfestTable[models.MANIFEST_PRO_FILE_DEMOGRAPHICS])
	// 	if err != nil {
	// 		fmt.Println(">>> (rollback) errer happen when ProcessDemographics err -> ",err)
	// 		err = orProcess.RollBackOneRoster(orgDistrict)
	// 		return err
	// 	}
	// }

	return nil
}

func ReadManifestCSV(filename string) ([]models.OrManifest, error) {
	// Open CSV file
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	var manifestValues []models.OrManifest

	err = gocsv.UnmarshalFile(f, &manifestValues)
	if err != nil {
		return nil, err
	}
	return manifestValues, nil
}

func ProcessAcademicSessionsCSV(dirPath string, orProcess models.ORProcess, importType string) error {

	academicSessionsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_ACADEMICSESSIONS)

	f, err := os.Open(academicSessionsPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var academicSessions []models.ORAcademicSessions

	err = gocsv.UnmarshalFile(f, &academicSessions)
	if err != nil {
		return err
	}
	if importType == models.IMPORT_TYPE_BULK {

		err := orProcess.HandleAddAcademicSessions(academicSessions)
		if err != nil {
			return err
		}
	} else if importType == models.IMPORT_TYPE_DELTA {
		orAcademicSessionToEdit := []models.ORAcademicSessions{}
		orAcademicSessionIDsToDelete := []string{}
		for _, orAcademicSession := range academicSessions {
			if orAcademicSession.Status == models.STATUS_TYPE_ACTIVE {
				orAcademicSessionToEdit = append(orAcademicSessionToEdit, orAcademicSession)
				// err = orProcess.HandleEditClass(orClass)
			} else if orAcademicSession.Status == models.STATUS_TYPE_TOBEDELETED {
				orAcademicSessionIDsToDelete = append(orAcademicSessionIDsToDelete, orAcademicSession.SourcedId)
			}
			if err != nil {
				return err
			}
		}
		err = orProcess.HandleEditAcademicSessions(orAcademicSessionToEdit)
		err = orProcess.HandleDeleteAcademicSessions(orAcademicSessionIDsToDelete)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessOrgsCSV(dirPath string, orProcess models.ORProcess, importType string) ([]models.OROrg, []bson.ObjectId, bool, error) {

	var orgDistricts []models.OROrg
	var districtIDs []bson.ObjectId
	// do rollback for all district or not
	rollback := true
	orgsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_ORGS)

	f, err := os.Open(orgsPath)
	if err != nil {
		return orgDistricts, districtIDs, rollback, err
	}
	defer f.Close()
	var orgs []models.OROrg

	err = gocsv.UnmarshalFile(f, &orgs)
	if err != nil {
		fmt.Println(err)
		return orgDistricts, districtIDs, rollback, err
	}

	if importType == models.IMPORT_TYPE_BULK {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.ORG_TYPE_DISTRICT {
				// collect all district
				orgDistricts = append(orgDistricts, org)

				rollback, err = orProcess.HandleAddDistrict(org)
				if err != nil {
					return orgDistricts, districtIDs, rollback, err
				}
			}
		}

		// get the mongo IDs for the district to use for edit and delete other files data
		districtIDs, err := orProcess.GetDistrictsIDs(orgDistricts)
		if err != nil {
			return orgDistricts, districtIDs, true, err
		}

		for _, org := range orgs {

			if org.OrgType == models.ORG_TYPE_SCHOOL {
				err = orProcess.HandleAddSchool(org, districtIDs)
				if err != nil {
					return orgDistricts, districtIDs, true, err
				}
			}
		}
	} else if importType == models.IMPORT_TYPE_DELTA {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.ORG_TYPE_DISTRICT {
				// // collect all district
				orgDistricts = append(orgDistricts, org)

				if org.Status == models.STATUS_TYPE_ACTIVE {
					// err = orProcess.HandleEditDistrict(org)
					err = orProcess.HandleAddOrEditDistrict(org)
				} else if org.Status == models.STATUS_TYPE_TOBEDELETED {
					err = orProcess.HandleDeleteDistrict(org)
				}

				if err != nil {
					return orgDistricts, districtIDs, false, err
				}

			}
		}

		// get the mongo IDs for the district to use for edit and delete other files data
		districtIDs, err := orProcess.GetDistrictsIDs(orgDistricts)
		if err != nil {
			return orgDistricts, districtIDs, true, err
		}

		for _, org := range orgs {
			if org.OrgType == models.ORG_TYPE_SCHOOL {
				if org.Status == models.STATUS_TYPE_ACTIVE {
					err = orProcess.HandleAddOrEditSchool(org, districtIDs)
				} else if org.Status == models.STATUS_TYPE_TOBEDELETED {
					err = orProcess.HandleDeleteSchool(org, districtIDs)
				}

				if err != nil {
					return orgDistricts, districtIDs, false, err
				}
			}
		}

	}

	return orgDistricts, districtIDs, false, nil
}

func ProcessOrgsClassLinkCSV(dirPath string, orProcess models.ORProcess, importType string) ([]models.OROrg, []bson.ObjectId, bool, error) {

	var orgDistricts []models.OROrg
	var districtIDs []bson.ObjectId
	// do rollback for all district or not
	rollback := true
	orgsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_ORGS)

	f, err := os.Open(orgsPath)
	if err != nil {
		return orgDistricts, districtIDs, rollback, err
	}
	defer f.Close()
	var orgs []models.OROrg

	err = gocsv.UnmarshalFile(f, &orgs)
	if err != nil {
		fmt.Println(err)
		return orgDistricts, districtIDs, rollback, err
	}

	if importType == models.IMPORT_TYPE_BULK {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.ORG_TYPE_DISTRICT {
				// collect all district
				orgDistricts = append(orgDistricts, org)

				rollback, err = orProcess.HandleAddDistrict(org)
				if err != nil {
					return orgDistricts, districtIDs, rollback, err
				}
			}
		}

		districtIDs, err := orProcess.GetDistrictsIDs(orgDistricts)
		if err != nil {
			return orgDistricts, districtIDs, true, err
		}

		for _, org := range orgs {
			if org.OrgType == models.ORG_TYPE_SCHOOL {
				if len(orgDistricts) == 1 && org.ParentSourcedId == "" {
					org.ParentSourcedId = orgDistricts[0].SourcedId
				}
				err = orProcess.HandleAddSchool(org, districtIDs)
				if err != nil {
					return orgDistricts, districtIDs, true, err
				}
			}
		}
	} else if importType == models.IMPORT_TYPE_DELTA {
		for _, org := range orgs {
			var err error = nil
			if org.OrgType == models.ORG_TYPE_DISTRICT {
				// collect all district
				orgDistricts = append(orgDistricts, org)

				if org.Status == models.STATUS_TYPE_ACTIVE {
					// err = orProcess.HandleEditDistrict(org)
					err = orProcess.HandleAddOrEditDistrict(org)
				} else if org.Status == models.STATUS_TYPE_TOBEDELETED {
					err = orProcess.HandleDeleteDistrict(org)
				}

				if err != nil {
					return orgDistricts, districtIDs, false, err
				}

			}
		}

		districtIDs, err := orProcess.GetDistrictsIDs(orgDistricts)
		if err != nil {
			return orgDistricts, districtIDs, true, err
		}

		for _, org := range orgs {
			if org.OrgType == models.ORG_TYPE_SCHOOL {
				if org.Status == models.STATUS_TYPE_ACTIVE {
					err = orProcess.HandleAddOrEditSchool(org, districtIDs)
				} else if org.Status == models.STATUS_TYPE_TOBEDELETED {
					err = orProcess.HandleDeleteSchool(org, districtIDs)
				}

				if err != nil {
					return orgDistricts, districtIDs, false, err
				}
			}
		}
	}

	return orgDistricts, districtIDs, false, nil
}

func ProcessCoursesCSV(dirPath string, orProcess models.ORProcess, importType string, districtIDs []bson.ObjectId) error {

	coursesPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_COURSES)

	f, err := os.Open(coursesPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orCourses []models.ORCourse
	err = gocsv.UnmarshalFile(f, &orCourses)
	if err != nil {
		return err
	}

	if importType == models.IMPORT_TYPE_BULK {
		err := orProcess.HandleAddCourses(orCourses, districtIDs)
		if err != nil {
			return err
		}
	} else if importType == models.IMPORT_TYPE_DELTA {
		orCourseToEdit := []models.ORCourse{}
		orCoursesIDsToDelete := []string{}
		for _, orCourse := range orCourses {
			if orCourse.Status == models.STATUS_TYPE_ACTIVE {
				// err = orProcess.HandleEditCourse(orCourse)
				orCourseToEdit = append(orCourseToEdit, orCourse)
			} else if orCourse.Status == models.STATUS_TYPE_TOBEDELETED {
				// err = orProcess.HandleDeleteCourse(orCourse)
				orCoursesIDsToDelete = append(orCoursesIDsToDelete, orCourse.SourcedId)
			}
			if err != nil {
				return err
			}
		}
		err = orProcess.HandleAddOrEditCourse(orCourseToEdit, districtIDs)
		err = orProcess.HandleDeleteCourses(orCoursesIDsToDelete, districtIDs)
		if err != nil {
			return err
		}

	}
	return nil
}

func ProcessClassesCSV(dirPath string, orProcess models.ORProcess, importType string, districtIDs []bson.ObjectId) error {

	classesPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_CLASSES)

	f, err := os.Open(classesPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orClasses []models.ORClass
	err = gocsv.UnmarshalFile(f, &orClasses)
	if err != nil {
		return err
	}

	if importType == models.IMPORT_TYPE_BULK {
		err := orProcess.HandleAddClasses(orClasses, districtIDs)
		if err != nil {
			return err
		}
	} else if importType == models.IMPORT_TYPE_DELTA {
		orClassesToEdit := []models.ORClass{}
		orClassIDsToDelete := []string{}
		for _, orClass := range orClasses {
			if orClass.Status == models.STATUS_TYPE_ACTIVE {
				orClassesToEdit = append(orClassesToEdit, orClass)
				// err = orProcess.HandleEditClass(orClass)
			} else if orClass.Status == models.STATUS_TYPE_TOBEDELETED {
				orClassIDsToDelete = append(orClassIDsToDelete, orClass.SourcedId)
			}
			if err != nil {
				return err
			}
		}
		err = orProcess.HandleAddOrEditClass(orClassesToEdit, districtIDs)
		err = orProcess.HandleDeleteClasses(orClassIDsToDelete, districtIDs)
		if err != nil {
			return err
		}
	}
	return nil
}

func ProcessUsersCSV(dirPath string, orProcess models.ORProcess, importType string, districtIDs []bson.ObjectId) error {

	usersPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_USERS)

	f, err := os.Open(usersPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orUsers []models.ORUser

	err = gocsv.UnmarshalFile(f, &orUsers)
	if err != nil {
		return err
	}
	if importType == models.IMPORT_TYPE_BULK {
		err := orProcess.HandleAddUsers(orUsers, districtIDs)
		if err != nil {
			return err
		}
	} else if importType == models.IMPORT_TYPE_DELTA {
		orUsersToEdit := []models.ORUser{}
		orUsersIDsToDelete := []string{}
		for _, orUser := range orUsers {
			if orUser.Status == models.STATUS_TYPE_ACTIVE {
				orUsersToEdit = append(orUsersToEdit, orUser)
			} else if orUser.Status == models.STATUS_TYPE_TOBEDELETED {
				orUsersIDsToDelete = append(orUsersIDsToDelete, orUser.SourcedId)
			}
			if err != nil {
				return err
			}
		}
		err = orProcess.HandleAddOrEditUsers(orUsersToEdit, districtIDs)
		err = orProcess.HandleDeleteUsers(orUsersIDsToDelete, districtIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessEntrollmentCSV(dirPath string, orProcess models.ORProcess, importType string, districtIDs []bson.ObjectId) error {

	entrollmentsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_ENROLLMENTS)

	f, err := os.Open(entrollmentsPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orEntrollments []models.OREnrollment

	err = gocsv.UnmarshalFile(f, &orEntrollments)
	if err != nil {
		return err
	}
	if importType == models.IMPORT_TYPE_BULK {
		err := orProcess.HandleAddEnrollment(orEntrollments, districtIDs)
		if err != nil {
			return err
		}
	} else if importType == models.IMPORT_TYPE_DELTA {

		orEntrollmentsToEdit := []models.OREnrollment{}
		orEntrollmentsIDsToDelete := []models.OREnrollment{}
		for _, orEntrollment := range orEntrollments {
			if orEntrollment.Status == models.STATUS_TYPE_ACTIVE {
				orEntrollmentsToEdit = append(orEntrollmentsToEdit, orEntrollment)
			} else if orEntrollment.Status == models.STATUS_TYPE_TOBEDELETED {
				orEntrollmentsIDsToDelete = append(orEntrollmentsIDsToDelete, orEntrollment)
			}
			if err != nil {
				return err
			}
		}
		err = orProcess.HandleDeleteEnrollments(orEntrollmentsIDsToDelete, districtIDs)
		err = orProcess.HandleAddOrEditEnrollments(orEntrollmentsToEdit, districtIDs)
		if err != nil {
			return err
		}
	}

	return nil
}

func ProcessDemographicsCSV(dirPath string, orProcess models.ORProcess, importType string) error {

	demographicsPath := fmt.Sprintf("%s/%s", dirPath, models.CSV_NAME_DEMOGRAPHICS)

	f, err := os.Open(demographicsPath)
	if err != nil {
		return err
	}
	defer f.Close()
	var orDemographics []models.ORDemographics

	err = gocsv.UnmarshalFile(f, &orDemographics)
	if err != nil {
		return err
	}
	if importType == models.IMPORT_TYPE_BULK {
		// err := orProcess.HandleAddorDemographics(orEntrollments)
		// if err != nil {
		// 	fmt.Println(">>> ProcessEntrollments error ",err)
		// 	return err
		// }
	} else if importType == models.IMPORT_TYPE_DELTA {
		fmt.Println(">> *** Delta *** ProcessDemographics")
	}

	return nil
}
