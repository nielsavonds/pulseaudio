package pulseaudio

type command uint32

//go:generate stringer -type=command
const (
	/* Generic commands */
	commandError command = iota
	commandTimeout
	commandReply // 2

	/* CLIENT->SERVER */
	commandCreatePlaybackStream // 3
	commandDeletePlaybackStream
	commandCreateRecordStream
	commandDeleteRecordStream
	commandExit
	commandAuth // 8
	commandSetClientName
	commandLookupSink
	commandLookupSource
	commandDrainPlaybackStream
	commandStat
	commandGetPlaybackLatency
	commandCreateUploadStream
	commandDeleteUploadStream
	commandFinishUploadStream
	commandPlaySample
	commandRemoveSample // 19

	commandGetServerInfo
	commandGetSinkInfo
	commandGetSinkInfoList
	commandGetSourceInfo
	commandGetSourceInfoList
	commandGetModuleInfo
	commandGetModuleInfoList
	commandGetClientInfo
	commandGetClientInfoList
	commandGetSinkInputInfo
	commandGetSinkInputInfoList
	commandGetSourceOutputInfo
	commandGetSourceOutputInfoList
	commandGetSampleInfo
	commandGetSampleInfoList
	commandSubscribe

	commandSetSinkVolume
	commandSetSinkInputVolume
	commandSetSourceVolume

	commandSetSinkMute
	commandSetSourceMute // 40

	commandCorkPlaybackStream
	commandFlushPlaybackStream
	commandTriggerPlaybackStream // 43

	commandSetDefaultSink
	commandSetDefaultSource // 45

	commandSetPlaybackStreamName
	commandSetRecordStreamName // 47

	commandKillClient
	commandKillSinkInput
	commandKillSourceOutput // 50

	commandLoadModule
	commandUnloadModule // 52

	commandAddAutoloadObsolete
	commandRemoveAutoloadObsolete
	commandGetAutoloadInfoObsolete
	commandGetAutoloadInfoListObsolete //56

	commandGetRecordLatency
	commandCorkRecordStream
	commandFlushRecordStream
	commandPrebufPlaybackStream // 60

	/* SERVER->CLIENT */
	commandRequest // 61
	commandOverflow
	commandUnderflow
	commandPlaybackStreamKilled
	commandRecordStreamKilled
	commandSubscribeEvent

	/* A few more client->server commands */

	commandMoveSinkInput
	commandMoveSourceOutput
	commandSetSinkInputMute
	commandSuspendSink
	commandSuspendSource

	commandSetPlaybackStreamBufferAttr
	commandSetRecordStreamBufferAttr

	commandUpdatePlaybackStreamSampleRate
	commandUpdateRecordStreamSampleRate

	/* SERVER->CLIENT */
	commandPlaybackStreamSuspended
	commandRecordStreamSuspended
	commandPlaybackStreamMoved
	commandRecordStreamMoved

	commandUpdateRecordStreamProplist
	commandUpdatePlaybackStreamProplist
	commandUpdateClientProplist
	commandRemoveRecordStreamProplist
	commandRemovePlaybackStreamProplist
	commandRemoveClientProplist

	/* SERVER->CLIENT */
	commandStarted

	commandExtension

	commandGetCardInfo
	commandGetCardInfoList
	commandSetCardProfile

	commandClientEvent
	commandPlaybackStreamEvent
	commandRecordStreamEvent

	/* SERVER->CLIENT */
	commandPlaybackBufferAttrChanged
	commandRecordBufferAttrChanged

	commandSetSinkPort
	commandSetSourcePort

	commandSetSourceOutputVolume
	commandSetSourceOutputMute

	commandSetPortLatencyOffset

	/* BOTH DIRECTIONS */
	commandEnableSrbchannel
	commandDisableSrbchannel

	/* BOTH DIRECTIONS */
	commandRegisterMemfdShmid

	commandMax
)

