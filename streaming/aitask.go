// File generated from our OpenAPI spec by Stainless. See CONTRIBUTING.md for details.

package streaming

import (
	"context"
	"encoding/json"
	"errors"
	"fmt"
	"net/http"
	"net/url"
	"slices"
	"time"

	"github.com/G-Core/gcore-go/internal/apijson"
	"github.com/G-Core/gcore-go/internal/apiquery"
	"github.com/G-Core/gcore-go/internal/requestconfig"
	"github.com/G-Core/gcore-go/option"
	"github.com/G-Core/gcore-go/packages/pagination"
	"github.com/G-Core/gcore-go/packages/param"
	"github.com/G-Core/gcore-go/packages/respjson"
)

// AITaskService contains methods and other services that help with interacting
// with the gcore API.
//
// Note, unlike clients, this service does not read variables from the environment
// automatically. You should not instantiate this service directly, and instead use
// the [NewAITaskService] method instead.
type AITaskService struct {
	Options []option.RequestOption
}

// NewAITaskService generates a new service that applies the given options to each
// request. These options are applied after the parent client's options (if there
// is one), and before any request-specific options.
func NewAITaskService(opts ...option.RequestOption) (r AITaskService) {
	r = AITaskService{}
	r.Options = opts
	return
}

// Creating an AI task.
//
// This method allows you to create an AI task for VOD video processing:
//
// - ASR: Transcribe video
// - ASR: Translate subtitles
// - CM: Sports detection
// - CM: Not Safe For Work (NSFW) content detection
// - CM: Soft nudity detection
// - CM: Hard nudity detection
// - CM: Objects recognition (soon)
//
// ![Auto generated subtitles example](https://demo-files.gvideo.io/apidocs/captions.gif)
//
// How to use:
//
// - Create an AI task, specify algorithm to use
// - Get `task_id`
// - Check a result using `.../ai/tasks/{task_id}` method
//
// For more detailed information, see the algorithm-specific sections below.
//
// **AI Automatic Speech Recognition (ASR)**
//
// AI is instrumental in automatic video processing for subtitles creation by using
// Automatic Speech Recognition (ASR) technology to transcribe spoken words into
// text, which can then be translated into multiple languages for broader
// accessibility.
//
// Categories:
//
//   - `transcription` – to create subtitles/captions from audio in the original
//     language.
//   - `translation` – to translate subtitles/captions from the original language to
//     99+ other languages.
//
// AI subtitle transcription and translation tools are highly efficient, processing
// large volumes of audio-visual content quickly and providing accurate
// transcriptions and translations with minimal human intervention. Additionally,
// AI-driven solutions can significantly reduce costs and turnaround times compared
// to traditional methods, making them an invaluable resource for content creators
// and broadcasters aiming to reach global audiences.
//
// Example response with positive result:
//
// ```
//
//	{
//	  "status": "SUCCESS",
//	  "result": {
//	    "subtitles": [
//	      {
//	          "start_time": "00:00:00.031",
//	          "end_time": "00:00:03.831",
//	          "text": "Come on team, ..."
//	      }, ...
//	    ]
//	    "vttContent": "WEBVTT\n\n1\n00:00:00.031 --> 00:00:03.831\nCome on team, ...",
//	    "concatenated_text": "Come on team, ...",
//	    "languages": [ "eng" ],
//	    "speech_detected": true
//	    }
//	  }, ...
//	}
//
// ```
//
// **AI Content Moderation (CM)**
//
// The AI Content Moderation API offers a powerful solution for analyzing video
// content to detect various categories of inappropriate material. Leveraging
// state-of-the-art AI models, this API ensures real-time analysis and flagging of
// sensitive or restricted content types, making it an essential tool for platforms
// requiring stringent content moderation.
//
// Categories:
//
//   - `nsfw`: Quick algorithm to detect pornographic material, ensuring content is
//     "not-safe-for-work" or normal.
//   - `hard_nudity`: Detailed analysis of video which detects explicit nudity
//     involving genitalia.
//   - `soft_nudity`: Detailed video analysis that reveals both explicit and partial
//     nudity, including the presence of male and female faces and other uncovered
//     body parts.
//   - `sport`: Recognizes various sporting activities.
//
// The AI Content Moderation API is an invaluable tool for managing and controlling
// the type of content being shared or streamed on your platform. By implementing
// this API, you can ensure compliance with community guidelines and legal
// requirements, as well as provide a safer environment for your users.
//
// Important notes:
//
//   - It's allowed to analyze still images too (where applicable). Format of image:
//     JPEG, PNG. In that case one image is the same as video of 1 second duration.
//   - Not all frames in the video are used for analysis, but only key frames
//     (Iframe). For example, if a key frame in a video is set every ±2 seconds, then
//     detection will only occur at these timestamps. If an object appears and
//     disappears between these time stamps, it will not be detected. We are working
//     on a version to analyze more frames, please contact your manager or our
//     support team to enable this method.
//
// Example response with positive result:
//
// ```
//
//	{
//	  "status": "SUCCESS",
//	  "result": {
//	      "nsfw_detected": true,
//	      "detection_results": [ "nsfw" ],
//	      "frames": [
//	          {
//	              "label": "nsfw",
//	              "confidence": 1.0,
//	              "frame_number": 24
//	          },...
//	      ]
//	  }
//	}
//
// ```
//
// **Additional information**
//
// Billing takes into account the duration of the analyzed video. Or the duration
// until the stop tag(where applicable), if the condition was triggered during the
// analysis.
//
// The heart of content moderation is AI, with additional services. They run on our
// own infrastructure, so the files/data are not transferred anywhere to external
// services. After processing, original files are also deleted from local storage
// of AI.
//
// Read more detailed information about our solution, and architecture, and
// benefits in the knowledge base and blog.
//
// **Algorithm-specific details**
//
// # Create AI ASR task
//
// Transcribing is the process of writing down the words you hear in an audio. Our
// solution allows you to transcribe audio from your video and get subtitles
// automatically. To do this, we use modern AI models.
//
// The result:
//
//   - Transcription – subtitles in the original language. I.e. audio is in English –
//     subtitles are in English too, audio is in German – subtitles are in German
//     too.
//   - Translation – subtitles is translated from the original language to any other
//     language.
//
// **How to use?**
//
//   - Explicit call to this AI method. Applicable for any file stored with us or
//     located on the Internet.
//   - Standard video upload but with automatic subtitle generation. Look at
//     ["VOD uploading"](/api-reference/streaming/videos/create-video).
//
// **What language will the subtitles be in?**
//
// You can specify the language explicitly, then it will be used to create
// subtitles: the source language in the audio, the resulting subtitle language. If
// this is not set, the system will run auto language identification and the
// subtitles will be in the detected language. The method also works based on AI
// analysis.
//
// Additionally, when this is not set, we also support recognition of alternate
// languages in the video (code-switching). For example, when in a video different
// speakers speak several languages, or when they switch from their native language
// to English and back. Thus when you have multiple languages in the video it is
// better to not specify an "audio_language" otherwise AI may force the system to
// recognize gibberish.
//
// **What can be transcribed?**
//
// Service uses additional methods to detect presence of speech in audio track,
// thus improving the detection of any human conversations:
//
// - Speech of one speaker,
// - Speech of several speakers,
// - Speech in different languages,
// - etc
//
// Restriction on music, lyrics most likely will not be created.
//
// **What about translation?**
//
// It is also possible to automatically translate from the original language to
// another you need.
//
// To create a translation, specify the desired language explicitly in
// "subtitles_language" parameter. Otherwise, the subtitles will be in the original
// language. Translation into different languages should be done by creating
// separate tasks.
//
// ![Auto generated subtitles example](https://demo-files.gvideo.io/apidocs/captions.gif)
//
// Use MP4 videos to process. This method is not tied to videos that are stored
// only in our video hosting (look at how get a link to MP4 rendition), so you can
// use links to any other external file with HTTP/HTTPS access.
//
// For now, only the first audio track can be processed; later this functionality
// will be improved to allow to use any. Also, not all language pairs are currently
// supported. If a language pair is not supported for automatic translation, the
// task status will be FAILURE with description of the reason. Example:
// `eng => uzb`. You can request to add the language pair you need for automatic
// translation. Contact our support.
//
// Example of modes to transcribe and/or translate:
//
//   - Auto language detection: `{ "url":"..." }`
//   - From German language explicitly : `{ "url":"...", "audio_language":"ger" }`
//   - From any auto-detected to English language explicitly:
//     `{ "url":"...", "subtitles_language":"eng" }`
//   - From German language to English language explicitly:
//     `{ "url":"...", "audio_language":"ger", "subtitles_language":"eng" }`
//
// Example of setting a task to process MP4 file (animated gif from above):
//
// ```
// curl -L 'https://api.gcore.com/streaming/ai/tasks' \
// -H 'Content-Type: application/json' \
// -H 'Authorization: APIKey 1234$abcd...' \
//
//	-d '{
//	    "url": "https://demo-files.gvideo.io/apidocs/spritefright-blender-cut30sec.mp4"
//	}'
//
// ```
//
// As described above, transcription is done automatically using AI. Therefore, the
// quality may differ from a manual transcription by a professional person. If this
// happens to you, then you can download subtitles and change them in an external
// editor.
//
// Transcription and translation are 2 different AI tasks:
//
// - Transcription is set only for transcription.
// - Translation, if non-original languages are set for translation.
//
// Billing takes into account the duration of the analyzed original video.
//
// The heart for transcribing is the AI model Whisper from OpenAI, with additional
// optimizations and services. The AI models run on our own infrastructure, so the
// files/data are not transferred anywhere to external services. After processing,
// original files are also deleted from local storage of AI.
//
// Read more detailed information about our solution, and architecture, and
// benefits in the knowledge base and blog.
//
// Create AI CM:nsfw task
//
// This algorithm allows to quickly detect inappropriate content, determining that
// the content is NSFW ("Not Safe For Work") or normal. Generic info about all
// capabilities and limits see in the generic
// ["Content Moderation"](/api-reference/streaming/ai/create-ai-task) method.
//
// **What is "Not Safe For Work"?**
//
// The algorithm has recognized inappropriate content in a video and it might not
// be suitable to view in public places. The solution provides its confidence level
// (in percentage) of how sure it is that the content is NSFW, or it most likely
// does not contain any sexual or similar content.
//
// Different to soft-nudity-detection and hard-nudity-detection, this model will
// only check for sensitive material that can be considered not-safe-for-work.
//
// ![AI Content Moderation: NSFW detection visual example](https://demo-files.gvideo.io/apidocs/nsfw-detection.gif)
//
// **How to use?**
//
// Frames within the specified video are analyzed.
//
// Response will contain only frames for which the class nsfw is detected with a
// confidence of more than 50%.
//
// Example of detected NSFW:
//
// ```
//
//	{
//	  "nsfw_detected": true,
//	  "detection_results": [ "nsfw" ],
//	  "frames": [
//	      {
//	          "label": "nsfw",
//	          "confidence": 0.93,
//	          "frame_number": 1
//	      },..
//	  ]
//	}
//
// ```
//
// Example of a response without detecting inappropriate content:
//
// ```
//
//	{
//	  "nsfw_detected": false,
//	  "detection_results": [],
//	  "frames": []
//	}
//
// ```
//
// Please note that the API only provides a set of data (json) about the objects
// found, so no video is generated. The demo video video (above ^) was specially
// created based on json from the API for visual demonstration and better
// perception of the possibilities.
//
// Create AI CM:`hard_nudity` task
//
// This algorithm allows to detect explicit nudity of the human body (involving
// genitals) in a video. Generic info about all capabilities and limits see in the
// generic ["Content Moderation"](/api-reference/streaming/ai/create-ai-task)
// method.
//
// **What is Hard nudity detection?**
//
// This method is often used to analyze UGC to determine whether videos can be
// published to all users, or to prohibit publication due to offensive and
// inappropriate content.
//
// Objects that can be detected:
//
// - `ANUS_EXPOSED`
// - `BUTTOCKS_EXPOSED`
// - `FEMALE_BREAST_EXPOSED`
// - `FEMALE_GENITALIA_EXPOSED`
// - `MALE_BREAST_EXPOSED`
// - `MALE_GENITALIA_EXPOSED`
//
// Please note that the number of objects is less than in the
// soft-nudity-detection. This method works faster and better if only exposed body
// parts detection is required.
//
// ![AI Content Moderation: hard nudity detection visual example](https://demo-files.gvideo.io/apidocs/nudity_detection.gif)
//
// **How to use?**
//
// The information is returned with the video frame number where it was found and
// probability of the detected object. Nudity detection is done using AI, so for
// each object a probability percentage is applied; objects with a probability of
// at least 30% are included in the response.
//
// Video processing speed is approximately 1:5.
//
// Example of detected nudity or body parts:
//
// ```
//
//	{
//	  "nudity_detected": true,
//	  "detection_results": [ "MALE_GENITALIA_EXPOSED" ]
//	  "frames": [
//	      {
//	          "confidence": 0.75,
//	          "frame_number": 35,
//	          "label": "MALE_GENITALIA_EXPOSED"
//	      },...
//	  ]
//	}
//
// ```
//
// Example response when nudity or body parts were not found:
//
// ```
//
//	{
//	  "nudity_detected": false,
//	  "detection_results": []
//	  "frames": []
//	}
//
// ```
//
// There is no universal recipe under which a video can be considered unacceptable,
// since different services host different types of videos for different audiences:
// adult content, children's content, educational content, etc. You can determine
// the probability threshold at which you consider a video inappropriate. The
// easiest option is to run several of your videos and analyze the resulting
// probability coefficient.
//
// Sometimes a detected object at the beginning of the video immediately makes it
// clear that there is no need to further analyze the video. For such cases, you
// can use stop tags. Use parameter "stop_objects" to specify comma separated stop
// tags. It is also possible to specify % probability threshold value, above which
// the stop tag will be triggered.
//
// ```
//
//	{
//	    "url": "...",
//	    "stop_objects": "MALE_GENITALIA_EXPOSED:0.8,FEMALE_GENITALIA_EXPOSED"
//	}
//
// ```
//
// Please note that the API only provides a set of data (json) about the objects
// found, so no video is generated. The demo video video (above ^) was specially
// created based on json from the API for visual demonstration and better
// perception of the possibilities.
//
// Create AI CM:`soft_nudity` task
//
// This algorithm allows to identify explicit nudity and partial nudity too
// (including the presence of male and female faces and other uncovered body parts)
// in a video. Generic info about all capabilities and limits see in the generic
// ["Content Moderation"](/api-reference/streaming/ai/create-ai-task) method.
//
// **What is Soft nudity detection?**
//
// This method is often used to analyze UGC to determine whether videos can be
// published to all users, or to prohibit publication due to offensive and
// inappropriate content.
//
// Objects that can be detected:
//
// - `ANUS_COVERED`
// - `ANUS_EXPOSED`
// - `ARMPITS_COVERED`
// - `ARMPITS_EXPOSED`
// - `BELLY_COVERED`
// - `BELLY_EXPOSED`
// - `BUTTOCKS_COVERED`
// - `BUTTOCKS_EXPOSED`
// - `FACE_FEMALE`
// - `FACE_MALE`
// - `FEET_COVERED`
// - `FEET_EXPOSED`
// - `FEMALE_BREAST_COVERED`
// - `FEMALE_BREAST_EXPOSED`
// - `FEMALE_GENITALIA_COVERED`
// - `FEMALE_GENITALIA_EXPOSED`
// - `MALE_BREAST_EXPOSED`
// - `MALE_GENITALIA_EXPOSED`
//
// This method allows you to identify faces and other body parts. Used to find
// complex combinations of what is happening in a video. Please note that the
// number of objects is more than in the hard-nudity-detection. The method is
// slower.
//
// ![AI Content Moderation: hard nudity detection visual example](https://demo-files.gvideo.io/apidocs/soft_nudity_detection.gif)
//
// **How to use?**
//
// The information is returned with the video frame number where it was found and
// probability of the detected object. Nudity detection is done using AI, so for
// each object a probability percentage is applied; objects with a probability of
// at least 30% are included in the response.
//
// Video processing speed is approximately 1:5.
//
// Example of detected nudity or body parts:
//
// ```
//
//	{
//	  "nudity_detected": true,
//	  "detection_results": [ "FACE_FEMALE", "BELLY_COVERED" ]
//	  "frames": [
//	      {
//	          "confidence": 0.82,
//	          "frame_number": 1,
//	          "label": "BELLY_COVERED"
//	      },...
//	  ]
//	}
//
// ```
//
// Example response when nudity or body parts were not found:
//
// ```
//
//	{
//	  "nudity_detected": false,
//	  "detection_results": []
//	  "frames": []
//	}
//
// ```
//
// There is no universal recipe under which a video can be considered unacceptable,
// since different services host different types of videos for different audiences:
// adult content, children's content, educational content, etc. You can determine
// the probability threshold at which you consider a video inappropriate. The
// easiest option is to run several of your videos and analyze the resulting
// probability coefficient.
//
// Sometimes a detected object at the beginning of the video immediately makes it
// clear that there is no need to further analyze the video. For such cases, you
// can use stop tags. Use parameter "stop_objects" to specify comma separated stop
// tags. It is also possible to specify % probability threshold value, above which
// the stop tag will be triggered.
//
// ```
//
//	{
//	    "url": "...",
//	    "stop_objects": "BELLY_COVERED:0.9,FEMALE_GENITALIA_COVERED"
//	}
//
// ```
//
// Please note that the API only provides a set of data (json) about the objects
// found, so no video is generated. The demo video video (above ^) was specially
// created based on json from the API for visual demonstration and better
// perception of the possibilities.
//
// Create AI CM:sport task
//
// This algorithm allows to identify various sporting activities in a video.
// Generic info about all capabilities and limits see in the generic
// ["Content Moderation"](/api-reference/streaming/ai/create-ai-task) method.
//
// **What is Sports activity detection?**
//
// Sports activity detection by AI involves using machine learning and computer
// vision technologies to automatically identify, analyze, and interpret various
// activities within sports and generic videos. This can include detecting specific
// types, actions, events, and moments.
//
// This model operates on a video sequence (and not on images as most of the used
// computer vision models). Make sure your video is at least 10-15 seconds long.
//
// Sports activities can be detected:
//
// - archery
// - arm wrestling
// - playing badminton
// - playing baseball
// - basketball dunk
// - bowling
// - boxing punch
// - boxing speed bag
// - catching or throwing baseball
// - catching or throwing softball
// - cricket
// - curling
// - disc golfing
// - dodgeball
// - fencing
// - football
// - golf chipping
// - golf driving
// - golf putting
// - hitting baseball
// - hockey stop
// - ice skating
// - javelin throw
// - juggling soccer ball
// - kayaking
// - kicking field goal
// - kicking soccer ball
// - playing cricket
// - playing field hockey
// - playing ice hockey
// - playing kickball
// - playing lacrosse
// - playing ping pong
// - playing polo
// - playing squash or racquetball
// - playing tennis
// - playing volleyball
// - pole vault
// - riding a bike
// - riding or walking with horse
// - roller skating
// - rowing
// - sailing
// - shooting goal (soccer)
// - skateboarding
// - skiing
//
// Use cases:
//
//   - Sports leagues and content creators can use AI to monitor UGC for unauthorized
//     publications of their content. This can include detecting specific sporting
//     events or activities that are part of copyrighted content.
//   - Sports fans often miss live games and rely on highlight reels. AI can
//     automatically detect key moments like goals, touchdowns, or game-winning shots
//     in uploaded UGC videos and compile them into personalized highlight reels.
//
// ![AI Content Moderation: sports activity detection visual example](https://demo-files.gvideo.io/apidocs/sports_football_detection.gif)
//
// **How to use?**
//
// The information is returned with the video frame number where it was found and
// probability of the detected activity. Identification is done using AI, so for
// each activity a probability percentage is applied; activities with a probability
// of at least 30% are included in the response.
//
// Video processing speed is approximately 1:5.
//
// Example of detected sports activity:
//
// ```
//
//	{
//	  "sport_detected": true,
//	  "detection_results": [ "shooting goal (soccer)" ],
//	  "frames": [
//	      {
//	          "label": "shooting goal (soccer)",
//	          "frame_number": 98,
//	          "confidence": 0.99
//	      },...
//	  ]
//	}
//
// ```
//
// Example response when sports activities were not found:
//
// ```
//
//	{
//	  "sport_detected": false,
//	  "detection_results": []
//	  "frames": []
//	}
//
// ```
//
// Please note that the API only provides a set of data (json) about the objects
// found, so no video is generated. The demo video video (above ^) was specially
// created based on json from the API for visual demonstration and better
// perception of the possibilities.
func (r *AITaskService) New(ctx context.Context, body AITaskNewParams, opts ...option.RequestOption) (res *AITaskNewResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/ai/tasks"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, body, &res, opts...)
	return res, err
}

