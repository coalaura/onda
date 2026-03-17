package types

type KindID uint16
type TypeID uint16

const (
	KindUnknown KindID = iota
	Kind7ZipArchive
	KindAACAudio
	KindACEArchive
	KindAcrobatFormsDataFormat
	KindAdobeInDesignDocument
	KindALZArchive
	KindAMRAudio
	KindAMRWBAudio
	KindAndroidART
	KindAndroidBootImage
	KindAndroidOAT
	KindAndroidODEX
	KindAndroidSparseImage
	KindAndroidVDEX
	KindApacheArrowFile
	KindApacheParquet
	KindAPFSFilesystem
	KindAppImage
	KindAppleBinaryPropertyList
	KindAppleDesktopServicesStore
	KindAppleDouble
	KindAppleiWorkDocument
	KindAppleSingle
	KindARArchive
	KindARJArchive
	KindASFContainer
	KindASTCTexture
	KindAUAudio
	KindAutoCADDrawing
	KindAvroObjectContainer
	KindBAMData
	KindBcachefsFilesystem
	KindBinkVideo
	KindBitLockerDiskEncryption
	KindBlenderFile
	KindBMPImage
	KindBPGImage
	KindBtrfsFilesystem
	KindBzip2Archive
	KindCabinetArchive
	KindCAFAudio
	KindCanonRAWImage
	KindCBORData
	KindCHMDocument
	KindCineonImage
	KindCPIOArchive
	KindCRAMData
	KindCramFS
	KindCRXBrowserExtension
	KindDalvikExecutable
	KindDDSImage
	KindDebianPackage
	KindDICOMMedicalImage
	KindDjVuDocument
	KindDOSExecutable
	KindDPXImage
	KindDSDIFFAudio
	KindDSFAudio
	KindDTSAudio
	KindEBMLContainer
	KindEncapsulatedPostScript
	KindEOTFont
	KindESRIShapefile
	KindExecutableAndLinkableFormat
	KindExFATFilesystem
	KindExtFilesystem
	KindF2FSFilesystem
	KindFarbfeldImage
	KindFastTrackerModule
	KindFBXModel
	KindFilesystemEntry
	KindFITSAstronomicalImage
	KindFLACAudio
	KindFLIFImage
	KindFLVVideo
	KindFujifilmRAWImage
	KindGameBoyAdvanceROM
	KindGameBoyROM
	KindGameCubeROM
	KindGIFImage
	KindGIMPXCFImage
	KindGitIndex
	KindGitPack
	KindGLTFBinary
	KindGRIBData
	KindGzipArchive
	KindGzipData
	KindHDF5Data
	KindHFSPlusFilesystem
	KindHTMLDocument
	KindICalendar
	KindICCProfile
	KindICNSIcon
	KindICOCURImage
	KindIFFContainer
	KindImpulseTrackerModule
	KindISO9660Image
	KindISOBaseMedia
	KindIVFVideo
	KindJavaClass
	KindJavaKeyStore
	KindJavaModule
	KindJFFS2Filesystem
	KindJNGImage
	KindJPEG2000Image
	KindJPEGImage
	KindJPEGXLImage
	KindJPEGXRImage
	KindKeePassDatabase
	KindKTXTexture
	KindLHAArchive
	KindLLVMBitcode
	KindLuaBytecode
	KindLUKSDiskEncryption
	KindLZ4Frame
	KindLZFSEData
	KindLZIPArchive
	KindLZMAData
	KindLZOPArchive
	KindM3U8Playlist
	KindMacBinary
	KindMachOBinary
	KindMachOUniversalBinary
	KindMacOSAlias
	KindMaterialExchangeFormat
	KindMayaASCII
	KindMayaBinary
	KindMetafileImage
	KindMicrosoftAccessDatabase
	KindMicrosoftOutlookEmailFolder
	KindMIDISequence
	KindMNGImage
	KindMOBIDocument
	KindMonkeysAudio
	KindMPEG2TransportStream
	KindMPEGAudio
	KindMPEGAudioFrame
	KindMPEGProgramStream
	KindMPEGTransportStream
	KindMPEGVideo
	KindMusepackAudio
	KindNESROM
	KindNetCDFData
	KindNetpbmImage
	KindNILFS2Filesystem
	KindNintendo3DSROM
	KindNintendo64ROM
	KindNintendoDSROM
	KindNintendoSwitchPackage
	KindNintendoSwitchROM
	KindNTFSFilesystem
	KindOggContainer
	KindOLECompoundDocument
	KindOlympusRAWImage
	KindOpenEXRImage
	KindOpenTypeFont
	KindOptimFROGAudio
	KindORCColumnarData
	KindOutlookExpressDatabase
	KindPCAPCapture
	KindPCAPNGCapture
	KindPCXImage
	KindPDFDocument
	KindPEMCertificate
	KindPEMPrivateKey
	KindPGPMessage
	KindPhotoshopDocument
	KindPKCS12
	KindPlayStationPortableISO
	KindPLYModel
	KindPNGImage
	KindPortableExecutable
	KindPostgreSQLCustomDump
	KindPostScriptDocument
	KindPuttyPrivateKey
	KindPVRTexture
	KindPythonBytecode
	KindQCOWDiskImage
	KindQOIImage
	KindRadianceHDRImage
	KindRARArchive
	KindRealMedia
	KindRedisDatabase
	KindRichTextFormatDocument
	KindRIFFContainer
	KindRKAudio
	KindROMFS
	KindRPMPackage
	KindRubyGemPackage
	KindScreamTrackerModule
	KindSegaMegaDriveROM
	KindSGIImage
	KindShebangScript
	KindShockwaveFlash
	KindSketchDocument
	KindSketchUpModel
	KindSmackerVideo
	KindSnappyFramedData
	KindSQLite3WriteAheadLog
	KindSQLiteDatabase
	KindSquashFSFilesystem
	KindStuffItArchive
	KindSunRasterImage
	KindSVGImage
	KindSymbianInstallationFormat
	KindTAKAudio
	KindTARArchive
	KindTextFile
	KindTIFFImage
	KindTimezoneData
	KindTorrentFile
	KindTrueTypeCollection
	KindTrueTypeFont
	KindTTAAudio
	KindU3DModel
	KindUBIFSFilesystem
	KindUBootImage
	KindUniversalDiskFormat
	KindUnixCompressArchive
	KindValveTextureFormat
	KindVCard
	KindVHDDiskImage
	KindVHDXDiskImage
	KindVirtualBoxDiskImage
	KindVMwareDiskImage
	KindVOCAudio
	KindWADArchive
	KindWARCFile
	KindWavPackAudio
	KindWebAssemblyModule
	KindWebVTT
	KindWiiBackupFileSystem
	KindWiiROM
	KindWindowsEventLog
	KindWindowsImagingFormat
	KindWindowsRegistryHive
	KindWindowsShortcut
	KindWOFF2Font
	KindWOFFFont
	KindWTVVideo
	KindXARArchive
	KindXboxISO
	KindXFSFilesystem
	KindXMLDocument
	KindXPMImage
	KindXZArchive
	KindZFSFilesystem
	KindZIPArchive
	KindZlibData
	KindZOOArchive
	KindZPAQArchive
	KindZstandardArchive
	KindZstandardDictionary
	KindZXTape
)

