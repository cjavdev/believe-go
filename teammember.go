// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package believe

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"

	"github.com/stainless-sdks/believe-go/internal/apijson"
	"github.com/stainless-sdks/believe-go/internal/apiquery"
	"github.com/stainless-sdks/believe-go/internal/requestconfig"
	"github.com/stainless-sdks/believe-go/option"
	"github.com/stainless-sdks/believe-go/packages/pagination"
	"github.com/stainless-sdks/believe-go/packages/param"
	"github.com/stainless-sdks/believe-go/packages/respjson"
)

// TeamMemberService contains methods and other services that help with interacting
// with the believe API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewTeamMemberService] method instead.
type TeamMemberService struct {
	Options []option.RequestOption
}

// NewTeamMemberService generates a new service that applies the given options to
// each request. These options are applied after the parent client's options (if
// there is one), and before any request-specific options.
func NewTeamMemberService(opts ...option.RequestOption) (r TeamMemberService) {
	r = TeamMemberService{}
	r.Options = opts
	return
}

// Add a new team member to a team.
//
// The request body is a **union type (oneOf)** - you must include the
// `member_type` discriminator field:
//
//   - `"member_type": "player"` - Creates a player (requires position,
//     jersey_number, etc.)
//   - `"member_type": "coach"` - Creates a coach (requires specialty, etc.)
//   - `"member_type": "medical_staff"` - Creates medical staff (requires medical
//     specialty, etc.)
//   - `"member_type": "equipment_manager"` - Creates equipment manager (requires
//     responsibilities, etc.)
//
// The `character_id` field references an existing character from
// `/characters/{id}`.
//
// **Example for creating a player:**
//
// ```json
//
//	{
//	  "member_type": "player",
//	  "character_id": "sam-obisanya",
//	  "team_id": "afc-richmond",
//	  "years_with_team": 2,
//	  "position": "midfielder",
//	  "jersey_number": 24,
//	  "goals_scored": 12,
//	  "assists": 15
//	}
//
// ```
func (r *TeamMemberService) New(ctx context.Context, body TeamMemberNewParams, opts ...option.RequestOption) (res *TeamMemberNewResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "team-members"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return
}

