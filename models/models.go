package models

import (
	"github.com/globalsign/mgo/bson"
)

type ORProcess interface {

	// AcademicSessions
	HandleAddAcademicSessions(orAcademicSessions []ORAcademicSessions) error
	HandleDeleteAcademicSessions(orAcademicSessionsIDs []string) error
	HandleEditAcademicSessions(orAcademicSessions []ORAcademicSessions) error
	HandleAddOrEditAcademicSessions(orAcademicSessions []ORAcademicSessions) error

	// Users
	HandleAddUsers(orUser []ORUser, districtIDs []bson.ObjectId) error
	HandleDeleteUsers(oruserIDs []string, districtIDs []bson.ObjectId) error
	HandleEditUsers(orUser []ORUser) error
	HandleAddOrEditUsers(orUser []ORUser, districtIDs []bson.ObjectId) error

	// Districts
	HandleAddDistrict(orOrg OROrg) (bool, error)
	HandleDeleteDistrict(orOrg OROrg) error
	HandleEditDistrict(orOrg OROrg, districtId bson.ObjectId) error
	HandleAddOrEditDistrict(orOrg OROrg) error

	// Schools
	HandleAddSchool(orOrg OROrg, districtIDs []bson.ObjectId) error
	HandleDeleteSchool(orOrg OROrg, districtIDs []bson.ObjectId) error
	HandleEditSchool(orOrg OROrg) error
	HandleAddOrEditSchool(orOrg OROrg, districtIDs []bson.ObjectId) error

	// Classes
	HandleAddClasses(orClass []ORClass, districtIDs []bson.ObjectId) error
	HandleDeleteClasses(orClassIDs []string, districtIDs []bson.ObjectId) error
	HandleEditClass(orClass []ORClass) error
	HandleAddOrEditClass(orClass []ORClass, districtIDs []bson.ObjectId) error

	// Courses
	HandleAddCourses(orCourse []ORCourse, districtIDs []bson.ObjectId) error
	HandleDeleteCourses(orCourseIDs []string, districtIDs []bson.ObjectId) error
	HandleEditCourse(orCourse []ORCourse) error
	HandleAddOrEditCourse(orCourse []ORCourse, districtIDs []bson.ObjectId) error

	// Enrollments
	HandleAddEnrollment(orEnrollment []OREnrollment, districtIDs []bson.ObjectId) error
	HandleDeleteEnrollments(orEnrollment []OREnrollment, districtIDs []bson.ObjectId) error
	HandleAddOrEditEnrollments(orEnrollment []OREnrollment, districtIDs []bson.ObjectId) error

	RollBackOneRoster(orgDistrict []OROrg) error

	GetDistrictsIDs(orOrgs []OROrg) ([]bson.ObjectId, error) 

}


type OrManifest struct {
	PropertyName			string `csv:"propertyName"`
	Value					string `csv:"value"`
}

type ORAcademicSessions struct{
	SourcedId 			string `csv:"sourcedId" json:"sourcedId"`//GUID
	Status 				string `csv:"status" json:"status"`//Enumeration
	DateLastModified 	string `csv:"dateLastModified" json:"dateLastModified"`//DateTime
	Title 				string `csv:"title" json:"title"`
	SessionType 		string `csv:"type" json:"type"`//Enumeration
	StartDate 			string `csv:"startDate" json:"startDate"`//date
	EndDate 			string `csv:"endDate" json:"endDate"`//date
	ParentSourcedId 	string `csv:"parentSourcedId"`//GUID Reference
	Parent 				GUIDRef `json:"parent"`
	Children 			[]GUIDRef `json:"children"`
	SchoolYear 			string `csv:"schoolYear" json:"schoolYear"`//year
}

type ORClass struct{
	SourcedId 			string 			`csv:"sourcedId" json:"sourcedId"`	//GUID
	Status 				string 			`csv:"status" json:"status"`//Enumeration
	DateLastModified 	string 			`csv:"dateLastModified" json:"dateLastModified"`//DateTime
	Title 				string 			`csv:"title" json:"title"`
	Grades 				string 			`csv:"grades" json:""` //[]string
	CourseSourcedId 	string 			`csv:"courseSourcedId"`//GUID Reference
	Course 				GUIDRef 		`json:"course"`
	ClassCode 			string			`csv:"classCode" json:"classCode"`
	ClassType 			string 			`csv:"classType" json:"classType"`//Enumeration
	Location 			string 			`csv:"location" json:"location"`
	SchoolSourcedId 	string 			`csv:"schoolSourcedId" `//GUID Reference
	School				GUIDRef			`json:"school"`	
	TermSourcedIds 		string 		`csv:"termSourcedIds" `//List of GUID Reference
	Terms				[]GUIDRef	`json:"terms"`
	Subjects 			string 		`csv:"subjects" json:"subjects"`
	SubjectCodes 		string		`csv:"subjectCodes" json:"subjectCodes"`
	Periods 			string		`csv:"periods" json:"periods"`
	Resources			[]GUIDRef	`json:"resources"`
}


