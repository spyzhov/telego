package telego

import "context"

// getUpdates
// https://core.telegram.org/bots/api#getupdates
// Use this method to receive incoming updates using long polling (wiki). An Array of Update objects
// is returned.
func (b *Bot) GetUpdates(ctx context.Context, request *GetUpdatesRequest) (result []*Update, err error) {
	result = make([]*Update, 0)
	return result, b.postResult(ctx, "getUpdates", request, &result)
}

// setWebhook
// https://core.telegram.org/bots/api#setwebhook
// Use this method to specify a url and receive incoming updates via an outgoing webhook. Whenever
// there is an update for the bot, we will send an HTTPS POST request to the specified url, containing a
// JSON-serialized Update. In case of an unsuccessful request, we will give up after a reasonable
// amount of attempts. Returns True on success.
func (b *Bot) SetWebhook(ctx context.Context, request *SetWebhookRequest) (result bool, err error) {
	return result, b.postResult(ctx, "setWebhook", request, &result)
}

// deleteWebhook
// https://core.telegram.org/bots/api#deletewebhook
// Use this method to remove webhook integration if you decide to switch back to getUpdates. Returns
// True on success. Requires no parameters.
func (b *Bot) DeleteWebhook(ctx context.Context) (result bool, err error) {
	return result, b.postResult(ctx, "deleteWebhook", nil, &result)
}

// getWebhookInfo
// https://core.telegram.org/bots/api#getwebhookinfo
// Use this method to get current webhook status. Requires no parameters. On success, returns a
// WebhookInfo object. If the bot is using getUpdates, will return an object with the url field empty.
func (b *Bot) GetWebhookInfo(ctx context.Context) (result *WebhookInfo, err error) {
	result = new(WebhookInfo)
	return result, b.postResult(ctx, "getWebhookInfo", nil, &result)
}

// getMe
// https://core.telegram.org/bots/api#getme
// A simple method for testing your bot's auth token. Requires no parameters. Returns basic
// information about the bot in form of a User object.
func (b *Bot) GetMe(ctx context.Context) (result *User, err error) {
	result = new(User)
	return result, b.postResult(ctx, "getMe", nil, &result)
}

// sendMessage
// https://core.telegram.org/bots/api#sendmessage
// Use this method to send text messages. On success, the sent Message is returned.
func (b *Bot) SendMessage(ctx context.Context, request *SendMessageRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendMessage", request, &result)
}

// forwardMessage
// https://core.telegram.org/bots/api#forwardmessage
// Use this method to forward messages of any kind. On success, the sent Message is returned.
func (b *Bot) ForwardMessage(ctx context.Context, request *ForwardMessageRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "forwardMessage", request, &result)
}

// sendPhoto
// https://core.telegram.org/bots/api#sendphoto
// Use this method to send photos. On success, the sent Message is returned.
func (b *Bot) SendPhoto(ctx context.Context, request *SendPhotoRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendPhoto", request, &result)
}

// sendAudio
// https://core.telegram.org/bots/api#sendaudio
// Use this method to send audio files, if you want Telegram clients to display them in the music
// player. Your audio must be in the .mp3 format. On success, the sent Message is returned. Bots can
// currently send audio files of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendAudio(ctx context.Context, request *SendAudioRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendAudio", request, &result)
}

// sendDocument
// https://core.telegram.org/bots/api#senddocument
// Use this method to send general files. On success, the sent Message is returned. Bots can
// currently send files of any type of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendDocument(ctx context.Context, request *SendDocumentRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendDocument", request, &result)
}

// sendVideo
// https://core.telegram.org/bots/api#sendvideo
// Use this method to send video files, Telegram clients support mp4 videos (other formats may be
// sent as Document). On success, the sent Message is returned. Bots can currently send video files of up
// to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendVideo(ctx context.Context, request *SendVideoRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendVideo", request, &result)
}

// sendAnimation
// https://core.telegram.org/bots/api#sendanimation
// Use this method to send animation files (GIF or H.264/MPEG-4 AVC video without sound). On success,
// the sent Message is returned. Bots can currently send animation files of up to 50 MB in size, this
// limit may be changed in the future.
func (b *Bot) SendAnimation(ctx context.Context, request *SendAnimationRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendAnimation", request, &result)
}