// Retrieve detailed information about a specific team member.
//
// The response is a **union type (oneOf)** - the actual shape depends on the
// member's type:
//
//   - **player**: Includes position, jersey_number, goals_scored, assists,
//     is_captain
//   - **coach**: Includes specialty, certifications, win_rate
//   - **medical_staff**: Includes specialty, qualifications, license_number
//   - **equipment_manager**: Includes responsibilities, is_head_kitman
//
// Use `character_id` to fetch full character details from
// `/characters/{character_id}`.
func (r *TeamMemberService) Get(ctx context.Context, memberID string, opts ...option.RequestOption) (res *TeamMemberGetResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("team-members/%s", url.PathEscape(memberID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return
}

// Update specific fields of an existing team member. Fields vary by member type.
func (r *TeamMemberService) Update(ctx context.Context, memberID string, body TeamMemberUpdateParams, opts ...option.RequestOption) (res *TeamMemberUpdateResponseUnion, err error) {
	opts = slices.Concat(r.Options, opts)
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("team-members/%s", url.PathEscape(memberID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPatch, path, body, &res, opts...)
	return
}

// Get a paginated list of all team members.
//
// This endpoint demonstrates **union types (oneOf)** in the response. Each team
// member can be one of: Player, Coach, MedicalStaff, or EquipmentManager. The
// `member_type` field acts as a discriminator to determine the shape of each
// object.
func (r *TeamMemberService) List(ctx context.Context, query TeamMemberListParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[TeamMemberListResponseUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "team-members"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get a paginated list of all team members.
//
// This endpoint demonstrates **union types (oneOf)** in the response. Each team
// member can be one of: Player, Coach, MedicalStaff, or EquipmentManager. The
// `member_type` field acts as a discriminator to determine the shape of each
// object.
func (r *TeamMemberService) ListAutoPaging(ctx context.Context, query TeamMemberListParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[TeamMemberListResponseUnion] {
	return pagination.NewSkipLimitPageAutoPager(r.List(ctx, query, opts...))
}

// Remove a team member from the roster.
func (r *TeamMemberService) Delete(ctx context.Context, memberID string, opts ...option.RequestOption) (err error) {
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithHeader("Accept", "*/*")}, opts...)
	if memberID == "" {
		err = errors.New("missing required member_id parameter")
		return
	}
	path := fmt.Sprintf("team-members/%s", url.PathEscape(memberID))
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodDelete, path, nil, nil, opts...)
	return
}

// Get only coaches (filtered subset of team members).
func (r *TeamMemberService) ListCoaches(ctx context.Context, query TeamMemberListCoachesParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Coach], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "team-members/coaches/"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get only coaches (filtered subset of team members).
func (r *TeamMemberService) ListCoachesAutoPaging(ctx context.Context, query TeamMemberListCoachesParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Coach] {
	return pagination.NewSkipLimitPageAutoPager(r.ListCoaches(ctx, query, opts...))
}

// Get only players (filtered subset of team members).
func (r *TeamMemberService) ListPlayers(ctx context.Context, query TeamMemberListPlayersParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[Player], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "team-members/players/"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get only players (filtered subset of team members).
func (r *TeamMemberService) ListPlayersAutoPaging(ctx context.Context, query TeamMemberListPlayersParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[Player] {
	return pagination.NewSkipLimitPageAutoPager(r.ListPlayers(ctx, query, opts...))
}

// Get all staff members (medical staff and equipment managers).
//
// This demonstrates a **narrower union type** - the response is oneOf MedicalStaff
// or EquipmentManager.
func (r *TeamMemberService) ListStaff(ctx context.Context, query TeamMemberListStaffParams, opts ...option.RequestOption) (res *pagination.SkipLimitPage[TeamMemberListStaffResponseUnion], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "team-members/staff/"
	cfg, err := requestconfig.NewRequestConfig(ctx, http.MethodGet, path, query, &res, opts...)
	if err != nil {
		return nil, err
	}
	err = cfg.Execute()
	if err != nil {
		return nil, err
	}
	res.SetPageConfig(cfg, raw)
	return res, nil
}

// Get all staff members (medical staff and equipment managers).
//
// This demonstrates a **narrower union type** - the response is oneOf MedicalStaff
// or EquipmentManager.
func (r *TeamMemberService) ListStaffAutoPaging(ctx context.Context, query TeamMemberListStaffParams, opts ...option.RequestOption) *pagination.SkipLimitPageAutoPager[TeamMemberListStaffResponseUnion] {
	return pagination.NewSkipLimitPageAutoPager(r.ListStaff(ctx, query, opts...))
}

// Full coach model with ID.
type Coach struct {
	// Unique identifier for this team membership
	ID string `json:"id,required"`
	// ID of the character (references /characters/{id})
	CharacterID string `json:"character_id,required"`
	// Coaching specialty/role
	//
	// Any of "head_coach", "assistant_coach", "goalkeeping_coach", "fitness_coach",
	// "tactical_analyst".
	Specialty CoachSpecialty `json:"specialty,required"`
	// ID of the team they belong to
	TeamID string `json:"team_id,required"`
	// Number of years with the current team
	YearsWithTeam int64 `json:"years_with_team,required"`
	// Coaching certifications and licenses
	Certifications []string `json:"certifications"`
	// Discriminator field indicating this is a coach
	//
	// Any of "coach".
	MemberType CoachMemberType `json:"member_type"`
	// Career win rate (0.0 to 1.0)
	WinRate float64 `json:"win_rate,nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CharacterID    respjson.Field
		Specialty      respjson.Field
		TeamID         respjson.Field
		YearsWithTeam  respjson.Field
		Certifications respjson.Field
		MemberType     respjson.Field
		WinRate        respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Coach) RawJSON() string { return r.JSON.raw }
func (r *Coach) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator field indicating this is a coach
type CoachMemberType string

const (
	CoachMemberTypeCoach CoachMemberType = "coach"
)

// Coaching specialties.
type CoachSpecialty string

const (
	CoachSpecialtyHeadCoach        CoachSpecialty = "head_coach"
	CoachSpecialtyAssistantCoach   CoachSpecialty = "assistant_coach"
	CoachSpecialtyGoalkeepingCoach CoachSpecialty = "goalkeeping_coach"
	CoachSpecialtyFitnessCoach     CoachSpecialty = "fitness_coach"
	CoachSpecialtyTacticalAnalyst  CoachSpecialty = "tactical_analyst"
)

// Full equipment manager model with ID.
type EquipmentManager struct {
	// Unique identifier for this team membership
	ID string `json:"id,required"`
	// ID of the character (references /characters/{id})
	CharacterID string `json:"character_id,required"`
	// ID of the team they belong to
	TeamID string `json:"team_id,required"`
	// Number of years with the current team
	YearsWithTeam int64 `json:"years_with_team,required"`
	// Whether this is the head equipment manager
	IsHeadKitman bool `json:"is_head_kitman"`
	// Discriminator field indicating this is an equipment manager
	//
	// Any of "equipment_manager".
	MemberType EquipmentManagerMemberType `json:"member_type"`
	// List of responsibilities
	Responsibilities []string `json:"responsibilities"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID               respjson.Field
		CharacterID      respjson.Field
		TeamID           respjson.Field
		YearsWithTeam    respjson.Field
		IsHeadKitman     respjson.Field
		MemberType       respjson.Field
		Responsibilities respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r EquipmentManager) RawJSON() string { return r.JSON.raw }
func (r *EquipmentManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator field indicating this is an equipment manager
type EquipmentManagerMemberType string

const (
	EquipmentManagerMemberTypeEquipmentManager EquipmentManagerMemberType = "equipment_manager"
)

// Medical staff specialties.
type MedicalSpecialty string

const (
	MedicalSpecialtyTeamDoctor         MedicalSpecialty = "team_doctor"
	MedicalSpecialtyPhysiotherapist    MedicalSpecialty = "physiotherapist"
	MedicalSpecialtySportsPsychologist MedicalSpecialty = "sports_psychologist"
	MedicalSpecialtyNutritionist       MedicalSpecialty = "nutritionist"
	MedicalSpecialtyMassageTherapist   MedicalSpecialty = "massage_therapist"
)

// Full medical staff model with ID.
type MedicalStaff struct {
	// Unique identifier for this team membership
	ID string `json:"id,required"`
	// ID of the character (references /characters/{id})
	CharacterID string `json:"character_id,required"`
	// Medical specialty
	//
	// Any of "team_doctor", "physiotherapist", "sports_psychologist", "nutritionist",
	// "massage_therapist".
	Specialty MedicalSpecialty `json:"specialty,required"`
	// ID of the team they belong to
	TeamID string `json:"team_id,required"`
	// Number of years with the current team
	YearsWithTeam int64 `json:"years_with_team,required"`
	// Professional license number
	LicenseNumber string `json:"license_number,nullable"`
	// Discriminator field indicating this is medical staff
	//
	// Any of "medical_staff".
	MemberType MedicalStaffMemberType `json:"member_type"`
	// Medical qualifications and degrees
	Qualifications []string `json:"qualifications"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID             respjson.Field
		CharacterID    respjson.Field
		Specialty      respjson.Field
		TeamID         respjson.Field
		YearsWithTeam  respjson.Field
		LicenseNumber  respjson.Field
		MemberType     respjson.Field
		Qualifications respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r MedicalStaff) RawJSON() string { return r.JSON.raw }
func (r *MedicalStaff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator field indicating this is medical staff
type MedicalStaffMemberType string

const (
	MedicalStaffMemberTypeMedicalStaff MedicalStaffMemberType = "medical_staff"
)

// Full player model with ID.
type Player struct {
	// Unique identifier for this team membership
	ID string `json:"id,required"`
	// ID of the character (references /characters/{id})
	CharacterID string `json:"character_id,required"`
	// Jersey/shirt number
	JerseyNumber int64 `json:"jersey_number,required"`
	// Playing position on the field
	//
	// Any of "goalkeeper", "defender", "midfielder", "forward".
	Position Position `json:"position,required"`
	// ID of the team they belong to
	TeamID string `json:"team_id,required"`
	// Number of years with the current team
	YearsWithTeam int64 `json:"years_with_team,required"`
	// Total assists for the team
	Assists int64 `json:"assists"`
	// Total goals scored for the team
	GoalsScored int64 `json:"goals_scored"`
	// Whether this player is team captain
	IsCaptain bool `json:"is_captain"`
	// Discriminator field indicating this is a player
	//
	// Any of "player".
	MemberType PlayerMemberType `json:"member_type"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ID            respjson.Field
		CharacterID   respjson.Field
		JerseyNumber  respjson.Field
		Position      respjson.Field
		TeamID        respjson.Field
		YearsWithTeam respjson.Field
		Assists       respjson.Field
		GoalsScored   respjson.Field
		IsCaptain     respjson.Field
		MemberType    respjson.Field
		ExtraFields   map[string]respjson.Field
		raw           string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r Player) RawJSON() string { return r.JSON.raw }
func (r *Player) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Discriminator field indicating this is a player
type PlayerMemberType string

const (
	PlayerMemberTypePlayer PlayerMemberType = "player"
)

// Football positions for players.
type Position string

const (
	PositionGoalkeeper Position = "goalkeeper"
	PositionDefender   Position = "defender"
	PositionMidfielder Position = "midfielder"
	PositionForward    Position = "forward"
)

// TeamMemberNewResponseUnion contains all possible properties and values from
// [Player], [Coach], [MedicalStaff], [EquipmentManager].
//
// Use the [TeamMemberNewResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TeamMemberNewResponseUnion struct {
	ID          string `json:"id"`
	CharacterID string `json:"character_id"`
	// This field is from variant [Player].
	JerseyNumber int64 `json:"jersey_number"`
	// This field is from variant [Player].
	Position      Position `json:"position"`
	TeamID        string   `json:"team_id"`
	YearsWithTeam int64    `json:"years_with_team"`
	// This field is from variant [Player].
	Assists int64 `json:"assists"`
	// This field is from variant [Player].
	GoalsScored int64 `json:"goals_scored"`
	// This field is from variant [Player].
	IsCaptain bool `json:"is_captain"`
	// Any of "player", "coach", "medical_staff", "equipment_manager".
	MemberType string `json:"member_type"`
	Specialty  string `json:"specialty"`
	// This field is from variant [Coach].
	Certifications []string `json:"certifications"`
	// This field is from variant [Coach].
	WinRate float64 `json:"win_rate"`
	// This field is from variant [MedicalStaff].
	LicenseNumber string `json:"license_number"`
	// This field is from variant [MedicalStaff].
	Qualifications []string `json:"qualifications"`
	// This field is from variant [EquipmentManager].
	IsHeadKitman bool `json:"is_head_kitman"`
	// This field is from variant [EquipmentManager].
	Responsibilities []string `json:"responsibilities"`
	JSON             struct {
		ID               respjson.Field
		CharacterID      respjson.Field
		JerseyNumber     respjson.Field
		Position         respjson.Field
		TeamID           respjson.Field
		YearsWithTeam    respjson.Field
		Assists          respjson.Field
		GoalsScored      respjson.Field
		IsCaptain        respjson.Field
		MemberType       respjson.Field
		Specialty        respjson.Field
		Certifications   respjson.Field
		WinRate          respjson.Field
		LicenseNumber    respjson.Field
		Qualifications   respjson.Field
		IsHeadKitman     respjson.Field
		Responsibilities respjson.Field
		raw              string
	} `json:"-"`
}

// anyTeamMemberNewResponse is implemented by each variant of
// [TeamMemberNewResponseUnion] to add type safety for the return type of
// [TeamMemberNewResponseUnion.AsAny]
type anyTeamMemberNewResponse interface {
	implTeamMemberNewResponseUnion()
}

func (Player) implTeamMemberNewResponseUnion()           {}
func (Coach) implTeamMemberNewResponseUnion()            {}
func (MedicalStaff) implTeamMemberNewResponseUnion()     {}
func (EquipmentManager) implTeamMemberNewResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TeamMemberNewResponseUnion.AsAny().(type) {
//	case believe.Player:
//	case believe.Coach:
//	case believe.MedicalStaff:
//	case believe.EquipmentManager:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TeamMemberNewResponseUnion) AsAny() anyTeamMemberNewResponse {
	switch u.MemberType {
	case "player":
		return u.AsPlayer()
	case "coach":
		return u.AsCoach()
	case "medical_staff":
		return u.AsMedicalStaff()
	case "equipment_manager":
		return u.AsEquipmentManager()
	}
	return nil
}

func (u TeamMemberNewResponseUnion) AsPlayer() (v Player) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberNewResponseUnion) AsCoach() (v Coach) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberNewResponseUnion) AsMedicalStaff() (v MedicalStaff) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberNewResponseUnion) AsEquipmentManager() (v EquipmentManager) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TeamMemberNewResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *TeamMemberNewResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TeamMemberGetResponseUnion contains all possible properties and values from
// [Player], [Coach], [MedicalStaff], [EquipmentManager].
//
// Use the [TeamMemberGetResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TeamMemberGetResponseUnion struct {
	ID          string `json:"id"`
	CharacterID string `json:"character_id"`
	// This field is from variant [Player].
	JerseyNumber int64 `json:"jersey_number"`
	// This field is from variant [Player].
	Position      Position `json:"position"`
	TeamID        string   `json:"team_id"`
	YearsWithTeam int64    `json:"years_with_team"`
	// This field is from variant [Player].
	Assists int64 `json:"assists"`
	// This field is from variant [Player].
	GoalsScored int64 `json:"goals_scored"`
	// This field is from variant [Player].
	IsCaptain bool `json:"is_captain"`
	// Any of "player", "coach", "medical_staff", "equipment_manager".
	MemberType string `json:"member_type"`
	Specialty  string `json:"specialty"`
	// This field is from variant [Coach].
	Certifications []string `json:"certifications"`
	// This field is from variant [Coach].
	WinRate float64 `json:"win_rate"`
	// This field is from variant [MedicalStaff].
	LicenseNumber string `json:"license_number"`
	// This field is from variant [MedicalStaff].
	Qualifications []string `json:"qualifications"`
	// This field is from variant [EquipmentManager].
	IsHeadKitman bool `json:"is_head_kitman"`
	// This field is from variant [EquipmentManager].
	Responsibilities []string `json:"responsibilities"`
	JSON             struct {
		ID               respjson.Field
		CharacterID      respjson.Field
		JerseyNumber     respjson.Field
		Position         respjson.Field
		TeamID           respjson.Field
		YearsWithTeam    respjson.Field
		Assists          respjson.Field
		GoalsScored      respjson.Field
		IsCaptain        respjson.Field
		MemberType       respjson.Field
		Specialty        respjson.Field
		Certifications   respjson.Field
		WinRate          respjson.Field
		LicenseNumber    respjson.Field
		Qualifications   respjson.Field
		IsHeadKitman     respjson.Field
		Responsibilities respjson.Field
		raw              string
	} `json:"-"`
}

// anyTeamMemberGetResponse is implemented by each variant of
// [TeamMemberGetResponseUnion] to add type safety for the return type of
// [TeamMemberGetResponseUnion.AsAny]
type anyTeamMemberGetResponse interface {
	implTeamMemberGetResponseUnion()
}

func (Player) implTeamMemberGetResponseUnion()           {}
func (Coach) implTeamMemberGetResponseUnion()            {}
func (MedicalStaff) implTeamMemberGetResponseUnion()     {}
func (EquipmentManager) implTeamMemberGetResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TeamMemberGetResponseUnion.AsAny().(type) {
//	case believe.Player:
//	case believe.Coach:
//	case believe.MedicalStaff:
//	case believe.EquipmentManager:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TeamMemberGetResponseUnion) AsAny() anyTeamMemberGetResponse {
	switch u.MemberType {
	case "player":
		return u.AsPlayer()
	case "coach":
		return u.AsCoach()
	case "medical_staff":
		return u.AsMedicalStaff()
	case "equipment_manager":
		return u.AsEquipmentManager()
	}
	return nil
}

func (u TeamMemberGetResponseUnion) AsPlayer() (v Player) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberGetResponseUnion) AsCoach() (v Coach) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberGetResponseUnion) AsMedicalStaff() (v MedicalStaff) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberGetResponseUnion) AsEquipmentManager() (v EquipmentManager) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TeamMemberGetResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *TeamMemberGetResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TeamMemberUpdateResponseUnion contains all possible properties and values from
// [Player], [Coach], [MedicalStaff], [EquipmentManager].
//
// Use the [TeamMemberUpdateResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TeamMemberUpdateResponseUnion struct {
	ID          string `json:"id"`
	CharacterID string `json:"character_id"`
	// This field is from variant [Player].
	JerseyNumber int64 `json:"jersey_number"`
	// This field is from variant [Player].
	Position      Position `json:"position"`
	TeamID        string   `json:"team_id"`
	YearsWithTeam int64    `json:"years_with_team"`
	// This field is from variant [Player].
	Assists int64 `json:"assists"`
	// This field is from variant [Player].
	GoalsScored int64 `json:"goals_scored"`
	// This field is from variant [Player].
	IsCaptain bool `json:"is_captain"`
	// Any of "player", "coach", "medical_staff", "equipment_manager".
	MemberType string `json:"member_type"`
	Specialty  string `json:"specialty"`
	// This field is from variant [Coach].
	Certifications []string `json:"certifications"`
	// This field is from variant [Coach].
	WinRate float64 `json:"win_rate"`
	// This field is from variant [MedicalStaff].
	LicenseNumber string `json:"license_number"`
	// This field is from variant [MedicalStaff].
	Qualifications []string `json:"qualifications"`
	// This field is from variant [EquipmentManager].
	IsHeadKitman bool `json:"is_head_kitman"`
	// This field is from variant [EquipmentManager].
	Responsibilities []string `json:"responsibilities"`
	JSON             struct {
		ID               respjson.Field
		CharacterID      respjson.Field
		JerseyNumber     respjson.Field
		Position         respjson.Field
		TeamID           respjson.Field
		YearsWithTeam    respjson.Field
		Assists          respjson.Field
		GoalsScored      respjson.Field
		IsCaptain        respjson.Field
		MemberType       respjson.Field
		Specialty        respjson.Field
		Certifications   respjson.Field
		WinRate          respjson.Field
		LicenseNumber    respjson.Field
		Qualifications   respjson.Field
		IsHeadKitman     respjson.Field
		Responsibilities respjson.Field
		raw              string
	} `json:"-"`
}

