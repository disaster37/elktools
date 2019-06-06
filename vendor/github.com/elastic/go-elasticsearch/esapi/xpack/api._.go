// Code generated from specification version 7.0.0-SNAPSHOT (3499078): DO NOT EDIT

package xpack

// API contains the Elasticsearch APIs
//
type XAPI struct {
	Ccr        *Ccr
	Ilm        *Ilm
	Indices    *Indices
	License    *License
	Migration  *Migration
	Ml         *Ml
	Monitoring *Monitoring
	Rollup     *Rollup
	Security   *Security
	Sql        *Sql
	Ssl        *Ssl
	Watcher    *Watcher
	Xpack      *Xpack

	GraphExplore GraphExplore
}

// Ccr contains the Ccr APIs
type Ccr struct {
	DeleteAutoFollowPattern CcrDeleteAutoFollowPattern
	Follow                  CcrFollow
	FollowInfo              CcrFollowInfo
	FollowStats             CcrFollowStats
	ForgetFollower          CcrForgetFollower
	GetAutoFollowPattern    CcrGetAutoFollowPattern
	PauseFollow             CcrPauseFollow
	PutAutoFollowPattern    CcrPutAutoFollowPattern
	ResumeFollow            CcrResumeFollow
	Stats                   CcrStats
	Unfollow                CcrUnfollow
}

// Ilm contains the Ilm APIs
type Ilm struct {
	DeleteLifecycle  IlmDeleteLifecycle
	ExplainLifecycle IlmExplainLifecycle
	GetLifecycle     IlmGetLifecycle
	GetStatus        IlmGetStatus
	MoveToStep       IlmMoveToStep
	PutLifecycle     IlmPutLifecycle
	RemovePolicy     IlmRemovePolicy
	Retry            IlmRetry
	Start            IlmStart
	Stop             IlmStop
}

// Indices contains the Indices APIs
type Indices struct {
	Freeze   IndicesFreeze
	Unfreeze IndicesUnfreeze
}

// License contains the License APIs
type License struct {
	Delete         LicenseDelete
	Get            LicenseGet
	GetBasicStatus LicenseGetBasicStatus
	GetTrialStatus LicenseGetTrialStatus
	Post           LicensePost
	PostStartBasic LicensePostStartBasic
	PostStartTrial LicensePostStartTrial
}

// Migration contains the Migration APIs
type Migration struct {
	Deprecations MigrationDeprecations
}

// Ml contains the Ml APIs
type Ml struct {
	CloseJob            MlCloseJob
	DeleteCalendar      MlDeleteCalendar
	DeleteCalendarEvent MlDeleteCalendarEvent
	DeleteCalendarJob   MlDeleteCalendarJob
	DeleteDatafeed      MlDeleteDatafeed
	DeleteExpiredData   MlDeleteExpiredData
	DeleteFilter        MlDeleteFilter
	DeleteForecast      MlDeleteForecast
	DeleteJob           MlDeleteJob
	DeleteModelSnapshot MlDeleteModelSnapshot
	FindFileStructure   MlFindFileStructure
	FlushJob            MlFlushJob
	Forecast            MlForecast
	GetBuckets          MlGetBuckets
	GetCalendarEvents   MlGetCalendarEvents
	GetCalendars        MlGetCalendars
	GetCategories       MlGetCategories
	GetDatafeedStats    MlGetDatafeedStats
	GetDatafeeds        MlGetDatafeeds
	GetFilters          MlGetFilters
	GetInfluencers      MlGetInfluencers
	GetJobStats         MlGetJobStats
	GetJobs             MlGetJobs
	GetModelSnapshots   MlGetModelSnapshots
	GetOverallBuckets   MlGetOverallBuckets
	GetRecords          MlGetRecords
	Info                MlInfo
	OpenJob             MlOpenJob
	PostCalendarEvents  MlPostCalendarEvents
	PostData            MlPostData
	PreviewDatafeed     MlPreviewDatafeed
	PutCalendar         MlPutCalendar
	PutCalendarJob      MlPutCalendarJob
	PutDatafeed         MlPutDatafeed
	PutFilter           MlPutFilter
	PutJob              MlPutJob
	RevertModelSnapshot MlRevertModelSnapshot
	SetUpgradeMode      MlSetUpgradeMode
	StartDatafeed       MlStartDatafeed
	StopDatafeed        MlStopDatafeed
	UpdateDatafeed      MlUpdateDatafeed
	UpdateFilter        MlUpdateFilter
	UpdateJob           MlUpdateJob
	UpdateModelSnapshot MlUpdateModelSnapshot
	Validate            MlValidate
	ValidateDetector    MlValidateDetector
}