type ORCourse struct {
	SourcedId 			string  	`csv:"sourcedId" json:"sourcedId"`		//GUID
	Status				string 		`csv:"status" json:"status"`		//Enumeration
	DateLastModified 	string 		`csv:"dateLastModified" json:"dateLastModified"`//DateTime
	SchoolYearSourcedId string 		`csv:"schoolYearSourcedId"`//GUID Reference
	SchoolYear			GUIDRef 	`json:"schoolYear"`
	Title 				string 		`csv:"title" json:"title"`
	CourseCode			string		`csv:"courseCode" json:"courseCode"`
	// Grades				*[]string	`csv:"grades"`
	OrgSourcedId		string 		`csv:"orgSourcedId"`//GUID Reference
	Org					GUIDRef 	`json:"org"`
	Subjects			string		`csv:"subjects" json:"subjects"`
	SubjectCodes		string		`csv:"subjectCodes" json:"subjectCodes"`
}

type ORDemographics struct{
	SourcedId 			string 						`csv:"sourcedId" json:"sourcedId"`//GUID
	Status				string 						`csv:"status" json:"status"`//Enumeration
	DateLastModified 	string 						`csv:"dateLastModified" json:"dateLastModified"`//DateTime
	BirthDate			string 						`csv:"birthDate" json:"birthDate"`//date
	Sex					string 						`csv:"sex" json:"sex"`//Enumeration
	AmericanIndianOrAlaskaNative 			string 	`csv:"americanIndianOrAlaskaNative" json:"americanIndianOrAlaskaNative"`//Enumeration
	Asian									string 	`csv:"asian" json:"asian"`//Enumeration
	BlackOrAfricanAmerican					string 	`csv:"blackOrAfricanAmerican" json:"blackOrAfricanAmerican"`//Enumeration
	NativeHawaiianOrOtherPacificIslander 	string 	`csv:"nativeHawaiianOrOtherPacificIslander" json:"nativeHawaiianOrOtherPacificIslander"`//Enumeration
	White									string 	`csv:"white" json:"white"`//Enumeration
	DemographicRaceTwoOrMoreRaces			string 	`csv:"demographicRaceTwoOrMoreRaces" json:"demographicRaceTwoOrMoreRaces"`//Enumeration
	HispanicOrLatinoEthnicity				string 	`csv:"hispanicOrLatinoEthnicity" json:"hispanicOrLatinoEthnicity"`//Enumeration
	CountryOfBirthCode						string	`csv:"countryOfBirthCode" json:"countryOfBirthCode"`
	StateOfBirthAbbreviation				string	`csv:"stateOfBirthAbbreviation" json:"stateOfBirthAbbreviation"`
	CityOfBirth								string	`csv:"cityOfBirth" json:"cityOfBirth"`
	PublicSchoolResidenceStatus 			string	`csv:"publicSchoolResidenceStatus" json:"publicSchoolResidenceStatus"`
}

type OREnrollment struct{
	SourcedId 			string			`csv:"sourcedId" json:"sourcedId"` //GUID
	Status				string 			`csv:"status" json:"status"`//Enumeration
	DateLastModified 	string 			`csv:"dateLastModified" json:"dateLastModified"`//DateTime
	ClassSourcedId		string 			`csv:"classSourcedId"`//GUID Reference
	Class				GUIDRef 		`json:"class"`
	SchoolSourcedId		string 			`csv:"schoolSourcedId"`//GUID Reference
	School				GUIDRef 		`json:"school"`
	UserSourcedId 		string 			`csv:"userSourcedId"`//GUID Reference
	User		 		GUIDRef 		`json:"user"`
	Role				string 			`csv:"role" json:"role"`//Enumeration
	Primary				bool			`csv:"primary" json:"primary"`
	BeginDate			string 			`csv:"beginDate" json:"beginDate"`//date
	EndDate				string 			`csv:"endDate" json:"endDate"`//date
}