// anyTeamMemberUpdateResponse is implemented by each variant of
// [TeamMemberUpdateResponseUnion] to add type safety for the return type of
// [TeamMemberUpdateResponseUnion.AsAny]
type anyTeamMemberUpdateResponse interface {
	implTeamMemberUpdateResponseUnion()
}

func (Player) implTeamMemberUpdateResponseUnion()           {}
func (Coach) implTeamMemberUpdateResponseUnion()            {}
func (MedicalStaff) implTeamMemberUpdateResponseUnion()     {}
func (EquipmentManager) implTeamMemberUpdateResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TeamMemberUpdateResponseUnion.AsAny().(type) {
//	case believe.Player:
//	case believe.Coach:
//	case believe.MedicalStaff:
//	case believe.EquipmentManager:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TeamMemberUpdateResponseUnion) AsAny() anyTeamMemberUpdateResponse {
	switch u.MemberType {
	case "player":
		return u.AsPlayer()
	case "coach":
		return u.AsCoach()
	case "medical_staff":
		return u.AsMedicalStaff()
	case "equipment_manager":
		return u.AsEquipmentManager()
	}
	return nil
}

func (u TeamMemberUpdateResponseUnion) AsPlayer() (v Player) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberUpdateResponseUnion) AsCoach() (v Coach) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberUpdateResponseUnion) AsMedicalStaff() (v MedicalStaff) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberUpdateResponseUnion) AsEquipmentManager() (v EquipmentManager) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TeamMemberUpdateResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *TeamMemberUpdateResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TeamMemberListResponseUnion contains all possible properties and values from
// [Player], [Coach], [MedicalStaff], [EquipmentManager].
//
// Use the [TeamMemberListResponseUnion.AsAny] method to switch on the variant.
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TeamMemberListResponseUnion struct {
	ID          string `json:"id"`
	CharacterID string `json:"character_id"`
	// This field is from variant [Player].
	JerseyNumber int64 `json:"jersey_number"`
	// This field is from variant [Player].
	Position      Position `json:"position"`
	TeamID        string   `json:"team_id"`
	YearsWithTeam int64    `json:"years_with_team"`
	// This field is from variant [Player].
	Assists int64 `json:"assists"`
	// This field is from variant [Player].
	GoalsScored int64 `json:"goals_scored"`
	// This field is from variant [Player].
	IsCaptain bool `json:"is_captain"`
	// Any of "player", "coach", "medical_staff", "equipment_manager".
	MemberType string `json:"member_type"`
	Specialty  string `json:"specialty"`
	// This field is from variant [Coach].
	Certifications []string `json:"certifications"`
	// This field is from variant [Coach].
	WinRate float64 `json:"win_rate"`
	// This field is from variant [MedicalStaff].
	LicenseNumber string `json:"license_number"`
	// This field is from variant [MedicalStaff].
	Qualifications []string `json:"qualifications"`
	// This field is from variant [EquipmentManager].
	IsHeadKitman bool `json:"is_head_kitman"`
	// This field is from variant [EquipmentManager].
	Responsibilities []string `json:"responsibilities"`
	JSON             struct {
		ID               respjson.Field
		CharacterID      respjson.Field
		JerseyNumber     respjson.Field
		Position         respjson.Field
		TeamID           respjson.Field
		YearsWithTeam    respjson.Field
		Assists          respjson.Field
		GoalsScored      respjson.Field
		IsCaptain        respjson.Field
		MemberType       respjson.Field
		Specialty        respjson.Field
		Certifications   respjson.Field
		WinRate          respjson.Field
		LicenseNumber    respjson.Field
		Qualifications   respjson.Field
		IsHeadKitman     respjson.Field
		Responsibilities respjson.Field
		raw              string
	} `json:"-"`
}