// Monitoring contains the Monitoring APIs
type Monitoring struct {
	Bulk MonitoringBulk
}

// Rollup contains the Rollup APIs
type Rollup struct {
	DeleteJob          RollupDeleteJob
	GetJobs            RollupGetJobs
	GetRollupCaps      RollupGetRollupCaps
	GetRollupIndexCaps RollupGetRollupIndexCaps
	PutJob             RollupPutJob
	RollupSearch       RollupRollupSearch
	StartJob           RollupStartJob
	StopJob            RollupStopJob
}

// Security contains the Security APIs
type Security struct {
	Authenticate      SecurityAuthenticate
	ChangePassword    SecurityChangePassword
	ClearCachedRealms SecurityClearCachedRealms
	ClearCachedRoles  SecurityClearCachedRoles
	CreateApiKey      SecurityCreateApiKey
	DeletePrivileges  SecurityDeletePrivileges
	DeleteRole        SecurityDeleteRole
	DeleteRoleMapping SecurityDeleteRoleMapping
	DeleteUser        SecurityDeleteUser
	DisableUser       SecurityDisableUser
	EnableUser        SecurityEnableUser
	GetApiKey         SecurityGetApiKey
	GetPrivileges     SecurityGetPrivileges
	GetRole           SecurityGetRole
	GetRoleMapping    SecurityGetRoleMapping
	GetToken          SecurityGetToken
	GetUser           SecurityGetUser
	GetUserPrivileges SecurityGetUserPrivileges
	HasPrivileges     SecurityHasPrivileges
	InvalidateApiKey  SecurityInvalidateApiKey
	InvalidateToken   SecurityInvalidateToken
	PutPrivileges     SecurityPutPrivileges
	PutRole           SecurityPutRole
	PutRoleMapping    SecurityPutRoleMapping
	PutUser           SecurityPutUser
}

// Sql contains the Sql APIs
type Sql struct {
	ClearCursor SqlClearCursor
	Query       SqlQuery
	Translate   SqlTranslate
}

// Ssl contains the Ssl APIs
type Ssl struct {
	Certificates SslCertificates
}

// Watcher contains the Watcher APIs
type Watcher struct {
	AckWatch        WatcherAckWatch
	ActivateWatch   WatcherActivateWatch
	DeactivateWatch WatcherDeactivateWatch
	DeleteWatch     WatcherDeleteWatch
	ExecuteWatch    WatcherExecuteWatch
	GetWatch        WatcherGetWatch
	PutWatch        WatcherPutWatch
	Start           WatcherStart
	Stats           WatcherStats
	Stop            WatcherStop
}

// Xpack contains the Xpack APIs
type Xpack struct {
	Info  XpackInfo
	Usage XpackUsage
}