// sendVoice
// https://core.telegram.org/bots/api#sendvoice
// Use this method to send audio files, if you want Telegram clients to display the file as a
// playable voice message. For this to work, your audio must be in an .ogg file encoded with OPUS (other
// formats may be sent as Audio or Document). On success, the sent Message is returned. Bots can currently
// send voice messages of up to 50 MB in size, this limit may be changed in the future.
func (b *Bot) SendVoice(ctx context.Context, request *SendVoiceRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendVoice", request, &result)
}

// sendVideoNote
// https://core.telegram.org/bots/api#sendvideonote
// As of v.4.0, Telegram clients support rounded square mp4 videos of up to 1 minute long. Use this
// method to send video messages. On success, the sent Message is returned.
func (b *Bot) SendVideoNote(ctx context.Context, request *SendVideoNoteRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendVideoNote", request, &result)
}

// sendMediaGroup
// https://core.telegram.org/bots/api#sendmediagroup
// Use this method to send a group of photos or videos as an album. On success, an array of the sent
// Messages is returned.
func (b *Bot) SendMediaGroup(ctx context.Context, request *SendMediaGroupRequest) (result []*Message, err error) {
	result = make([]*Message, 0)
	return result, b.postResult(ctx, "sendMediaGroup", request, &result)
}

// sendLocation
// https://core.telegram.org/bots/api#sendlocation
// Use this method to send point on the map. On success, the sent Message is returned.
func (b *Bot) SendLocation(ctx context.Context, request *SendLocationRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendLocation", request, &result)
}

// editMessageLiveLocation
// https://core.telegram.org/bots/api#editmessagelivelocation
// Use this method to edit live location messages. A location can be edited until its live_period
// expires or editing is explicitly disabled by a call to stopMessageLiveLocation. On success, if the
// edited message was sent by the bot, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageLiveLocation(ctx context.Context, request *EditMessageLiveLocationRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "editMessageLiveLocation", request, &result)
}

// stopMessageLiveLocation
// https://core.telegram.org/bots/api#stopmessagelivelocation
// Use this method to stop updating a live location message before live_period expires. On success,
// if the message was sent by the bot, the sent Message is returned, otherwise True is returned.
func (b *Bot) StopMessageLiveLocation(ctx context.Context, request *StopMessageLiveLocationRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "stopMessageLiveLocation", request, &result)
}

// sendVenue
// https://core.telegram.org/bots/api#sendvenue
// Use this method to send information about a venue. On success, the sent Message is returned.
func (b *Bot) SendVenue(ctx context.Context, request *SendVenueRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendVenue", request, &result)
}

// sendContact
// https://core.telegram.org/bots/api#sendcontact
// Use this method to send phone contacts. On success, the sent Message is returned.
func (b *Bot) SendContact(ctx context.Context, request *SendContactRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendContact", request, &result)
}

// sendPoll
// https://core.telegram.org/bots/api#sendpoll
// Use this method to send a native poll. A native poll can't be sent to a private chat. On success,
// the sent Message is returned.
func (b *Bot) SendPoll(ctx context.Context, request *SendPollRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendPoll", request, &result)
}

// sendChatAction
// https://core.telegram.org/bots/api#sendchataction
// Use this method when you need to tell the user that something is happening on the bot's side. The
// status is set for 5 seconds or less (when a message arrives from your bot, Telegram clients clear
// its typing status). Returns True on success.
func (b *Bot) SendChatAction(ctx context.Context, request *SendChatActionRequest) (result bool, err error) {
	return result, b.postResult(ctx, "sendChatAction", request, &result)
}

// getUserProfilePhotos
// https://core.telegram.org/bots/api#getuserprofilephotos
// Use this method to get a list of profile pictures for a user. Returns a UserProfilePhotos object.
func (b *Bot) GetUserProfilePhotos(ctx context.Context, request *GetUserProfilePhotosRequest) (result interface{}, err error) {
	return result, b.postResult(ctx, "getUserProfilePhotos", request, &result)
}

// getFile
// https://core.telegram.org/bots/api#getfile
// Use this method to get basic info about a file and prepare it for downloading. For the moment,
// bots can download files of up to 20MB in size. On success, a File object is returned. The file can then
// be downloaded via the link https://api.telegram.org/file/bot<token>/<file_path>, where <file_path>
// is taken from the response. It is guaranteed that the link will be valid for at least 1 hour. When
// the link expires, a new one can be requested by calling getFile again.
func (b *Bot) GetFile(ctx context.Context, request *GetFileRequest) (result *File, err error) {
	result = new(File)
	return result, b.postResult(ctx, "getFile", request, &result)
}