// Returns a list of previously created and processed AI tasks.
//
// The list contains brief information about the task and its execution status.
// Data is displayed page by page.
func (r *AITaskService) List(ctx context.Context, query AITaskListParams, opts ...option.RequestOption) (res *pagination.PageStreamingAI[AITask], err error) {
	var raw *http.Response
	opts = slices.Concat(r.Options, opts)
	opts = append([]option.RequestOption{option.WithResponseInto(&raw)}, opts...)
	path := "streaming/ai/tasks"
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

// Returns a list of previously created and processed AI tasks.
//
// The list contains brief information about the task and its execution status.
// Data is displayed page by page.
func (r *AITaskService) ListAutoPaging(ctx context.Context, query AITaskListParams, opts ...option.RequestOption) *pagination.PageStreamingAIAutoPager[AITask] {
	return pagination.NewPageStreamingAIAutoPager(r.List(ctx, query, opts...))
}

// Stopping a previously launched AI-task without waiting for it to be fully
// completed.
//
// The task will be moved to "REVOKED" status.
func (r *AITaskService) Cancel(ctx context.Context, taskID string, opts ...option.RequestOption) (res *AITaskCancelResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("streaming/ai/tasks/%s/cancel", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodPost, path, nil, &res, opts...)
	return res, err
}

// This is the single method to check the execution status of an AI task, and
// obtain the result of any type of AI task.
//
// Based on the results of processing, the “result” field will contain an answer
// corresponding to the type of the initially created task:
//
// - ASR: Transcribe video
// - ASR: Translate subtitles
// - CM: Sports detection
// - CM: Not Safe For Work (NSFW) content detection
// - CM: Soft nudity detection
// - CM: Hard nudity detection
// - CM: Objects recognition (soon)
// - etc... (see other methods from /ai/ domain)
//
// A queue is used to process videos. The waiting time depends on the total number
// of requests in the system, so sometimes you will have to wait.
//
// Statuses:
//
//   - RECEIVED – the task is accepted by the system
//   - PENDING – the task is received and it is pending for available resources
//   - STARTED – processing has started
//   - SUCCESS – processing has completed successfully
//   - FAILURE – processing failed
//   - REVOKED – processing was cancelled by the user (or the system)
//   - RETRY – the task execution failed due to internal reasons, the task is queued
//     for re-execution (up to 3 times)
//
// Each task is processed in sub-stages, for example, original language is first
// determined in a video, and then transcription is performed. In such cases, the
// video processing status may change from "STARTED" to "PENDING", and back. This
// is due to waiting for resources for a specific processing sub-stage. In this
// case, the overall percentage "progress" of video processing will reflect the
// full picture.
//
// The result data is stored for 1 month, after which it is deleted.
//
// For billing conditions see the corresponding methods in /ai/ domain. The task is
// billed only after successful completion of the task and transition to "SUCCESS"
// status.
func (r *AITaskService) Get(ctx context.Context, taskID string, opts ...option.RequestOption) (res *AITaskGetResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	if taskID == "" {
		err = errors.New("missing required task_id parameter")
		return nil, err
	}
	path := fmt.Sprintf("streaming/ai/tasks/%s", taskID)
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, nil, &res, opts...)
	return res, err
}