// anyTeamMemberListResponse is implemented by each variant of
// [TeamMemberListResponseUnion] to add type safety for the return type of
// [TeamMemberListResponseUnion.AsAny]
type anyTeamMemberListResponse interface {
	implTeamMemberListResponseUnion()
}

func (Player) implTeamMemberListResponseUnion()           {}
func (Coach) implTeamMemberListResponseUnion()            {}
func (MedicalStaff) implTeamMemberListResponseUnion()     {}
func (EquipmentManager) implTeamMemberListResponseUnion() {}

// Use the following switch statement to find the correct variant
//
//	switch variant := TeamMemberListResponseUnion.AsAny().(type) {
//	case believe.Player:
//	case believe.Coach:
//	case believe.MedicalStaff:
//	case believe.EquipmentManager:
//	default:
//	  fmt.Errorf("no variant present")
//	}
func (u TeamMemberListResponseUnion) AsAny() anyTeamMemberListResponse {
	switch u.MemberType {
	case "player":
		return u.AsPlayer()
	case "coach":
		return u.AsCoach()
	case "medical_staff":
		return u.AsMedicalStaff()
	case "equipment_manager":
		return u.AsEquipmentManager()
	}
	return nil
}

func (u TeamMemberListResponseUnion) AsPlayer() (v Player) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberListResponseUnion) AsCoach() (v Coach) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberListResponseUnion) AsMedicalStaff() (v MedicalStaff) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberListResponseUnion) AsEquipmentManager() (v EquipmentManager) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TeamMemberListResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *TeamMemberListResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// TeamMemberListStaffResponseUnion contains all possible properties and values
// from [MedicalStaff], [EquipmentManager].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type TeamMemberListStaffResponseUnion struct {
	ID          string `json:"id"`
	CharacterID string `json:"character_id"`
	// This field is from variant [MedicalStaff].
	Specialty     MedicalSpecialty `json:"specialty"`
	TeamID        string           `json:"team_id"`
	YearsWithTeam int64            `json:"years_with_team"`
	// This field is from variant [MedicalStaff].
	LicenseNumber string `json:"license_number"`
	MemberType    string `json:"member_type"`
	// This field is from variant [MedicalStaff].
	Qualifications []string `json:"qualifications"`
	// This field is from variant [EquipmentManager].
	IsHeadKitman bool `json:"is_head_kitman"`
	// This field is from variant [EquipmentManager].
	Responsibilities []string `json:"responsibilities"`
	JSON             struct {
		ID               respjson.Field
		CharacterID      respjson.Field
		Specialty        respjson.Field
		TeamID           respjson.Field
		YearsWithTeam    respjson.Field
		LicenseNumber    respjson.Field
		MemberType       respjson.Field
		Qualifications   respjson.Field
		IsHeadKitman     respjson.Field
		Responsibilities respjson.Field
		raw              string
	} `json:"-"`
}