const (
	TypeNone TypeID = iota
	Type32BitBigEndian
	Type32BitLittleEndian
	Type3G2Media
	Type3GPPMedia
	Type3MFDocument
	Type64Bit
	Type64BitBigEndian
	Type64BitLittleEndian
	TypeAC3
	TypeAdobeDNGDNG
	TypeADTS
	TypeAIFCAudio
	TypeAIFFAudio
	TypeAndroidAppBundleAAB
	TypeAndroidArchiveAAR
	TypeAndroidPackageAPK
	TypeAndroidPackageXAPK
	TypeAndroidSplitAPKS
	TypeAndroidSystemPackageAPEX
	TypeAPPXPackage
	TypeASCIIText
	TypeAVIFImage
	TypeAVIFImageSequence
	TypeAVIVideo
	TypeBGZF
	TypeBigEndian
	TypeBinaryBigEndian
	TypeBinaryLittleEndian
	TypeBlockDevice
	TypeByteSwapped
	TypeCanonRAW3CR3
	TypeCDAAudio
	TypeCharacterDevice
	TypeCodestream
	TypeCondaPackage
	TypeContainer
	TypeCorelDRAWDocumentCDR
	TypeCRXVersion2
	TypeCRXVersion3
	TypeDEX035
	TypeDEX036
	TypeDEX037
	TypeDEX038
	TypeDEX039
	TypeDEX040
	TypeDEX041
	TypeDirectory
	TypeDownloadableSounds
	TypeEAC3
	TypeELF
	TypeELF32
	TypeELF32BigEndian
	TypeELF32LittleEndian
	TypeELF64
	TypeELF64BigEndian
	TypeELF64LittleEndian
	TypeEmpty
	TypeEnhancedMetafileEMF
	TypeEPUBDocument
	TypeExt2
	TypeExt3
	TypeExt4
	TypeF4VVideo
	TypeFirefoxExtensionXPI
	TypeFLACAudio
	TypeGIF87a
	TypeGIF89a
	TypeHEIFImage
	TypeIOSApplicationArchiveIPA
	TypeIWAD
	TypeJavaArchiveJAR
	TypeJavaEnterpriseArchiveEAR
	TypeJavaWebArchiveWAR
	TypeJMOD
	TypeKDB
	TypeKDBX
	TypeKMZArchive
	TypeKTX
	TypeKTX2
	TypeLittleEndian
	TypeLZ4Legacy
	TypeLZMACompressed
	TypeM2TS
	TypeM2TSBDAV
	TypeM4VVideo
	TypeMatroska
	TypeMicrosoftExcelAddInXLAM
	TypeMicrosoftExcelMacroEnabledTemplateXLTM
	TypeMicrosoftExcelMacroEnabledWorkbookXLSM
	TypeMicrosoftExcelSpreadsheetXLSX
	TypeMicrosoftExcelTemplateXLTX
	TypeMicrosoftExcelWorkbookXLS
	TypeMicrosoftInstallerMSI
	TypeMicrosoftOutlookMessageMSG
	TypeMicrosoftPowerPointAddInPPAM
	TypeMicrosoftPowerPointMacroEnabledPresentationPPTM
	TypeMicrosoftPowerPointMacroEnabledSlideshowPPSM
	TypeMicrosoftPowerPointMacroEnabledTemplatePOTM
	TypeMicrosoftPowerPointPresentationPPT
	TypeMicrosoftPowerPointPresentationPPTX
	TypeMicrosoftPowerPointSlideshowPPSX
	TypeMicrosoftPowerPointTemplatePOTX
	TypeMicrosoftWordDocumentDOC
	TypeMicrosoftWordDocumentDOCX
	TypeMicrosoftWordMacroEnabledDocumentDOCM
	TypeMicrosoftWordMacroEnabledTemplateDOTM
	TypeMicrosoftWordTemplateDOTX
	TypeMP3
	TypeMP3ID3Tagged
	TypeMP4Video
	TypeMPEG12Video
	TypeMPEG4AudioM4AFamily
	TypeMPEGLayer2
	TypeMPEGLayer3
	TypeMSIXPackage
	TypeNamedPipe
	TypeNanosecondBigEndian
	TypeNanosecondLittleEndian
	TypeNewASCII
	TypeNewASCIIWithCRC
	TypeNikonRAWNEF
	TypeNpmPackageTarball
	TypeNuGetPackageNUPKG
	TypeOCIImageLayoutTar
	TypeOldASCII
	TypeOpenDocumentPresentationODP
	TypeOpenDocumentSpreadsheetODS
	TypeOpenDocumentTextODT
	TypeOpenRasterImageORA
	TypeOpusAudio
	TypePAM
	TypePanasonicRAWRW2
	TypePBMASCII
	TypePBMBinary
	TypePE32ARM
	TypePE32ARM64
	TypePE32ARMv7
	TypePE32Itanium
	TypePE32PlusARM64
	TypePE32PlusUnknown
	TypePE32PlusX8664
	TypePE32Unknown
	TypePE32X86
	TypePE32X8664
	TypePentaxRAWPEF
	TypePGMASCII
	TypePGMBinary
	TypePPMASCII
	TypePPMBinary
	TypePSB
	TypePSD
	TypePWAD
	TypePythonSourceDistributionSDist
	TypePythonWheelWHL
	TypeQCOW2
	TypeQCPAudio
	TypeQuickTimeMovie
	TypeRAR4
	TypeRAR5
	TypeRSAPrivateKey
	TypeSketchDocument
	TypeSocket
	TypeSonyRAWARW
	TypeSonyRAWSR2
	TypeSoundFont2
	TypeSpanned
	TypeSpecial
	TypeSpeexAudio
	TypeStreamVersion7
	TypeStreamVersion8
	TypeSymbolicLink
	TypeTheoraVideo
	TypeTS
	TypeUncompressed
	TypeUTF8Text
	TypeVisualStudioExtensionVSIX
	TypeVMDK
	TypeVorbisAudio
	TypeWAVAudio
	TypeWebM
	TypeWebPImage
	TypeWindowsCursor
	TypeWindowsIcon
	TypeWindowsMetafileWMF
	TypeWrapper
	TypeZlibCompressed
	TypeZstandardSkinnableFrame
)