// The method for revealing basic information and advanced underlying settings that
// are used when performing AI-tasks.
//
// Parameter sections:
//
//   - "language_support" – AI Translation: check if a language pair is supported or
//     not for AI translation.
//   - this list will expand as new AI methods are added.
//
// **`language_support`**
//
// There are many languages available for transcription. But not all languages can
// be automatically translated to and from with good quality. In order to determine
// the availability of translation from the audio language to the desired subtitle
// language, you can use this type of "language_support".
//
// AI models are constantly improving, so this method can be used for dynamic
// determination.
//
// Example:
//
// ```
// curl -L 'https://api.gcore.com/streaming/ai/info?type=language_support&audio_language=eng&subtitles_language=fre'
//
// { "supported": true }
// ```
//
// Today we provide the following capabilities as below.
//
// These are the 100 languages for which we support only transcription and
// translation to English. The iso639-2b codes for these are:
// `afr, sqi, amh, ara, hye, asm, aze, bak, eus, bel, ben, bos, bre, bul, mya, cat, zho, hrv, ces, dan, nld, eng, est, fao, fin, fra, glg, kat, deu, guj, hat, hau, haw, heb, hin, hun, isl, ind, ita, jpn, jav, kan, kaz, khm, kor, lao, lat, lav, lin, lit, ltz, mkd, mlg, msa, mal, mlt, mri, mar, ell, mon, nep, nor, nno, oci, pan, fas, pol, por, pus, ron, rus, san, srp, sna, snd, sin, slk, slv, som, spa, sun, swa, swe, tgl, tgk, tam, tat, tel, tha, bod, tur, tuk, ukr, urd, uzb, vie, cym, yid, yor`.
//
// These are the 77 languages for which we support translation to other languages
// and translation to:
// `afr, amh, ara, hye, asm, aze, eus, bel, ben, bos, bul, mya, cat, zho, hrv, ces, dan, nld, eng, est, fin, fra, glg, kat, deu, guj, heb, hin, hun, isl, ind, ita, jpn, jav, kan, kaz, khm, kor, lao, lav, lit, mkd, mal, mlt, mar, ell, mon, nep, nno, pan, fas, pol, por, pus, ron, rus, srp, sna, snd, slk, slv, som, spa, swa, swe, tgl, tgk, tam, tel, tha, tur, ukr, urd, vie, cym, yor`.
func (r *AITaskService) GetAISettings(ctx context.Context, query AITaskGetAISettingsParams, opts ...option.RequestOption) (res *AITaskGetAISettingsResponse, err error) {
	opts = slices.Concat(r.Options, opts)
	path := "streaming/ai/info"
	err = requestconfig.ExecuteNewRequest(ctx, http.MethodGet, path, query, &res, opts...)
	return res, err
}

type AIContentmoderationHardnudity struct {
	// AI content moderation with "hard_nudity" algorithm
	//
	// Any of "hard_nudity".
	Category AIContentmoderationHardnudityCategory `json:"category" api:"required"`
	// Name of the task to be performed
	//
	// Any of "content-moderation".
	TaskName AIContentmoderationHardnudityTaskName `json:"task_name" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (nudity detection, etc), then the ID of the
	// associated video for which the task was performed will be explicitly indicated
	// here.
	ClientEntityData string `json:"client_entity_data"`
	// Client ID associated with the task.
	ClientID int64 `json:"client_id"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID string `json:"client_user_id"`
	// Comma separated objects, and probabilities, that will cause the processing to
	// stop immediately after finding.
	//
	// Any of "ANUS_EXPOSED", "BUTTOCKS_EXPOSED", "FEMALE_BREAST_EXPOSED",
	// "FEMALE_GENITALIA_EXPOSED", "MALE_BREAST_EXPOSED", "MALE_GENITALIA_EXPOSED".
	StopObjects AIContentmoderationHardnudityStopObjects `json:"stop_objects"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category         respjson.Field
		TaskName         respjson.Field
		URL              respjson.Field
		ClientEntityData respjson.Field
		ClientID         respjson.Field
		ClientUserID     respjson.Field
		StopObjects      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIContentmoderationHardnudity) RawJSON() string { return r.JSON.raw }
func (r *AIContentmoderationHardnudity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIContentmoderationHardnudity to a
// AIContentmoderationHardnudityParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIContentmoderationHardnudityParam.Overrides()
func (r AIContentmoderationHardnudity) ToParam() AIContentmoderationHardnudityParam {
	return param.Override[AIContentmoderationHardnudityParam](json.RawMessage(r.RawJSON()))
}

// AI content moderation with "hard_nudity" algorithm
type AIContentmoderationHardnudityCategory string

const (
	AIContentmoderationHardnudityCategoryHardNudity AIContentmoderationHardnudityCategory = "hard_nudity"
)

// Name of the task to be performed
type AIContentmoderationHardnudityTaskName string

const (
	AIContentmoderationHardnudityTaskNameContentModeration AIContentmoderationHardnudityTaskName = "content-moderation"
)

// Comma separated objects, and probabilities, that will cause the processing to
// stop immediately after finding.
type AIContentmoderationHardnudityStopObjects string

const (
	AIContentmoderationHardnudityStopObjectsAnusExposed            AIContentmoderationHardnudityStopObjects = "ANUS_EXPOSED"
	AIContentmoderationHardnudityStopObjectsButtocksExposed        AIContentmoderationHardnudityStopObjects = "BUTTOCKS_EXPOSED"
	AIContentmoderationHardnudityStopObjectsFemaleBreastExposed    AIContentmoderationHardnudityStopObjects = "FEMALE_BREAST_EXPOSED"
	AIContentmoderationHardnudityStopObjectsFemaleGenitaliaExposed AIContentmoderationHardnudityStopObjects = "FEMALE_GENITALIA_EXPOSED"
	AIContentmoderationHardnudityStopObjectsMaleBreastExposed      AIContentmoderationHardnudityStopObjects = "MALE_BREAST_EXPOSED"
	AIContentmoderationHardnudityStopObjectsMaleGenitaliaExposed   AIContentmoderationHardnudityStopObjects = "MALE_GENITALIA_EXPOSED"
)

// The properties Category, TaskName, URL are required.
type AIContentmoderationHardnudityParam struct {
	// AI content moderation with "hard_nudity" algorithm
	//
	// Any of "hard_nudity".
	Category AIContentmoderationHardnudityCategory `json:"category,omitzero" api:"required"`
	// Name of the task to be performed
	//
	// Any of "content-moderation".
	TaskName AIContentmoderationHardnudityTaskName `json:"task_name,omitzero" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (nudity detection, etc), then the ID of the
	// associated video for which the task was performed will be explicitly indicated
	// here.
	ClientEntityData param.Opt[string] `json:"client_entity_data,omitzero"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID param.Opt[string] `json:"client_user_id,omitzero"`
	// Comma separated objects, and probabilities, that will cause the processing to
	// stop immediately after finding.
	//
	// Any of "ANUS_EXPOSED", "BUTTOCKS_EXPOSED", "FEMALE_BREAST_EXPOSED",
	// "FEMALE_GENITALIA_EXPOSED", "MALE_BREAST_EXPOSED", "MALE_GENITALIA_EXPOSED".
	StopObjects AIContentmoderationHardnudityStopObjects `json:"stop_objects,omitzero"`
	paramObj
}

func (r AIContentmoderationHardnudityParam) MarshalJSON() (data []byte, err error) {
	type shadow AIContentmoderationHardnudityParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIContentmoderationHardnudityParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIContentmoderationNsfw struct {
	// AI content moderation with NSFW detection algorithm
	//
	// Any of "nsfw".
	Category AIContentmoderationNsfwCategory `json:"category" api:"required"`
	// Name of the task to be performed
	//
	// Any of "content-moderation".
	TaskName AIContentmoderationNsfwTaskName `json:"task_name" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (nudity detection, etc), then the ID of the
	// associated video for which the task was performed will be explicitly indicated
	// here.
	ClientEntityData string `json:"client_entity_data"`
	// Client ID associated with the task.
	ClientID int64 `json:"client_id"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID string `json:"client_user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category         respjson.Field
		TaskName         respjson.Field
		URL              respjson.Field
		ClientEntityData respjson.Field
		ClientID         respjson.Field
		ClientUserID     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIContentmoderationNsfw) RawJSON() string { return r.JSON.raw }