func (u TeamMemberListStaffResponseUnion) AsMedicalStaff() (v MedicalStaff) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u TeamMemberListStaffResponseUnion) AsEquipmentManager() (v EquipmentManager) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u TeamMemberListStaffResponseUnion) RawJSON() string { return u.JSON.raw }

func (r *TeamMemberListStaffResponseUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamMemberNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. A
	// football player on the team.
	OfPlayer *TeamMemberNewParamsMemberPlayer `json:",inline"`
	// This field is a request body variant, only one variant field can be set. A coach
	// or coaching staff member.
	OfCoach *TeamMemberNewParamsMemberCoach `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Medical
	// and wellness staff member.
	OfMedicalStaff *TeamMemberNewParamsMemberMedicalStaff `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	// Equipment and kit management staff.
	OfEquipmentManager *TeamMemberNewParamsMemberEquipmentManager `json:",inline"`

	paramObj
}

func (u TeamMemberNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPlayer, u.OfCoach, u.OfMedicalStaff, u.OfEquipmentManager)
}
func (r *TeamMemberNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// A football player on the team.
//
// The properties CharacterID, JerseyNumber, Position, TeamID, YearsWithTeam are
// required.
type TeamMemberNewParamsMemberPlayer struct {
	// ID of the character (references /characters/{id})
	CharacterID string `json:"character_id,required"`
	// Jersey/shirt number
	JerseyNumber int64 `json:"jersey_number,required"`
	// Playing position on the field
	//
	// Any of "goalkeeper", "defender", "midfielder", "forward".
	Position Position `json:"position,omitzero,required"`
	// ID of the team they belong to
	TeamID string `json:"team_id,required"`
	// Number of years with the current team
	YearsWithTeam int64 `json:"years_with_team,required"`
	// Total assists for the team
	Assists param.Opt[int64] `json:"assists,omitzero"`
	// Total goals scored for the team
	GoalsScored param.Opt[int64] `json:"goals_scored,omitzero"`
	// Whether this player is team captain
	IsCaptain param.Opt[bool] `json:"is_captain,omitzero"`
	// Discriminator field indicating this is a player
	//
	// Any of "player".
	MemberType string `json:"member_type,omitzero"`
	paramObj
}

func (r TeamMemberNewParamsMemberPlayer) MarshalJSON() (data []byte, err error) {
	type shadow TeamMemberNewParamsMemberPlayer
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamMemberNewParamsMemberPlayer) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TeamMemberNewParamsMemberPlayer](
		"member_type", "player",
	)
}

