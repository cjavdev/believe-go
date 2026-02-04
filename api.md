# believe

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#GetWelcomeResponse">GetWelcomeResponse</a>

Methods:

- <code title="get /">client.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BelieveService.GetWelcome">GetWelcome</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#GetWelcomeResponse">GetWelcomeResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Characters

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterRole">CharacterRole</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EmotionalStatsParam">EmotionalStatsParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#GrowthArcParam">GrowthArcParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Character">Character</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterRole">CharacterRole</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EmotionalStats">EmotionalStats</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#GrowthArc">GrowthArc</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterGetAllCharactersResponse">CharacterGetAllCharactersResponse</a>

Methods:

- <code title="post /characters">client.Characters.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterNewParams">CharacterNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Character">Character</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /characters/{character_id}">client.Characters.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, characterID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Character">Character</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /characters/{character_id}">client.Characters.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, characterID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterUpdateParams">CharacterUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Character">Character</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /characters/{character_id}">client.Characters.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, characterID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /characters">client.Characters.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterService.GetAllCharacters">GetAllCharacters</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterGetAllCharactersParams">CharacterGetAllCharactersParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterGetAllCharactersResponse">CharacterGetAllCharactersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /characters/{character_id}/quotes">client.Characters.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CharacterService.GetQuotes">GetQuotes</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, characterID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/builtin#string">string</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Teams

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#GeoLocationParam">GeoLocationParam</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#League">League</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamValuesParam">TeamValuesParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#GeoLocation">GeoLocation</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#League">League</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Team">Team</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamValues">TeamValues</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamListResponse">TeamListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamGetCultureResponse">TeamGetCultureResponse</a>

Methods:

- <code title="post /teams">client.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamNewParams">TeamNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Team">Team</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /teams/{team_id}">client.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Team">Team</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /teams/{team_id}">client.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamUpdateParams">TeamUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Team">Team</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /teams">client.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamListParams">TeamListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamListResponse">TeamListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /teams/{team_id}">client.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /teams/{team_id}/culture">client.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamService.GetCulture">GetCulture</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamGetCultureResponse">TeamGetCultureResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /teams/{team_id}/rivals">client.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamService.GetRivals">GetRivals</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Team">Team</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /teams/{team_id}/logos">client.Teams.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamService.ListLogos">ListLogos</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#FileUpload">FileUpload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Logo

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#FileUpload">FileUpload</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamLogoDownloadResponse">TeamLogoDownloadResponse</a>

Methods:

- <code title="delete /teams/{team_id}/logo/{file_id}">client.Teams.Logo.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamLogoService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamLogoDeleteParams">TeamLogoDeleteParams</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /teams/{team_id}/logo/{file_id}">client.Teams.Logo.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamLogoService.Download">Download</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, fileID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamLogoDownloadParams">TeamLogoDownloadParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamLogoDownloadResponse">TeamLogoDownloadResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /teams/{team_id}/logo">client.Teams.Logo.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamLogoService.Upload">Upload</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, teamID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamLogoUploadParams">TeamLogoUploadParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#FileUpload">FileUpload</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Matches

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchResult">MatchResult</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchType">MatchType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TurningPointParam">TurningPointParam</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Match">Match</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchResult">MatchResult</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchType">MatchType</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TurningPoint">TurningPoint</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchListResponse">MatchListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchGetLessonResponse">MatchGetLessonResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchGetTurningPointsResponse">MatchGetTurningPointsResponse</a>

Methods:

- <code title="post /matches">client.Matches.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchNewParams">MatchNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Match">Match</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /matches/{match_id}">client.Matches.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, matchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Match">Match</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /matches/{match_id}">client.Matches.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, matchID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchUpdateParams">MatchUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Match">Match</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /matches">client.Matches.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchListParams">MatchListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchListResponse">MatchListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /matches/{match_id}">client.Matches.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, matchID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /matches/{match_id}/lesson">client.Matches.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchService.GetLesson">GetLesson</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, matchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchGetLessonResponse">MatchGetLessonResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /matches/{match_id}/turning-points">client.Matches.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchService.GetTurningPoints">GetTurningPoints</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, matchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchGetTurningPointsResponse">MatchGetTurningPointsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

## Commentary

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchCommentaryStreamResponse">MatchCommentaryStreamResponse</a>

Methods:

- <code title="post /matches/{match_id}/commentary/stream">client.Matches.Commentary.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchCommentaryService.Stream">Stream</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, matchID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MatchCommentaryStreamResponse">MatchCommentaryStreamResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Episodes

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Episode">Episode</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PaginatedResponse">PaginatedResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeGetWisdomResponse">EpisodeGetWisdomResponse</a>

Methods:

- <code title="post /episodes">client.Episodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeNewParams">EpisodeNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Episode">Episode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /episodes/{episode_id}">client.Episodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, episodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Episode">Episode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /episodes/{episode_id}">client.Episodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, episodeID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeUpdateParams">EpisodeUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Episode">Episode</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /episodes">client.Episodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeListParams">EpisodeListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PaginatedResponse">PaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /episodes/{episode_id}">client.Episodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, episodeID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /episodes/{episode_id}/wisdom">client.Episodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeService.GetWisdom">GetWisdom</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, episodeID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeGetWisdomResponse">EpisodeGetWisdomResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /episodes/seasons/{season_number}">client.Episodes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeService.ListBySeason">ListBySeason</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, seasonNumber <a href="https://pkg.go.dev/builtin#int64">int64</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EpisodeListBySeasonParams">EpisodeListBySeasonParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PaginatedResponse">PaginatedResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Quotes

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteMoment">QuoteMoment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteTheme">QuoteTheme</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PaginatedResponseQuote">PaginatedResponseQuote</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Quote">Quote</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteMoment">QuoteMoment</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteTheme">QuoteTheme</a>

Methods:

- <code title="post /quotes">client.Quotes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteNewParams">QuoteNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Quote">Quote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /quotes/{quote_id}">client.Quotes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, quoteID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Quote">Quote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /quotes/{quote_id}">client.Quotes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, quoteID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteUpdateParams">QuoteUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Quote">Quote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /quotes">client.Quotes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteListParams">QuoteListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PaginatedResponseQuote">PaginatedResponseQuote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /quotes/{quote_id}">client.Quotes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, quoteID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /quotes/random">client.Quotes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteService.GetRandom">GetRandom</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteGetRandomParams">QuoteGetRandomParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Quote">Quote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /quotes/characters/{character_id}">client.Quotes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteService.ListByCharacter">ListByCharacter</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, characterID <a href="https://pkg.go.dev/builtin#string">string</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteListByCharacterParams">QuoteListByCharacterParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PaginatedResponseQuote">PaginatedResponseQuote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /quotes/themes/{theme}">client.Quotes.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteService.ListByTheme">ListByTheme</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, theme <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteTheme">QuoteTheme</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#QuoteListByThemeParams">QuoteListByThemeParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PaginatedResponseQuote">PaginatedResponseQuote</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Believe

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BelieveSubmitResponse">BelieveSubmitResponse</a>

Methods:

- <code title="post /believe">client.Believe.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BelieveService.Submit">Submit</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BelieveSubmitParams">BelieveSubmitParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BelieveSubmitResponse">BelieveSubmitResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Conflicts

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#ConflictResolveResponse">ConflictResolveResponse</a>

Methods:

- <code title="post /conflicts/resolve">client.Conflicts.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#ConflictService.Resolve">Resolve</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#ConflictResolveParams">ConflictResolveParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#ConflictResolveResponse">ConflictResolveResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Reframe

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#ReframeTransformNegativeThoughtsResponse">ReframeTransformNegativeThoughtsResponse</a>

Methods:

- <code title="post /reframe">client.Reframe.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#ReframeService.TransformNegativeThoughts">TransformNegativeThoughts</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#ReframeTransformNegativeThoughtsParams">ReframeTransformNegativeThoughtsParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#ReframeTransformNegativeThoughtsResponse">ReframeTransformNegativeThoughtsResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Press

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PressSimulateResponse">PressSimulateResponse</a>

Methods:

- <code title="post /press">client.Press.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PressService.Simulate">Simulate</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PressSimulateParams">PressSimulateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PressSimulateResponse">PressSimulateResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Coaching

## Principles

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrinciple">CoachingPrinciple</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrincipleListResponse">CoachingPrincipleListResponse</a>

Methods:

- <code title="get /coaching/principles/{principle_id}">client.Coaching.Principles.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrincipleService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, principleID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrinciple">CoachingPrinciple</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /coaching/principles">client.Coaching.Principles.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrincipleService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrincipleListParams">CoachingPrincipleListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrincipleListResponse">CoachingPrincipleListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /coaching/principles/random">client.Coaching.Principles.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrincipleService.GetRandom">GetRandom</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachingPrinciple">CoachingPrinciple</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Biscuits

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Biscuit">Biscuit</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BiscuitListResponse">BiscuitListResponse</a>

Methods:

- <code title="get /biscuits/{biscuit_id}">client.Biscuits.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BiscuitService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, biscuitID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Biscuit">Biscuit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /biscuits">client.Biscuits.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BiscuitService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BiscuitListParams">BiscuitListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BiscuitListResponse">BiscuitListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /biscuits/fresh">client.Biscuits.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#BiscuitService.GetFresh">GetFresh</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Biscuit">Biscuit</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# PepTalk

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PepTalkGetResponse">PepTalkGetResponse</a>