type OROrg struct{
	SourcedId 			string `csv:"sourcedId" json:"sourcedId"`	//GUID
	Status 				string `csv:"status" json:"status"`//Enumeration
	DateLastModified 	string `csv:"dateLastModified" json:"dateLastModified"`//DateTime
	Name 				string `csv:"name" json:"name"`
	OrgType 			string `csv:"type" json:"type"`// type Enumeration
	Identifier 			string `csv:"identifier" json:"identifier"`
	ParentSourcedId 	string `csv:"parentSourcedId"`	//GUID Reference
	Parent 				GUIDRef 	`json:"parent"`
	Children 			[]GUIDRef 	`json:"children"`
}


type ORUser struct {
	SourcedId 			string 	`csv:"sourcedId" json:"sourcedId"`	//GUID
	Status 				string 	`csv:"status" json:"status"`		//Enumeration
	DateLastModified 	string 	`csv:"dateLastModified" json:"dateLastModified"`//DateTime
	EnabledUser 		bool	`csv:"enabledUser" json:"enabledUser"`
	OrgSourcedIds 		string 	`csv:"orgSourcedIds"`//List of GUID References.
	Orgs				[]GUIDRef	`json:"orgs"`
	Role 				string 	`csv:"role" json:"role"`		//Enumeration
	Username 			string	`csv:"username" json:"username"`
	UserIds 			string  `csv:"userIds"` //[] string
	UserIdsIdentifer 	[]UserIdentifer		`json:"userIds"`
	GivenName 			string	`csv:"givenName" json:"givenName"`
	FamilyName 			string	`csv:"familyName" json:"familyName"`
	MiddleName 			string	`csv:"middleName" json:"middleName"`
	Identifier 			string 	`csv:"identifier" json:"identifier"`
	Email 				string	`csv:"email" json:"email"`
	Sms 				string	`csv:"sms" json:"sms"`
	Phone 				string	`csv:"phone" json:"phone"`
	AgentSourcedIds 	string 	`csv:"agentSourcedIds"`//List of GUID References
	Agents			 	[]GUIDRef 	`json:"agents"`
	Grades 				string 	`csv:"grades" json:"grades"`
	Password 			string	`csv:"password" json:"password"`

}

type ORCategory struct{
	SourcedId 			string 	//GUID
	Status 				string 	//Enumeration
	DateLastModified 	string  //DateTime
	Title 				string 
}

type ORClassResources struct{
	SourcedId 			string 		//GUID
	Status				string 			//Enumeration
	DateLastModified 	string //DateTime
	Title 				string 
	ClassSourcedId 		string //GUID Reference
	ResourceSourcedId 	string //GUID Reference
}

type ORCourseResources struct{
	SourcedId 			string 		//GUID
	Status				string 			//Enumeration
	DateLastModified 	string //DateTime
	Title 				string 
	CourseSourcedId 	string //GUID Reference
	ResourceSourcedId 	string //GUID Reference
}

type ORResource struct{
	SourcedId 			string 		//GUID
	Status 				string 			//Enumeration
	DateLastModified 	string //DateTime
	VendorResourceId 	string //id
	Title				string 
	Roles				[]string //Enumeration List
	Importance			string 
	VendorId			string //id
	ApplicationId		string //id
}

type ORResult struct {
	SourcedId string 		//GUID
	Status string 			//Enumeration
	DateLastModified string //DateTime
	LineItemSourcedId string//GUID Reference
	StudentSourcedId string //GUID Reference
	ScoreStatus string 		//Enumeration
	Score float64 			//float
	ScoreDate string 		//date
	Comment string

}

type ORLineItems struct {
	SourcedId 			string 			//GUID
	Status				string 			//Enumeration
	DateLastModified 	string //DateTime
	Title 				string 
	Description			string
	SssignDate			string //date
	DueDate				string //date
	ClassSourcedId		string // GUID References
	CategorySourcedId	string // GUID References
	GradingPeriodSourcedId	string // GUID References
	ResultValueMin 		float64
	ResultValueMax		float64
}

// import type 
const (
	IMPORT_TYPE_BULK = "bulk"
	IMPORT_TYPE_DELTA = "delta"
	IMPORT_TYPE_ABSENT = "absent"
)