// A coach or coaching staff member.
//
// The properties CharacterID, Specialty, TeamID, YearsWithTeam are required.
type TeamMemberNewParamsMemberCoach struct {
	// ID of the character (references /characters/{id})
	CharacterID string `json:"character_id,required"`
	// Coaching specialty/role
	//
	// Any of "head_coach", "assistant_coach", "goalkeeping_coach", "fitness_coach",
	// "tactical_analyst".
	Specialty CoachSpecialty `json:"specialty,omitzero,required"`
	// ID of the team they belong to
	TeamID string `json:"team_id,required"`
	// Number of years with the current team
	YearsWithTeam int64 `json:"years_with_team,required"`
	// Career win rate (0.0 to 1.0)
	WinRate param.Opt[float64] `json:"win_rate,omitzero"`
	// Coaching certifications and licenses
	Certifications []string `json:"certifications,omitzero"`
	// Discriminator field indicating this is a coach
	//
	// Any of "coach".
	MemberType string `json:"member_type,omitzero"`
	paramObj
}

func (r TeamMemberNewParamsMemberCoach) MarshalJSON() (data []byte, err error) {
	type shadow TeamMemberNewParamsMemberCoach
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamMemberNewParamsMemberCoach) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TeamMemberNewParamsMemberCoach](
		"member_type", "coach",
	)
}