func (r *AIContentmoderationNsfw) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIContentmoderationNsfw to a AIContentmoderationNsfwParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIContentmoderationNsfwParam.Overrides()
func (r AIContentmoderationNsfw) ToParam() AIContentmoderationNsfwParam {
	return param.Override[AIContentmoderationNsfwParam](json.RawMessage(r.RawJSON()))
}

// AI content moderation with NSFW detection algorithm
type AIContentmoderationNsfwCategory string

const (
	AIContentmoderationNsfwCategoryNsfw AIContentmoderationNsfwCategory = "nsfw"
)

// Name of the task to be performed
type AIContentmoderationNsfwTaskName string

const (
	AIContentmoderationNsfwTaskNameContentModeration AIContentmoderationNsfwTaskName = "content-moderation"
)

// The properties Category, TaskName, URL are required.
type AIContentmoderationNsfwParam struct {
	// AI content moderation with NSFW detection algorithm
	//
	// Any of "nsfw".
	Category AIContentmoderationNsfwCategory `json:"category,omitzero" api:"required"`
	// Name of the task to be performed
	//
	// Any of "content-moderation".
	TaskName AIContentmoderationNsfwTaskName `json:"task_name,omitzero" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (nudity detection, etc), then the ID of the
	// associated video for which the task was performed will be explicitly indicated
	// here.
	ClientEntityData param.Opt[string] `json:"client_entity_data,omitzero"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID param.Opt[string] `json:"client_user_id,omitzero"`
	paramObj
}

