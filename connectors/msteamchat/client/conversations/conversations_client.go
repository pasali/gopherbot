package conversations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"
)

// New creates a new conversations API client.
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Client {
	return &Client{transport: transport, formats: formats}
}

/*
Client for conversations API
*/
type Client struct {
	transport runtime.ClientTransport
	formats   strfmt.Registry
}

/*
ConversationsCreateConversation creates conversation

Create a new Conversation.

POST to this method with a
* Bot being the bot creating the conversation
* IsGroup set to true if this is not a direct message (default is false)
* Members array contining the members you want to have be in the conversation.

The return value is a ResourceResponse which contains a conversation id which is suitable for use
in the message payload and REST API uris.

Most channels only support the semantics of bots initiating a direct message conversation.  An example of how to do that would be:

```
var resource = await connector.conversations.CreateConversation(new ConversationParameters(){ Bot = bot, members = new ChannelAccount[] { new ChannelAccount("user1") } );
await connect.Conversations.SendToConversationAsync(resource.Id, new Activity() ... ) ;

```
*/
func (a *Client) ConversationsCreateConversation(params *ConversationsCreateConversationParams) (*ConversationsCreateConversationOK, *ConversationsCreateConversationCreated, *ConversationsCreateConversationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsCreateConversationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Conversations_CreateConversation",
		Method:             "POST",
		PathPattern:        "/v3/conversations",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml", "text/json", "text/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsCreateConversationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ConversationsCreateConversationOK:
		return value, nil, nil, nil
	case *ConversationsCreateConversationCreated:
		return nil, value, nil, nil
	case *ConversationsCreateConversationAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ConversationsDeleteActivity deletes activity

Delete an existing activity.

Some channels allow you to delete an existing activity, and if successful this method will remove the specified activity.
*/
func (a *Client) ConversationsDeleteActivity(params *ConversationsDeleteActivityParams) (*ConversationsDeleteActivityOK, *ConversationsDeleteActivityAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsDeleteActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Conversations_DeleteActivity",
		Method:             "DELETE",
		PathPattern:        "/v3/conversations/{conversationId}/activities/{activityId}",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsDeleteActivityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, err
	}
	switch value := result.(type) {
	case *ConversationsDeleteActivityOK:
		return value, nil, nil
	case *ConversationsDeleteActivityAccepted:
		return nil, value, nil
	}
	return nil, nil, nil

}

/*
ConversationsGetActivityMembers gets activity members

Enumerate the members of an activity.

This REST API takes a ConversationId and a ActivityId, returning an array of ChannelAccount objects representing the members of the particular activity in the conversation.
*/
func (a *Client) ConversationsGetActivityMembers(params *ConversationsGetActivityMembersParams) (*ConversationsGetActivityMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsGetActivityMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Conversations_GetActivityMembers",
		Method:             "GET",
		PathPattern:        "/v3/conversations/{conversationId}/activities/{activityId}/members",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsGetActivityMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConversationsGetActivityMembersOK), nil

}

/*
ConversationsGetConversationMembers gets conversation members

Enumerate the members of a converstion.

This REST API takes a ConversationId and returns an array of ChannelAccount objects representing the members of the conversation.
*/
func (a *Client) ConversationsGetConversationMembers(params *ConversationsGetConversationMembersParams) (*ConversationsGetConversationMembersOK, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsGetConversationMembersParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Conversations_GetConversationMembers",
		Method:             "GET",
		PathPattern:        "/v3/conversations/{conversationId}/members",
		ProducesMediaTypes: []string{"application/json", "application/xml", "text/json", "text/xml"},
		ConsumesMediaTypes: []string{""},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsGetConversationMembersReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, err
	}
	return result.(*ConversationsGetConversationMembersOK), nil

}