// manifest property names 
const (
	MANIFEST_PRO_VERSION = "manifest.version"
	MANIFEST_PRO_ONEROSTER_VERSION = "oneroster.version"
	MANIFEST_PRO_FILE_ACADEMICSESSIONS = "file.academicSessions"
	MANIFEST_PRO_FILE_CATEGORIES = "file.categories"
	MANIFEST_PRO_FILE_CLASSES = "file.classes"
	MANIFEST_PRO_FILE_CLASSRESOURCES = "file.classResources"
	MANIFEST_PRO_FILE_COURSES = "file.courses"
	MANIFEST_PRO_FILE_COURSERESOURCES = "file.courseResources"
	MANIFEST_PRO_FILE_DEMOGRAPHICS = "file.demographics"
	MANIFEST_PRO_FILE_ENROLLMENTS = "file.enrollments"
	MANIFEST_PRO_FILE_LINEITEMS = "file.lineItems"
	MANIFEST_PRO_FILE_ORGS = "file.orgs"
	MANIFEST_PRO_FILE_RESOURCES = "file.resources"
	MANIFEST_PRO_FILE_RESULTS = "file.results"
	MANIFEST_PRO_FILE_USERS = "file.users"
	MANIFEST_PRO_SOURCE_SYSTEMNAME = "source.systemName"
	MANIFEST_PRO_SOURCE_SYSTEMCODE = "source.systemCode"
)

//csv files name 
const (
	CSV_NAME_MANIFEST = "manifest.csv"
	CSV_NAME_ACADEMICSESSIONS = "academicSessions.csv"
	CSV_NAME_CATEGORIES = "categories.csv"
	CSV_NAME_CLASSES = "classes.csv"
	CSV_NAME_COURSES = "courses.csv"
	CSV_NAME_CLASSRESOURCES = "classResources.csv"
	CSV_NAME_DEMOGRAPHICS = "demographics.csv"
	CSV_NAME_ENROLLMENTS = "enrollments.csv"
	CSV_NAME_ORGS = "orgs.csv"
	CSV_NAME_RESOURCES = "resources.csv"
	CSV_NAME_LINEITEMS = "lineItems.csv"
	CSV_NAME_RESULTS = "results.csv"
	CSV_NAME_USERS = "users.csv"
)

//orgs types
const (
	ORG_TYPE_DISTRICT = "district"
	ORG_TYPE_SCHOOL = "school"

)

//Status types
const (
	STATUS_TYPE_ACTIVE = "Active"
	STATUS_TYPE_TOBEDELETED = "ToBeDeleted"
)



////// JSON ///////
type GUIDRef struct{
	Href		string			`json:"href"`
	SourcedId	string 			`json:"sourcedId"`
	GUIDType	string 			`json:"type"`
}

type UserIdentifer struct {
	Type 		string 			`json:"type"`
	Identifier 	string 			`json:"identifier"`
}

const (
	GUIDTYPE_ACADEMICSESSION = "academicSession"
	GUIDTYPE_CATEGORY = "category"
	GUIDTYPE_CLASS = "class"
	GUIDTYPE_COURSE = "course"
	GUIDTYPE_DEMOGRAPHICS = "demographics"
	GUIDTYPE_ENROLLMENT = "enrollment"
	GUIDTYPE_ORG = "org"
	GUIDTYPE_RESOURCE = "resource"
	GUIDTYPE_LINEITEM = "lineItem"
	GUIDTYPE_RESULT = "result"
	GUIDTYPE_USER = "user"
	GUIDTYPE_STUDENT = "student"
	GUIDTYPE_TEACHER = "teacher"
	GUIDTYPE_TERM = "term"
	GUIDTYPE_GRADINGPERIOD = "gardingPeriod"
)


//// Rest API Responses //// 

type OrgsResponse struct {
    Orgs []OROrg `json:"orgs"`
}

type AcademicSessionsResponse struct {
    AcademicSessions []ORAcademicSessions `json:"academicSessions"`
}
type ClassesResponse struct {
    Classes []ORClass `json:"classes"`
}
type CoursesResponse struct {
    Courses []ORCourse `json:"courses"`
}
type EnrollmentsResponse struct {
    Enrollments []OREnrollment `json:"enrollments"`
}

type UsersResponse struct {
    Users []ORUser `json:"users"`
}