Methods:

- <code title="get /pep-talk">client.PepTalk.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PepTalkService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PepTalkGetParams">PepTalkGetParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#PepTalkGetResponse">PepTalkGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Stream

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#StreamTestConnectionResponse">StreamTestConnectionResponse</a>

Methods:

- <code title="get /stream/test">client.Stream.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#StreamService.TestConnection">TestConnection</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#StreamTestConnectionResponse">StreamTestConnectionResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# TeamMembers

Params Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachSpecialty">CoachSpecialty</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MedicalSpecialty">MedicalSpecialty</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Position">Position</a>

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Coach">Coach</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#CoachSpecialty">CoachSpecialty</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#EquipmentManager">EquipmentManager</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MedicalSpecialty">MedicalSpecialty</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#MedicalStaff">MedicalStaff</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Player">Player</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#Position">Position</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberNewResponseUnion">TeamMemberNewResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberGetResponseUnion">TeamMemberGetResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberUpdateResponseUnion">TeamMemberUpdateResponseUnion</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListResponse">TeamMemberListResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListCoachesResponse">TeamMemberListCoachesResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListPlayersResponse">TeamMemberListPlayersResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListStaffResponse">TeamMemberListStaffResponse</a>

Methods:

- <code title="post /team-members">client.TeamMembers.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberNewParams">TeamMemberNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberNewResponseUnion">TeamMemberNewResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team-members/{member_id}">client.TeamMembers.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memberID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberGetResponseUnion">TeamMemberGetResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="patch /team-members/{member_id}">client.TeamMembers.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberService.Update">Update</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memberID <a href="https://pkg.go.dev/builtin#string">string</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberUpdateParams">TeamMemberUpdateParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberUpdateResponseUnion">TeamMemberUpdateResponseUnion</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team-members">client.TeamMembers.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListParams">TeamMemberListParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListResponse">TeamMemberListResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /team-members/{member_id}">client.TeamMembers.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, memberID <a href="https://pkg.go.dev/builtin#string">string</a>) <a href="https://pkg.go.dev/builtin#error">error</a></code>
- <code title="get /team-members/coaches/">client.TeamMembers.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberService.ListCoaches">ListCoaches</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListCoachesParams">TeamMemberListCoachesParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListCoachesResponse">TeamMemberListCoachesResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team-members/players/">client.TeamMembers.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberService.ListPlayers">ListPlayers</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListPlayersParams">TeamMemberListPlayersParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListPlayersResponse">TeamMemberListPlayersResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /team-members/staff/">client.TeamMembers.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberService.ListStaff">ListStaff</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, query <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListStaffParams">TeamMemberListStaffParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#TeamMemberListStaffResponse">TeamMemberListStaffResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Webhooks

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#RegisteredWebhook">RegisteredWebhook</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookNewResponse">WebhookNewResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookDeleteResponse">WebhookDeleteResponse</a>
- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookTriggerEventResponse">WebhookTriggerEventResponse</a>

Methods:

- <code title="post /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookService.New">New</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookNewParams">WebhookNewParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookNewResponse">WebhookNewResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks/{webhook_id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#RegisteredWebhook">RegisteredWebhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="get /webhooks">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookService.List">List</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*[]<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#RegisteredWebhook">RegisteredWebhook</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="delete /webhooks/{webhook_id}">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookService.Delete">Delete</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, webhookID <a href="https://pkg.go.dev/builtin#string">string</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookDeleteResponse">WebhookDeleteResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
- <code title="post /webhooks/trigger">client.Webhooks.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookService.TriggerEvent">TriggerEvent</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>, body <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookTriggerEventParams">WebhookTriggerEventParams</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#WebhookTriggerEventResponse">WebhookTriggerEventResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Health

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#HealthCheckResponse">HealthCheckResponse</a>

Methods:

- <code title="get /health">client.Health.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#HealthService.Check">Check</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#HealthCheckResponse">HealthCheckResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>

# Version

Response Types:

- <a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#VersionGetResponse">VersionGetResponse</a>

Methods:

- <code title="get /version">client.Version.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#VersionService.Get">Get</a>(ctx <a href="https://pkg.go.dev/context">context</a>.<a href="https://pkg.go.dev/context#Context">Context</a>) (\*<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go">believe</a>.<a href="https://pkg.go.dev/github.com/stainless-sdks/believe-go#VersionGetResponse">VersionGetResponse</a>, <a href="https://pkg.go.dev/builtin#error">error</a>)</code>