var kindNames = [...]string{
	KindUnknown:                     "Unknown",
	Kind7ZipArchive:                 "7-Zip Archive",
	KindAACAudio:                    "AAC Audio",
	KindACEArchive:                  "ACE Archive",
	KindAcrobatFormsDataFormat:      "Acrobat Forms Data Format",
	KindAdobeInDesignDocument:       "Adobe InDesign Document",
	KindALZArchive:                  "ALZ Archive",
	KindAMRAudio:                    "AMR Audio",
	KindAMRWBAudio:                  "AMR-WB Audio",
	KindAndroidART:                  "Android ART",
	KindAndroidBootImage:            "Android Boot Image",
	KindAndroidOAT:                  "Android OAT",
	KindAndroidODEX:                 "Android ODEX",
	KindAndroidSparseImage:          "Android Sparse Image",
	KindAndroidVDEX:                 "Android VDEX",
	KindApacheArrowFile:             "Apache Arrow File",
	KindApacheParquet:               "Apache Parquet",
	KindAPFSFilesystem:              "APFS Filesystem",
	KindAppImage:                    "AppImage",
	KindAppleBinaryPropertyList:     "Apple Binary Property List",
	KindAppleDesktopServicesStore:   "Apple Desktop Services Store",
	KindAppleDouble:                 "AppleDouble File",
	KindAppleiWorkDocument:          "Apple iWork Document",
	KindAppleSingle:                 "AppleSingle File",
	KindARArchive:                   "AR Archive",
	KindARJArchive:                  "ARJ Archive",
	KindASFContainer:                "ASF Container",
	KindASTCTexture:                 "ASTC Texture",
	KindAUAudio:                     "AU Audio",
	KindAutoCADDrawing:              "AutoCAD Drawing",
	KindAvroObjectContainer:         "Avro Object Container",
	KindBAMData:                     "BAM Data",
	KindBcachefsFilesystem:          "Bcachefs Filesystem",
	KindBinkVideo:                   "Bink Video",
	KindBitLockerDiskEncryption:     "BitLocker Disk Encryption",
	KindBlenderFile:                 "Blender File",
	KindBMPImage:                    "BMP Image",
	KindBPGImage:                    "BPG Image",
	KindBtrfsFilesystem:             "Btrfs Filesystem",
	KindBzip2Archive:                "Bzip2 Archive",
	KindCabinetArchive:              "Cabinet Archive",
	KindCAFAudio:                    "CAF Audio",
	KindCanonRAWImage:               "Canon RAW Image",
	KindCBORData:                    "CBOR Data",
	KindCHMDocument:                 "CHM Document",
	KindCineonImage:                 "Cineon Image",
	KindCPIOArchive:                 "CPIO Archive",
	KindCRAMData:                    "CRAM Data",
	KindCramFS:                      "CramFS",
	KindCRXBrowserExtension:         "CRX Browser Extension",
	KindDalvikExecutable:            "Dalvik Executable",
	KindDDSImage:                    "DDS Image",
	KindDebianPackage:               "Debian Package",
	KindDICOMMedicalImage:           "DICOM Medical Image",
	KindDjVuDocument:                "DjVu Document",
	KindDOSExecutable:               "DOS Executable",
	KindDPXImage:                    "DPX Image",
	KindDSDIFFAudio:                 "DSDIFF Audio",
	KindDSFAudio:                    "DSF Audio",
	KindDTSAudio:                    "DTS Audio",
	KindEBMLContainer:               "EBML Container",
	KindEncapsulatedPostScript:      "Encapsulated PostScript",
	KindEOTFont:                     "EOT Font",
	KindESRIShapefile:               "ESRI Shapefile",
	KindExecutableAndLinkableFormat: "Executable and Linkable Format",
	KindExFATFilesystem:             "exFAT Filesystem",
	KindExtFilesystem:               "ext Filesystem",
	KindF2FSFilesystem:              "F2FS Filesystem",
	KindFarbfeldImage:               "Farbfeld Image",
	KindFastTrackerModule:           "FastTracker Module",
	KindFBXModel:                    "FBX Model",
	KindFilesystemEntry:             "Filesystem Entry",
	KindFITSAstronomicalImage:       "FITS Astronomical Image",
	KindFLACAudio:                   "FLAC Audio",
	KindFLIFImage:                   "FLIF Image",
	KindFLVVideo:                    "FLV Video",
	KindFujifilmRAWImage:            "Fujifilm RAW Image",
	KindGameBoyAdvanceROM:           "Game Boy Advance ROM",
	KindGameBoyROM:                  "Game Boy ROM",
	KindGameCubeROM:                 "GameCube ROM",
	KindGIFImage:                    "GIF Image",
	KindGIMPXCFImage:                "GIMP XCF Image",
	KindGitIndex:                    "Git Index",
	KindGitPack:                     "Git Pack",
	KindGLTFBinary:                  "glTF Binary",
	KindGRIBData:                    "GRIB Data",
	KindGzipArchive:                 "Gzip Archive",
	KindGzipData:                    "Gzip Data",
	KindHDF5Data:                    "HDF5 Data",
	KindHFSPlusFilesystem:           "HFS+ Filesystem",
	KindHTMLDocument:                "HTML Document",
	KindICalendar:                   "iCalendar",
	KindICCProfile:                  "ICC Profile",
	KindICNSIcon:                    "ICNS Icon",
	KindICOCURImage:                 "ICO/CUR Image",
	KindIFFContainer:                "IFF Container",
	KindImpulseTrackerModule:        "Impulse Tracker Module",
	KindISO9660Image:                "ISO 9660 Image",
	KindISOBaseMedia:                "ISO Base Media",
	KindIVFVideo:                    "IVF Video",
	KindJavaClass:                   "Java Class",
	KindJavaKeyStore:                "Java KeyStore",
	KindJavaModule:                  "Java Module",
	KindJFFS2Filesystem:             "JFFS2 Filesystem",
	KindJNGImage:                    "JNG Image",
	KindJPEG2000Image:               "JPEG 2000 Image",
	KindJPEGImage:                   "JPEG Image",
	KindJPEGXLImage:                 "JPEG XL Image",
	KindJPEGXRImage:                 "JPEG XR Image",
	KindKeePassDatabase:             "KeePass Database",
	KindKTXTexture:                  "KTX Texture",
	KindLHAArchive:                  "LHA Archive",
	KindLLVMBitcode:                 "LLVM Bitcode",
	KindLuaBytecode:                 "Lua Bytecode",
	KindLUKSDiskEncryption:          "LUKS Disk Encryption",
	KindLZ4Frame:                    "LZ4 Frame",
	KindLZFSEData:                   "LZFSE Data",
	KindLZIPArchive:                 "LZIP Archive",
	KindLZMAData:                    "LZMA Data",
	KindLZOPArchive:                 "LZOP Archive",
	KindM3U8Playlist:                "M3U8 Playlist",
	KindMacBinary:                   "MacBinary",
	KindMachOBinary:                 "Mach-O Binary",
	KindMachOUniversalBinary:        "Mach-O Universal Binary",
	KindMacOSAlias:                  "macOS Alias",
	KindMaterialExchangeFormat:      "Material Exchange Format",
	KindMayaASCII:                   "Maya ASCII",
	KindMayaBinary:                  "Maya Binary",
	KindMetafileImage:               "Metafile Image",
	KindMicrosoftAccessDatabase:     "Microsoft Access Database",
	KindMicrosoftOutlookEmailFolder: "Microsoft Outlook Email Folder",
	KindMIDISequence:                "MIDI Sequence",
	KindMNGImage:                    "MNG Image",
	KindMOBIDocument:                "MOBI Document",
	KindMonkeysAudio:                "Monkey's Audio",
	KindMPEG2TransportStream:        "MPEG Transport Stream",
	KindMPEGAudio:                   "MPEG Audio",
	KindMPEGAudioFrame:              "MPEG Audio",
	KindMPEGProgramStream:           "MPEG Program Stream",
	KindMPEGTransportStream:         "MPEG Transport Stream",
	KindMPEGVideo:                   "MPEG Video",
	KindMusepackAudio:               "Musepack Audio",
	KindNESROM:                      "NES ROM",
	KindNetCDFData:                  "NetCDF Data",
	KindNetpbmImage:                 "Netpbm Image",
	KindNILFS2Filesystem:            "NILFS2 Filesystem",
	KindNintendo3DSROM:              "Nintendo 3DS ROM",
	KindNintendo64ROM:               "Nintendo 64 ROM",
	KindNintendoDSROM:               "Nintendo DS ROM",
	KindNintendoSwitchPackage:       "Nintendo Switch Package",
	KindNintendoSwitchROM:           "Nintendo Switch ROM",
	KindNTFSFilesystem:              "NTFS Filesystem",
	KindOggContainer:                "Ogg Container",
	KindOLECompoundDocument:         "OLE Compound Document",
	KindOlympusRAWImage:             "Olympus RAW Image",
	KindOpenEXRImage:                "OpenEXR Image",
	KindOpenTypeFont:                "OpenType Font",
	KindOptimFROGAudio:              "OptimFROG Audio",
	KindORCColumnarData:             "ORC Columnar Data",
	KindOutlookExpressDatabase:      "Outlook Express Database",
	KindPCAPCapture:                 "PCAP Capture",
	KindPCAPNGCapture:               "PCAPNG Capture",
	KindPCXImage:                    "PCX Image",
	KindPDFDocument:                 "PDF Document",
	KindPEMCertificate:              "PEM Certificate",
	KindPEMPrivateKey:               "PEM Private Key",
	KindPGPMessage:                  "PGP Message",
	KindPhotoshopDocument:           "Photoshop Document",
	KindPKCS12:                      "PKCS#12",
	KindPlayStationPortableISO:      "PlayStation Portable ISO",
	KindPLYModel:                    "PLY Model",
	KindPNGImage:                    "PNG Image",
	KindPortableExecutable:          "Portable Executable",
	KindPostgreSQLCustomDump:        "PostgreSQL Custom Dump",
	KindPostScriptDocument:          "PostScript Document",
	KindPuttyPrivateKey:             "PuTTY Private Key",
	KindPVRTexture:                  "PVR Texture",
	KindPythonBytecode:              "Python Bytecode",
	KindQCOWDiskImage:               "QCOW Disk Image",
	KindQOIImage:                    "QOI Image",
	KindRadianceHDRImage:            "Radiance HDR Image",
	KindRARArchive:                  "RAR Archive",
	KindRealMedia:                   "RealMedia",
	KindRedisDatabase:               "Redis Database",
	KindRichTextFormatDocument:      "Rich Text Format Document",
	KindRIFFContainer:               "RIFF Container",
	KindRKAudio:                     "RKAudio",
	KindROMFS:                       "ROMFS",
	KindRPMPackage:                  "RPM Package",
	KindRubyGemPackage:              "RubyGem Package",
	KindScreamTrackerModule:         "Scream Tracker Module",
	KindSegaMegaDriveROM:            "Sega Mega Drive ROM",
	KindSGIImage:                    "SGI Image",
	KindShebangScript:               "Shebang Script",
	KindShockwaveFlash:              "Shockwave Flash",
	KindSketchDocument:              "Sketch Document",
	KindSketchUpModel:               "SketchUp Model",
	KindSmackerVideo:                "Smacker Video",
	KindSnappyFramedData:            "Snappy Framed Data",
	KindSQLite3WriteAheadLog:        "SQLite Write-Ahead Log",
	KindSQLiteDatabase:              "SQLite Database",
	KindSquashFSFilesystem:          "SquashFS Filesystem",
	KindStuffItArchive:              "StuffIt Archive",
	KindSunRasterImage:              "Sun Raster Image",
	KindSVGImage:                    "SVG Image",
	KindSymbianInstallationFormat:   "Symbian Installation Format",
	KindTAKAudio:                    "TAK Audio",
	KindTARArchive:                  "TAR Archive",
	KindTextFile:                    "Text File",
	KindTIFFImage:                   "TIFF Image",
	KindTimezoneData:                "Timezone Data",
	KindTorrentFile:                 "Torrent File",
	KindTrueTypeCollection:          "TrueType Collection",
	KindTrueTypeFont:                "TrueType Font",
	KindTTAAudio:                    "TTA Audio",
	KindU3DModel:                    "U3D Model",
	KindUBIFSFilesystem:             "UBIFS Filesystem",
	KindUBootImage:                  "U-Boot Image",
	KindUniversalDiskFormat:         "Universal Disk Format",
	KindUnixCompressArchive:         "Unix Compress Archive",
	KindValveTextureFormat:          "Valve Texture Format",
	KindVCard:                       "vCard",
	KindVHDDiskImage:                "VHD Disk Image",
	KindVHDXDiskImage:               "VHDX Disk Image",
	KindVirtualBoxDiskImage:         "VirtualBox Disk Image",
	KindVMwareDiskImage:             "VMware Disk Image",
	KindVOCAudio:                    "VOC Audio",
	KindWADArchive:                  "WAD Archive",
	KindWARCFile:                    "WARC File",
	KindWavPackAudio:                "WavPack Audio",
	KindWebAssemblyModule:           "WebAssembly Module",
	KindWebVTT:                      "WebVTT",
	KindWiiBackupFileSystem:         "Wii Backup File System",
	KindWiiROM:                      "Wii ROM",
	KindWindowsEventLog:             "Windows Event Log",
	KindWindowsImagingFormat:        "Windows Imaging Format",
	KindWindowsRegistryHive:         "Windows Registry Hive",
	KindWindowsShortcut:             "Windows Shortcut",
	KindWOFF2Font:                   "WOFF2 Font",
	KindWOFFFont:                    "WOFF Font",
	KindWTVVideo:                    "WTV Video",
	KindXARArchive:                  "XAR Archive",
	KindXboxISO:                     "Xbox ISO",
	KindXFSFilesystem:               "XFS Filesystem",
	KindXMLDocument:                 "XML Document",
	KindXPMImage:                    "XPM Image",
	KindXZArchive:                   "XZ Archive",
	KindZFSFilesystem:               "ZFS Filesystem",
	KindZIPArchive:                  "ZIP Archive",
	KindZlibData:                    "Zlib Data",
	KindZOOArchive:                  "ZOO Archive",
	KindZPAQArchive:                 "ZPAQ Archive",
	KindZstandardArchive:            "Zstandard Archive",
	KindZstandardDictionary:         "Zstandard Dictionary",
	KindZXTape:                      "ZX Spectrum Tape",
}

