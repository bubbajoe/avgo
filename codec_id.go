package avgo

//#cgo pkg-config: libavcodec libavformat
//#include <libavcodec/avcodec.h>
//#include <libavformat/avformat.h>
import "C"

// https://github.com/FFmpeg/FFmpeg/blob/n4.4/libavcodec/codec_id.h#L47
type CodecID C.enum_AVCodecID

const (
	CodecID012V                     = CodecID(C.AV_CODEC_ID_012V)
	CodecID4Xm                      = CodecID(C.AV_CODEC_ID_4XM)
	CodecID8Bps                     = CodecID(C.AV_CODEC_ID_8BPS)
	CodecID8SvxExp                  = CodecID(C.AV_CODEC_ID_8SVX_EXP)
	CodecID8SvxFib                  = CodecID(C.AV_CODEC_ID_8SVX_FIB)
	CodecIDA64Multi                 = CodecID(C.AV_CODEC_ID_A64_MULTI)
	CodecIDA64Multi5                = CodecID(C.AV_CODEC_ID_A64_MULTI5)
	CodecIDAac                      = CodecID(C.AV_CODEC_ID_AAC)
	CodecIDAacLatm                  = CodecID(C.AV_CODEC_ID_AAC_LATM)
	CodecIDAasc                     = CodecID(C.AV_CODEC_ID_AASC)
	CodecIDAc3                      = CodecID(C.AV_CODEC_ID_AC3)
	CodecIDAdpcm4Xm                 = CodecID(C.AV_CODEC_ID_ADPCM_4XM)
	CodecIDAdpcmAdx                 = CodecID(C.AV_CODEC_ID_ADPCM_ADX)
	CodecIDAdpcmAfc                 = CodecID(C.AV_CODEC_ID_ADPCM_AFC)
	CodecIDAdpcmCt                  = CodecID(C.AV_CODEC_ID_ADPCM_CT)
	CodecIDAdpcmDtk                 = CodecID(C.AV_CODEC_ID_ADPCM_DTK)
	CodecIDAdpcmEa                  = CodecID(C.AV_CODEC_ID_ADPCM_EA)
	CodecIDAdpcmEaMaxisXa           = CodecID(C.AV_CODEC_ID_ADPCM_EA_MAXIS_XA)
	CodecIDAdpcmEaR1                = CodecID(C.AV_CODEC_ID_ADPCM_EA_R1)
	CodecIDAdpcmEaR2                = CodecID(C.AV_CODEC_ID_ADPCM_EA_R2)
	CodecIDAdpcmEaR3                = CodecID(C.AV_CODEC_ID_ADPCM_EA_R3)
	CodecIDAdpcmEaXas               = CodecID(C.AV_CODEC_ID_ADPCM_EA_XAS)
	CodecIDAdpcmG722                = CodecID(C.AV_CODEC_ID_ADPCM_G722)
	CodecIDAdpcmG726                = CodecID(C.AV_CODEC_ID_ADPCM_G726)
	CodecIDAdpcmG726Le              = CodecID(C.AV_CODEC_ID_ADPCM_G726LE)
	CodecIDAdpcmImaAmv              = CodecID(C.AV_CODEC_ID_ADPCM_IMA_AMV)
	CodecIDAdpcmImaApc              = CodecID(C.AV_CODEC_ID_ADPCM_IMA_APC)
	CodecIDAdpcmImaDk3              = CodecID(C.AV_CODEC_ID_ADPCM_IMA_DK3)
	CodecIDAdpcmImaDk4              = CodecID(C.AV_CODEC_ID_ADPCM_IMA_DK4)
	CodecIDAdpcmImaEaEacs           = CodecID(C.AV_CODEC_ID_ADPCM_IMA_EA_EACS)
	CodecIDAdpcmImaEaSead           = CodecID(C.AV_CODEC_ID_ADPCM_IMA_EA_SEAD)
	CodecIDAdpcmImaIss              = CodecID(C.AV_CODEC_ID_ADPCM_IMA_ISS)
	CodecIDAdpcmImaOki              = CodecID(C.AV_CODEC_ID_ADPCM_IMA_OKI)
	CodecIDAdpcmImaQt               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_QT)
	CodecIDAdpcmImaRad              = CodecID(C.AV_CODEC_ID_ADPCM_IMA_RAD)
	CodecIDAdpcmImaSmjpeg           = CodecID(C.AV_CODEC_ID_ADPCM_IMA_SMJPEG)
	CodecIDAdpcmImaWav              = CodecID(C.AV_CODEC_ID_ADPCM_IMA_WAV)
	CodecIDAdpcmImaWs               = CodecID(C.AV_CODEC_ID_ADPCM_IMA_WS)
	CodecIDAdpcmMs                  = CodecID(C.AV_CODEC_ID_ADPCM_MS)
	CodecIDAdpcmSbpro2              = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_2)
	CodecIDAdpcmSbpro3              = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_3)
	CodecIDAdpcmSbpro4              = CodecID(C.AV_CODEC_ID_ADPCM_SBPRO_4)
	CodecIDAdpcmSwf                 = CodecID(C.AV_CODEC_ID_ADPCM_SWF)
	CodecIDAdpcmThp                 = CodecID(C.AV_CODEC_ID_ADPCM_THP)
	CodecIDAdpcmVima                = CodecID(C.AV_CODEC_ID_ADPCM_VIMA)
	CodecIDAdpcmVimaDeprecated      = CodecID(C.AV_CODEC_ID_ADPCM_VIMA)
	CodecIDAdpcmXa                  = CodecID(C.AV_CODEC_ID_ADPCM_XA)
	CodecIDAdpcmYamaha              = CodecID(C.AV_CODEC_ID_ADPCM_YAMAHA)
	CodecIDAic                      = CodecID(C.AV_CODEC_ID_AIC)
	CodecIDAlac                     = CodecID(C.AV_CODEC_ID_ALAC)
	CodecIDAliasPix                 = CodecID(C.AV_CODEC_ID_ALIAS_PIX)
	CodecIDAmrNb                    = CodecID(C.AV_CODEC_ID_AMR_NB)
	CodecIDAmrWb                    = CodecID(C.AV_CODEC_ID_AMR_WB)
	CodecIDAmv                      = CodecID(C.AV_CODEC_ID_AMV)
	CodecIDAnm                      = CodecID(C.AV_CODEC_ID_ANM)
	CodecIDAnsi                     = CodecID(C.AV_CODEC_ID_ANSI)
	CodecIDApe                      = CodecID(C.AV_CODEC_ID_APE)
	CodecIDAss                      = CodecID(C.AV_CODEC_ID_ASS)
	CodecIDAsv1                     = CodecID(C.AV_CODEC_ID_ASV1)
	CodecIDAsv2                     = CodecID(C.AV_CODEC_ID_ASV2)
	CodecIDAtrac1                   = CodecID(C.AV_CODEC_ID_ATRAC1)
	CodecIDAtrac3                   = CodecID(C.AV_CODEC_ID_ATRAC3)
	CodecIDAtrac3P                  = CodecID(C.AV_CODEC_ID_ATRAC3P)
	CodecIDAura                     = CodecID(C.AV_CODEC_ID_AURA)
	CodecIDAura2                    = CodecID(C.AV_CODEC_ID_AURA2)
	CodecIDAvrn                     = CodecID(C.AV_CODEC_ID_AVRN)
	CodecIDAvrp                     = CodecID(C.AV_CODEC_ID_AVRP)
	CodecIDAvs                      = CodecID(C.AV_CODEC_ID_AVS)
	CodecIDAvui                     = CodecID(C.AV_CODEC_ID_AVUI)
	CodecIDAyuv                     = CodecID(C.AV_CODEC_ID_AYUV)
	CodecIDBethsoftvid              = CodecID(C.AV_CODEC_ID_BETHSOFTVID)
	CodecIDBfi                      = CodecID(C.AV_CODEC_ID_BFI)
	CodecIDBinData                  = CodecID(C.AV_CODEC_ID_BIN_DATA)
	CodecIDBinkaudioDct             = CodecID(C.AV_CODEC_ID_BINKAUDIO_DCT)
	CodecIDBinkaudioRdft            = CodecID(C.AV_CODEC_ID_BINKAUDIO_RDFT)
	CodecIDBinkvideo                = CodecID(C.AV_CODEC_ID_BINKVIDEO)
	CodecIDBCodecIDext              = CodecID(C.AV_CODEC_ID_BINTEXT)
	CodecIDBmp                      = CodecID(C.AV_CODEC_ID_BMP)
	CodecIDBmvAudio                 = CodecID(C.AV_CODEC_ID_BMV_AUDIO)
	CodecIDBmvVideo                 = CodecID(C.AV_CODEC_ID_BMV_VIDEO)
	CodecIDBrenderPix               = CodecID(C.AV_CODEC_ID_BRENDER_PIX)
	CodecIDBrenderPixDeprecated     = CodecID(C.AV_CODEC_ID_BRENDER_PIX)
	CodecIDC93                      = CodecID(C.AV_CODEC_ID_C93)
	CodecIDCavs                     = CodecID(C.AV_CODEC_ID_CAVS)
	CodecIDCdgraphics               = CodecID(C.AV_CODEC_ID_CDGRAPHICS)
	CodecIDCdxl                     = CodecID(C.AV_CODEC_ID_CDXL)
	CodecIDCelt                     = CodecID(C.AV_CODEC_ID_CELT)
	CodecIDCinepak                  = CodecID(C.AV_CODEC_ID_CINEPAK)
	CodecIDCljr                     = CodecID(C.AV_CODEC_ID_CLJR)
	CodecIDCllc                     = CodecID(C.AV_CODEC_ID_CLLC)
	CodecIDCmv                      = CodecID(C.AV_CODEC_ID_CMV)
	CodecIDComfortNoise             = CodecID(C.AV_CODEC_ID_COMFORT_NOISE)
	CodecIDCook                     = CodecID(C.AV_CODEC_ID_COOK)
	CodecIDCpia                     = CodecID(C.AV_CODEC_ID_CPIA)
	CodecIDCscd                     = CodecID(C.AV_CODEC_ID_CSCD)
	CodecIDCyuv                     = CodecID(C.AV_CODEC_ID_CYUV)
	CodecIDDfa                      = CodecID(C.AV_CODEC_ID_DFA)
	CodecIDDirac                    = CodecID(C.AV_CODEC_ID_DIRAC)
	CodecIDDnxhd                    = CodecID(C.AV_CODEC_ID_DNXHD)
	CodecIDDpx                      = CodecID(C.AV_CODEC_ID_DPX)
	CodecIDDsdLsbf                  = CodecID(C.AV_CODEC_ID_DSD_LSBF)
	CodecIDDsdLsbfPlanar            = CodecID(C.AV_CODEC_ID_DSD_LSBF_PLANAR)
	CodecIDDsdMsbf                  = CodecID(C.AV_CODEC_ID_DSD_MSBF)
	CodecIDDsdMsbfPlanar            = CodecID(C.AV_CODEC_ID_DSD_MSBF_PLANAR)
	CodecIDDsicinaudio              = CodecID(C.AV_CODEC_ID_DSICINAUDIO)
	CodecIDDsicinvideo              = CodecID(C.AV_CODEC_ID_DSICINVIDEO)
	CodecIDDts                      = CodecID(C.AV_CODEC_ID_DTS)
	CodecIDDvaudio                  = CodecID(C.AV_CODEC_ID_DVAUDIO)
	CodecIDDvbSubtitle              = CodecID(C.AV_CODEC_ID_DVB_SUBTITLE)
	CodecIDDvbTeletext              = CodecID(C.AV_CODEC_ID_DVB_TELETEXT)
	CodecIDDvdNav                   = CodecID(C.AV_CODEC_ID_DVD_NAV)
	CodecIDDvdSubtitle              = CodecID(C.AV_CODEC_ID_DVD_SUBTITLE)
	CodecIDDvvideo                  = CodecID(C.AV_CODEC_ID_DVVIDEO)
	CodecIDDxa                      = CodecID(C.AV_CODEC_ID_DXA)
	CodecIDDxtory                   = CodecID(C.AV_CODEC_ID_DXTORY)
	CodecIDEac3                     = CodecID(C.AV_CODEC_ID_EAC3)
	CodecIDEia608                   = CodecID(C.AV_CODEC_ID_EIA_608)
	CodecIDEscape124                = CodecID(C.AV_CODEC_ID_ESCAPE124)
	CodecIDEscape130                = CodecID(C.AV_CODEC_ID_ESCAPE130)
	CodecIDEscape130Deprecated      = CodecID(C.AV_CODEC_ID_ESCAPE130)
	CodecIDEvrc                     = CodecID(C.AV_CODEC_ID_EVRC)
	CodecIDExr                      = CodecID(C.AV_CODEC_ID_EXR)
	CodecIDExrDeprecated            = CodecID(C.AV_CODEC_ID_EXR)
	CodecIDFfmetadata               = CodecID(C.AV_CODEC_ID_FFMETADATA)
	CodecIDFfv1                     = CodecID(C.AV_CODEC_ID_FFV1)
	CodecIDFfvhuff                  = CodecID(C.AV_CODEC_ID_FFVHUFF)
	CodecIDFfwavesynth              = CodecID(C.AV_CODEC_ID_FFWAVESYNTH)
	CodecIDFic                      = CodecID(C.AV_CODEC_ID_FIC)
	CodecIDFirstAudio               = CodecID(C.AV_CODEC_ID_FIRST_AUDIO)
	CodecIDFirstSubtitle            = CodecID(C.AV_CODEC_ID_FIRST_SUBTITLE)
	CodecIDFirstUnknown             = CodecID(C.AV_CODEC_ID_FIRST_UNKNOWN)
	CodecIDFlac                     = CodecID(C.AV_CODEC_ID_FLAC)
	CodecIDFlashsv                  = CodecID(C.AV_CODEC_ID_FLASHSV)
	CodecIDFlashsv2                 = CodecID(C.AV_CODEC_ID_FLASHSV2)
	CodecIDFlic                     = CodecID(C.AV_CODEC_ID_FLIC)
	CodecIDFlv1                     = CodecID(C.AV_CODEC_ID_FLV1)
	CodecIDFraps                    = CodecID(C.AV_CODEC_ID_FRAPS)
	CodecIDFrwu                     = CodecID(C.AV_CODEC_ID_FRWU)
	CodecIDG2M                      = CodecID(C.AV_CODEC_ID_G2M)
	CodecIDG2MDeprecated            = CodecID(C.AV_CODEC_ID_G2M)
	CodecIDG7231                    = CodecID(C.AV_CODEC_ID_G723_1)
	CodecIDG729                     = CodecID(C.AV_CODEC_ID_G729)
	CodecIDGif                      = CodecID(C.AV_CODEC_ID_GIF)
	CodecIDGsm                      = CodecID(C.AV_CODEC_ID_GSM)
	CodecIDGsmMs                    = CodecID(C.AV_CODEC_ID_GSM_MS)
	CodecIDH261                     = CodecID(C.AV_CODEC_ID_H261)
	CodecIDH263                     = CodecID(C.AV_CODEC_ID_H263)
	CodecIDH263I                    = CodecID(C.AV_CODEC_ID_H263I)
	CodecIDH263P                    = CodecID(C.AV_CODEC_ID_H263P)
	CodecIDH264                     = CodecID(C.AV_CODEC_ID_H264)
	CodecIDHdmvPgsSubtitle          = CodecID(C.AV_CODEC_ID_HDMV_PGS_SUBTITLE)
	CodecIDHevc                     = CodecID(C.AV_CODEC_ID_HEVC)
	CodecIDHevcDeprecated           = CodecID(C.AV_CODEC_ID_HEVC)
	CodecIDHnm4Video                = CodecID(C.AV_CODEC_ID_HNM4_VIDEO)
	CodecIDHuffyuv                  = CodecID(C.AV_CODEC_ID_HUFFYUV)
	CodecIDIac                      = CodecID(C.AV_CODEC_ID_IAC)
	CodecIDIdcin                    = CodecID(C.AV_CODEC_ID_IDCIN)
	CodecIDIdf                      = CodecID(C.AV_CODEC_ID_IDF)
	CodecIDIffByterun1              = CodecID(C.AV_CODEC_ID_IFF_BYTERUN1)
	CodecIDIffIlbm                  = CodecID(C.AV_CODEC_ID_IFF_ILBM)
	CodecIDIlbc                     = CodecID(C.AV_CODEC_ID_ILBC)
	CodecIDImc                      = CodecID(C.AV_CODEC_ID_IMC)
	CodecIDIndeo2                   = CodecID(C.AV_CODEC_ID_INDEO2)
	CodecIDIndeo3                   = CodecID(C.AV_CODEC_ID_INDEO3)
	CodecIDIndeo4                   = CodecID(C.AV_CODEC_ID_INDEO4)
	CodecIDIndeo5                   = CodecID(C.AV_CODEC_ID_INDEO5)
	CodecIDInterplayDpcm            = CodecID(C.AV_CODEC_ID_INTERPLAY_DPCM)
	CodecIDInterplayVideo           = CodecID(C.AV_CODEC_ID_INTERPLAY_VIDEO)
	CodecIDJacosub                  = CodecID(C.AV_CODEC_ID_JACOSUB)
	CodecIDJpeg2000                 = CodecID(C.AV_CODEC_ID_JPEG2000)
	CodecIDJpegls                   = CodecID(C.AV_CODEC_ID_JPEGLS)
	CodecIDJv                       = CodecID(C.AV_CODEC_ID_JV)
	CodecIDKgv1                     = CodecID(C.AV_CODEC_ID_KGV1)
	CodecIDKmvc                     = CodecID(C.AV_CODEC_ID_KMVC)
	CodecIDLagarith                 = CodecID(C.AV_CODEC_ID_LAGARITH)
	CodecIDLjpeg                    = CodecID(C.AV_CODEC_ID_LJPEG)
	CodecIDLoco                     = CodecID(C.AV_CODEC_ID_LOCO)
	CodecIDMace3                    = CodecID(C.AV_CODEC_ID_MACE3)
	CodecIDMace6                    = CodecID(C.AV_CODEC_ID_MACE6)
	CodecIDMad                      = CodecID(C.AV_CODEC_ID_MAD)
	CodecIDMdec                     = CodecID(C.AV_CODEC_ID_MDEC)
	CodecIDMetasound                = CodecID(C.AV_CODEC_ID_METASOUND)
	CodecIDMicrodvd                 = CodecID(C.AV_CODEC_ID_MICRODVD)
	CodecIDMimic                    = CodecID(C.AV_CODEC_ID_MIMIC)
	CodecIDMjpeg                    = CodecID(C.AV_CODEC_ID_MJPEG)
	CodecIDMjpegb                   = CodecID(C.AV_CODEC_ID_MJPEGB)
	CodecIDMlp                      = CodecID(C.AV_CODEC_ID_MLP)
	CodecIDMmvideo                  = CodecID(C.AV_CODEC_ID_MMVIDEO)
	CodecIDMotionpixels             = CodecID(C.AV_CODEC_ID_MOTIONPIXELS)
	CodecIDMovText                  = CodecID(C.AV_CODEC_ID_MOV_TEXT)
	CodecIDMp1                      = CodecID(C.AV_CODEC_ID_MP1)
	CodecIDMp2                      = CodecID(C.AV_CODEC_ID_MP2)
	CodecIDMp3                      = CodecID(C.AV_CODEC_ID_MP3)
	CodecIDMp3Adu                   = CodecID(C.AV_CODEC_ID_MP3ADU)
	CodecIDMp3On4                   = CodecID(C.AV_CODEC_ID_MP3ON4)
	CodecIDMp4Als                   = CodecID(C.AV_CODEC_ID_MP4ALS)
	CodecIDMpeg1Video               = CodecID(C.AV_CODEC_ID_MPEG1VIDEO)
	CodecIDMpeg2Ts                  = CodecID(C.AV_CODEC_ID_MPEG2TS)
	CodecIDMpeg2Video               = CodecID(C.AV_CODEC_ID_MPEG2VIDEO)
	CodecIDMpeg4                    = CodecID(C.AV_CODEC_ID_MPEG4)
	CodecIDMpeg4Systems             = CodecID(C.AV_CODEC_ID_MPEG4SYSTEMS)
	CodecIDMpl2                     = CodecID(C.AV_CODEC_ID_MPL2)
	CodecIDMsa1                     = CodecID(C.AV_CODEC_ID_MSA1)
	CodecIDMsmpeg4V1                = CodecID(C.AV_CODEC_ID_MSMPEG4V1)
	CodecIDMsmpeg4V2                = CodecID(C.AV_CODEC_ID_MSMPEG4V2)
	CodecIDMsmpeg4V3                = CodecID(C.AV_CODEC_ID_MSMPEG4V3)
	CodecIDMsrle                    = CodecID(C.AV_CODEC_ID_MSRLE)
	CodecIDMss1                     = CodecID(C.AV_CODEC_ID_MSS1)
	CodecIDMss2                     = CodecID(C.AV_CODEC_ID_MSS2)
	CodecIDMsvideo1                 = CodecID(C.AV_CODEC_ID_MSVIDEO1)
	CodecIDMszh                     = CodecID(C.AV_CODEC_ID_MSZH)
	CodecIDMts2                     = CodecID(C.AV_CODEC_ID_MTS2)
	CodecIDMusepack7                = CodecID(C.AV_CODEC_ID_MUSEPACK7)
	CodecIDMusepack8                = CodecID(C.AV_CODEC_ID_MUSEPACK8)
	CodecIDMvc1                     = CodecID(C.AV_CODEC_ID_MVC1)
	CodecIDMvc1Deprecated           = CodecID(C.AV_CODEC_ID_MVC1)
	CodecIDMvc2                     = CodecID(C.AV_CODEC_ID_MVC2)
	CodecIDMvc2Deprecated           = CodecID(C.AV_CODEC_ID_MVC2)
	CodecIDMxpeg                    = CodecID(C.AV_CODEC_ID_MXPEG)
	CodecIDNellymoser               = CodecID(C.AV_CODEC_ID_NELLYMOSER)
	CodecIDNone                     = CodecID(C.AV_CODEC_ID_NONE)
	CodecIDNuv                      = CodecID(C.AV_CODEC_ID_NUV)
	CodecIDOn2Avc                   = CodecID(C.AV_CODEC_ID_ON2AVC)
	CodecIDOpus                     = CodecID(C.AV_CODEC_ID_OPUS)
	CodecIDOpusDeprecated           = CodecID(C.AV_CODEC_ID_OPUS)
	CodecIDOtf                      = CodecID(C.AV_CODEC_ID_OTF)
	CodecIDPafAudio                 = CodecID(C.AV_CODEC_ID_PAF_AUDIO)
	CodecIDPafAudioDeprecated       = CodecID(C.AV_CODEC_ID_PAF_AUDIO)
	CodecIDPafVideo                 = CodecID(C.AV_CODEC_ID_PAF_VIDEO)
	CodecIDPafVideoDeprecated       = CodecID(C.AV_CODEC_ID_PAF_VIDEO)
	CodecIDPam                      = CodecID(C.AV_CODEC_ID_PAM)
	CodecIDPbm                      = CodecID(C.AV_CODEC_ID_PBM)
	CodecIDPcmAlaw                  = CodecID(C.AV_CODEC_ID_PCM_ALAW)
	CodecIDPcmBluray                = CodecID(C.AV_CODEC_ID_PCM_BLURAY)
	CodecIDPcmDvd                   = CodecID(C.AV_CODEC_ID_PCM_DVD)
	CodecIDPcmF32Be                 = CodecID(C.AV_CODEC_ID_PCM_F32BE)
	CodecIDPcmF32Le                 = CodecID(C.AV_CODEC_ID_PCM_F32LE)
	CodecIDPcmF64Be                 = CodecID(C.AV_CODEC_ID_PCM_F64BE)
	CodecIDPcmF64Le                 = CodecID(C.AV_CODEC_ID_PCM_F64LE)
	CodecIDPcmLxf                   = CodecID(C.AV_CODEC_ID_PCM_LXF)
	CodecIDPcmMulaw                 = CodecID(C.AV_CODEC_ID_PCM_MULAW)
	CodecIDPcmS16Be                 = CodecID(C.AV_CODEC_ID_PCM_S16BE)
	CodecIDPcmS16BePlanar           = CodecID(C.AV_CODEC_ID_PCM_S16BE_PLANAR)
	CodecIDPcmS16Le                 = CodecID(C.AV_CODEC_ID_PCM_S16LE)
	CodecIDPcmS16LePlanar           = CodecID(C.AV_CODEC_ID_PCM_S16LE_PLANAR)
	CodecIDPcmS24Be                 = CodecID(C.AV_CODEC_ID_PCM_S24BE)
	CodecIDPcmS24Daud               = CodecID(C.AV_CODEC_ID_PCM_S24DAUD)
	CodecIDPcmS24Le                 = CodecID(C.AV_CODEC_ID_PCM_S24LE)
	CodecIDPcmS24LePlanar           = CodecID(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	CodecIDPcmS24LePlanarDeprecated = CodecID(C.AV_CODEC_ID_PCM_S24LE_PLANAR)
	CodecIDPcmS32Be                 = CodecID(C.AV_CODEC_ID_PCM_S32BE)
	CodecIDPcmS32Le                 = CodecID(C.AV_CODEC_ID_PCM_S32LE)
	CodecIDPcmS32LePlanar           = CodecID(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	CodecIDPcmS32LePlanarDeprecated = CodecID(C.AV_CODEC_ID_PCM_S32LE_PLANAR)
	CodecIDPcmS8                    = CodecID(C.AV_CODEC_ID_PCM_S8)
	CodecIDPcmS8Planar              = CodecID(C.AV_CODEC_ID_PCM_S8_PLANAR)
	CodecIDPcmU16Be                 = CodecID(C.AV_CODEC_ID_PCM_U16BE)
	CodecIDPcmU16Le                 = CodecID(C.AV_CODEC_ID_PCM_U16LE)
	CodecIDPcmU24Be                 = CodecID(C.AV_CODEC_ID_PCM_U24BE)
	CodecIDPcmU24Le                 = CodecID(C.AV_CODEC_ID_PCM_U24LE)
	CodecIDPcmU32Be                 = CodecID(C.AV_CODEC_ID_PCM_U32BE)
	CodecIDPcmU32Le                 = CodecID(C.AV_CODEC_ID_PCM_U32LE)
	CodecIDPcmU8                    = CodecID(C.AV_CODEC_ID_PCM_U8)
	CodecIDPcmZork                  = CodecID(C.AV_CODEC_ID_PCM_ZORK)
	CodecIDPcx                      = CodecID(C.AV_CODEC_ID_PCX)
	CodecIDPgm                      = CodecID(C.AV_CODEC_ID_PGM)
	CodecIDPgmyuv                   = CodecID(C.AV_CODEC_ID_PGMYUV)
	CodecIDPictor                   = CodecID(C.AV_CODEC_ID_PICTOR)
	CodecIDPjs                      = CodecID(C.AV_CODEC_ID_PJS)
	CodecIDPng                      = CodecID(C.AV_CODEC_ID_PNG)
	CodecIDPpm                      = CodecID(C.AV_CODEC_ID_PPM)
	CodecIDProbe                    = CodecID(C.AV_CODEC_ID_PROBE)
	CodecIDProres                   = CodecID(C.AV_CODEC_ID_PRORES)
	CodecIDPtx                      = CodecID(C.AV_CODEC_ID_PTX)
	CodecIDQcelp                    = CodecID(C.AV_CODEC_ID_QCELP)
	CodecIDQdm2                     = CodecID(C.AV_CODEC_ID_QDM2)
	CodecIDQdmc                     = CodecID(C.AV_CODEC_ID_QDMC)
	CodecIDQdraw                    = CodecID(C.AV_CODEC_ID_QDRAW)
	CodecIDQpeg                     = CodecID(C.AV_CODEC_ID_QPEG)
	CodecIDQtrle                    = CodecID(C.AV_CODEC_ID_QTRLE)
	CodecIDR10K                     = CodecID(C.AV_CODEC_ID_R10K)
	CodecIDR210                     = CodecID(C.AV_CODEC_ID_R210)
	CodecIDRa144                    = CodecID(C.AV_CODEC_ID_RA_144)
	CodecIDRa288                    = CodecID(C.AV_CODEC_ID_RA_288)
	CodecIDRalf                     = CodecID(C.AV_CODEC_ID_RALF)
	CodecIDRawvideo                 = CodecID(C.AV_CODEC_ID_RAWVIDEO)
	CodecIDRealtext                 = CodecID(C.AV_CODEC_ID_REALTEXT)
	CodecIDRl2                      = CodecID(C.AV_CODEC_ID_RL2)
	CodecIDRoq                      = CodecID(C.AV_CODEC_ID_ROQ)
	CodecIDRoqDpcm                  = CodecID(C.AV_CODEC_ID_ROQ_DPCM)
	CodecIDRpza                     = CodecID(C.AV_CODEC_ID_RPZA)
	CodecIDRv10                     = CodecID(C.AV_CODEC_ID_RV10)
	CodecIDRv20                     = CodecID(C.AV_CODEC_ID_RV20)
	CodecIDRv30                     = CodecID(C.AV_CODEC_ID_RV30)
	CodecIDRv40                     = CodecID(C.AV_CODEC_ID_RV40)
	CodecIDS302M                    = CodecID(C.AV_CODEC_ID_S302M)
	CodecIDSami                     = CodecID(C.AV_CODEC_ID_SAMI)
	CodecIDSanm                     = CodecID(C.AV_CODEC_ID_SANM)
	CodecIDSanmDeprecated           = CodecID(C.AV_CODEC_ID_SANM)
	CodecIDSgi                      = CodecID(C.AV_CODEC_ID_SGI)
	CodecIDSgirle                   = CodecID(C.AV_CODEC_ID_SGIRLE)
	CodecIDSgirleDeprecated         = CodecID(C.AV_CODEC_ID_SGIRLE)
	CodecIDShorten                  = CodecID(C.AV_CODEC_ID_SHORTEN)
	CodecIDSipr                     = CodecID(C.AV_CODEC_ID_SIPR)
	CodecIDSmackaudio               = CodecID(C.AV_CODEC_ID_SMACKAUDIO)
	CodecIDSmackvideo               = CodecID(C.AV_CODEC_ID_SMACKVIDEO)
	CodecIDSmc                      = CodecID(C.AV_CODEC_ID_SMC)
	CodecIDSmpteKlv                 = CodecID(C.AV_CODEC_ID_SMPTE_KLV)
	CodecIDSmv                      = CodecID(C.AV_CODEC_ID_SMV)
	CodecIDSmvjpeg                  = CodecID(C.AV_CODEC_ID_SMVJPEG)
	CodecIDSnow                     = CodecID(C.AV_CODEC_ID_SNOW)
	CodecIDSolDpcm                  = CodecID(C.AV_CODEC_ID_SOL_DPCM)
	CodecIDSonic                    = CodecID(C.AV_CODEC_ID_SONIC)
	CodecIDSonicLs                  = CodecID(C.AV_CODEC_ID_SONIC_LS)
	CodecIDSp5X                     = CodecID(C.AV_CODEC_ID_SP5X)
	CodecIDSpeex                    = CodecID(C.AV_CODEC_ID_SPEEX)
	CodecIDSrt                      = CodecID(C.AV_CODEC_ID_SRT)
	CodecIDSsa                      = CodecID(C.AV_CODEC_ID_SSA)
	CodecIDSubrip                   = CodecID(C.AV_CODEC_ID_SUBRIP)
	CodecIDSubviewer                = CodecID(C.AV_CODEC_ID_SUBVIEWER)
	CodecIDSubviewer1               = CodecID(C.AV_CODEC_ID_SUBVIEWER1)
	CodecIDSunrast                  = CodecID(C.AV_CODEC_ID_SUNRAST)
	CodecIDSvq1                     = CodecID(C.AV_CODEC_ID_SVQ1)
	CodecIDSvq3                     = CodecID(C.AV_CODEC_ID_SVQ3)
	CodecIDTak                      = CodecID(C.AV_CODEC_ID_TAK)
	CodecIDTakDeprecated            = CodecID(C.AV_CODEC_ID_TAK)
	CodecIDTarga                    = CodecID(C.AV_CODEC_ID_TARGA)
	CodecIDTargaY216                = CodecID(C.AV_CODEC_ID_TARGA_Y216)
	CodecIDText                     = CodecID(C.AV_CODEC_ID_TEXT)
	CodecIDTgq                      = CodecID(C.AV_CODEC_ID_TGQ)
	CodecIDTgv                      = CodecID(C.AV_CODEC_ID_TGV)
	CodecIDTheora                   = CodecID(C.AV_CODEC_ID_THEORA)
	CodecIDThp                      = CodecID(C.AV_CODEC_ID_THP)
	CodecIDTiertexseqvideo          = CodecID(C.AV_CODEC_ID_TIERTEXSEQVIDEO)
	CodecIDTiff                     = CodecID(C.AV_CODEC_ID_TIFF)
	CodecIDTimedId3                 = CodecID(C.AV_CODEC_ID_TIMED_ID3)
	CodecIDTmv                      = CodecID(C.AV_CODEC_ID_TMV)
	CodecIDTqi                      = CodecID(C.AV_CODEC_ID_TQI)
	CodecIDTruehd                   = CodecID(C.AV_CODEC_ID_TRUEHD)
	CodecIDTruemotion1              = CodecID(C.AV_CODEC_ID_TRUEMOTION1)
	CodecIDTruemotion2              = CodecID(C.AV_CODEC_ID_TRUEMOTION2)
	CodecIDTruespeech               = CodecID(C.AV_CODEC_ID_TRUESPEECH)
	CodecIDTscc                     = CodecID(C.AV_CODEC_ID_TSCC)
	CodecIDTscc2                    = CodecID(C.AV_CODEC_ID_TSCC2)
	CodecIDTta                      = CodecID(C.AV_CODEC_ID_TTA)
	CodecIDTtf                      = CodecID(C.AV_CODEC_ID_TTF)
	CodecIDTwinvq                   = CodecID(C.AV_CODEC_ID_TWINVQ)
	CodecIDTxd                      = CodecID(C.AV_CODEC_ID_TXD)
	CodecIDUlti                     = CodecID(C.AV_CODEC_ID_ULTI)
	CodecIDUtvideo                  = CodecID(C.AV_CODEC_ID_UTVIDEO)
	CodecIDV210                     = CodecID(C.AV_CODEC_ID_V210)
	CodecIDV210X                    = CodecID(C.AV_CODEC_ID_V210X)
	CodecIDV308                     = CodecID(C.AV_CODEC_ID_V308)
	CodecIDV408                     = CodecID(C.AV_CODEC_ID_V408)
	CodecIDV410                     = CodecID(C.AV_CODEC_ID_V410)
	CodecIDVb                       = CodecID(C.AV_CODEC_ID_VB)
	CodecIDVble                     = CodecID(C.AV_CODEC_ID_VBLE)
	CodecIDVc1                      = CodecID(C.AV_CODEC_ID_VC1)
	CodecIDVc1Image                 = CodecID(C.AV_CODEC_ID_VC1IMAGE)
	CodecIDVcr1                     = CodecID(C.AV_CODEC_ID_VCR1)
	CodecIDVixl                     = CodecID(C.AV_CODEC_ID_VIXL)
	CodecIDVmdaudio                 = CodecID(C.AV_CODEC_ID_VMDAUDIO)
	CodecIDVmdvideo                 = CodecID(C.AV_CODEC_ID_VMDVIDEO)
	CodecIDVmnc                     = CodecID(C.AV_CODEC_ID_VMNC)
	CodecIDVorbis                   = CodecID(C.AV_CODEC_ID_VORBIS)
	CodecIDVp3                      = CodecID(C.AV_CODEC_ID_VP3)
	CodecIDVp5                      = CodecID(C.AV_CODEC_ID_VP5)
	CodecIDVp6                      = CodecID(C.AV_CODEC_ID_VP6)
	CodecIDVp6A                     = CodecID(C.AV_CODEC_ID_VP6A)
	CodecIDVp6F                     = CodecID(C.AV_CODEC_ID_VP6F)
	CodecIDVp7                      = CodecID(C.AV_CODEC_ID_VP7)
	CodecIDVp7Deprecated            = CodecID(C.AV_CODEC_ID_VP7)
	CodecIDVp8                      = CodecID(C.AV_CODEC_ID_VP8)
	CodecIDVp9                      = CodecID(C.AV_CODEC_ID_VP9)
	CodecIDVplayer                  = CodecID(C.AV_CODEC_ID_VPLAYER)
	CodecIDWavpack                  = CodecID(C.AV_CODEC_ID_WAVPACK)
	CodecIDWebp                     = CodecID(C.AV_CODEC_ID_WEBP)
	CodecIDWebpDeprecated           = CodecID(C.AV_CODEC_ID_WEBP)
	CodecIDWebvtt                   = CodecID(C.AV_CODEC_ID_WEBVTT)
	CodecIDWestwoodSnd1             = CodecID(C.AV_CODEC_ID_WESTWOOD_SND1)
	CodecIDWmalossless              = CodecID(C.AV_CODEC_ID_WMALOSSLESS)
	CodecIDWmapro                   = CodecID(C.AV_CODEC_ID_WMAPRO)
	CodecIDWmav1                    = CodecID(C.AV_CODEC_ID_WMAV1)
	CodecIDWmav2                    = CodecID(C.AV_CODEC_ID_WMAV2)
	CodecIDWmavoice                 = CodecID(C.AV_CODEC_ID_WMAVOICE)
	CodecIDWmv1                     = CodecID(C.AV_CODEC_ID_WMV1)
	CodecIDWmv2                     = CodecID(C.AV_CODEC_ID_WMV2)
	CodecIDWmv3                     = CodecID(C.AV_CODEC_ID_WMV3)
	CodecIDWmv3Image                = CodecID(C.AV_CODEC_ID_WMV3IMAGE)
	CodecIDWnv1                     = CodecID(C.AV_CODEC_ID_WNV1)
	CodecIDWsVqa                    = CodecID(C.AV_CODEC_ID_WS_VQA)
	CodecIDXanDpcm                  = CodecID(C.AV_CODEC_ID_XAN_DPCM)
	CodecIDXanWc3                   = CodecID(C.AV_CODEC_ID_XAN_WC3)
	CodecIDXanWc4                   = CodecID(C.AV_CODEC_ID_XAN_WC4)
	CodecIDXbin                     = CodecID(C.AV_CODEC_ID_XBIN)
	CodecIDXbm                      = CodecID(C.AV_CODEC_ID_XBM)
	CodecIDXface                    = CodecID(C.AV_CODEC_ID_XFACE)
	CodecIDXsub                     = CodecID(C.AV_CODEC_ID_XSUB)
	CodecIDXwd                      = CodecID(C.AV_CODEC_ID_XWD)
	CodecIDY41P                     = CodecID(C.AV_CODEC_ID_Y41P)
	CodecIDYop                      = CodecID(C.AV_CODEC_ID_YOP)
	CodecIDYuv4                     = CodecID(C.AV_CODEC_ID_YUV4)
	CodecIDZerocodec                = CodecID(C.AV_CODEC_ID_ZEROCODEC)
	CodecIDZlib                     = CodecID(C.AV_CODEC_ID_ZLIB)
	CodecIDZmbv                     = CodecID(C.AV_CODEC_ID_ZMBV)
)

func (c CodecID) MediaType() MediaType {
	return MediaType(C.avcodec_get_type((C.enum_AVCodecID)(c)))
}

func (c CodecID) Name() string {
	return C.GoString(C.avcodec_get_name((C.enum_AVCodecID)(c)))
}

func (c CodecID) String() string {
	return c.Name()
}
