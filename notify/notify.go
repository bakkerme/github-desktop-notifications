package notify

// NotificationAction is a callback invoked when the
// user interacts with the Notification
type NotificationAction func(n *Notification)

// Notification represents the information to be conveyed
// in a notification.
type Notification struct {
	Summary string
	Body    string
}

// Notifier is an interface for sending notifications to the
// operating system
type Notifier interface {
	RegisterActionHandler(nAction NotificationAction)
	SendNotification(n *Notification)
}

// Notify is a system that allows notifications to be sent
// to the host operating system
type Notify struct {
	actionHandler NotificationAction
}

// New creates a new Notify instance
func New() *Notify {
	return &Notify{}
}

// RegisterActionHandler registers a handler for any user
// actions made on the notification
func (notify Notify) RegisterActionHandler(nAction NotificationAction) {
	notify.actionHandler = nAction
}

// SendNotification sends a notification to the user
func (notify Notify) SendNotification(n *Notification) {

}
