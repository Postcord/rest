// Code generated by generate_interface.go; DO NOT EDIT.

package rest

import (
	"context"
	"image"

	"github.com/Postcord/objects"
)

// RESTClient is the interface that contains all functions in *rest.Client.
type RESTClient interface {
	AddGuildCommand(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *objects.ApplicationCommand) (*objects.ApplicationCommand, error)
	AddGuildMember(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *AddGuildMemberParams) (*objects.GuildMember, error)
	AddGuildMemberRole(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, objects.SnowflakeObject, string) error
	AddPinnedMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) error
	AddThreadMember(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) error
	BatchEditApplicationCommandPermissions(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, []*objects.GuildApplicationCommandPermissions) ([]*objects.GuildApplicationCommandPermissions, error)
	BeginGuildPrune(context.Context, objects.SnowflakeObject, *BeginGuildPruneParams) (int, error)
	BulkDeleteMessages(context.Context, objects.SnowflakeObject, *DeleteMessagesParams) error
	BulkOverwriteGlobalCommands(context.Context, objects.SnowflakeObject, []*objects.ApplicationCommand) ([]*objects.ApplicationCommand, error)
	BulkOverwriteGuildCommands(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, []*objects.ApplicationCommand) ([]*objects.ApplicationCommand, error)
	CreateBan(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *CreateGuildBanParams) error
	CreateChannelInvite(context.Context, objects.SnowflakeObject, *CreateInviteParams) (*objects.Invite, error)
	CreateCommand(context.Context, objects.SnowflakeObject, *objects.ApplicationCommand) (*objects.ApplicationCommand, error)
	CreateDM(context.Context, *CreateDMParams) (*objects.Channel, error)
	CreateFollowupMessage(context.Context, objects.SnowflakeObject, string, *CreateFollowupMessageParams) (*objects.Message, error)
	CreateGroupDM(context.Context, *CreateGroupDMParams) (*objects.Channel, error)
	CreateGuild(context.Context, *CreateGuildParams) (*objects.Guild, error)
	CreateGuildChannel(context.Context, objects.SnowflakeObject, *ChannelCreateParams) (*objects.Channel, error)
	CreateGuildFromTemplate(context.Context, string, string) (*objects.Guild, error)
	CreateGuildRole(context.Context, objects.SnowflakeObject, *CreateGuildRoleParams) (*objects.Role, error)
	CreateGuildScheduledEvent(context.Context, objects.SnowflakeObject, *CreateGuildScheduledEventParams) (*objects.GuildScheduledEvent, error)
	CreateGuildSticker(context.Context, objects.SnowflakeObject, *CreateGuildStickerParams) (*objects.Sticker, error)
	CreateGuildTemplate(context.Context, objects.SnowflakeObject, *CreateGuildTemplateParams) (*objects.Template, error)
	CreateInteractionResponse(context.Context, objects.SnowflakeObject, string, *objects.InteractionResponse) error
	CreateMessage(context.Context, objects.SnowflakeObject, *CreateMessageParams) (*objects.Message, error)
	CreateReaction(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, interface {}) error
	CreateWebhook(context.Context, objects.SnowflakeObject, *CreateWebhookParams) (*objects.Webhook, error)
	CrossPostMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) (*objects.Message, error)
	DeleteAllReactions(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) error
	DeleteChannel(context.Context, objects.SnowflakeObject, string) (*objects.Channel, error)
	DeleteChannelPermission(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, string) error
	DeleteCommand(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) error
	DeleteEmojiReactions(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, interface {}) error
	DeleteFollowupMessage(context.Context, objects.SnowflakeObject, string, objects.SnowflakeObject) error
	DeleteGuild(context.Context, objects.SnowflakeObject) error
	DeleteGuildCommand(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, objects.SnowflakeObject) error
	DeleteGuildIntegration(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, string) error
	DeleteGuildRole(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, string) error
	DeleteGuildScheduledEvent(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) error
	DeleteGuildSticker(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, ...string) error
	DeleteGuildTemplate(context.Context, objects.SnowflakeObject, string, string) (*objects.Template, error)
	DeleteInvite(context.Context, string, string) (*objects.Invite, error)
	DeleteMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) error
	DeleteOriginalInteractionResponse(context.Context, objects.SnowflakeObject, string) error
	DeleteOwnReaction(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, interface {}) error
	DeletePinnedMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) error
	DeleteUserReaction(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, objects.SnowflakeObject, interface {}) error
	DeleteWebhook(context.Context, objects.SnowflakeObject) error
	DeleteWebhookMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, string) error
	DeleteWebhookWithToken(context.Context, objects.SnowflakeObject, string) error
	EditApplicationCommandPermissions(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, objects.SnowflakeObject, []*objects.ApplicationCommandPermissions) (*objects.GuildApplicationCommandPermissions, error)
	EditChannelPermissions(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *EditChannelParams) error
	EditFollowupMessage(context.Context, objects.SnowflakeObject, string, objects.SnowflakeObject, *EditWebhookMessageParams) (*objects.Message, error)
	EditMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *EditMessageParams) (*objects.Message, error)
	EditOriginalInteractionResponse(context.Context, objects.SnowflakeObject, string, *EditWebhookMessageParams) (*objects.Message, error)
	EditWebhookMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, string, *EditWebhookMessageParams) (*objects.Message, error)
	ExecuteWebhook(context.Context, objects.SnowflakeObject, string, *ExecuteWebhookParams) (*objects.Message, error)
	FollowNewsChannel(context.Context, objects.SnowflakeObject) (*objects.FollowedChannel, error)
	Gateway(context.Context) (*objects.Gateway, error)
	GatewayBot(context.Context) (*objects.Gateway, error)
	GetApplicationCommandPermissions(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, objects.SnowflakeObject) (*objects.GuildApplicationCommandPermissions, error)
	GetAuditLogs(context.Context, objects.SnowflakeObject, *GetAuditLogParams) (*objects.AuditLog, error)
	GetChannel(context.Context, objects.SnowflakeObject) (*objects.Channel, error)
	GetChannelInvites(context.Context, objects.SnowflakeObject) ([]*objects.Invite, error)
	GetChannelMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) (*objects.Message, error)
	GetChannelMessages(context.Context, objects.SnowflakeObject, *GetChannelMessagesParams) ([]*objects.Message, error)
	GetChannelWebhooks(context.Context, objects.SnowflakeObject) ([]*objects.Webhook, error)
	GetCommand(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) (*objects.ApplicationCommand, error)
	GetCommands(context.Context, objects.SnowflakeObject) ([]*objects.ApplicationCommand, error)
	GetCurrentUser(context.Context) (*objects.User, error)
	GetCurrentUserGuildMember(context.Context, objects.SnowflakeObject) (*objects.GuildMember, error)
	GetCurrentUserGuilds(context.Context, *CurrentUserGuildsParams) ([]*objects.Guild, error)
	GetFollowupMessage(context.Context, objects.SnowflakeObject, string, objects.SnowflakeObject) (*objects.Message, error)
	GetGuild(context.Context, objects.SnowflakeObject) (*objects.Guild, error)
	GetGuildApplicationCommandPermissions(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) ([]*objects.GuildApplicationCommandPermissions, error)
	GetGuildBan(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) (*objects.Ban, error)
	GetGuildBans(context.Context, objects.SnowflakeObject) ([]*objects.Ban, error)
	GetGuildChannels(context.Context, objects.SnowflakeObject) ([]*objects.Channel, error)
	GetGuildCommand(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, objects.SnowflakeObject) (*objects.ApplicationCommand, error)
	GetGuildCommands(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) ([]*objects.ApplicationCommand, error)
	GetGuildIntegrations(context.Context, objects.SnowflakeObject) ([]*objects.Integration, error)
	GetGuildInvites(context.Context, objects.SnowflakeObject) ([]*objects.Invite, error)
	GetGuildMember(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) (*objects.GuildMember, error)
	GetGuildPreview(context.Context, objects.SnowflakeObject) (*objects.GuildPreview, error)
	GetGuildPruneCount(context.Context, objects.SnowflakeObject, *GetGuildPruneCountParams) (int, error)
	GetGuildRoles(context.Context, objects.SnowflakeObject) ([]*objects.Role, error)
	GetGuildScheduledEvent(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, ...*GetGuildScheduledEventParams) (*objects.GuildScheduledEvent, error)
	GetGuildScheduledEventUsers(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, ...*GetGuildScheduledEventUsersParams) ([]*objects.GuildScheduledEventUser, error)
	GetGuildSticker(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) (*objects.Sticker, error)
	GetGuildTemplates(context.Context, objects.SnowflakeObject) ([]*objects.Template, error)
	GetGuildVanityURL(context.Context, objects.SnowflakeObject) (*objects.Invite, error)
	GetGuildVoiceRegions(context.Context, objects.SnowflakeObject) ([]*objects.VoiceRegion, error)
	GetGuildWebhooks(context.Context, objects.SnowflakeObject) ([]*objects.Webhook, error)
	GetGuildWelcomeScreen(context.Context, objects.SnowflakeObject) (*objects.MembershipScreening, error)
	GetGuildWidget(context.Context, objects.SnowflakeObject) (*objects.GuildWidgetJSON, error)
	GetGuildWidgetImage(context.Context, objects.SnowflakeObject, *GuildWidgetImageParams) (image.Image, error)
	GetGuildWidgetSettings(context.Context, objects.SnowflakeObject) (*objects.GuildWidget, error)
	GetInvite(context.Context, string, *GetInviteParams) (*objects.Invite, error)
	GetOriginalInteractionResponse(context.Context, objects.SnowflakeObject, string) (*objects.Message, error)
	GetPinnedMessages(context.Context, objects.SnowflakeObject) ([]*objects.Message, error)
	GetReactions(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, interface {}, *GetReactionsParams) ([]*objects.User, error)
	GetSticker(context.Context, objects.SnowflakeObject) (*objects.Sticker, error)
	GetTemplate(context.Context, string) (*objects.Template, error)
	GetUser(context.Context, objects.SnowflakeObject) (*objects.User, error)
	GetUserConnections(context.Context) ([]*objects.Connection, error)
	GetVoiceRegions(context.Context) ([]*objects.VoiceRegion, error)
	GetWebhook(context.Context, objects.SnowflakeObject) (*objects.Webhook, error)
	GetWebhookWithToken(context.Context, objects.SnowflakeObject, string) (*objects.Webhook, error)
	JoinThread(context.Context, objects.SnowflakeObject) error
	LeaveGuild(context.Context, objects.SnowflakeObject) error
	LeaveThread(context.Context, objects.SnowflakeObject) error
	ListActiveThreads(context.Context, objects.SnowflakeObject) ([]*ListThreadsResponse, error)
	ListGuildMembers(context.Context, objects.SnowflakeObject, *ListGuildMembersParams) ([]*objects.GuildMember, error)
	ListGuildStickers(context.Context, objects.SnowflakeObject) ([]*objects.Sticker, error)
	ListJoinedPrivateArchivedThreads(context.Context, objects.SnowflakeObject, ...*ListThreadsParams) (*ListThreadsResponse, error)
	ListNitroStickerPacks(context.Context) ([]*objects.StickerPack, error)
	ListPrivateArchivedThreads(context.Context, objects.SnowflakeObject, ...*ListThreadsParams) (*ListThreadsResponse, error)
	ListPublicArchivedThreads(context.Context, objects.SnowflakeObject, ...*ListThreadsParams) (*ListThreadsResponse, error)
	ListThreadMembers(context.Context, objects.SnowflakeObject) ([]*objects.ThreadMember, error)
	ModifyChannel(context.Context, objects.SnowflakeObject, *ModifyChannelParams) (*objects.Channel, error)
	ModifyCurrentUser(context.Context, *ModifyCurrentUserParams) (*objects.User, error)
	ModifyCurrentUserNick(context.Context, objects.SnowflakeObject, *ModifyCurrentUserNickParams) (*ModifyCurrentUserNickParams, error)
	ModifyGuild(context.Context, objects.SnowflakeObject, *ModifyGuildParams) (*objects.Guild, error)
	ModifyGuildChannelPositions(context.Context, objects.SnowflakeObject, []*ModifyChannelPositionParams, string) error
	ModifyGuildMember(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *ModifyGuildMemberParams) (*objects.GuildMember, error)
	ModifyGuildRole(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *ModifyGuildRoleParams) (*objects.Role, error)
	ModifyGuildRolePositions(context.Context, objects.SnowflakeObject, []*ModifyGuildRolePositionsParams) ([]*objects.Role, error)
	ModifyGuildScheduledEvent(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *ModifyGuildScheduledEventParams) (*objects.GuildScheduledEvent, error)
	ModifyGuildSticker(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *BaseStickerParams) (*objects.Sticker, error)
	ModifyGuildTemplate(context.Context, objects.SnowflakeObject, string, *ModifyGuildTemplateParams) (*objects.Template, error)
	ModifyGuildWelcomeScreen(context.Context, objects.SnowflakeObject, *ModifyGuildMembershipScreeningParams) (*objects.MembershipScreening, error)
	ModifyGuildWidget(context.Context, objects.SnowflakeObject, *GuildWidgetParams) (*objects.GuildWidget, error)
	ModifyWebhook(context.Context, objects.SnowflakeObject, *ModifyWebhookParams) (*objects.Webhook, error)
	ModifyWebhookWithToken(context.Context, objects.SnowflakeObject, string, *ModifyWebhookWithTokenParams) (*objects.Webhook, error)
	RemoveGuildBan(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, string) error
	RemoveGuildMember(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, string) error
	RemoveGuildMemberRole(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, objects.SnowflakeObject, string) error
	RemoveThreadMember(context.Context, objects.SnowflakeObject, objects.SnowflakeObject) error
	StartThread(context.Context, objects.SnowflakeObject, *StartThreadParams) (*objects.Channel, error)
	StartThreadWithMessage(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *StartThreadParams) (*objects.Channel, error)
	StartTyping(context.Context, objects.SnowflakeObject) error
	SyncGuildTemplate(context.Context, objects.SnowflakeObject, string) (*objects.Template, error)
	UpdateCommand(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, *objects.ApplicationCommand) (*objects.ApplicationCommand, error)
	UpdateGuildCommand(context.Context, objects.SnowflakeObject, objects.SnowflakeObject, objects.SnowflakeObject, *objects.ApplicationCommand) (*objects.ApplicationCommand, error)
}

var _ RESTClient = (*Client)(nil)