func (r AIContentmoderationNsfwParam) MarshalJSON() (data []byte, err error) {
	type shadow AIContentmoderationNsfwParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIContentmoderationNsfwParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIContentmoderationSoftnudity struct {
	// AI content moderation with "soft_nudity" algorithm
	//
	// Any of "soft_nudity".
	Category AIContentmoderationSoftnudityCategory `json:"category" api:"required"`
	// Name of the task to be performed
	//
	// Any of "content-moderation".
	TaskName AIContentmoderationSoftnudityTaskName `json:"task_name" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (nudity detection, etc), then the ID of the
	// associated video for which the task was performed will be explicitly indicated
	// here.
	ClientEntityData string `json:"client_entity_data"`
	// Client ID associated with the task.
	ClientID int64 `json:"client_id"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID string `json:"client_user_id"`
	// Comma separated objects, and probabilities, that will cause the processing to
	// stop immediately after finding.
	//
	// Any of "ANUS_COVERED", "ANUS_EXPOSED", "ARMPITS_COVERED", "ARMPITS_EXPOSED",
	// "BELLY_COVERED", "BELLY_EXPOSED", "BUTTOCKS_COVERED", "BUTTOCKS_EXPOSED",
	// "FACE_FEMALE", "FACE_MALE", "FEET_COVERED", "FEET_EXPOSED",
	// "FEMALE_BREAST_COVERED", "FEMALE_BREAST_EXPOSED", "FEMALE_GENITALIA_COVERED",
	// "FEMALE_GENITALIA_EXPOSED", "MALE_BREAST_EXPOSED", "MALE_GENITALIA_EXPOSED".
	StopObjects AIContentmoderationSoftnudityStopObjects `json:"stop_objects"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category         respjson.Field
		TaskName         respjson.Field
		URL              respjson.Field
		ClientEntityData respjson.Field
		ClientID         respjson.Field
		ClientUserID     respjson.Field
		StopObjects      respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIContentmoderationSoftnudity) RawJSON() string { return r.JSON.raw }
func (r *AIContentmoderationSoftnudity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIContentmoderationSoftnudity to a
// AIContentmoderationSoftnudityParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIContentmoderationSoftnudityParam.Overrides()
func (r AIContentmoderationSoftnudity) ToParam() AIContentmoderationSoftnudityParam {
	return param.Override[AIContentmoderationSoftnudityParam](json.RawMessage(r.RawJSON()))
}

// AI content moderation with "soft_nudity" algorithm
type AIContentmoderationSoftnudityCategory string

const (
	AIContentmoderationSoftnudityCategorySoftNudity AIContentmoderationSoftnudityCategory = "soft_nudity"
)

// Name of the task to be performed
type AIContentmoderationSoftnudityTaskName string

const (
	AIContentmoderationSoftnudityTaskNameContentModeration AIContentmoderationSoftnudityTaskName = "content-moderation"
)

// Comma separated objects, and probabilities, that will cause the processing to
// stop immediately after finding.
type AIContentmoderationSoftnudityStopObjects string

const (
	AIContentmoderationSoftnudityStopObjectsAnusCovered            AIContentmoderationSoftnudityStopObjects = "ANUS_COVERED"
	AIContentmoderationSoftnudityStopObjectsAnusExposed            AIContentmoderationSoftnudityStopObjects = "ANUS_EXPOSED"
	AIContentmoderationSoftnudityStopObjectsArmpitsCovered         AIContentmoderationSoftnudityStopObjects = "ARMPITS_COVERED"
	AIContentmoderationSoftnudityStopObjectsArmpitsExposed         AIContentmoderationSoftnudityStopObjects = "ARMPITS_EXPOSED"
	AIContentmoderationSoftnudityStopObjectsBellyCovered           AIContentmoderationSoftnudityStopObjects = "BELLY_COVERED"
	AIContentmoderationSoftnudityStopObjectsBellyExposed           AIContentmoderationSoftnudityStopObjects = "BELLY_EXPOSED"
	AIContentmoderationSoftnudityStopObjectsButtocksCovered        AIContentmoderationSoftnudityStopObjects = "BUTTOCKS_COVERED"
	AIContentmoderationSoftnudityStopObjectsButtocksExposed        AIContentmoderationSoftnudityStopObjects = "BUTTOCKS_EXPOSED"
	AIContentmoderationSoftnudityStopObjectsFaceFemale             AIContentmoderationSoftnudityStopObjects = "FACE_FEMALE"
	AIContentmoderationSoftnudityStopObjectsFaceMale               AIContentmoderationSoftnudityStopObjects = "FACE_MALE"
	AIContentmoderationSoftnudityStopObjectsFeetCovered            AIContentmoderationSoftnudityStopObjects = "FEET_COVERED"
	AIContentmoderationSoftnudityStopObjectsFeetExposed            AIContentmoderationSoftnudityStopObjects = "FEET_EXPOSED"
	AIContentmoderationSoftnudityStopObjectsFemaleBreastCovered    AIContentmoderationSoftnudityStopObjects = "FEMALE_BREAST_COVERED"
	AIContentmoderationSoftnudityStopObjectsFemaleBreastExposed    AIContentmoderationSoftnudityStopObjects = "FEMALE_BREAST_EXPOSED"
	AIContentmoderationSoftnudityStopObjectsFemaleGenitaliaCovered AIContentmoderationSoftnudityStopObjects = "FEMALE_GENITALIA_COVERED"
	AIContentmoderationSoftnudityStopObjectsFemaleGenitaliaExposed AIContentmoderationSoftnudityStopObjects = "FEMALE_GENITALIA_EXPOSED"
	AIContentmoderationSoftnudityStopObjectsMaleBreastExposed      AIContentmoderationSoftnudityStopObjects = "MALE_BREAST_EXPOSED"
	AIContentmoderationSoftnudityStopObjectsMaleGenitaliaExposed   AIContentmoderationSoftnudityStopObjects = "MALE_GENITALIA_EXPOSED"
)

// The properties Category, TaskName, URL are required.
type AIContentmoderationSoftnudityParam struct {
	// AI content moderation with "soft_nudity" algorithm
	//
	// Any of "soft_nudity".
	Category AIContentmoderationSoftnudityCategory `json:"category,omitzero" api:"required"`
	// Name of the task to be performed
	//
	// Any of "content-moderation".
	TaskName AIContentmoderationSoftnudityTaskName `json:"task_name,omitzero" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (nudity detection, etc), then the ID of the
	// associated video for which the task was performed will be explicitly indicated
	// here.
	ClientEntityData param.Opt[string] `json:"client_entity_data,omitzero"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID param.Opt[string] `json:"client_user_id,omitzero"`
	// Comma separated objects, and probabilities, that will cause the processing to
	// stop immediately after finding.
	//
	// Any of "ANUS_COVERED", "ANUS_EXPOSED", "ARMPITS_COVERED", "ARMPITS_EXPOSED",
	// "BELLY_COVERED", "BELLY_EXPOSED", "BUTTOCKS_COVERED", "BUTTOCKS_EXPOSED",
	// "FACE_FEMALE", "FACE_MALE", "FEET_COVERED", "FEET_EXPOSED",
	// "FEMALE_BREAST_COVERED", "FEMALE_BREAST_EXPOSED", "FEMALE_GENITALIA_COVERED",
	// "FEMALE_GENITALIA_EXPOSED", "MALE_BREAST_EXPOSED", "MALE_GENITALIA_EXPOSED".
	StopObjects AIContentmoderationSoftnudityStopObjects `json:"stop_objects,omitzero"`
	paramObj
}

func (r AIContentmoderationSoftnudityParam) MarshalJSON() (data []byte, err error) {
	type shadow AIContentmoderationSoftnudityParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIContentmoderationSoftnudityParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AIContentmoderationSport struct {
	// AI content moderation with types of sports activity detection
	//
	// Any of "sport".
	Category AIContentmoderationSportCategory `json:"category" api:"required"`
	// Name of the task to be performed
	//
	// Any of "content-moderation".
	TaskName AIContentmoderationSportTaskName `json:"task_name" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (nudity detection, etc), then the ID of the
	// associated video for which the task was performed will be explicitly indicated
	// here.
	ClientEntityData string `json:"client_entity_data"`
	// Client ID associated with the task.
	ClientID int64 `json:"client_id"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID string `json:"client_user_id"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Category         respjson.Field
		TaskName         respjson.Field
		URL              respjson.Field
		ClientEntityData respjson.Field
		ClientID         respjson.Field
		ClientUserID     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AIContentmoderationSport) RawJSON() string { return r.JSON.raw }
func (r *AIContentmoderationSport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// ToParam converts this AIContentmoderationSport to a
// AIContentmoderationSportParam.
//
// Warning: the fields of the param type will not be present. ToParam should only
// be used at the last possible moment before sending a request. Test for this with
// AIContentmoderationSportParam.Overrides()
func (r AIContentmoderationSport) ToParam() AIContentmoderationSportParam {
	return param.Override[AIContentmoderationSportParam](json.RawMessage(r.RawJSON()))
}

// AI content moderation with types of sports activity detection
type AIContentmoderationSportCategory string

const (
	AIContentmoderationSportCategorySport AIContentmoderationSportCategory = "sport"
)

// Name of the task to be performed
type AIContentmoderationSportTaskName string

const (
	AIContentmoderationSportTaskNameContentModeration AIContentmoderationSportTaskName = "content-moderation"
)

// The properties Category, TaskName, URL are required.
type AIContentmoderationSportParam struct {
	// AI content moderation with types of sports activity detection
	//
	// Any of "sport".
	Category AIContentmoderationSportCategory `json:"category,omitzero" api:"required"`
	// Name of the task to be performed
	//
	// Any of "content-moderation".
	TaskName AIContentmoderationSportTaskName `json:"task_name,omitzero" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (nudity detection, etc), then the ID of the
	// associated video for which the task was performed will be explicitly indicated
	// here.
	ClientEntityData param.Opt[string] `json:"client_entity_data,omitzero"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID param.Opt[string] `json:"client_user_id,omitzero"`
	paramObj
}

func (r AIContentmoderationSportParam) MarshalJSON() (data []byte, err error) {
	type shadow AIContentmoderationSportParam
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AIContentmoderationSportParam) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITask struct {
	ProcessingTime AITaskProcessingTime `json:"processing_time" api:"required"`
	// Percentage of task completed. A value greater than 0 means that it has been
	// taken into operation and is being processed.
	Progress int64 `json:"progress" api:"required"`
	// Status of processing the AI task. See GET /ai/tasks/{`task_id`} method for
	// description.
	//
	// Any of "PENDING", "STARTED", "SUCCESS", "FAILURE", "RECEIVED", "REVOKED",
	// "RETRY".
	Status AITaskStatus `json:"status" api:"required"`
	// The object will correspond to the task type that was specified in the original
	// request. There will be one object for transcription, another for searching for
	// nudity, and so on.
	TaskData AITaskTaskDataUnion `json:"task_data" api:"required"`
	// ID of the AI task
	TaskID string `json:"task_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		ProcessingTime respjson.Field
		Progress       respjson.Field
		Status         respjson.Field
		TaskData       respjson.Field
		TaskID         respjson.Field
		ExtraFields    map[string]respjson.Field
		raw            string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITask) RawJSON() string { return r.JSON.raw }
func (r *AITask) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskProcessingTime struct {
	// Video processing end time. Format is date time in ISO 8601
	CompletedAt time.Time `json:"completed_at" api:"nullable" format:"date-time"`
	// Video processing start time. Format is date time in ISO 8601
	StartedAt time.Time `json:"started_at" format:"date-time"`
	// Duration of video processing in seconds
	TotalTimeSec float64 `json:"total_time_sec" api:"nullable" format:"decimal"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		CompletedAt  respjson.Field
		StartedAt    respjson.Field
		TotalTimeSec respjson.Field
		ExtraFields  map[string]respjson.Field
		raw          string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskProcessingTime) RawJSON() string { return r.JSON.raw }
func (r *AITaskProcessingTime) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// Status of processing the AI task. See GET /ai/tasks/{`task_id`} method for
// description.
type AITaskStatus string

const (
	AITaskStatusPending  AITaskStatus = "PENDING"
	AITaskStatusStarted  AITaskStatus = "STARTED"
	AITaskStatusSuccess  AITaskStatus = "SUCCESS"
	AITaskStatusFailure  AITaskStatus = "FAILURE"
	AITaskStatusReceived AITaskStatus = "RECEIVED"
	AITaskStatusRevoked  AITaskStatus = "REVOKED"
	AITaskStatusRetry    AITaskStatus = "RETRY"
)

// AITaskTaskDataUnion contains all possible properties and values from
// [AITaskTaskDataAITranscriptionTaskData], [AIContentmoderationNsfw],
// [AIContentmoderationHardnudity], [AIContentmoderationSoftnudity],
// [AIContentmoderationSport].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AITaskTaskDataUnion struct {
	TaskName string `json:"task_name"`
	URL      string `json:"url"`
	// This field is from variant [AITaskTaskDataAITranscriptionTaskData].
	AudioLanguage    string `json:"audio_language"`
	ClientEntityData string `json:"client_entity_data"`
	ClientID         int64  `json:"client_id"`
	ClientUserID     string `json:"client_user_id"`
	// This field is from variant [AITaskTaskDataAITranscriptionTaskData].
	SubtitlesLanguage string `json:"subtitles_language"`
	Category          string `json:"category"`
	StopObjects       string `json:"stop_objects"`
	JSON              struct {
		TaskName          respjson.Field
		URL               respjson.Field
		AudioLanguage     respjson.Field
		ClientEntityData  respjson.Field
		ClientID          respjson.Field
		ClientUserID      respjson.Field
		SubtitlesLanguage respjson.Field
		Category          respjson.Field
		StopObjects       respjson.Field
		raw               string
	} `json:"-"`
}

func (u AITaskTaskDataUnion) AsAITranscriptionTaskData() (v AITaskTaskDataAITranscriptionTaskData) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskTaskDataUnion) AsAIContentModerationNsfwTaskData() (v AIContentmoderationNsfw) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskTaskDataUnion) AsAIContentModerationHardNudityTaskData() (v AIContentmoderationHardnudity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskTaskDataUnion) AsAIContentModerationSoftNudityTaskData() (v AIContentmoderationSoftnudity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskTaskDataUnion) AsAIContentModerationSportTaskData() (v AIContentmoderationSport) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITaskTaskDataUnion) RawJSON() string { return u.JSON.raw }