// kickChatMember
// https://core.telegram.org/bots/api#kickchatmember
// Use this method to kick a user from a group, a supergroup or a channel. In the case of supergroups
// and channels, the user will not be able to return to the group on their own using invite links,
// etc., unless unbanned first. The bot must be an administrator in the chat for this to work and must
// have the appropriate admin rights. Returns True on success.
func (b *Bot) KickChatMember(ctx context.Context, request *KickChatMemberRequest) (result bool, err error) {
	return result, b.postResult(ctx, "kickChatMember", request, &result)
}

// unbanChatMember
// https://core.telegram.org/bots/api#unbanchatmember
// Use this method to unban a previously kicked user in a supergroup or channel. The user will not
// return to the group or channel automatically, but will be able to join via link, etc. The bot must be
// an administrator for this to work. Returns True on success.
func (b *Bot) UnbanChatMember(ctx context.Context, request *UnbanChatMemberRequest) (result bool, err error) {
	return result, b.postResult(ctx, "unbanChatMember", request, &result)
}

// restrictChatMember
// https://core.telegram.org/bots/api#restrictchatmember
// Use this method to restrict a user in a supergroup. The bot must be an administrator in the
// supergroup for this to work and must have the appropriate admin rights. Pass True for all boolean
// parameters to lift restrictions from a user. Returns True on success.
func (b *Bot) RestrictChatMember(ctx context.Context, request *RestrictChatMemberRequest) (result bool, err error) {
	return result, b.postResult(ctx, "restrictChatMember", request, &result)
}

// promoteChatMember
// https://core.telegram.org/bots/api#promotechatmember
// Use this method to promote or demote a user in a supergroup or a channel. The bot must be an
// administrator in the chat for this to work and must have the appropriate admin rights. Pass False for all
// boolean parameters to demote a user. Returns True on success.
func (b *Bot) PromoteChatMember(ctx context.Context, request *PromoteChatMemberRequest) (result bool, err error) {
	return result, b.postResult(ctx, "promoteChatMember", request, &result)
}

// exportChatInviteLink
// https://core.telegram.org/bots/api#exportchatinvitelink
// Use this method to generate a new invite link for a chat; any previously generated link is
// revoked. The bot must be an administrator in the chat for this to work and must have the appropriate admin
// rights. Returns the new invite link as String on success.
func (b *Bot) ExportChatInviteLink(ctx context.Context, request *ExportChatInviteLinkRequest) (result string, err error) {
	return result, b.postResult(ctx, "exportChatInviteLink", request, &result)
}

// setChatPhoto
// https://core.telegram.org/bots/api#setchatphoto
// Use this method to set a new profile photo for the chat. Photos can't be changed for private
// chats. The bot must be an administrator in the chat for this to work and must have the appropriate admin
// rights. Returns True on success.
func (b *Bot) SetChatPhoto(ctx context.Context, request *SetChatPhotoRequest) (result bool, err error) {
	return result, b.postResult(ctx, "setChatPhoto", request, &result)
}

// deleteChatPhoto
// https://core.telegram.org/bots/api#deletechatphoto
// Use this method to delete a chat photo. Photos can't be changed for private chats. The bot must be
// an administrator in the chat for this to work and must have the appropriate admin rights. Returns
// True on success.
func (b *Bot) DeleteChatPhoto(ctx context.Context, request *DeleteChatPhotoRequest) (result bool, err error) {
	return result, b.postResult(ctx, "deleteChatPhoto", request, &result)
}

// setChatTitle
// https://core.telegram.org/bots/api#setchattitle
// Use this method to change the title of a chat. Titles can't be changed for private chats. The bot
// must be an administrator in the chat for this to work and must have the appropriate admin rights.
// Returns True on success.
func (b *Bot) SetChatTitle(ctx context.Context, request *SetChatTitleRequest) (result bool, err error) {
	return result, b.postResult(ctx, "setChatTitle", request, &result)
}

// setChatDescription
// https://core.telegram.org/bots/api#setchatdescription
// Use this method to change the description of a supergroup or a channel. The bot must be an
// administrator in the chat for this to work and must have the appropriate admin rights. Returns True on
// success.
func (b *Bot) SetChatDescription(ctx context.Context, request *SetChatDescriptionRequest) (result bool, err error) {
	return result, b.postResult(ctx, "setChatDescription", request, &result)
}