var typeNames = [...]string{
	TypeNone:                     "",
	Type32BitBigEndian:           "32-bit Big-Endian",
	Type32BitLittleEndian:        "32-bit Little-Endian",
	Type3G2Media:                 "3G2 Media",
	Type3GPPMedia:                "3GPP Media",
	Type3MFDocument:              "3MF Document",
	Type64Bit:                    "64-bit",
	Type64BitBigEndian:           "64-bit Big-Endian",
	Type64BitLittleEndian:        "64-bit Little-Endian",
	TypeAC3:                      "AC-3",
	TypeAdobeDNGDNG:              "Adobe DNG (DNG)",
	TypeADTS:                     "ADTS",
	TypeAIFCAudio:                "AIFC Audio",
	TypeAIFFAudio:                "AIFF Audio",
	TypeAndroidAppBundleAAB:      "Android App Bundle (AAB)",
	TypeAndroidArchiveAAR:        "Android Archive (AAR)",
	TypeAndroidPackageAPK:        "Android Package (APK)",
	TypeAndroidPackageXAPK:       "Android Package (XAPK)",
	TypeAndroidSplitAPKS:         "Android Split APK Set (APKS)",
	TypeAndroidSystemPackageAPEX: "Android System Package (APEX)",
	TypeAPPXPackage:              "APPX Package",
	TypeASCIIText:                "ASCII Text",
	TypeAVIFImage:                "AVIF Image",
	TypeAVIFImageSequence:        "AVIF Image Sequence",
	TypeAVIVideo:                 "AVI Video",
	TypeBGZF:                     "BGZF",
	TypeBigEndian:                "Big-Endian",
	TypeBinaryBigEndian:          "Binary Big-Endian",
	TypeBinaryLittleEndian:       "Binary Little-Endian",
	TypeBlockDevice:              "Block Device",
	TypeByteSwapped:              "Byte-Swapped",
	TypeCanonRAW3CR3:             "Canon RAW 3 (CR3)",
	TypeCDAAudio:                 "CD Audio",
	TypeCharacterDevice:          "Character Device",
	TypeCodestream:               "Codestream",
	TypeCondaPackage:             "Conda Package",
	TypeContainer:                "Container",
	TypeCorelDRAWDocumentCDR:     "CorelDRAW Document (CDR)",
	TypeCRXVersion2:              "Version 2",
	TypeCRXVersion3:              "Version 3",
	TypeDEX035:                   "DEX 035",
	TypeDEX036:                   "DEX 036",
	TypeDEX037:                   "DEX 037",
	TypeDEX038:                   "DEX 038",
	TypeDEX039:                   "DEX 039",
	TypeDEX040:                   "DEX 040",
	TypeDEX041:                   "DEX 041",
	TypeDirectory:                "Directory",
	TypeDownloadableSounds:       "Downloadable Sounds",
	TypeEAC3:                     "E-AC-3",
	TypeELF:                      "ELF",
	TypeELF32:                    "ELF32",
	TypeELF32BigEndian:           "ELF32 Big-Endian",
	TypeELF32LittleEndian:        "ELF32 Little-Endian",
	TypeELF64:                    "ELF64",
	TypeELF64BigEndian:           "ELF64 Big-Endian",
	TypeELF64LittleEndian:        "ELF64 Little-Endian",
	TypeEmpty:                    "Empty",
	TypeEnhancedMetafileEMF:      "Enhanced Metafile (EMF)",
	TypeEPUBDocument:             "EPUB Document",
	TypeExt2:                     "ext2",
	TypeExt3:                     "ext3",
	TypeExt4:                     "ext4",
	TypeF4VVideo:                 "F4V Video",
	TypeFirefoxExtensionXPI:      "Firefox Extension (XPI)",
	TypeFLACAudio:                "FLAC Audio",
	TypeGIF87a:                   "GIF87a",
	TypeGIF89a:                   "GIF89a",
	TypeHEIFImage:                "HEIF Image",
	TypeIOSApplicationArchiveIPA: "iOS Application Archive (IPA)",
	TypeIWAD:                     "IWAD",
	TypeJavaArchiveJAR:           "Java Archive (JAR)",
	TypeJavaEnterpriseArchiveEAR: "Java Enterprise Archive (EAR)",
	TypeJavaWebArchiveWAR:        "Java Web Archive (WAR)",
	TypeJMOD:                     "JMOD",
	TypeKDB:                      "KDB",
	TypeKDBX:                     "KDBX",
	TypeKMZArchive:               "KMZ Archive",
	TypeKTX:                      "KTX",
	TypeKTX2:                     "KTX2",
	TypeLittleEndian:             "Little-Endian",
	TypeLZ4Legacy:                "LZ4 Legacy",
	TypeLZMACompressed:           "LZMA Compressed",
	TypeM2TS:                     "M2TS",
	TypeM2TSBDAV:                 "M2TS/BDAV",
	TypeM4VVideo:                 "M4V Video",
	TypeMatroska:                 "Matroska",
	TypeMicrosoftExcelAddInXLAM:  "Microsoft Excel Add-In (XLAM)",
	TypeMicrosoftExcelMacroEnabledTemplateXLTM:          "Microsoft Excel Macro-Enabled Template (XLTM)",
	TypeMicrosoftExcelMacroEnabledWorkbookXLSM:          "Microsoft Excel Macro-Enabled Workbook (XLSM)",
	TypeMicrosoftExcelSpreadsheetXLSX:                   "Microsoft Excel Spreadsheet (XLSX)",
	TypeMicrosoftExcelTemplateXLTX:                      "Microsoft Excel Template (XLTX)",
	TypeMicrosoftExcelWorkbookXLS:                       "Microsoft Excel Workbook (XLS)",
	TypeMicrosoftInstallerMSI:                           "Microsoft Installer (MSI)",
	TypeMicrosoftOutlookMessageMSG:                      "Microsoft Outlook Message (MSG)",
	TypeMicrosoftPowerPointAddInPPAM:                    "Microsoft PowerPoint Add-In (PPAM)",
	TypeMicrosoftPowerPointMacroEnabledPresentationPPTM: "Microsoft PowerPoint Macro-Enabled Presentation (PPTM)",
	TypeMicrosoftPowerPointMacroEnabledSlideshowPPSM:    "Microsoft PowerPoint Macro-Enabled Slideshow (PPSM)",
	TypeMicrosoftPowerPointMacroEnabledTemplatePOTM:     "Microsoft PowerPoint Macro-Enabled Template (POTM)",
	TypeMicrosoftPowerPointPresentationPPT:              "Microsoft PowerPoint Presentation (PPT)",
	TypeMicrosoftPowerPointPresentationPPTX:             "Microsoft PowerPoint Presentation (PPTX)",
	TypeMicrosoftPowerPointSlideshowPPSX:                "Microsoft PowerPoint Slideshow (PPSX)",
	TypeMicrosoftPowerPointTemplatePOTX:                 "Microsoft PowerPoint Template (POTX)",
	TypeMicrosoftWordDocumentDOC:                        "Microsoft Word Document (DOC)",
	TypeMicrosoftWordDocumentDOCX:                       "Microsoft Word Document (DOCX)",
	TypeMicrosoftWordMacroEnabledDocumentDOCM:           "Microsoft Word Macro-Enabled Document (DOCM)",
	TypeMicrosoftWordMacroEnabledTemplateDOTM:           "Microsoft Word Macro-Enabled Template (DOTM)",
	TypeMicrosoftWordTemplateDOTX:                       "Microsoft Word Template (DOTX)",
	TypeMP3:                                             "MP3",
	TypeMP3ID3Tagged:                                    "MP3 (ID3 Tagged)",
	TypeMP4Video:                                        "MP4 Video",
	TypeMPEG12Video:                                     "MPEG-1/2 Video",
	TypeMPEG4AudioM4AFamily:                             "MPEG-4 Audio (M4A Family)",
	TypeMPEGLayer2:                                      "MPEG Layer II",
	TypeMPEGLayer3:                                      "MPEG Layer III",
	TypeMSIXPackage:                                     "MSIX Package",
	TypeNamedPipe:                                       "Named Pipe",
	TypeNanosecondBigEndian:                             "Nanosecond Big-Endian",
	TypeNanosecondLittleEndian:                          "Nanosecond Little-Endian",
	TypeNewASCII:                                        "New ASCII",
	TypeNewASCIIWithCRC:                                 "New ASCII with CRC",
	TypeNikonRAWNEF:                                     "Nikon RAW (NEF)",
	TypeNpmPackageTarball:                               "npm Package Tarball",
	TypeNuGetPackageNUPKG:                               "NuGet Package (NUPKG)",
	TypeOCIImageLayoutTar:                               "OCI Image Layout (TAR)",
	TypeOldASCII:                                        "Old ASCII",
	TypeOpenDocumentPresentationODP:                     "OpenDocument Presentation (ODP)",
	TypeOpenDocumentSpreadsheetODS:                      "OpenDocument Spreadsheet (ODS)",
	TypeOpenDocumentTextODT:                             "OpenDocument Text (ODT)",
	TypeOpenRasterImageORA:                              "OpenRaster Image (ORA)",
	TypeOpusAudio:                                       "Opus Audio",
	TypePAM:                                             "PAM",
	TypePanasonicRAWRW2:                                 "Panasonic RAW (RW2)",
	TypePBMASCII:                                        "PBM ASCII",
	TypePBMBinary:                                       "PBM binary",
	TypePE32ARM:                                         "PE32 ARM",
	TypePE32ARM64:                                       "PE32 ARM64",
	TypePE32ARMv7:                                       "PE32 ARMv7",
	TypePE32Itanium:                                     "PE32 Itanium",
	TypePE32PlusARM64:                                   "PE32+ ARM64",
	TypePE32PlusUnknown:                                 "PE32+ Unknown",
	TypePE32PlusX8664:                                   "PE32+ x86-64",
	TypePE32Unknown:                                     "PE32 Unknown",
	TypePE32X86:                                         "PE32 x86",
	TypePE32X8664:                                       "PE32 x86-64",
	TypePentaxRAWPEF:                                    "Pentax RAW (PEF)",
	TypePGMASCII:                                        "PGM ASCII",
	TypePGMBinary:                                       "PGM binary",
	TypePPMASCII:                                        "PPM ASCII",
	TypePPMBinary:                                       "PPM binary",
	TypePSB:                                             "PSB",
	TypePSD:                                             "PSD",
	TypePWAD:                                            "PWAD",
	TypePythonSourceDistributionSDist:                   "Python Source Distribution (sdist)",
	TypePythonWheelWHL:                                  "Python Wheel (WHL)",
	TypeQCOW2:                                           "QCOW2",
	TypeQCPAudio:                                        "QCP Audio",
	TypeQuickTimeMovie:                                  "QuickTime Movie",
	TypeRAR4:                                            "RAR4",
	TypeRAR5:                                            "RAR5",
	TypeRSAPrivateKey:                                   "RSA Private Key",
	TypeSketchDocument:                                  "Sketch Document",
	TypeSocket:                                          "Socket",
	TypeSonyRAWARW:                                      "Sony RAW (ARW)",
	TypeSonyRAWSR2:                                      "Sony RAW (SR2)",
	TypeSoundFont2:                                      "SoundFont 2",
	TypeSpanned:                                         "Spanned",
	TypeSpecial:                                         "Special",
	TypeSpeexAudio:                                      "Speex Audio",
	TypeStreamVersion7:                                  "Stream Version 7",
	TypeStreamVersion8:                                  "Stream Version 8",
	TypeSymbolicLink:                                    "Symbolic Link",
	TypeTheoraVideo:                                     "Theora Video",
	TypeTS:                                              "TS",
	TypeUncompressed:                                    "Uncompressed",
	TypeUTF8Text:                                        "UTF-8 Text",
	TypeVisualStudioExtensionVSIX:                       "Visual Studio Extension (VSIX)",
	TypeVMDK:                                            "VMDK",
	TypeVorbisAudio:                                     "Vorbis Audio",
	TypeWAVAudio:                                        "WAV Audio",
	TypeWebM:                                            "WebM",
	TypeWebPImage:                                       "WebP Image",
	TypeWindowsCursor:                                   "Windows Cursor",
	TypeWindowsIcon:                                     "Windows Icon",
	TypeWindowsMetafileWMF:                              "Windows Metafile (WMF)",
	TypeWrapper:                                         "Wrapper",
	TypeZlibCompressed:                                  "Zlib Compressed",
	TypeZstandardSkinnableFrame:                         "Zstandard Skinnable Frame",
}

func (k KindID) String() string {
	if int(k) >= 0 && int(k) < len(kindNames) {
		name := kindNames[k]
		if name != "" {
			return name
		}
	}

	return "Unknown"
}

func (t TypeID) String() string {
	if int(t) >= 0 && int(t) < len(typeNames) {
		return typeNames[t]
	}

	return ""
}