// New creates new API
//
func New(t Transport) *XAPI {
	return &XAPI{
		GraphExplore: newGraphExploreFunc(t),
		Ccr: &Ccr{
			DeleteAutoFollowPattern: newCcrDeleteAutoFollowPatternFunc(t),
			Follow:                  newCcrFollowFunc(t),
			FollowInfo:              newCcrFollowInfoFunc(t),
			FollowStats:             newCcrFollowStatsFunc(t),
			ForgetFollower:          newCcrForgetFollowerFunc(t),
			GetAutoFollowPattern:    newCcrGetAutoFollowPatternFunc(t),
			PauseFollow:             newCcrPauseFollowFunc(t),
			PutAutoFollowPattern:    newCcrPutAutoFollowPatternFunc(t),
			ResumeFollow:            newCcrResumeFollowFunc(t),
			Stats:                   newCcrStatsFunc(t),
			Unfollow:                newCcrUnfollowFunc(t),
		},
		Ilm: &Ilm{
			DeleteLifecycle:  newIlmDeleteLifecycleFunc(t),
			ExplainLifecycle: newIlmExplainLifecycleFunc(t),
			GetLifecycle:     newIlmGetLifecycleFunc(t),
			GetStatus:        newIlmGetStatusFunc(t),
			MoveToStep:       newIlmMoveToStepFunc(t),
			PutLifecycle:     newIlmPutLifecycleFunc(t),
			RemovePolicy:     newIlmRemovePolicyFunc(t),
			Retry:            newIlmRetryFunc(t),
			Start:            newIlmStartFunc(t),
			Stop:             newIlmStopFunc(t),
		},
		Indices: &Indices{
			Freeze:   newIndicesFreezeFunc(t),
			Unfreeze: newIndicesUnfreezeFunc(t),
		},
		License: &License{
			Delete:         newLicenseDeleteFunc(t),
			Get:            newLicenseGetFunc(t),
			GetBasicStatus: newLicenseGetBasicStatusFunc(t),
			GetTrialStatus: newLicenseGetTrialStatusFunc(t),
			Post:           newLicensePostFunc(t),
			PostStartBasic: newLicensePostStartBasicFunc(t),
			PostStartTrial: newLicensePostStartTrialFunc(t),
		},
		Migration: &Migration{
			Deprecations: newMigrationDeprecationsFunc(t),
		},
		Ml: &Ml{
			CloseJob:            newMlCloseJobFunc(t),
			DeleteCalendar:      newMlDeleteCalendarFunc(t),
			DeleteCalendarEvent: newMlDeleteCalendarEventFunc(t),
			DeleteCalendarJob:   newMlDeleteCalendarJobFunc(t),
			DeleteDatafeed:      newMlDeleteDatafeedFunc(t),
			DeleteExpiredData:   newMlDeleteExpiredDataFunc(t),
			DeleteFilter:        newMlDeleteFilterFunc(t),
			DeleteForecast:      newMlDeleteForecastFunc(t),
			DeleteJob:           newMlDeleteJobFunc(t),
			DeleteModelSnapshot: newMlDeleteModelSnapshotFunc(t),
			FindFileStructure:   newMlFindFileStructureFunc(t),
			FlushJob:            newMlFlushJobFunc(t),
			Forecast:            newMlForecastFunc(t),
			GetBuckets:          newMlGetBucketsFunc(t),
			GetCalendarEvents:   newMlGetCalendarEventsFunc(t),
			GetCalendars:        newMlGetCalendarsFunc(t),
			GetCategories:       newMlGetCategoriesFunc(t),
			GetDatafeedStats:    newMlGetDatafeedStatsFunc(t),
			GetDatafeeds:        newMlGetDatafeedsFunc(t),
			GetFilters:          newMlGetFiltersFunc(t),
			GetInfluencers:      newMlGetInfluencersFunc(t),
			GetJobStats:         newMlGetJobStatsFunc(t),
			GetJobs:             newMlGetJobsFunc(t),
			GetModelSnapshots:   newMlGetModelSnapshotsFunc(t),
			GetOverallBuckets:   newMlGetOverallBucketsFunc(t),
			GetRecords:          newMlGetRecordsFunc(t),
			Info:                newMlInfoFunc(t),
			OpenJob:             newMlOpenJobFunc(t),
			PostCalendarEvents:  newMlPostCalendarEventsFunc(t),
			PostData:            newMlPostDataFunc(t),
			PreviewDatafeed:     newMlPreviewDatafeedFunc(t),
			PutCalendar:         newMlPutCalendarFunc(t),
			PutCalendarJob:      newMlPutCalendarJobFunc(t),
			PutDatafeed:         newMlPutDatafeedFunc(t),
			PutFilter:           newMlPutFilterFunc(t),
			PutJob:              newMlPutJobFunc(t),
			RevertModelSnapshot: newMlRevertModelSnapshotFunc(t),
			SetUpgradeMode:      newMlSetUpgradeModeFunc(t),
			StartDatafeed:       newMlStartDatafeedFunc(t),
			StopDatafeed:        newMlStopDatafeedFunc(t),
			UpdateDatafeed:      newMlUpdateDatafeedFunc(t),
			UpdateFilter:        newMlUpdateFilterFunc(t),
			UpdateJob:           newMlUpdateJobFunc(t),
			UpdateModelSnapshot: newMlUpdateModelSnapshotFunc(t),
			Validate:            newMlValidateFunc(t),
			ValidateDetector:    newMlValidateDetectorFunc(t),
		},
		Monitoring: &Monitoring{
			Bulk: newMonitoringBulkFunc(t),
		},
		Rollup: &Rollup{
			DeleteJob:          newRollupDeleteJobFunc(t),
			GetJobs:            newRollupGetJobsFunc(t),
			GetRollupCaps:      newRollupGetRollupCapsFunc(t),
			GetRollupIndexCaps: newRollupGetRollupIndexCapsFunc(t),
			PutJob:             newRollupPutJobFunc(t),
			RollupSearch:       newRollupRollupSearchFunc(t),
			StartJob:           newRollupStartJobFunc(t),
			StopJob:            newRollupStopJobFunc(t),
		},
		Security: &Security{
			Authenticate:      newSecurityAuthenticateFunc(t),
			ChangePassword:    newSecurityChangePasswordFunc(t),
			ClearCachedRealms: newSecurityClearCachedRealmsFunc(t),
			ClearCachedRoles:  newSecurityClearCachedRolesFunc(t),
			CreateApiKey:      newSecurityCreateApiKeyFunc(t),
			DeletePrivileges:  newSecurityDeletePrivilegesFunc(t),
			DeleteRole:        newSecurityDeleteRoleFunc(t),
			DeleteRoleMapping: newSecurityDeleteRoleMappingFunc(t),
			DeleteUser:        newSecurityDeleteUserFunc(t),
			DisableUser:       newSecurityDisableUserFunc(t),
			EnableUser:        newSecurityEnableUserFunc(t),
			GetApiKey:         newSecurityGetApiKeyFunc(t),
			GetPrivileges:     newSecurityGetPrivilegesFunc(t),
			GetRole:           newSecurityGetRoleFunc(t),
			GetRoleMapping:    newSecurityGetRoleMappingFunc(t),
			GetToken:          newSecurityGetTokenFunc(t),
			GetUser:           newSecurityGetUserFunc(t),
			GetUserPrivileges: newSecurityGetUserPrivilegesFunc(t),
			HasPrivileges:     newSecurityHasPrivilegesFunc(t),
			InvalidateApiKey:  newSecurityInvalidateApiKeyFunc(t),
			InvalidateToken:   newSecurityInvalidateTokenFunc(t),
			PutPrivileges:     newSecurityPutPrivilegesFunc(t),
			PutRole:           newSecurityPutRoleFunc(t),
			PutRoleMapping:    newSecurityPutRoleMappingFunc(t),
			PutUser:           newSecurityPutUserFunc(t),
		},
		Sql: &Sql{
			ClearCursor: newSqlClearCursorFunc(t),
			Query:       newSqlQueryFunc(t),
			Translate:   newSqlTranslateFunc(t),
		},
		Ssl: &Ssl{
			Certificates: newSslCertificatesFunc(t),
		},
		Watcher: &Watcher{
			AckWatch:        newWatcherAckWatchFunc(t),
			ActivateWatch:   newWatcherActivateWatchFunc(t),
			DeactivateWatch: newWatcherDeactivateWatchFunc(t),
			DeleteWatch:     newWatcherDeleteWatchFunc(t),
			ExecuteWatch:    newWatcherExecuteWatchFunc(t),
			GetWatch:        newWatcherGetWatchFunc(t),
			PutWatch:        newWatcherPutWatchFunc(t),
			Start:           newWatcherStartFunc(t),
			Stats:           newWatcherStatsFunc(t),
			Stop:            newWatcherStopFunc(t),
		},
		Xpack: &Xpack{
			Info:  newXpackInfoFunc(t),
			Usage: newXpackUsageFunc(t),
		},
	}
}