// pinChatMessage
// https://core.telegram.org/bots/api#pinchatmessage
// Use this method to pin a message in a group, a supergroup, or a channel. The bot must be an
// administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the
// supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
func (b *Bot) PinChatMessage(ctx context.Context, request *PinChatMessageRequest) (result bool, err error) {
	return result, b.postResult(ctx, "pinChatMessage", request, &result)
}

// unpinChatMessage
// https://core.telegram.org/bots/api#unpinchatmessage
// Use this method to unpin a message in a group, a supergroup, or a channel. The bot must be an
// administrator in the chat for this to work and must have the ‘can_pin_messages’ admin right in the
// supergroup or ‘can_edit_messages’ admin right in the channel. Returns True on success.
func (b *Bot) UnpinChatMessage(ctx context.Context, request *UnpinChatMessageRequest) (result bool, err error) {
	return result, b.postResult(ctx, "unpinChatMessage", request, &result)
}

// leaveChat
// https://core.telegram.org/bots/api#leavechat
// Use this method for your bot to leave a group, supergroup or channel. Returns True on success.
func (b *Bot) LeaveChat(ctx context.Context, request *LeaveChatRequest) (result bool, err error) {
	return result, b.postResult(ctx, "leaveChat", request, &result)
}

// getChat
// https://core.telegram.org/bots/api#getchat
// Use this method to get up to date information about the chat (current name of the user for
// one-on-one conversations, current username of a user, group or channel, etc.). Returns a Chat object on
// success.
func (b *Bot) GetChat(ctx context.Context, request *GetChatRequest) (result *Chat, err error) {
	result = new(Chat)
	return result, b.postResult(ctx, "getChat", request, &result)
}

// getChatAdministrators
// https://core.telegram.org/bots/api#getchatadministrators
// Use this method to get a list of administrators in a chat. On success, returns an Array of
// ChatMember objects that contains information about all chat administrators except other bots. If the chat
// is a group or a supergroup and no administrators were appointed, only the creator will be returned.
func (b *Bot) GetChatAdministrators(ctx context.Context, request *GetChatAdministratorsRequest) (result []*ChatMember, err error) {
	result = make([]*ChatMember, 0)
	return result, b.postResult(ctx, "getChatAdministrators", request, &result)
}

// getChatMembersCount
// https://core.telegram.org/bots/api#getchatmemberscount
// Use this method to get the number of members in a chat. Returns Int on success.
func (b *Bot) GetChatMembersCount(ctx context.Context, request *GetChatMembersCountRequest) (result int, err error) {
	return result, b.postResult(ctx, "getChatMembersCount", request, &result)
}

// getChatMember
// https://core.telegram.org/bots/api#getchatmember
// Use this method to get information about a member of a chat. Returns a ChatMember object on
// success.
func (b *Bot) GetChatMember(ctx context.Context, request *GetChatMemberRequest) (result *ChatMember, err error) {
	result = new(ChatMember)
	return result, b.postResult(ctx, "getChatMember", request, &result)
}

// setChatStickerSet
// https://core.telegram.org/bots/api#setchatstickerset
// Use this method to set a new group sticker set for a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True
// on success.
func (b *Bot) SetChatStickerSet(ctx context.Context, request *SetChatStickerSetRequest) (result bool, err error) {
	return result, b.postResult(ctx, "setChatStickerSet", request, &result)
}

// deleteChatStickerSet
// https://core.telegram.org/bots/api#deletechatstickerset
// Use this method to delete a group sticker set from a supergroup. The bot must be an administrator
// in the chat for this to work and must have the appropriate admin rights. Use the field
// can_set_sticker_set optionally returned in getChat requests to check if the bot can use this method. Returns True
// on success.
func (b *Bot) DeleteChatStickerSet(ctx context.Context, request *DeleteChatStickerSetRequest) (result bool, err error) {
	return result, b.postResult(ctx, "deleteChatStickerSet", request, &result)
}

// answerCallbackQuery
// https://core.telegram.org/bots/api#answercallbackquery
// Use this method to send answers to callback queries sent from inline keyboards. The answer will be
// displayed to the user as a notification at the top of the chat screen or as an alert. On success,
// True is returned.
func (b *Bot) AnswerCallbackQuery(ctx context.Context, request *AnswerCallbackQueryRequest) (result bool, err error) {
	return result, b.postResult(ctx, "answerCallbackQuery", request, &result)
}