/*
ConversationsReplyToActivity replies to activity

This method allows you to reply to an activity.

This is slightly different from SendToConversation().
* SendToConverstion(conversationId) - will append the activity to the end of the conversation according to the timestamp or semantics of the channel.
* ReplyToActivity(conversationId,ActivityId) - adds the activity as a reply to another activity, if the channel supports it. If the channel does not support nested replies, ReplyToActivity falls back to SendToConversation.

Use ReplyToActivity when replying to a specific activity in the conversation.

Use SendToConversation in all other cases.
*/
func (a *Client) ConversationsReplyToActivity(params *ConversationsReplyToActivityParams) (*ConversationsReplyToActivityOK, *ConversationsReplyToActivityCreated, *ConversationsReplyToActivityAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsReplyToActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Conversations_ReplyToActivity",
		Method:             "POST",
		PathPattern:        "/v3/conversations/{conversationId}/activities/{activityId}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml", "text/json", "text/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsReplyToActivityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ConversationsReplyToActivityOK:
		return value, nil, nil, nil
	case *ConversationsReplyToActivityCreated:
		return nil, value, nil, nil
	case *ConversationsReplyToActivityAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ConversationsSendToConversation sends to conversation

This method allows you to send an activity to the end of a conversation.

This is slightly different from ReplyToActivity().
* SendToConverstion(conversationId) - will append the activity to the end of the conversation according to the timestamp or semantics of the channel.
* ReplyToActivity(conversationId,ActivityId) - adds the activity as a reply to another activity, if the channel supports it. If the channel does not support nested replies, ReplyToActivity falls back to SendToConversation.

Use ReplyToActivity when replying to a specific activity in the conversation.

Use SendToConversation in all other cases.
*/
func (a *Client) ConversationsSendToConversation(params *ConversationsSendToConversationParams) (*ConversationsSendToConversationOK, *ConversationsSendToConversationCreated, *ConversationsSendToConversationAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsSendToConversationParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Conversations_SendToConversation",
		Method:             "POST",
		PathPattern:        "/v3/conversations/{conversationId}/activities",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml", "text/json", "text/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsSendToConversationReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ConversationsSendToConversationOK:
		return value, nil, nil, nil
	case *ConversationsSendToConversationCreated:
		return nil, value, nil, nil
	case *ConversationsSendToConversationAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ConversationsUpdateActivity updates activity

Edit an existing activity.

Some channels allow you to edit an existing activity to reflect the new state of a bot conversation.

For example, you can remove buttons after someone has clicked "Approve" button.
*/
func (a *Client) ConversationsUpdateActivity(params *ConversationsUpdateActivityParams) (*ConversationsUpdateActivityOK, *ConversationsUpdateActivityCreated, *ConversationsUpdateActivityAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsUpdateActivityParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Conversations_UpdateActivity",
		Method:             "PUT",
		PathPattern:        "/v3/conversations/{conversationId}/activities/{activityId}",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml", "text/json", "text/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsUpdateActivityReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ConversationsUpdateActivityOK:
		return value, nil, nil, nil
	case *ConversationsUpdateActivityCreated:
		return nil, value, nil, nil
	case *ConversationsUpdateActivityAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

/*
ConversationsUploadAttachment uploads attachment

Upload an attachment directly into a channel's blob storage.

This is useful because it allows you to store data in a compliant store when dealing with enterprises.

The response is a ResourceResponse which contains an AttachmentId which is suitable for using with the attachments API.
*/
func (a *Client) ConversationsUploadAttachment(params *ConversationsUploadAttachmentParams) (*ConversationsUploadAttachmentOK, *ConversationsUploadAttachmentCreated, *ConversationsUploadAttachmentAccepted, error) {
	// TODO: Validate the params before sending
	if params == nil {
		params = NewConversationsUploadAttachmentParams()
	}

	result, err := a.transport.Submit(&runtime.ClientOperation{
		ID:                 "Conversations_UploadAttachment",
		Method:             "POST",
		PathPattern:        "/v3/conversations/{conversationId}/attachments",
		ProducesMediaTypes: []string{"application/json", "text/json"},
		ConsumesMediaTypes: []string{"application/json", "application/x-www-form-urlencoded", "application/xml", "text/json", "text/xml"},
		Schemes:            []string{"https"},
		Params:             params,
		Reader:             &ConversationsUploadAttachmentReader{formats: a.formats},
		Context:            params.Context,
		Client:             params.HTTPClient,
	})
	if err != nil {
		return nil, nil, nil, err
	}
	switch value := result.(type) {
	case *ConversationsUploadAttachmentOK:
		return value, nil, nil, nil
	case *ConversationsUploadAttachmentCreated:
		return nil, value, nil, nil
	case *ConversationsUploadAttachmentAccepted:
		return nil, nil, value, nil
	}
	return nil, nil, nil, nil

}

// SetTransport changes the transport on the client
func (a *Client) SetTransport(transport runtime.ClientTransport) {
	a.transport = transport
}