// Medical and wellness staff member.
//
// The properties CharacterID, Specialty, TeamID, YearsWithTeam are required.
type TeamMemberNewParamsMemberMedicalStaff struct {
	// ID of the character (references /characters/{id})
	CharacterID string `json:"character_id,required"`
	// Medical specialty
	//
	// Any of "team_doctor", "physiotherapist", "sports_psychologist", "nutritionist",
	// "massage_therapist".
	Specialty MedicalSpecialty `json:"specialty,omitzero,required"`
	// ID of the team they belong to
	TeamID string `json:"team_id,required"`
	// Number of years with the current team
	YearsWithTeam int64 `json:"years_with_team,required"`
	// Professional license number
	LicenseNumber param.Opt[string] `json:"license_number,omitzero"`
	// Discriminator field indicating this is medical staff
	//
	// Any of "medical_staff".
	MemberType string `json:"member_type,omitzero"`
	// Medical qualifications and degrees
	Qualifications []string `json:"qualifications,omitzero"`
	paramObj
}

func (r TeamMemberNewParamsMemberMedicalStaff) MarshalJSON() (data []byte, err error) {
	type shadow TeamMemberNewParamsMemberMedicalStaff
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamMemberNewParamsMemberMedicalStaff) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TeamMemberNewParamsMemberMedicalStaff](
		"member_type", "medical_staff",
	)
}

// Equipment and kit management staff.
//
// The properties CharacterID, TeamID, YearsWithTeam are required.
type TeamMemberNewParamsMemberEquipmentManager struct {
	// ID of the character (references /characters/{id})
	CharacterID string `json:"character_id,required"`
	// ID of the team they belong to
	TeamID string `json:"team_id,required"`
	// Number of years with the current team
	YearsWithTeam int64 `json:"years_with_team,required"`
	// Whether this is the head equipment manager
	IsHeadKitman param.Opt[bool] `json:"is_head_kitman,omitzero"`
	// Discriminator field indicating this is an equipment manager
	//
	// Any of "equipment_manager".
	MemberType string `json:"member_type,omitzero"`
	// List of responsibilities
	Responsibilities []string `json:"responsibilities,omitzero"`
	paramObj
}

func (r TeamMemberNewParamsMemberEquipmentManager) MarshalJSON() (data []byte, err error) {
	type shadow TeamMemberNewParamsMemberEquipmentManager
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamMemberNewParamsMemberEquipmentManager) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[TeamMemberNewParamsMemberEquipmentManager](
		"member_type", "equipment_manager",
	)
}

type TeamMemberUpdateParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set. Update
	// model for players.
	OfPlayerUpdate *TeamMemberUpdateParamsUpdatesPlayerUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Update
	// model for coaches.
	OfCoachUpdate *TeamMemberUpdateParamsUpdatesCoachUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Update
	// model for medical staff.
	OfMedicalStaffUpdate *TeamMemberUpdateParamsUpdatesMedicalStaffUpdate `json:",inline"`
	// This field is a request body variant, only one variant field can be set. Update
	// model for equipment managers.
	OfEquipmentManagerUpdate *TeamMemberUpdateParamsUpdatesEquipmentManagerUpdate `json:",inline"`

	paramObj
}