// editMessageText
// https://core.telegram.org/bots/api#editmessagetext
// Use this method to edit text and game messages. On success, if edited message is sent by the bot,
// the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageText(ctx context.Context, request *EditMessageTextRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "editMessageText", request, &result)
}

// editMessageCaption
// https://core.telegram.org/bots/api#editmessagecaption
// Use this method to edit captions of messages. On success, if edited message is sent by the bot,
// the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageCaption(ctx context.Context, request *EditMessageCaptionRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "editMessageCaption", request, &result)
}

// editMessageMedia
// https://core.telegram.org/bots/api#editmessagemedia
// Use this method to edit animation, audio, document, photo, or video messages. If a message is a
// part of a message album, then it can be edited only to a photo or a video. Otherwise, message type can
// be changed arbitrarily. When inline message is edited, new file can't be uploaded. Use previously
// uploaded file via its file_id or specify a URL. On success, if the edited message was sent by the
// bot, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageMedia(ctx context.Context, request *EditMessageMediaRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "editMessageMedia", request, &result)
}

// editMessageReplyMarkup
// https://core.telegram.org/bots/api#editmessagereplymarkup
// Use this method to edit only the reply markup of messages. On success, if edited message is sent
// by the bot, the edited Message is returned, otherwise True is returned.
func (b *Bot) EditMessageReplyMarkup(ctx context.Context, request *EditMessageReplyMarkupRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "editMessageReplyMarkup", request, &result)
}

// stopPoll
// https://core.telegram.org/bots/api#stoppoll
// Use this method to stop a poll which was sent by the bot. On success, the stopped Poll with the
// final results is returned.
func (b *Bot) StopPoll(ctx context.Context, request *StopPollRequest) (result *Poll, err error) {
	result = new(Poll)
	return result, b.postResult(ctx, "stopPoll", request, &result)
}

// deleteMessage
// https://core.telegram.org/bots/api#deletemessage
// Use this method to delete a message, including service messages, with the following limitations:-
// A message can only be deleted if it was sent less than 48 hours ago.- Bots can delete outgoing
// messages in private chats, groups, and supergroups.- Bots can delete incoming messages in private chats.-
// Bots granted can_post_messages permissions can delete outgoing messages in channels.- If the bot is
// an administrator of a group, it can delete any message there.- If the bot has can_delete_messages
// permission in a supergroup or a channel, it can delete any message there.Returns True on success.
func (b *Bot) DeleteMessage(ctx context.Context, request *DeleteMessageRequest) (result bool, err error) {
	return result, b.postResult(ctx, "deleteMessage", request, &result)
}

// sendSticker
// https://core.telegram.org/bots/api#sendsticker
// Use this method to send .webp stickers. On success, the sent Message is returned.
func (b *Bot) SendSticker(ctx context.Context, request *SendStickerRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendSticker", request, &result)
}

// getStickerSet
// https://core.telegram.org/bots/api#getstickerset
// Use this method to get a sticker set. On success, a StickerSet object is returned.
func (b *Bot) GetStickerSet(ctx context.Context, request *GetStickerSetRequest) (result *StickerSet, err error) {
	result = new(StickerSet)
	return result, b.postResult(ctx, "getStickerSet", request, &result)
}

// uploadStickerFile
// https://core.telegram.org/bots/api#uploadstickerfile
// Use this method to upload a .png file with a sticker for later use in createNewStickerSet and
// addStickerToSet methods (can be used multiple times). Returns the uploaded File on success.
func (b *Bot) UploadStickerFile(ctx context.Context, request *UploadStickerFileRequest) (result *File, err error) {
	result = new(File)
	return result, b.postResult(ctx, "uploadStickerFile", request, &result)
}

// createNewStickerSet
// https://core.telegram.org/bots/api#createnewstickerset
// Use this method to create new sticker set owned by a user. The bot will be able to edit the
// created sticker set. Returns True on success.
func (b *Bot) CreateNewStickerSet(ctx context.Context, request *CreateNewStickerSetRequest) (result bool, err error) {
	return result, b.postResult(ctx, "createNewStickerSet", request, &result)
}

// addStickerToSet
// https://core.telegram.org/bots/api#addstickertoset
// Use this method to add a new sticker to a set created by the bot. Returns True on success.
func (b *Bot) AddStickerToSet(ctx context.Context, request *AddStickerToSetRequest) (result bool, err error) {
	return result, b.postResult(ctx, "addStickerToSet", request, &result)
}