func (r *AITaskTaskDataUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskTaskDataAITranscriptionTaskData struct {
	// Name of the task to be performed
	//
	// Any of "transcription".
	TaskName string `json:"task_name" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Language in original audio (transcription only). This value is used to determine
	// the language from which to transcribe.
	//
	// If this is not set, the system will run auto language identification and the
	// subtitles will be in the detected language. The method also works based on AI
	// analysis. It's fairly accurate, but if it's wrong, then set the language
	// explicitly.
	//
	// Additionally, when this is not set, we also support recognition of alternate
	// languages in the video (language code-switching).
	//
	// Language is set by 3-letter language code according to ISO-639-2 (bibliographic
	// code).
	//
	// We can process languages:
	//
	// - 'afr': Afrikaans
	// - 'alb': Albanian
	// - 'amh': Amharic
	// - 'ara': Arabic
	// - 'arm': Armenian
	// - 'asm': Assamese
	// - 'aze': Azerbaijani
	// - 'bak': Bashkir
	// - 'baq': Basque
	// - 'bel': Belarusian
	// - 'ben': Bengali
	// - 'bos': Bosnian
	// - 'bre': Breton
	// - 'bul': Bulgarian
	// - 'bur': Myanmar
	// - 'cat': Catalan
	// - 'chi': Chinese
	// - 'cze': Czech
	// - 'dan': Danish
	// - 'dut': Nynorsk
	// - 'eng': English
	// - 'est': Estonian
	// - 'fao': Faroese
	// - 'fin': Finnish
	// - 'fre': French
	// - 'geo': Georgian
	// - 'ger': German
	// - 'glg': Galician
	// - 'gre': Greek
	// - 'guj': Gujarati
	// - 'hat': Haitian creole
	// - 'hau': Hausa
	// - 'haw': Hawaiian
	// - 'heb': Hebrew
	// - 'hin': Hindi
	// - 'hrv': Croatian
	// - 'hun': Hungarian
	// - 'ice': Icelandic
	// - 'ind': Indonesian
	// - 'ita': Italian
	// - 'jav': Javanese
	// - 'jpn': Japanese
	// - 'kan': Kannada
	// - 'kaz': Kazakh
	// - 'khm': Khmer
	// - 'kor': Korean
	// - 'lao': Lao
	// - 'lat': Latin
	// - 'lav': Latvian
	// - 'lin': Lingala
	// - 'lit': Lithuanian
	// - 'ltz': Luxembourgish
	// - 'mac': Macedonian
	// - 'mal': Malayalam
	// - 'mao': Maori
	// - 'mar': Marathi
	// - 'may': Malay
	// - 'mlg': Malagasy
	// - 'mlt': Maltese
	// - 'mon': Mongolian
	// - 'nep': Nepali
	// - 'dut': Dutch
	// - 'nor': Norwegian
	// - 'oci': Occitan
	// - 'pan': Punjabi
	// - 'per': Persian
	// - 'pol': Polish
	// - 'por': Portuguese
	// - 'pus': Pashto
	// - 'rum': Romanian
	// - 'rus': Russian
	// - 'san': Sanskrit
	// - 'sin': Sinhala
	// - 'slo': Slovak
	// - 'slv': Slovenian
	// - 'sna': Shona
	// - 'snd': Sindhi
	// - 'som': Somali
	// - 'spa': Spanish
	// - 'srp': Serbian
	// - 'sun': Sundanese
	// - 'swa': Swahili
	// - 'swe': Swedish
	// - 'tam': Tamil
	// - 'tat': Tatar
	// - 'tel': Telugu
	// - 'tgk': Tajik
	// - 'tgl': Tagalog
	// - 'tha': Thai
	// - 'tib': Tibetan
	// - 'tuk': Turkmen
	// - 'tur': Turkish
	// - 'ukr': Ukrainian
	// - 'urd': Urdu
	// - 'uzb': Uzbek
	// - 'vie': Vietnamese
	// - 'wel': Welsh
	// - 'yid': Yiddish
	// - 'yor': Yoruba
	AudioLanguage string `json:"audio_language"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (transcribing, translation), then the ID of
	// the associated video for which the task was performed will be explicitly
	// indicated here.
	ClientEntityData string `json:"client_entity_data"`
	// Client ID associated with the task.
	ClientID int64 `json:"client_id"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID string `json:"client_user_id"`
	// Indicates which language it is clearly necessary to translate into. If this is
	// not set, the original language will be used from attribute "audio_language".
	//
	// Please note that:
	//
	//   - transcription into the original language is a free procedure,
	//   - and translation from the original language into any other languages is a
	//     "translation" procedure and is paid. More details in
	//     [POST /streaming/ai/tasks](/api-reference/streaming/ai/create-ai-task).
	//     Language is set by 3-letter language code according to ISO-639-2
	//     (bibliographic code).
	SubtitlesLanguage string `json:"subtitles_language"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaskName          respjson.Field
		URL               respjson.Field
		AudioLanguage     respjson.Field
		ClientEntityData  respjson.Field
		ClientID          respjson.Field
		ClientUserID      respjson.Field
		SubtitlesLanguage respjson.Field
		ExtraFields       map[string]respjson.Field
		raw               string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskTaskDataAITranscriptionTaskData) RawJSON() string { return r.JSON.raw }
func (r *AITaskTaskDataAITranscriptionTaskData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskNewResponse struct {
	// ID of the created AI task, from which you can get the execution result
	TaskID string `json:"task_id" api:"required" format:"uuid"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		TaskID      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskNewResponse) RawJSON() string { return r.JSON.raw }
func (r *AITaskNewResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskCancelResponse struct {
	// A textual explicit description of the result of the operation
	Result string `json:"result"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskCancelResponse) RawJSON() string { return r.JSON.raw }
func (r *AITaskCancelResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponse struct {
	Result AITaskGetResponseResultUnion `json:"result" api:"nullable"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Result      respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
	AITask
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponse) RawJSON() string { return r.JSON.raw }
func (r *AITaskGetResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITaskGetResponseResultUnion contains all possible properties and values from
// [AITaskGetResponseResultAIResultsTranscribe],
// [AITaskGetResponseResultAIResultsContentmoderationSport],
// [AITaskGetResponseResultAIResultsContentmoderationNsfw],
// [AITaskGetResponseResultAIResultsContentmoderationHardnudity],
// [AITaskGetResponseResultAIResultsContentmoderationSoftnudity],
// [AITaskGetResponseResultAIResultsFailure].
//
// Use the methods beginning with 'As' to cast the union to one of its variants.
type AITaskGetResponseResultUnion struct {
	// This field is from variant [AITaskGetResponseResultAIResultsTranscribe].
	SpeechDetected bool `json:"speech_detected"`
	// This field is from variant [AITaskGetResponseResultAIResultsTranscribe].
	ConcatenatedText string `json:"concatenated_text"`
	// This field is from variant [AITaskGetResponseResultAIResultsTranscribe].
	Languages []string `json:"languages"`
	// This field is from variant [AITaskGetResponseResultAIResultsTranscribe].
	Subtitles []AITaskGetResponseResultAIResultsTranscribeSubtitle `json:"subtitles"`
	// This field is from variant [AITaskGetResponseResultAIResultsTranscribe].
	VttContent       string   `json:"vttContent"`
	DetectionResults []string `json:"detection_results"`
	// This field is a union of
	// [[]AITaskGetResponseResultAIResultsContentmoderationSportFrame],
	// [[]AITaskGetResponseResultAIResultsContentmoderationNsfwFrame],
	// [[]AITaskGetResponseResultAIResultsContentmoderationHardnudityFrame],
	// [[]AITaskGetResponseResultAIResultsContentmoderationSoftnudityFrame]
	Frames AITaskGetResponseResultUnionFrames `json:"frames"`
	// This field is from variant
	// [AITaskGetResponseResultAIResultsContentmoderationSport].
	SportDetected bool `json:"sport_detected"`
	// This field is from variant
	// [AITaskGetResponseResultAIResultsContentmoderationNsfw].
	NsfwDetected   bool `json:"nsfw_detected"`
	NudityDetected bool `json:"nudity_detected"`
	// This field is from variant [AITaskGetResponseResultAIResultsFailure].
	Error string `json:"error"`
	JSON  struct {
		SpeechDetected   respjson.Field
		ConcatenatedText respjson.Field
		Languages        respjson.Field
		Subtitles        respjson.Field
		VttContent       respjson.Field
		DetectionResults respjson.Field
		Frames           respjson.Field
		SportDetected    respjson.Field
		NsfwDetected     respjson.Field
		NudityDetected   respjson.Field
		Error            respjson.Field
		raw              string
	} `json:"-"`
}

func (u AITaskGetResponseResultUnion) AsAITaskGetResponseResultAIResultsTranscribe() (v AITaskGetResponseResultAIResultsTranscribe) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskGetResponseResultUnion) AsAITaskGetResponseResultAIResultsContentmoderationSport() (v AITaskGetResponseResultAIResultsContentmoderationSport) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskGetResponseResultUnion) AsAITaskGetResponseResultAIResultsContentmoderationNsfw() (v AITaskGetResponseResultAIResultsContentmoderationNsfw) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskGetResponseResultUnion) AsAITaskGetResponseResultAIResultsContentmoderationHardnudity() (v AITaskGetResponseResultAIResultsContentmoderationHardnudity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskGetResponseResultUnion) AsAITaskGetResponseResultAIResultsContentmoderationSoftnudity() (v AITaskGetResponseResultAIResultsContentmoderationSoftnudity) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

func (u AITaskGetResponseResultUnion) AsAITaskGetResponseResultAIResultsFailure() (v AITaskGetResponseResultAIResultsFailure) {
	apijson.UnmarshalRoot(json.RawMessage(u.JSON.raw), &v)
	return
}

// Returns the unmodified JSON received from the API
func (u AITaskGetResponseResultUnion) RawJSON() string { return u.JSON.raw }

func (r *AITaskGetResponseResultUnion) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// AITaskGetResponseResultUnionFrames is an implicit subunion of
// [AITaskGetResponseResultUnion]. AITaskGetResponseResultUnionFrames provides
// convenient access to the sub-properties of the union.
//
// For type safety it is recommended to directly use a variant of the
// [AITaskGetResponseResultUnion].
//
// If the underlying value is not a json object, one of the following properties
// will be valid: OfAITaskGetResponseResultAIResultsContentmoderationSportFrames
// OfAITaskGetResponseResultAIResultsContentmoderationNsfwFrames
// OfAITaskGetResponseResultAIResultsContentmoderationHardnudityFrames
// OfAITaskGetResponseResultAIResultsContentmoderationSoftnudityFrames]
type AITaskGetResponseResultUnionFrames struct {
	// This field will be present if the value is a
	// [[]AITaskGetResponseResultAIResultsContentmoderationSportFrame] instead of an
	// object.
	OfAITaskGetResponseResultAIResultsContentmoderationSportFrames []AITaskGetResponseResultAIResultsContentmoderationSportFrame `json:",inline"`
	// This field will be present if the value is a
	// [[]AITaskGetResponseResultAIResultsContentmoderationNsfwFrame] instead of an
	// object.
	OfAITaskGetResponseResultAIResultsContentmoderationNsfwFrames []AITaskGetResponseResultAIResultsContentmoderationNsfwFrame `json:",inline"`
	// This field will be present if the value is a
	// [[]AITaskGetResponseResultAIResultsContentmoderationHardnudityFrame] instead of
	// an object.
	OfAITaskGetResponseResultAIResultsContentmoderationHardnudityFrames []AITaskGetResponseResultAIResultsContentmoderationHardnudityFrame `json:",inline"`
	// This field will be present if the value is a
	// [[]AITaskGetResponseResultAIResultsContentmoderationSoftnudityFrame] instead of
	// an object.
	OfAITaskGetResponseResultAIResultsContentmoderationSoftnudityFrames []AITaskGetResponseResultAIResultsContentmoderationSoftnudityFrame `json:",inline"`
	JSON                                                                struct {
		OfAITaskGetResponseResultAIResultsContentmoderationSportFrames      respjson.Field
		OfAITaskGetResponseResultAIResultsContentmoderationNsfwFrames       respjson.Field
		OfAITaskGetResponseResultAIResultsContentmoderationHardnudityFrames respjson.Field
		OfAITaskGetResponseResultAIResultsContentmoderationSoftnudityFrames respjson.Field
		raw                                                                 string
	} `json:"-"`
}

func (r *AITaskGetResponseResultUnionFrames) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsTranscribe struct {
	// Determines whether speech was detected or not.
	//
	// Please note: If the task is in "SUCCESS" status and speech was not found in the
	// entire file, then "false" will be indicated here and the `subtitles` field will
	// be empty.
	SpeechDetected bool `json:"speech_detected" api:"required"`
	// Full text of the analyzed video. The value is unstructured, unformatted text.
	ConcatenatedText string `json:"concatenated_text"`
	// An array of language codes that were discovered and/or used in transcription. If
	// the audio or subtitle language was explicitly specified in the initial
	// parameters, it will be copied here. For automatic detection the identified
	// languages will be displayed here. Also please note that for multilingual audio,
	// the first 5 languages are displayed in order of frequency of use.
	Languages []string `json:"languages"`
	// An array of phrases divided into time intervals, in the format "json". Suitable
	// when you need to display the result in chronometric form, or transfer the text
	// for further processing.
	Subtitles []AITaskGetResponseResultAIResultsTranscribeSubtitle `json:"subtitles"`
	// Auto generated subtitles in WebVTT format.
	VttContent string `json:"vttContent"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		SpeechDetected   respjson.Field
		ConcatenatedText respjson.Field
		Languages        respjson.Field
		Subtitles        respjson.Field
		VttContent       respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsTranscribe) RawJSON() string { return r.JSON.raw }
func (r *AITaskGetResponseResultAIResultsTranscribe) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsTranscribeSubtitle struct {
	// End time of the phrase, when it ends in the video. Format is "HH:mm:ss.fff".
	EndTime string `json:"end_time"`
	// Start time of the phrase when it is heard in the video. Format is
	// "HH:mm:ss.fff".
	StartTime string `json:"start_time"`
	// A complete phrase that sounds during a specified period of time.
	Text string `json:"text"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		EndTime     respjson.Field
		StartTime   respjson.Field
		Text        respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsTranscribeSubtitle) RawJSON() string { return r.JSON.raw }
func (r *AITaskGetResponseResultAIResultsTranscribeSubtitle) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsContentmoderationSport struct {
	// Any of "archery", "arm wrestling", "playing badminton", "playing baseball",
	// "basketball dunk", "bowling", "boxing punch", "boxing speed bag", "catching or
	// throwing baseball", "catching or throwing softball", "cricket", "curling", "disc
	// golfing", "dodgeball", "fencing", "football", "golf chipping", "golf driving",
	// "golf putting", "hitting baseball", "hockey stop", "ice skating", "javelin
	// throw", "juggling soccer ball", "kayaking", "kicking field goal", "kicking
	// soccer ball", "playing cricket", "playing field hockey", "playing ice hockey",
	// "playing kickball", "playing lacrosse", "playing ping pong", "playing polo",
	// "playing squash or racquetball", "playing tennis", "playing volleyball", "pole
	// vault", "riding a bike", "riding or walking with horse", "roller skating",
	// "rowing", "sailing", "shooting goal (soccer)", "skateboarding", "skiing".
	DetectionResults []string                                                      `json:"detection_results" api:"required"`
	Frames           []AITaskGetResponseResultAIResultsContentmoderationSportFrame `json:"frames" api:"required"`
	// A boolean value whether any sports were detected
	SportDetected bool `json:"sport_detected" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectionResults respjson.Field
		Frames           respjson.Field
		SportDetected    respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsContentmoderationSport) RawJSON() string { return r.JSON.raw }
func (r *AITaskGetResponseResultAIResultsContentmoderationSport) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsContentmoderationSportFrame struct {
	// Probability of identifying the object or activity
	Confidence float64 `json:"confidence" api:"required" format:"decimal"`
	// Video frame number where object or activity was found
	FrameNumber int64 `json:"frame_number" api:"required"`
	// Type of detected object or activity
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		FrameNumber respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsContentmoderationSportFrame) RawJSON() string {
	return r.JSON.raw
}
func (r *AITaskGetResponseResultAIResultsContentmoderationSportFrame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsContentmoderationNsfw struct {
	// Any of "nsfw".
	DetectionResults []string                                                     `json:"detection_results" api:"required"`
	Frames           []AITaskGetResponseResultAIResultsContentmoderationNsfwFrame `json:"frames" api:"required"`
	// A boolean value whether any Not Safe For Work content was detected
	NsfwDetected bool `json:"nsfw_detected" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectionResults respjson.Field
		Frames           respjson.Field
		NsfwDetected     respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsContentmoderationNsfw) RawJSON() string { return r.JSON.raw }
func (r *AITaskGetResponseResultAIResultsContentmoderationNsfw) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsContentmoderationNsfwFrame struct {
	// Probability of identifying the object or activity
	Confidence float64 `json:"confidence" api:"required" format:"decimal"`
	// Video frame number where object or activity was found
	FrameNumber int64 `json:"frame_number" api:"required"`
	// Type of detected object or activity
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		FrameNumber respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsContentmoderationNsfwFrame) RawJSON() string {
	return r.JSON.raw
}
func (r *AITaskGetResponseResultAIResultsContentmoderationNsfwFrame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsContentmoderationHardnudity struct {
	// Any of "ANUS_EXPOSED", "BUTTOCKS_EXPOSED", "FEMALE_BREAST_EXPOSED",
	// "FEMALE_GENITALIA_EXPOSED", "MALE_BREAST_EXPOSED", "MALE_GENITALIA_EXPOSED".
	DetectionResults []string                                                           `json:"detection_results" api:"required"`
	Frames           []AITaskGetResponseResultAIResultsContentmoderationHardnudityFrame `json:"frames" api:"required"`
	// A boolean value whether any nudity was detected
	NudityDetected bool `json:"nudity_detected" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectionResults respjson.Field
		Frames           respjson.Field
		NudityDetected   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsContentmoderationHardnudity) RawJSON() string {
	return r.JSON.raw
}
func (r *AITaskGetResponseResultAIResultsContentmoderationHardnudity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsContentmoderationHardnudityFrame struct {
	// Probability of identifying the object or activity
	Confidence float64 `json:"confidence" api:"required" format:"decimal"`
	// Video frame number where object or activity was found
	FrameNumber int64 `json:"frame_number" api:"required"`
	// Type of detected object or activity
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		FrameNumber respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsContentmoderationHardnudityFrame) RawJSON() string {
	return r.JSON.raw
}
func (r *AITaskGetResponseResultAIResultsContentmoderationHardnudityFrame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsContentmoderationSoftnudity struct {
	// Any of "ANUS_COVERED", "ANUS_EXPOSED", "ARMPITS_COVERED", "ARMPITS_EXPOSED",
	// "BELLY_COVERED", "BELLY_EXPOSED", "BUTTOCKS_COVERED", "BUTTOCKS_EXPOSED",
	// "FACE_FEMALE", "FACE_MALE", "FEET_COVERED", "FEET_EXPOSED",
	// "FEMALE_BREAST_COVERED", "FEMALE_BREAST_EXPOSED", "FEMALE_GENITALIA_COVERED",
	// "FEMALE_GENITALIA_EXPOSED", "MALE_BREAST_EXPOSED", "MALE_GENITALIA_EXPOSED".
	DetectionResults []string                                                           `json:"detection_results" api:"required"`
	Frames           []AITaskGetResponseResultAIResultsContentmoderationSoftnudityFrame `json:"frames" api:"required"`
	// A boolean value whether any nudity and other body part was detected
	NudityDetected bool `json:"nudity_detected" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		DetectionResults respjson.Field
		Frames           respjson.Field
		NudityDetected   respjson.Field
		ExtraFields      map[string]respjson.Field
		raw              string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsContentmoderationSoftnudity) RawJSON() string {
	return r.JSON.raw
}
func (r *AITaskGetResponseResultAIResultsContentmoderationSoftnudity) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsContentmoderationSoftnudityFrame struct {
	// Probability of identifying the object or activity
	Confidence float64 `json:"confidence" api:"required" format:"decimal"`
	// Video frame number where object or activity was found
	FrameNumber int64 `json:"frame_number" api:"required"`
	// Type of detected object or activity
	Label string `json:"label" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Confidence  respjson.Field
		FrameNumber respjson.Field
		Label       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsContentmoderationSoftnudityFrame) RawJSON() string {
	return r.JSON.raw
}
func (r *AITaskGetResponseResultAIResultsContentmoderationSoftnudityFrame) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetResponseResultAIResultsFailure struct {
	Error string `json:"error" api:"required"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Error       respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetResponseResultAIResultsFailure) RawJSON() string { return r.JSON.raw }
func (r *AITaskGetResponseResultAIResultsFailure) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskGetAISettingsResponse struct {
	// Is the given language pair supported for transcription and translation?
	Supported bool `json:"supported"`
	// JSON contains metadata for fields, check presence with [respjson.Field.Valid].
	JSON struct {
		Supported   respjson.Field
		ExtraFields map[string]respjson.Field
		raw         string
	} `json:"-"`
}

// Returns the unmodified JSON received from the API
func (r AITaskGetAISettingsResponse) RawJSON() string { return r.JSON.raw }
func (r *AITaskGetAISettingsResponse) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

type AITaskNewParams struct {

	//
	// Request body variants
	//

	// This field is a request body variant, only one variant field can be set.
	OfAITranscriptionTaskData *AITaskNewParamsBodyAITranscriptionTaskData `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAIContentModerationNsfwTaskData *AIContentmoderationNsfwParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAIContentModerationHardNudityTaskData *AIContentmoderationHardnudityParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAIContentModerationSoftNudityTaskData *AIContentmoderationSoftnudityParam `json:",inline"`
	// This field is a request body variant, only one variant field can be set.
	OfAIContentModerationSportTaskData *AIContentmoderationSportParam `json:",inline"`

	paramObj
}

func (u AITaskNewParams) MarshalJSON() ([]byte, error) {
	return param.MarshalUnion(u, u.OfAITranscriptionTaskData,
		u.OfAIContentModerationNsfwTaskData,
		u.OfAIContentModerationHardNudityTaskData,
		u.OfAIContentModerationSoftNudityTaskData,
		u.OfAIContentModerationSportTaskData)
}
func (r *AITaskNewParams) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

// The properties TaskName, URL are required.
type AITaskNewParamsBodyAITranscriptionTaskData struct {
	// Name of the task to be performed
	//
	// Any of "transcription".
	TaskName string `json:"task_name,omitzero" api:"required"`
	// URL to the MP4 file to analyze. File must be publicly accessible via HTTP/HTTPS.
	URL string `json:"url" api:"required"`
	// Language in original audio (transcription only). This value is used to determine
	// the language from which to transcribe.
	//
	// If this is not set, the system will run auto language identification and the
	// subtitles will be in the detected language. The method also works based on AI
	// analysis. It's fairly accurate, but if it's wrong, then set the language
	// explicitly.
	//
	// Additionally, when this is not set, we also support recognition of alternate
	// languages in the video (language code-switching).
	//
	// Language is set by 3-letter language code according to ISO-639-2 (bibliographic
	// code).
	//
	// We can process languages:
	//
	// - 'afr': Afrikaans
	// - 'alb': Albanian
	// - 'amh': Amharic
	// - 'ara': Arabic
	// - 'arm': Armenian
	// - 'asm': Assamese
	// - 'aze': Azerbaijani
	// - 'bak': Bashkir
	// - 'baq': Basque
	// - 'bel': Belarusian
	// - 'ben': Bengali
	// - 'bos': Bosnian
	// - 'bre': Breton
	// - 'bul': Bulgarian
	// - 'bur': Myanmar
	// - 'cat': Catalan
	// - 'chi': Chinese
	// - 'cze': Czech
	// - 'dan': Danish
	// - 'dut': Nynorsk
	// - 'eng': English
	// - 'est': Estonian
	// - 'fao': Faroese
	// - 'fin': Finnish
	// - 'fre': French
	// - 'geo': Georgian
	// - 'ger': German
	// - 'glg': Galician
	// - 'gre': Greek
	// - 'guj': Gujarati
	// - 'hat': Haitian creole
	// - 'hau': Hausa
	// - 'haw': Hawaiian
	// - 'heb': Hebrew
	// - 'hin': Hindi
	// - 'hrv': Croatian
	// - 'hun': Hungarian
	// - 'ice': Icelandic
	// - 'ind': Indonesian
	// - 'ita': Italian
	// - 'jav': Javanese
	// - 'jpn': Japanese
	// - 'kan': Kannada
	// - 'kaz': Kazakh
	// - 'khm': Khmer
	// - 'kor': Korean
	// - 'lao': Lao
	// - 'lat': Latin
	// - 'lav': Latvian
	// - 'lin': Lingala
	// - 'lit': Lithuanian
	// - 'ltz': Luxembourgish
	// - 'mac': Macedonian
	// - 'mal': Malayalam
	// - 'mao': Maori
	// - 'mar': Marathi
	// - 'may': Malay
	// - 'mlg': Malagasy
	// - 'mlt': Maltese
	// - 'mon': Mongolian
	// - 'nep': Nepali
	// - 'dut': Dutch
	// - 'nor': Norwegian
	// - 'oci': Occitan
	// - 'pan': Punjabi
	// - 'per': Persian
	// - 'pol': Polish
	// - 'por': Portuguese
	// - 'pus': Pashto
	// - 'rum': Romanian
	// - 'rus': Russian
	// - 'san': Sanskrit
	// - 'sin': Sinhala
	// - 'slo': Slovak
	// - 'slv': Slovenian
	// - 'sna': Shona
	// - 'snd': Sindhi
	// - 'som': Somali
	// - 'spa': Spanish
	// - 'srp': Serbian
	// - 'sun': Sundanese
	// - 'swa': Swahili
	// - 'swe': Swedish
	// - 'tam': Tamil
	// - 'tat': Tatar
	// - 'tel': Telugu
	// - 'tgk': Tajik
	// - 'tgl': Tagalog
	// - 'tha': Thai
	// - 'tib': Tibetan
	// - 'tuk': Turkmen
	// - 'tur': Turkish
	// - 'ukr': Ukrainian
	// - 'urd': Urdu
	// - 'uzb': Uzbek
	// - 'vie': Vietnamese
	// - 'wel': Welsh
	// - 'yid': Yiddish
	// - 'yor': Yoruba
	AudioLanguage param.Opt[string] `json:"audio_language,omitzero"`
	// Meta parameter, designed to store your own extra information about a video
	// entity: video source, video id, etc. It is not used in any way in video
	// processing.
	//
	// For example, if an AI-task was created automatically when you uploaded a video
	// with the AI auto-processing option (transcribing, translation), then the ID of
	// the associated video for which the task was performed will be explicitly
	// indicated here.
	ClientEntityData param.Opt[string] `json:"client_entity_data,omitzero"`
	// Meta parameter, designed to store your own identifier. Can be used by you to tag
	// requests from different end-users. It is not used in any way in video
	// processing.
	ClientUserID param.Opt[string] `json:"client_user_id,omitzero"`
	// Indicates which language it is clearly necessary to translate into. If this is
	// not set, the original language will be used from attribute "audio_language".
	//
	// Please note that:
	//
	//   - transcription into the original language is a free procedure,
	//   - and translation from the original language into any other languages is a
	//     "translation" procedure and is paid. More details in
	//     [POST /streaming/ai/tasks](/api-reference/streaming/ai/create-ai-task).
	//     Language is set by 3-letter language code according to ISO-639-2
	//     (bibliographic code).
	SubtitlesLanguage param.Opt[string] `json:"subtitles_language,omitzero"`
	paramObj
}

func (r AITaskNewParamsBodyAITranscriptionTaskData) MarshalJSON() (data []byte, err error) {
	type shadow AITaskNewParamsBodyAITranscriptionTaskData
	return param.MarshalObject(r, (*shadow)(&r))
}
func (r *AITaskNewParamsBodyAITranscriptionTaskData) UnmarshalJSON(data []byte) error {
	return apijson.UnmarshalRoot(data, r)
}

func init() {
	apijson.RegisterFieldValidator[AITaskNewParamsBodyAITranscriptionTaskData](
		"task_name", "transcription",
	)
}

type AITaskListParams struct {
	// Time when task was created. Datetime in ISO 8601 format.
	DateCreated param.Opt[string] `query:"date_created,omitzero" json:"-"`
	// Number of results to return per page.
	Limit param.Opt[int64] `query:"limit,omitzero" json:"-"`
	// Page to view from task list, starting from 1
	Page param.Opt[int64] `query:"page,omitzero" json:"-"`
	// This is an field for combined text search in the following fields: `task_id`,
	// `task_name`, status, and `task_data`.
	//
	// Both full and partial searches are possible inside specified above fields. For
	// example, you can filter tasks of a certain category, or tasks by a specific
	// original file.
	//
	// Example:
	//
	//   - To filter tasks of Content Moderation NSFW method:
	//     `GET /streaming/ai/tasks?search=nsfw`
	//   - To filter tasks of processing video from a specific origin:
	//     `GET /streaming/ai/tasks?search=s3.eu-west-1.amazonaws.com`
	Search param.Opt[string] `query:"search,omitzero" json:"-"`
	// The task unique identifier to find
	TaskID param.Opt[string] `query:"task_id,omitzero" format:"uuid" json:"-"`
	// Which field to use when ordering the results: `task_id`, status, and
	// `task_name`. Sorting is done in ascending (ASC) order.
	//
	// If parameter is omitted then "started_at DESC" is used for ordering by default.
	//
	// Any of "task_id", "status", "task_name", "started_at".
	Ordering AITaskListParamsOrdering `query:"ordering,omitzero" json:"-"`
	// Task status
	//
	// Any of "FAILURE", "PENDING", "RECEIVED", "RETRY", "REVOKED", "STARTED",
	// "SUCCESS".
	Status AITaskListParamsStatus `query:"status,omitzero" json:"-"`
	// Type of the AI task. Reflects the original API method that was used to create
	// the AI task.
	//
	// Any of "transcription", "content-moderation".
	TaskName AITaskListParamsTaskName `query:"task_name,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITaskListParams]'s query parameters as `url.Values`.
func (r AITaskListParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// Which field to use when ordering the results: `task_id`, status, and
// `task_name`. Sorting is done in ascending (ASC) order.
//
// If parameter is omitted then "started_at DESC" is used for ordering by default.
type AITaskListParamsOrdering string

const (
	AITaskListParamsOrderingTaskID    AITaskListParamsOrdering = "task_id"
	AITaskListParamsOrderingStatus    AITaskListParamsOrdering = "status"
	AITaskListParamsOrderingTaskName  AITaskListParamsOrdering = "task_name"
	AITaskListParamsOrderingStartedAt AITaskListParamsOrdering = "started_at"
)

// Task status
type AITaskListParamsStatus string

const (
	AITaskListParamsStatusFailure  AITaskListParamsStatus = "FAILURE"
	AITaskListParamsStatusPending  AITaskListParamsStatus = "PENDING"
	AITaskListParamsStatusReceived AITaskListParamsStatus = "RECEIVED"
	AITaskListParamsStatusRetry    AITaskListParamsStatus = "RETRY"
	AITaskListParamsStatusRevoked  AITaskListParamsStatus = "REVOKED"
	AITaskListParamsStatusStarted  AITaskListParamsStatus = "STARTED"
	AITaskListParamsStatusSuccess  AITaskListParamsStatus = "SUCCESS"
)

// Type of the AI task. Reflects the original API method that was used to create
// the AI task.
type AITaskListParamsTaskName string

const (
	AITaskListParamsTaskNameTranscription     AITaskListParamsTaskName = "transcription"
	AITaskListParamsTaskNameContentModeration AITaskListParamsTaskName = "content-moderation"
)

type AITaskGetAISettingsParams struct {
	// The parameters section for which parameters are requested
	//
	// Any of "language_support".
	Type AITaskGetAISettingsParamsType `query:"type,omitzero" api:"required" json:"-"`
	// The source language from which the audio will be transcribed. Required when
	// `type=language_support`. Value is 3-letter language code according to ISO-639-2
	// (bibliographic code), (e.g., fre for French).
	AudioLanguage param.Opt[string] `query:"audio_language,omitzero" json:"-"`
	// The target language the text will be translated into. If omitted, the API will
	// return whether the `audio_language` is supported for transcription only, instead
	// of translation. Value is 3-letter language code according to ISO-639-2
	// (bibliographic code), (e.g., fre for French).
	SubtitlesLanguage param.Opt[string] `query:"subtitles_language,omitzero" json:"-"`
	paramObj
}

// URLQuery serializes [AITaskGetAISettingsParams]'s query parameters as
// `url.Values`.
func (r AITaskGetAISettingsParams) URLQuery() (v url.Values, err error) {
	return apiquery.MarshalWithSettings(r, apiquery.QuerySettings{
		ArrayFormat:  apiquery.ArrayQueryFormatRepeat,
		NestedFormat: apiquery.NestedQueryFormatDots,
	})
}

// The parameters section for which parameters are requested
type AITaskGetAISettingsParamsType string

const (
	AITaskGetAISettingsParamsTypeLanguageSupport AITaskGetAISettingsParamsType = "language_support"
)