func (u TeamMemberUpdateParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfPlayerUpdate, u.OfCoachUpdate, u.OfMedicalStaffUpdate, u.OfEquipmentManagerUpdate)
}
func (r *TeamMemberUpdateParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update model for players.
type TeamMemberUpdateParamsUpdatesPlayerUpdate struct {
	Assists       param.Opt[int64]  `json:"assists,omitzero"`
	GoalsScored   param.Opt[int64]  `json:"goals_scored,omitzero"`
	IsCaptain     param.Opt[bool]   `json:"is_captain,omitzero"`
	JerseyNumber  param.Opt[int64]  `json:"jersey_number,omitzero"`
	TeamID        param.Opt[string] `json:"team_id,omitzero"`
	YearsWithTeam param.Opt[int64]  `json:"years_with_team,omitzero"`
	// Football positions for players.
	//
	// Any of "goalkeeper", "defender", "midfielder", "forward".
	Position Position `json:"position,omitzero"`
	paramObj
}

func (r TeamMemberUpdateParamsUpdatesPlayerUpdate) MarshalJSON() (data []byte, err error) {
	type shadow TeamMemberUpdateParamsUpdatesPlayerUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamMemberUpdateParamsUpdatesPlayerUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update model for coaches.
type TeamMemberUpdateParamsUpdatesCoachUpdate struct {
	TeamID         param.Opt[string]  `json:"team_id,omitzero"`
	WinRate        param.Opt[float64] `json:"win_rate,omitzero"`
	YearsWithTeam  param.Opt[int64]   `json:"years_with_team,omitzero"`
	Certifications []string           `json:"certifications,omitzero"`
	// Coaching specialties.
	//
	// Any of "head_coach", "assistant_coach", "goalkeeping_coach", "fitness_coach",
	// "tactical_analyst".
	Specialty CoachSpecialty `json:"specialty,omitzero"`
	paramObj
}

func (r TeamMemberUpdateParamsUpdatesCoachUpdate) MarshalJSON() (data []byte, err error) {
	type shadow TeamMemberUpdateParamsUpdatesCoachUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamMemberUpdateParamsUpdatesCoachUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update model for medical staff.
type TeamMemberUpdateParamsUpdatesMedicalStaffUpdate struct {
	LicenseNumber  param.Opt[string] `json:"license_number,omitzero"`
	TeamID         param.Opt[string] `json:"team_id,omitzero"`
	YearsWithTeam  param.Opt[int64]  `json:"years_with_team,omitzero"`
	Qualifications []string          `json:"qualifications,omitzero"`
	// Medical staff specialties.
	//
	// Any of "team_doctor", "physiotherapist", "sports_psychologist", "nutritionist",
	// "massage_therapist".
	Specialty MedicalSpecialty `json:"specialty,omitzero"`
	paramObj
}

func (r TeamMemberUpdateParamsUpdatesMedicalStaffUpdate) MarshalJSON() (data []byte, err error) {
	type shadow TeamMemberUpdateParamsUpdatesMedicalStaffUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamMemberUpdateParamsUpdatesMedicalStaffUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Update model for equipment managers.
type TeamMemberUpdateParamsUpdatesEquipmentManagerUpdate struct {
	IsHeadKitman     param.Opt[bool]   `json:"is_head_kitman,omitzero"`
	TeamID           param.Opt[string] `json:"team_id,omitzero"`
	YearsWithTeam    param.Opt[int64]  `json:"years_with_team,omitzero"`
	Responsibilities []string          `json:"responsibilities,omitzero"`
	paramObj
}

func (r TeamMemberUpdateParamsUpdatesEquipmentManagerUpdate) MarshalJSON() (data []byte, err error) {
	type shadow TeamMemberUpdateParamsUpdatesEquipmentManagerUpdate
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *TeamMemberUpdateParamsUpdatesEquipmentManagerUpdate) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type TeamMemberListParams struct {
	// Filter by team ID
	TeamID param.Opt[string] `query:"team_id,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Types of team members - used as discriminator.
	//
	// Any of "player", "coach", "medical_staff", "equipment_manager".
	MemberType TeamMemberListParamsMemberType `query:"member_type,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamMemberListParams]'s query parameters as `url.Values`.
func (r TeamMemberListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

// Types of team members - used as discriminator.
type TeamMemberListParamsMemberType string

const (
	TeamMemberListParamsMemberTypePlayer           TeamMemberListParamsMemberType = "player"
	TeamMemberListParamsMemberTypeCoach            TeamMemberListParamsMemberType = "coach"
	TeamMemberListParamsMemberTypeMedicalStaff     TeamMemberListParamsMemberType = "medical_staff"
	TeamMemberListParamsMemberTypeEquipmentManager TeamMemberListParamsMemberType = "equipment_manager"
)

type TeamMemberListCoachesParams struct {
	// Filter by team ID
	TeamID param.Opt[string] `query:"team_id,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Coaching specialties.
	//
	// Any of "head_coach", "assistant_coach", "goalkeeping_coach", "fitness_coach",
	// "tactical_analyst".
	Specialty CoachSpecialty `query:"specialty,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamMemberListCoachesParams]'s query parameters as
// `url.Values`.
func (r TeamMemberListCoachesParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamMemberListPlayersParams struct {
	// Filter by team ID
	TeamID param.Opt[string] `query:"team_id,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	// Football positions for players.
	//
	// Any of "goalkeeper", "defender", "midfielder", "forward".
	Position Position `query:"position,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamMemberListPlayersParams]'s query parameters as
// `url.Values`.
func (r TeamMemberListPlayersParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}

type TeamMemberListStaffParams struct {
	// Filter by team ID
	TeamID param.Opt[string] `query:"team_id,omitzero" json:"-"`
	// Maximum number of items to return (max: 100)
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Number of items to skip (offset)
	Skip param.Opt[int64] `query:"skip,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [TeamMemberListStaffParams]'s query parameters as
// `url.Values`.
func (r TeamMemberListStaffParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatComma,
		NestedFormat: apiquery.NestedQueryFormatBrackets,
	})
}