// setStickerPositionInSet
// https://core.telegram.org/bots/api#setstickerpositioninset
// Use this method to move a sticker in a set created by the bot to a specific position . Returns
// True on success.
func (b *Bot) SetStickerPositionInSet(ctx context.Context, request *SetStickerPositionInSetRequest) (result bool, err error) {
	return result, b.postResult(ctx, "setStickerPositionInSet", request, &result)
}

// deleteStickerFromSet
// https://core.telegram.org/bots/api#deletestickerfromset
// Use this method to delete a sticker from a set created by the bot. Returns True on success.
func (b *Bot) DeleteStickerFromSet(ctx context.Context, request *DeleteStickerFromSetRequest) (result bool, err error) {
	return result, b.postResult(ctx, "deleteStickerFromSet", request, &result)
}

// answerInlineQuery
// https://core.telegram.org/bots/api#answerinlinequery
// Use this method to send answers to an inline query. On success, True is returned.No more than 50
// results per query are allowed.
func (b *Bot) AnswerInlineQuery(ctx context.Context, request *AnswerInlineQueryRequest) (result bool, err error) {
	return result, b.postResult(ctx, "answerInlineQuery", request, &result)
}

// sendInvoice
// https://core.telegram.org/bots/api#sendinvoice
// Use this method to send invoices. On success, the sent Message is returned.
func (b *Bot) SendInvoice(ctx context.Context, request *SendInvoiceRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendInvoice", request, &result)
}

// answerShippingQuery
// https://core.telegram.org/bots/api#answershippingquery
// If you sent an invoice requesting a shipping address and the parameter is_flexible was specified,
// the Bot API will send an Update with a shipping_query field to the bot. Use this method to reply to
// shipping queries. On success, True is returned.
func (b *Bot) AnswerShippingQuery(ctx context.Context, request *AnswerShippingQueryRequest) (result bool, err error) {
	return result, b.postResult(ctx, "answerShippingQuery", request, &result)
}

// answerPreCheckoutQuery
// https://core.telegram.org/bots/api#answerprecheckoutquery
// Once the user has confirmed their payment and shipping details, the Bot API sends the final
// confirmation in the form of an Update with the field pre_checkout_query. Use this method to respond to
// such pre-checkout queries. On success, True is returned. Note: The Bot API must receive an answer
// within 10 seconds after the pre-checkout query was sent.
func (b *Bot) AnswerPreCheckoutQuery(ctx context.Context, request *AnswerPreCheckoutQueryRequest) (result bool, err error) {
	return result, b.postResult(ctx, "answerPreCheckoutQuery", request, &result)
}

// setPassportDataErrors
// https://core.telegram.org/bots/api#setpassportdataerrors
// Informs a user that some of the Telegram Passport elements they provided contains errors. The user
// will not be able to re-submit their Passport to you until the errors are fixed (the contents of the
// field for which you returned the error must change). Returns True on success.
func (b *Bot) SetPassportDataErrors(ctx context.Context, request *SetPassportDataErrorsRequest) (result bool, err error) {
	return result, b.postResult(ctx, "setPassportDataErrors", request, &result)
}

// sendGame
// https://core.telegram.org/bots/api#sendgame
// Use this method to send a game. On success, the sent Message is returned.
func (b *Bot) SendGame(ctx context.Context, request *SendGameRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "sendGame", request, &result)
}

// setGameScore
// https://core.telegram.org/bots/api#setgamescore
// Use this method to set the score of the specified user in a game. On success, if the message was
// sent by the bot, returns the edited Message, otherwise returns True. Returns an error, if the new
// score is not greater than the user's current score in the chat and force is False.
func (b *Bot) SetGameScore(ctx context.Context, request *SetGameScoreRequest) (result *Message, err error) {
	result = new(Message)
	return result, b.postResult(ctx, "setGameScore", request, &result)
}

// getGameHighScores
// https://core.telegram.org/bots/api#getgamehighscores
// Use this method to get data for high score tables. Will return the score of the specified user and
// several of his neighbors in a game. On success, returns an Array of GameHighScore objects.
func (b *Bot) GetGameHighScores(ctx context.Context, request *GetGameHighScoresRequest) (result []*GameHighScore, err error) {
	result = make([]*GameHighScore, 0)
	return result, b.postResult(ctx, "getGameHighScores", request, &result)
}
