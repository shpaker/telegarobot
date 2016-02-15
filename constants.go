package telegarobot

const (
	endpoint string = "https://api.telegram.org/bot"

	ParseModeHTML     string = "HTML"
	ParseModeMarkdown string = "Markdown"

	ActionTyping         string = "typing"          // for text messages
	ActionUploadPhoto    string = "upload_photo"    // for photos
	ActionRecordVideo    string = "record_video"    // for videos
	ActionUploadVideo    string = "upload_video"    // for videos
	ActionRecordAudio    string = "record_audio"    // for audio files
	ActionUploadAudio    string = "upload_audio"    // for audio files
	ActionUploadDocument string = "upload_document" // for general files
	ActionFindLocation   string = "find_location"   // for location data.
)