// Avoid unused errors
func init() {
	_ = commandError
	_ = commandTimeout
	_ = commandReply // 2

	/* CLIENT->SERVER */
	_ = commandCreatePlaybackStream // 3
	_ = commandDeletePlaybackStream
	_ = commandCreateRecordStream
	_ = commandDeleteRecordStream
	_ = commandExit
	_ = commandAuth // 8
	_ = commandSetClientName
	_ = commandLookupSink
	_ = commandLookupSource
	_ = commandDrainPlaybackStream
	_ = commandStat
	_ = commandGetPlaybackLatency
	_ = commandCreateUploadStream
	_ = commandDeleteUploadStream
	_ = commandFinishUploadStream
	_ = commandPlaySample
	_ = commandRemoveSample // 19

	_ = commandGetServerInfo
	_ = commandGetSinkInfo
	_ = commandGetSinkInfoList
	_ = commandGetSourceInfo
	_ = commandGetSourceInfoList
	_ = commandGetModuleInfo
	_ = commandGetModuleInfoList
	_ = commandGetClientInfo
	_ = commandGetClientInfoList
	_ = commandGetSinkInputInfo
	_ = commandGetSinkInputInfoList
	_ = commandGetSourceOutputInfo
	_ = commandGetSourceOutputInfoList
	_ = commandGetSampleInfo
	_ = commandGetSampleInfoList
	_ = commandSubscribe

	_ = commandSetSinkVolume
	_ = commandSetSinkInputVolume
	_ = commandSetSourceVolume

	_ = commandSetSinkMute
	_ = commandSetSourceMute // 40

	_ = commandCorkPlaybackStream
	_ = commandFlushPlaybackStream
	_ = commandTriggerPlaybackStream // 43

	_ = commandSetDefaultSink
	_ = commandSetDefaultSource // 45

	_ = commandSetPlaybackStreamName
	_ = commandSetRecordStreamName // 47

	_ = commandKillClient
	_ = commandKillSinkInput
	_ = commandKillSourceOutput // 50

	_ = commandLoadModule
	_ = commandUnloadModule // 52

	_ = commandAddAutoloadObsolete
	_ = commandRemoveAutoloadObsolete
	_ = commandGetAutoloadInfoObsolete
	_ = commandGetAutoloadInfoListObsolete //56

	_ = commandGetRecordLatency
	_ = commandCorkRecordStream
	_ = commandFlushRecordStream
	_ = commandPrebufPlaybackStream // 60

	/* SERVER->CLIENT */
	_ = commandRequest // 61
	_ = commandOverflow
	_ = commandUnderflow
	_ = commandPlaybackStreamKilled
	_ = commandRecordStreamKilled
	_ = commandSubscribeEvent

	/* A few more client->server _ = commands */

	_ = commandMoveSinkInput
	_ = commandMoveSourceOutput
	_ = commandSetSinkInputMute
	_ = commandSuspendSink
	_ = commandSuspendSource

	_ = commandSetPlaybackStreamBufferAttr
	_ = commandSetRecordStreamBufferAttr

	_ = commandUpdatePlaybackStreamSampleRate
	_ = commandUpdateRecordStreamSampleRate

	/* SERVER->CLIENT */
	_ = commandPlaybackStreamSuspended
	_ = commandRecordStreamSuspended
	_ = commandPlaybackStreamMoved
	_ = commandRecordStreamMoved

	_ = commandUpdateRecordStreamProplist
	_ = commandUpdatePlaybackStreamProplist
	_ = commandUpdateClientProplist
	_ = commandRemoveRecordStreamProplist
	_ = commandRemovePlaybackStreamProplist
	_ = commandRemoveClientProplist

	/* SERVER->CLIENT */
	_ = commandStarted

	_ = commandExtension

	_ = commandGetCardInfo
	_ = commandGetCardInfoList
	_ = commandSetCardProfile

	_ = commandClientEvent
	_ = commandPlaybackStreamEvent
	_ = commandRecordStreamEvent

	/* SERVER->CLIENT */
	_ = commandPlaybackBufferAttrChanged
	_ = commandRecordBufferAttrChanged

	_ = commandSetSinkPort
	_ = commandSetSourcePort

	_ = commandSetSourceOutputVolume
	_ = commandSetSourceOutputMute

	_ = commandSetPortLatencyOffset

	/* BOTH DIRECTIONS */
	_ = commandEnableSrbchannel
	_ = commandDisableSrbchannel

	/* BOTH DIRECTIONS */
	_ = commandRegisterMemfdShmid

	_ = commandMax
}
