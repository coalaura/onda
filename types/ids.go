package types

type KindID uint16
type TypeID uint16

const (
	KindUnknown KindID = iota
	Kind7ZipArchive
	KindAACAudio
	KindACEArchive
	KindAcrobatFormsDataFormat
	KindAcronisTrueImage
	KindAdobeInDesignDocument
	KindAdvancedForensicFormat
	KindALZArchive
	KindAmigaDiskFile
	KindAmigaHardDiskFile
	KindAmigaHunkExecutable
	KindAmigaOctaMED
	KindAMRAudio
	KindAMRWBAudio
	KindAndroidART
	KindAndroidBackup
	KindAndroidBinaryXML
	KindAndroidBootImage
	KindAndroidCompiledResources
	KindAndroidOAT
	KindAndroidODEX
	KindAndroidSparseImage
	KindAndroidVDEX
	KindApacheArrowFile
	KindApacheParquet
	KindAPFSFilesystem
	KindAppImage
	KindAppleArchive
	KindAppleBillOfMaterials
	KindAppleBinaryCrashReport
	KindAppleBinaryPropertyList
	KindAppleDesktopServicesStore
	KindAppleDiskImage
	KindAppleDouble
	KindAppleiWorkDocument
	KindAppleKeychain
	KindAppleSingle
	KindAppleSystemLog
	KindAppleTexture
	KindAppleXMLPropertyList
	KindARArchive
	KindARCArchive
	KindARJArchive
	KindARRIRAWImage
	KindASFContainer
	KindASTCTexture
	KindAtari7800ROM
	KindAtariLynxROM
	KindAtomFeed
	KindAUAudio
	KindAudacityBlockFile
	KindAutoCADDrawing
	KindAutoCADDXF
	KindAvroObjectContainer
	KindBAMData
	KindBcachefsFilesystem
	KindBCFData
	KindBerkeleyDB
	KindBinHex
	KindBinkVideo
	KindBitcoinBlock
	KindBitLockerDiskEncryption
	KindBlenderFile
	KindBlizzardTexture
	KindBluetoothSnoop
	KindBMPImage
	KindBochsDiskImage
	KindBouncyCastleKeystore
	KindBPGImage
	KindBroadbandEBook
	KindBtrfsFilesystem
	KindBUFR
	KindBzip2Archive
	KindBzip3Archive
	KindCabinetArchive
	KindCAFAudio
	KindCALSImage
	KindCanonRAWImage
	KindCBORData
	KindCHMDocument
	KindCinema4DModel
	KindCineonImage
	KindCloneCDControl
	KindComicBookArchive
	KindCommodore64Tape
	KindCommodore64TapeRAW
	KindCommodoreSID
	KindCompiledTerminfo
	KindCompressedSquareWave
	KindCorelPhotoPaintImage
	KindCPIOArchive
	KindCRAMData
	KindCramFS
	KindCRXBrowserExtension
	KindDalvikExecutable
	KindDDSImage
	KindDebianPackage
	KindDICOMMedicalImage
	KindDiracVideo
	KindDjVuDocument
	KindDolphinRVZ
	KindDOSExecutable
	KindDPXImage
	KindDSDIFFAudio
	KindDSFAudio
	KindDTSAudio
	KindDVRMSVideo
	KindEBMLContainer
	KindEGGArchive
	KindEncapsulatedPostScript
	KindEnCaseImage
	KindEOTFont
	KindErlangBEAM
	KindESRIShapefile
	KindExecutableAndLinkableFormat
	KindExFATFilesystem
	KindExtensibleMusicFormat
	KindExtFilesystem
	KindF2FSFilesystem
	KindFamicomDiskSystemROM
	KindFarbfeldImage
	KindFastTracker2ExtendedInstrument
	KindFastTrackerModule
	KindFATFilesystem
	KindFBXModel
	KindFeatherData
	KindFilesystemEntry
	KindFITActivity
	KindFITSAstronomicalImage
	KindFLACAudio
	KindFLIFImage
	KindFLStudioProject
	KindFLVVideo
	KindFreeArcArchive
	KindFujifilmRAWImage
	KindGameBoyAdvanceROM
	KindGameBoyROM
	KindGameBoySound
	KindGameCubeROM
	KindGDBMDatabase
	KindGHCiInterface
	KindGIFImage
	KindGIMPBrush
	KindGIMPPattern
	KindGIMPXCFImage
	KindGitBundle
	KindGitCommitGraph
	KindGitIndex
	KindGitPack
	KindGitPackBitmap
	KindGitPackIndex
	KindGitPackReverseIndex
	KindGlibcLocale
	KindGLTFBinary
	KindGlyphBitmapDistributionFormat
	KindGNOMEKeyring
	KindGNUGettextMachineCatalog
	KindGnuPGKeybox
	KindGodotPackage
	KindGPSExchangeFormat
	KindGRIBData
	KindGUIDPartitionTable
	KindGzipArchive
	KindGzipData
	KindHadoopRCFile
	KindHadoopSequenceFile
	KindHDF4Data
	KindHDF5Data
	KindHFSPlusFilesystem
	KindHTMLDocument
	KindICalendar
	KindICCProfile
	KindICNSIcon
	KindICOCURImage
	KindIFFContainer
	KindImpulseTrackerInstrument
	KindImpulseTrackerModule
	KindImpulseTrackerSample
	KindInstallShieldCAB
	KindISO9660Image
	KindISOBaseMedia
	KindIVFVideo
	KindJavaClass
	KindJavaCryptographyExtensionKeyStore
	KindJavaKeyStore
	KindJavaModule
	KindJavaPack200
	KindJavaSerialization
	KindJBIG2Image
	KindJFFS2Filesystem
	KindJNGImage
	KindJPEG2000Image
	KindJPEGImage
	KindJPEGLSImage
	KindJPEGXLImage
	KindJPEGXRImage
	KindJSONDocument
	KindKDEKWallet
	KindKeePassDatabase
	KindKeyholeMarkupLanguage
	KindKGBArchive
	KindKorgAudio
	KindKTXTexture
	KindKyotoCabinetDatabase
	KindLevelDB
	KindLHAArchive
	KindLidarData
	KindLightWaveScene
	KindLinuxKernelImage
	KindLLVMBitcode
	KindLottieAnimation
	KindLrzipArchive
	KindLuaBytecode
	KindLuceneIndex
	KindLUKSDiskEncryption
	KindLVM2PhysicalVolume
	KindLytroLightField
	KindLZ4Frame
	KindLZFArchive
	KindLZFSEData
	KindLZIPArchive
	KindLZMAData
	KindLZOPArchive
	KindLZXArchive
	KindM3U8Playlist
	KindMacBinary
	KindMachOBinary
	KindMachOUniversalBinary
	KindMacOSAlias
	KindMacOSXCodeSignature
	KindMacriumReflectImage
	KindMagicaVoxel
	KindMAMECHD
	KindMaterialExchangeFormat
	KindMATLABData
	KindMayaASCII
	KindMayaBinary
	KindMBOXEmailFolder
	KindMCAPCapture
	KindMediaDescriptor
	KindMetafileImage
	KindMicrosoftAccessDatabase
	KindMicrosoftCompress
	KindMicrosoftNetworkMonitor
	KindMicrosoftOutlookEmailFolder
	KindMicrosoftProgramDatabase
	KindMicrosoftReader
	KindMicrosoftSQLServerDatabase
	KindMicrosoftXbox360Profile
	KindMicrosoftXboxSave
	KindMIDISequence
	KindMinoltaRAWImage
	KindMNGImage
	KindMOBIDocument
	KindMonkeysAudio
	KindMoPaQArchive
	KindMozillaArchive
	KindMPEG2TransportStream
	KindMPEGAudio
	KindMPEGAudioFrame
	KindMPEGProgramStream
	KindMPEGTransportStream
	KindMPEGVideo
	KindMUGENSound
	KindMUGENSprite
	KindMusepackAudio
	KindMySQLMyISAM
	KindNeoGeoPocketROM
	KindNESROM
	KindNESSoundFormat
	KindNetCDFData
	KindNetpbmImage
	KindNIfTIMedicalImage
	KindNILFS2Filesystem
	KindNintendo3DSROM
	KindNintendo64ROM
	KindNintendoBCFNT
	KindNintendoBFLYT
	KindNintendoBRRES
	KindNintendoDSiROM
	KindNintendoDSROM
	KindNintendoRARC
	KindNintendoSwitchNRO
	KindNintendoSwitchNSO
	KindNintendoSwitchPackage
	KindNintendoSwitchROM
	KindNintendoU8Archive
	KindNintendoWiiUExecutable
	KindNintendoYay0
	KindNintendoYaz0
	KindNISTSPHEREAudio
	KindNTFSFilesystem
	KindNullsoftVideo
	KindNumPyArray
	KindOCamlObject
	KindOggContainer
	KindOLECompoundDocument
	KindOlympusRAWImage
	KindOpenEXRImage
	KindOpenSSHPrivateKey
	KindOpenTypeFont
	KindOptimFROGAudio
	KindORCColumnarData
	KindOutlookExpressDatabase
	KindPaintShopProImage
	KindParallelsDiskImage
	KindPCAPCapture
	KindPCAPNGCapture
	KindPCFFont
	KindPCXImage
	KindPDFDocument
	KindPeaZipArchive
	KindPEMCertificate
	KindPEMPrivateKey
	KindPGFImage
	KindPGPMessage
	KindPhotoshopDocument
	KindPico8Cartridge
	KindPIMArchive
	KindPKCS12
	KindPlaylistFile
	KindPlayStation1Executable
	KindPlayStation2MemoryCard
	KindPlayStationExecutable
	KindPlayStationPortableExecutable
	KindPlayStationPortableISO
	KindPlayStationSoundFormat
	KindPLYModel
	KindPNGImage
	KindPortableExecutable
	KindPostgreSQLCustomDump
	KindPostScriptDocument
	KindPowerISODAA
	KindPSFFont
	KindPuttyPrivateKey
	KindPVRTexture
	KindPythonBytecode
	KindPythonPickle
	KindQCOWDiskImage
	KindQDBMDatabase
	KindQEMUQEDDiskImage
	KindQOIImage
	KindQuakePAK
	KindRadianceHDRImage
	KindRARArchive
	KindRData
	KindRealMedia
	KindRedisDatabase
	KindREDRAWImage
	KindReiserFSFilesystem
	KindRenPyArchive
	KindRhino3DModel
	KindRichTextFormatDocument
	KindRIFFContainer
	KindRKAudio
	KindRNCArchive
	KindROMFS
	KindROOTData
	KindROSBag
	KindRPGMakerArchive
	KindRPMPackage
	KindRSSFeed
	KindRubyGemPackage
	KindRubyMarshal
	KindRzipArchive
	KindSASData
	KindScreamTrackerModule
	KindSegaCDImage
	KindSegaDreamcastROM
	KindSegaMasterSystemROM
	KindSegaMegaDriveROM
	KindSegaNaomiROM
	KindSegaSaturnROM
	KindSGIImage
	KindShebangScript
	KindShockwaveFlash
	KindShortenAudio
	KindSigmaRAWImage
	KindSiliconGraphicsMovie
	KindSketchDocument
	KindSketchUpModel
	KindSmackerVideo
	KindSnappyFramedData
	KindSNESSPC
	KindSnoopCapture
	KindSOAPMessage
	KindSonyOpenMG
	KindSonyWave64Audio
	KindSourceEngineBSP
	KindSPSSData
	KindSPSSPortableData
	KindSQLite3SharedMemory
	KindSQLite3WriteAheadLog
	KindSQLiteDatabase
	KindSquashFSFilesystem
	KindSQXArchive
	KindSSHPublicKey
	KindStataData
	KindSTEP3DModel
	KindStuffItArchive
	KindSubRipText
	KindSunRasterImage
	KindSVGImage
	KindSymantecGhostImage
	KindSymbianExecutable
	KindSymbianInstallationFormat
	KindSystemdJournal
	KindSzipArchive
	KindTAKAudio
	KindTARArchive
	KindTensorFlowLiteModel
	KindTextFile
	KindTIFFImage
	KindTimezoneData
	KindTokyoCabinetDatabase
	KindTorrentFile
	KindTransportNeutralEncapsulationFormat
	KindTrueTypeCollection
	KindTrueTypeFont
	KindTTAAudio
	KindU3DModel
	KindUBIFSFilesystem
	KindUBootImage
	KindUHAArchive
	KindUnityWebData
	KindUniversalDiskFormat
	KindUniversalSceneDescription
	KindUnixCompressArchive
	KindUnrealEnginePackage
	KindValvePak
	KindValveTextureFormat
	KindVCard
	KindVeeamBackup
	KindVHDDiskImage
	KindVHDXDiskImage
	KindVideoGameMusic
	KindVirtualBoxDiskImage
	KindVirtualBoxSavedState
	KindVMwareDiskImage
	KindVMwareNVRAM
	KindVMwareSnapshot
	KindVMwareSuspend
	KindVocaloidProject
	KindVOCAudio
	KindVRML3DModel
	KindVulkanSPIRV
	KindWADArchive
	KindWARCFile
	KindWarcraftIIIReplay
	KindWavPackAudio
	KindWavPackCorrection
	KindWebAssemblyModule
	KindWebVTT
	KindWiiBackupFileSystem
	KindWiiROM
	KindWiiUArchive
	KindWindowsEventLog
	KindWindowsHelp
	KindWindowsHibernation
	KindWindowsImagingFormat
	KindWindowsMinidump
	KindWindowsPrecompiledHeader
	KindWindowsPrefetch
	KindWindowsRegistryHive
	KindWindowsResourceFile
	KindWindowsShortcut
	KindWindowsTypeLibrary
	KindWiredTigerDatabase
	KindWOFF2Font
	KindWOFFFont
	KindWordPerfectGraphics
	KindWTVVideo
	KindXARArchive
	KindXBMImage
	KindXbox360Executable
	KindXboxExecutable
	KindXboxISO
	KindXCOFFExecutable
	KindXFSFilesystem
	KindXilinxBitstream
	KindXMLDocument
	KindXPMImage
	KindXZArchive
	KindYamahaSMAF
	KindYUV4MPEG2Video
	KindZBackup
	KindZchunk
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
	Type32Bit
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
	TypeAfterEffectsProjectAEP
	TypeAIFCAudio
	TypeAIFFAudio
	TypeAndroidAppBundleAAB
	TypeAndroidArchiveAAR
	TypeAndroidPackageAPK
	TypeAndroidPackageXAPK
	TypeAndroidSplitAPKS
	TypeAndroidSystemPackageAPEX
	TypeAPPXPackage
	TypeArchLinuxPackage
	TypeASCIIText
	TypeAVIFImage
	TypeAVIFImageSequence
	TypeAVIVideo
	TypeBGZF
	TypeBigEndian
	TypeBinaryBigEndian
	TypeBinaryLittleEndian
	TypeBlackmagicRAW
	TypeBlockDevice
	TypeByteSwapped
	TypeCanonRAW3CR3
	TypeCB7
	TypeCBR
	TypeCBT
	TypeCBZ
	TypeCDAAudio
	TypeCharacterDevice
	TypeCodestream
	TypeCondaPackage
	TypeContainer
	TypeCorelDRAWDocumentCDR
	TypeCRXVersion2
	TypeCRXVersion3
	TypeCubaseProjectCPR
	TypeDEX035
	TypeDEX036
	TypeDEX037
	TypeDEX038
	TypeDEX039
	TypeDEX040
	TypeDEX041
	TypeDirectory
	TypeDownloadableSounds
	TypeDSAPrivateKey
	TypeEAC3
	TypeECPrivateKey
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
	TypeFabricMod
	TypeFAT12
	TypeFAT16
	TypeFAT32
	TypeFirefoxExtensionXPI
	TypeFLACAudio
	TypeForgeMod
	TypeGIF87a
	TypeGIF89a
	TypeHEIFImage
	TypeILBMImage
	TypeIOSApplicationArchiveIPA
	TypeIWAD
	TypeJavaArchiveJAR
	TypeJavaEnterpriseArchiveEAR
	TypeJavaWebArchiveWAR
	TypeJMOD
	TypeKDB
	TypeKDBX
	TypeKerasModel
	TypeKMZArchive
	TypeKritaDocumentKRA
	TypeKTX
	TypeKTX2
	TypeLittleEndian
	TypeLOVEGame
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
	TypeMicrosoftProjectDocumentMPP
	TypeMicrosoftPublisherDocumentPUB
	TypeMicrosoftVisioDrawingVSD
	TypeMicrosoftWordDocumentDOC
	TypeMicrosoftWordDocumentDOCX
	TypeMicrosoftWordMacroEnabledDocumentDOCM
	TypeMicrosoftWordMacroEnabledTemplateDOTM
	TypeMicrosoftWordTemplateDOTX
	TypeMinecraftResourcePack
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
	TypeOpenDocumentChartODC
	TypeOpenDocumentDatabaseODB
	TypeOpenDocumentFormulaODF
	TypeOpenDocumentGraphicsODG
	TypeOpenDocumentImageODI
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
	TypePyTorchModel
	TypeQCOW1
	TypeQCOW2
	TypeQCPAudio
	TypeQuickTimeMovie
	TypeRAR4
	TypeRAR5
	TypeRIFFMIDI
	TypeRIFFPalette
	TypeRSAPrivateKey
	TypeSketchDocument
	TypeSlackwarePackage
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
	TypeVagrantBox
	TypeVisualStudioExtensionVSIX
	TypeVMDK
	TypeVorbisAudio
	TypeWAVAudio
	TypeWebM
	TypeWebPImage
	TypeWindowsAnimatedCursor
	TypeWindowsCursor
	TypeWindowsIcon
	TypeWindowsLE
	TypeWindowsLX
	TypeWindowsMetafileWMF
	TypeWindowsNE
	TypeWrapper
	TypeXMLPaperSpecificationXPS
	TypeZlibCompressed
	TypeZstandardSkinnableFrame
)

var kindNames = [...]string{
	KindUnknown:                             "Unknown",
	Kind7ZipArchive:                         "7-Zip Archive",
	KindAACAudio:                            "AAC Audio",
	KindACEArchive:                          "ACE Archive",
	KindAcrobatFormsDataFormat:              "Acrobat Forms Data Format",
	KindAcronisTrueImage:                    "Acronis True Image",
	KindAdobeInDesignDocument:               "Adobe InDesign Document",
	KindAdvancedForensicFormat:              "Advanced Forensic Format",
	KindALZArchive:                          "ALZ Archive",
	KindAmigaDiskFile:                       "Amiga Disk File",
	KindAmigaHardDiskFile:                   "Amiga Hard Disk File",
	KindAmigaHunkExecutable:                 "Amiga Hunk Executable",
	KindAmigaOctaMED:                        "Amiga OctaMED Audio",
	KindAMRAudio:                            "AMR Audio",
	KindAMRWBAudio:                          "AMR-WB Audio",
	KindAndroidART:                          "Android ART",
	KindAndroidBackup:                       "Android Backup",
	KindAndroidBinaryXML:                    "Android Binary XML",
	KindAndroidBootImage:                    "Android Boot Image",
	KindAndroidCompiledResources:            "Android Compiled Resources",
	KindAndroidOAT:                          "Android OAT",
	KindAndroidODEX:                         "Android ODEX",
	KindAndroidSparseImage:                  "Android Sparse Image",
	KindAndroidVDEX:                         "Android VDEX",
	KindApacheArrowFile:                     "Apache Arrow File",
	KindApacheParquet:                       "Apache Parquet",
	KindAPFSFilesystem:                      "APFS Filesystem",
	KindAppImage:                            "AppImage",
	KindAppleArchive:                        "Apple Archive",
	KindAppleBillOfMaterials:                "Apple Bill of Materials",
	KindAppleBinaryCrashReport:              "Apple Binary Crash Report",
	KindAppleBinaryPropertyList:             "Apple Binary Property List",
	KindAppleDesktopServicesStore:           "Apple Desktop Services Store",
	KindAppleDiskImage:                      "Apple Disk Image",
	KindAppleDouble:                         "AppleDouble File",
	KindAppleiWorkDocument:                  "Apple iWork Document",
	KindAppleKeychain:                       "Apple Keychain",
	KindAppleSingle:                         "AppleSingle File",
	KindAppleSystemLog:                      "Apple System Log",
	KindAppleTexture:                        "Apple Texture",
	KindAppleXMLPropertyList:                "Apple XML Property List",
	KindARArchive:                           "AR Archive",
	KindARCArchive:                          "ARC Archive",
	KindARJArchive:                          "ARJ Archive",
	KindARRIRAWImage:                        "ARRI RAW Image",
	KindASFContainer:                        "ASF Container",
	KindASTCTexture:                         "ASTC Texture",
	KindAtari7800ROM:                        "Atari 7800 ROM",
	KindAtariLynxROM:                        "Atari Lynx ROM",
	KindAtomFeed:                            "Atom Feed",
	KindAUAudio:                             "AU Audio",
	KindAudacityBlockFile:                   "Audacity Block File",
	KindAutoCADDrawing:                      "AutoCAD Drawing",
	KindAutoCADDXF:                          "AutoCAD DXF",
	KindAvroObjectContainer:                 "Avro Object Container",
	KindBAMData:                             "BAM Data",
	KindBcachefsFilesystem:                  "Bcachefs Filesystem",
	KindBCFData:                             "Binary Call Format (BCF)",
	KindBerkeleyDB:                          "Berkeley DB",
	KindBinHex:                              "BinHex Archive",
	KindBinkVideo:                           "Bink Video",
	KindBitcoinBlock:                        "Bitcoin Block Data",
	KindBitLockerDiskEncryption:             "BitLocker Disk Encryption",
	KindBlenderFile:                         "Blender File",
	KindBlizzardTexture:                     "Blizzard Texture",
	KindBluetoothSnoop:                      "Bluetooth Snoop Capture",
	KindBMPImage:                            "BMP Image",
	KindBochsDiskImage:                      "Bochs Disk Image",
	KindBouncyCastleKeystore:                "BouncyCastle Keystore",
	KindBPGImage:                            "BPG Image",
	KindBroadbandEBook:                      "Broadband eBook",
	KindBtrfsFilesystem:                     "Btrfs Filesystem",
	KindBUFR:                                "BUFR Data",
	KindBzip2Archive:                        "Bzip2 Archive",
	KindBzip3Archive:                        "Bzip3 Archive",
	KindCabinetArchive:                      "Cabinet Archive",
	KindCAFAudio:                            "CAF Audio",
	KindCALSImage:                           "CALS Raster Image",
	KindCanonRAWImage:                       "Canon RAW Image",
	KindCBORData:                            "CBOR Data",
	KindCHMDocument:                         "CHM Document",
	KindCinema4DModel:                       "Cinema 4D Model",
	KindCineonImage:                         "Cineon Image",
	KindCloneCDControl:                      "CloneCD Control File",
	KindComicBookArchive:                    "Comic Book Archive",
	KindCommodore64Tape:                     "Commodore 64 Tape",
	KindCommodore64TapeRAW:                  "Commodore 64 Tape RAW",
	KindCommodoreSID:                        "Commodore SID Audio",
	KindCompiledTerminfo:                    "Compiled Terminfo",
	KindCompressedSquareWave:                "Compressed Square Wave",
	KindCorelPhotoPaintImage:                "Corel Photo-Paint Image",
	KindCPIOArchive:                         "CPIO Archive",
	KindCRAMData:                            "CRAM Data",
	KindCramFS:                              "CramFS",
	KindCRXBrowserExtension:                 "CRX Browser Extension",
	KindDalvikExecutable:                    "Dalvik Executable",
	KindDDSImage:                            "DDS Image",
	KindDebianPackage:                       "Debian Package",
	KindDICOMMedicalImage:                   "DICOM Medical Image",
	KindDiracVideo:                          "Dirac Video",
	KindDjVuDocument:                        "DjVu Document",
	KindDolphinRVZ:                          "Dolphin RVZ",
	KindDOSExecutable:                       "DOS Executable",
	KindDPXImage:                            "DPX Image",
	KindDSDIFFAudio:                         "DSDIFF Audio",
	KindDSFAudio:                            "DSF Audio",
	KindDTSAudio:                            "DTS Audio",
	KindDVRMSVideo:                          "DVR-MS Video",
	KindEBMLContainer:                       "EBML Container",
	KindEGGArchive:                          "EGG Archive",
	KindEncapsulatedPostScript:              "Encapsulated PostScript",
	KindEnCaseImage:                         "EnCase Image",
	KindEOTFont:                             "EOT Font",
	KindErlangBEAM:                          "Erlang BEAM Bytecode",
	KindESRIShapefile:                       "ESRI Shapefile",
	KindExecutableAndLinkableFormat:         "Executable and Linkable Format",
	KindExFATFilesystem:                     "exFAT Filesystem",
	KindExtensibleMusicFormat:               "Extensible Music Format",
	KindExtFilesystem:                       "ext Filesystem",
	KindF2FSFilesystem:                      "F2FS Filesystem",
	KindFamicomDiskSystemROM:                "Famicom Disk System ROM",
	KindFarbfeldImage:                       "Farbfeld Image",
	KindFastTracker2ExtendedInstrument:      "FastTracker 2 Extended Instrument",
	KindFastTrackerModule:                   "FastTracker Module",
	KindFATFilesystem:                       "FAT Filesystem",
	KindFBXModel:                            "FBX Model",
	KindFeatherData:                         "Feather Data",
	KindFilesystemEntry:                     "Filesystem Entry",
	KindFITActivity:                         "FIT Activity",
	KindFITSAstronomicalImage:               "FITS Astronomical Image",
	KindFLACAudio:                           "FLAC Audio",
	KindFLIFImage:                           "FLIF Image",
	KindFLStudioProject:                     "FL Studio Project",
	KindFLVVideo:                            "FLV Video",
	KindFreeArcArchive:                      "FreeArc Archive",
	KindFujifilmRAWImage:                    "Fujifilm RAW Image",
	KindGameBoyAdvanceROM:                   "Game Boy Advance ROM",
	KindGameBoyROM:                          "Game Boy ROM",
	KindGameBoySound:                        "Game Boy Sound",
	KindGameCubeROM:                         "GameCube ROM",
	KindGDBMDatabase:                        "GDBM Database",
	KindGHCiInterface:                       "GHCi Interface",
	KindGIFImage:                            "GIF Image",
	KindGIMPBrush:                           "GIMP Brush",
	KindGIMPPattern:                         "GIMP Pattern",
	KindGIMPXCFImage:                        "GIMP XCF Image",
	KindGitBundle:                           "Git Bundle",
	KindGitCommitGraph:                      "Git Commit Graph",
	KindGitIndex:                            "Git Index",
	KindGitPack:                             "Git Pack",
	KindGitPackBitmap:                       "Git Pack Bitmap",
	KindGitPackIndex:                        "Git Pack Index",
	KindGitPackReverseIndex:                 "Git Pack Reverse Index",
	KindGlibcLocale:                         "Glibc Locale",
	KindGLTFBinary:                          "glTF Binary",
	KindGlyphBitmapDistributionFormat:       "Glyph Bitmap Distribution Format (BDF)",
	KindGNOMEKeyring:                        "GNOME Keyring",
	KindGNUGettextMachineCatalog:            "GNU Gettext Machine Catalog",
	KindGnuPGKeybox:                         "GnuPG Keybox",
	KindGodotPackage:                        "Godot Engine Package",
	KindGPSExchangeFormat:                   "GPS Exchange Format",
	KindGRIBData:                            "GRIB Data",
	KindGUIDPartitionTable:                  "GUID Partition Table",
	KindGzipArchive:                         "Gzip Archive",
	KindGzipData:                            "Gzip Data",
	KindHadoopRCFile:                        "Hadoop RCFile",
	KindHadoopSequenceFile:                  "Hadoop SequenceFile",
	KindHDF4Data:                            "HDF4 Data",
	KindHDF5Data:                            "HDF5 Data",
	KindHFSPlusFilesystem:                   "HFS+ Filesystem",
	KindHTMLDocument:                        "HTML Document",
	KindICalendar:                           "iCalendar",
	KindICCProfile:                          "ICC Profile",
	KindICNSIcon:                            "ICNS Icon",
	KindICOCURImage:                         "ICO/CUR Image",
	KindIFFContainer:                        "IFF Container",
	KindImpulseTrackerInstrument:            "Impulse Tracker Instrument",
	KindImpulseTrackerModule:                "Impulse Tracker Module",
	KindImpulseTrackerSample:                "Impulse Tracker Sample",
	KindInstallShieldCAB:                    "InstallShield Cabinet",
	KindISO9660Image:                        "ISO 9660 Image",
	KindISOBaseMedia:                        "ISO Base Media",
	KindIVFVideo:                            "IVF Video",
	KindJavaClass:                           "Java Class",
	KindJavaCryptographyExtensionKeyStore:   "JCE KeyStore",
	KindJavaKeyStore:                        "Java KeyStore",
	KindJavaModule:                          "Java Module",
	KindJavaPack200:                         "Java Pack200 Archive",
	KindJavaSerialization:                   "Java Serialization Data",
	KindJBIG2Image:                          "JBIG2 Image",
	KindJFFS2Filesystem:                     "JFFS2 Filesystem",
	KindJNGImage:                            "JNG Image",
	KindJPEG2000Image:                       "JPEG 2000 Image",
	KindJPEGImage:                           "JPEG Image",
	KindJPEGLSImage:                         "JPEG-LS Image",
	KindJPEGXLImage:                         "JPEG XL Image",
	KindJPEGXRImage:                         "JPEG XR Image",
	KindJSONDocument:                        "JSON Document",
	KindKDEKWallet:                          "KDE KWallet",
	KindKeePassDatabase:                     "KeePass Database",
	KindKeyholeMarkupLanguage:               "Keyhole Markup Language",
	KindKGBArchive:                          "KGB Archive",
	KindKorgAudio:                           "Korg Audio",
	KindKTXTexture:                          "KTX Texture",
	KindKyotoCabinetDatabase:                "Kyoto Cabinet Database",
	KindLevelDB:                             "LevelDB/RocksDB",
	KindLHAArchive:                          "LHA Archive",
	KindLidarData:                           "LiDAR Data (LAS)",
	KindLightWaveScene:                      "LightWave Scene",
	KindLinuxKernelImage:                    "Linux Kernel Image",
	KindLLVMBitcode:                         "LLVM Bitcode",
	KindLottieAnimation:                     "Lottie Animation",
	KindLrzipArchive:                        "Long Range ZIP Archive",
	KindLuaBytecode:                         "Lua Bytecode",
	KindLuceneIndex:                         "Lucene Index",
	KindLUKSDiskEncryption:                  "LUKS Disk Encryption",
	KindLVM2PhysicalVolume:                  "LVM2 Physical Volume",
	KindLytroLightField:                     "Lytro Light Field",
	KindLZ4Frame:                            "LZ4 Frame",
	KindLZFArchive:                          "LZF Archive",
	KindLZFSEData:                           "LZFSE Data",
	KindLZIPArchive:                         "LZIP Archive",
	KindLZMAData:                            "LZMA Data",
	KindLZOPArchive:                         "LZOP Archive",
	KindLZXArchive:                          "LZX Archive",
	KindM3U8Playlist:                        "M3U8 Playlist",
	KindMacBinary:                           "MacBinary",
	KindMachOBinary:                         "Mach-O Binary",
	KindMachOUniversalBinary:                "Mach-O Universal Binary",
	KindMacOSAlias:                          "macOS Alias",
	KindMacOSXCodeSignature:                 "macOS Code Signature",
	KindMacriumReflectImage:                 "Macrium Reflect Image",
	KindMagicaVoxel:                         "MagicaVoxel Model",
	KindMAMECHD:                             "MAME Compressed Hunks of Data",
	KindMaterialExchangeFormat:              "Material Exchange Format",
	KindMATLABData:                          "MATLAB Data",
	KindMayaASCII:                           "Maya ASCII",
	KindMayaBinary:                          "Maya Binary",
	KindMBOXEmailFolder:                     "MBOX Email Folder",
	KindMCAPCapture:                         "MCAP Capture",
	KindMediaDescriptor:                     "Media Descriptor",
	KindMetafileImage:                       "Metafile Image",
	KindMicrosoftAccessDatabase:             "Microsoft Access Database",
	KindMicrosoftCompress:                   "Microsoft Compress Archive",
	KindMicrosoftNetworkMonitor:             "Microsoft Network Monitor Capture",
	KindMicrosoftOutlookEmailFolder:         "Microsoft Outlook Email Folder",
	KindMicrosoftProgramDatabase:            "Microsoft Program Database (PDB)",
	KindMicrosoftReader:                     "Microsoft Reader eBook",
	KindMicrosoftSQLServerDatabase:          "Microsoft SQL Server Database",
	KindMicrosoftXbox360Profile:             "Xbox 360 Profile",
	KindMicrosoftXboxSave:                   "Xbox Save",
	KindMIDISequence:                        "MIDI Sequence",
	KindMinoltaRAWImage:                     "Minolta RAW Image",
	KindMNGImage:                            "MNG Image",
	KindMOBIDocument:                        "MOBI Document",
	KindMonkeysAudio:                        "Monkey's Audio",
	KindMoPaQArchive:                        "MoPaQ Archive",
	KindMozillaArchive:                      "Mozilla Archive (MAR)",
	KindMPEG2TransportStream:                "MPEG Transport Stream",
	KindMPEGAudio:                           "MPEG Audio",
	KindMPEGAudioFrame:                      "MPEG Audio",
	KindMPEGProgramStream:                   "MPEG Program Stream",
	KindMPEGTransportStream:                 "MPEG Transport Stream",
	KindMPEGVideo:                           "MPEG Video",
	KindMUGENSound:                          "M.U.G.E.N Sound",
	KindMUGENSprite:                         "M.U.G.E.N Sprite",
	KindMusepackAudio:                       "Musepack Audio",
	KindMySQLMyISAM:                         "MySQL MyISAM",
	KindNeoGeoPocketROM:                     "Neo Geo Pocket ROM",
	KindNESROM:                              "NES ROM",
	KindNESSoundFormat:                      "NES Sound Format",
	KindNetCDFData:                          "NetCDF Data",
	KindNetpbmImage:                         "Netpbm Image",
	KindNIfTIMedicalImage:                   "NIfTI Medical Image",
	KindNILFS2Filesystem:                    "NILFS2 Filesystem",
	KindNintendo3DSROM:                      "Nintendo 3DS ROM",
	KindNintendo64ROM:                       "Nintendo 64 ROM",
	KindNintendoBCFNT:                       "Nintendo BCFNT Font",
	KindNintendoBFLYT:                       "Nintendo BFLYT Layout",
	KindNintendoBRRES:                       "Nintendo BRRES Resource",
	KindNintendoDSiROM:                      "Nintendo DSi ROM",
	KindNintendoDSROM:                       "Nintendo DS ROM",
	KindNintendoRARC:                        "Nintendo RARC Archive",
	KindNintendoSwitchNRO:                   "Nintendo Switch NRO",
	KindNintendoSwitchNSO:                   "Nintendo Switch NSO",
	KindNintendoSwitchPackage:               "Nintendo Switch Package",
	KindNintendoSwitchROM:                   "Nintendo Switch ROM",
	KindNintendoU8Archive:                   "Nintendo U8 Archive",
	KindNintendoWiiUExecutable:              "Nintendo Wii U Executable",
	KindNintendoYay0:                        "Nintendo Yay0 Compressed",
	KindNintendoYaz0:                        "Nintendo Yaz0 Compressed",
	KindNISTSPHEREAudio:                     "NIST SPHERE Audio",
	KindNTFSFilesystem:                      "NTFS Filesystem",
	KindNullsoftVideo:                       "Nullsoft Video",
	KindNumPyArray:                          "NumPy Array",
	KindOCamlObject:                         "OCaml Object",
	KindOggContainer:                        "Ogg Container",
	KindOLECompoundDocument:                 "OLE Compound Document",
	KindOlympusRAWImage:                     "Olympus RAW Image",
	KindOpenEXRImage:                        "OpenEXR Image",
	KindOpenSSHPrivateKey:                   "OpenSSH Private Key",
	KindOpenTypeFont:                        "OpenType Font",
	KindOptimFROGAudio:                      "OptimFROG Audio",
	KindORCColumnarData:                     "ORC Columnar Data",
	KindOutlookExpressDatabase:              "Outlook Express Database",
	KindPaintShopProImage:                   "PaintShop Pro Image",
	KindParallelsDiskImage:                  "Parallels Disk Image",
	KindPCAPCapture:                         "PCAP Capture",
	KindPCAPNGCapture:                       "PCAPNG Capture",
	KindPCFFont:                             "Portable Compiled Format Font",
	KindPCXImage:                            "PCX Image",
	KindPDFDocument:                         "PDF Document",
	KindPeaZipArchive:                       "PeaZip Archive",
	KindPEMCertificate:                      "PEM Certificate",
	KindPEMPrivateKey:                       "PEM Private Key",
	KindPGFImage:                            "Progressive Graphics File",
	KindPGPMessage:                          "PGP Message",
	KindPhotoshopDocument:                   "Photoshop Document",
	KindPico8Cartridge:                      "Pico-8 Cartridge",
	KindPIMArchive:                          "PIM Archive",
	KindPKCS12:                              "PKCS#12",
	KindPlaylistFile:                        "Playlist File",
	KindPlayStation1Executable:              "PlayStation 1 Executable",
	KindPlayStation2MemoryCard:              "PlayStation 2 Memory Card",
	KindPlayStationExecutable:               "PlayStation Executable",
	KindPlayStationPortableExecutable:       "PlayStation Portable Executable",
	KindPlayStationPortableISO:              "PlayStation Portable ISO",
	KindPlayStationSoundFormat:              "PlayStation Sound Format",
	KindPLYModel:                            "PLY Model",
	KindPNGImage:                            "PNG Image",
	KindPortableExecutable:                  "Portable Executable",
	KindPostgreSQLCustomDump:                "PostgreSQL Custom Dump",
	KindPostScriptDocument:                  "PostScript Document",
	KindPowerISODAA:                         "PowerISO DAA",
	KindPSFFont:                             "PC Screen Font",
	KindPuttyPrivateKey:                     "PuTTY Private Key",
	KindPVRTexture:                          "PVR Texture",
	KindPythonBytecode:                      "Python Bytecode",
	KindPythonPickle:                        "Python Pickle",
	KindQCOWDiskImage:                       "QCOW Disk Image",
	KindQDBMDatabase:                        "QDBM Database",
	KindQEMUQEDDiskImage:                    "QEMU QED Disk Image",
	KindQOIImage:                            "QOI Image",
	KindQuakePAK:                            "Quake PAK Archive",
	KindRadianceHDRImage:                    "Radiance HDR Image",
	KindRARArchive:                          "RAR Archive",
	KindRData:                               "RData",
	KindRealMedia:                           "RealMedia",
	KindRedisDatabase:                       "Redis Database",
	KindREDRAWImage:                         "RED RAW Image",
	KindReiserFSFilesystem:                  "ReiserFS Filesystem",
	KindRenPyArchive:                        "Ren'Py Archive",
	KindRhino3DModel:                        "Rhino 3D Model",
	KindRichTextFormatDocument:              "Rich Text Format Document",
	KindRIFFContainer:                       "RIFF Container",
	KindRKAudio:                             "RKAudio",
	KindRNCArchive:                          "Rob Northen Compression Archive",
	KindROMFS:                               "ROMFS",
	KindROOTData:                            "ROOT Data",
	KindROSBag:                              "ROS Bag",
	KindRPGMakerArchive:                     "RPG Maker Archive",
	KindRPMPackage:                          "RPM Package",
	KindRSSFeed:                             "RSS Feed",
	KindRubyGemPackage:                      "RubyGem Package",
	KindRubyMarshal:                         "Ruby Marshal Data",
	KindRzipArchive:                         "rzip Archive",
	KindSASData:                             "SAS Data",
	KindScreamTrackerModule:                 "Scream Tracker Module",
	KindSegaCDImage:                         "Sega CD Image",
	KindSegaDreamcastROM:                    "Sega Dreamcast ROM",
	KindSegaMasterSystemROM:                 "Sega Master System ROM",
	KindSegaMegaDriveROM:                    "Sega Mega Drive ROM",
	KindSegaNaomiROM:                        "Sega NAOMI ROM",
	KindSegaSaturnROM:                       "Sega Saturn ROM",
	KindSGIImage:                            "SGI Image",
	KindShebangScript:                       "Shebang Script",
	KindShockwaveFlash:                      "Shockwave Flash",
	KindShortenAudio:                        "Shorten Audio",
	KindSigmaRAWImage:                       "Sigma RAW Image",
	KindSiliconGraphicsMovie:                "Silicon Graphics Movie",
	KindSketchDocument:                      "Sketch Document",
	KindSketchUpModel:                       "SketchUp Model",
	KindSmackerVideo:                        "Smacker Video",
	KindSnappyFramedData:                    "Snappy Framed Data",
	KindSNESSPC:                             "SNES SPC Audio",
	KindSnoopCapture:                        "Snoop Capture",
	KindSOAPMessage:                         "SOAP Message",
	KindSonyOpenMG:                          "Sony OpenMG Audio",
	KindSonyWave64Audio:                     "Sony Wave64 Audio",
	KindSourceEngineBSP:                     "Source Engine BSP",
	KindSPSSData:                            "SPSS Data",
	KindSPSSPortableData:                    "SPSS Portable Data",
	KindSQLite3SharedMemory:                 "SQLite Shared Memory",
	KindSQLite3WriteAheadLog:                "SQLite Write-Ahead Log",
	KindSQLiteDatabase:                      "SQLite Database",
	KindSquashFSFilesystem:                  "SquashFS Filesystem",
	KindSQXArchive:                          "SQX Archive",
	KindSSHPublicKey:                        "SSH Public Key",
	KindStataData:                           "Stata Data",
	KindSTEP3DModel:                         "STEP 3D Model",
	KindStuffItArchive:                      "StuffIt Archive",
	KindSubRipText:                          "SubRip Text (SRT)",
	KindSunRasterImage:                      "Sun Raster Image",
	KindSVGImage:                            "SVG Image",
	KindSymantecGhostImage:                  "Symantec Ghost Image",
	KindSymbianExecutable:                   "Symbian Executable",
	KindSymbianInstallationFormat:           "Symbian Installation Format",
	KindSystemdJournal:                      "Systemd Journal",
	KindSzipArchive:                         "Szip Archive",
	KindTAKAudio:                            "TAK Audio",
	KindTARArchive:                          "TAR Archive",
	KindTensorFlowLiteModel:                 "TensorFlow Lite Model",
	KindTextFile:                            "Text File",
	KindTIFFImage:                           "TIFF Image",
	KindTimezoneData:                        "Timezone Data",
	KindTokyoCabinetDatabase:                "Tokyo Cabinet Database",
	KindTorrentFile:                         "Torrent File",
	KindTransportNeutralEncapsulationFormat: "Transport Neutral Encapsulation Format (TNEF)",
	KindTrueTypeCollection:                  "TrueType Collection",
	KindTrueTypeFont:                        "TrueType Font",
	KindTTAAudio:                            "TTA Audio",
	KindU3DModel:                            "U3D Model",
	KindUBIFSFilesystem:                     "UBIFS Filesystem",
	KindUBootImage:                          "U-Boot Image",
	KindUHAArchive:                          "UHA Archive",
	KindUnityWebData:                        "Unity Web Data",
	KindUniversalDiskFormat:                 "Universal Disk Format",
	KindUniversalSceneDescription:           "Universal Scene Description",
	KindUnixCompressArchive:                 "Unix Compress Archive",
	KindUnrealEnginePackage:                 "Unreal Engine Package",
	KindValvePak:                            "Valve Pak",
	KindValveTextureFormat:                  "Valve Texture Format",
	KindVCard:                               "vCard",
	KindVeeamBackup:                         "Veeam Backup",
	KindVHDDiskImage:                        "VHD Disk Image",
	KindVHDXDiskImage:                       "VHDX Disk Image",
	KindVideoGameMusic:                      "Video Game Music (VGM)",
	KindVirtualBoxDiskImage:                 "VirtualBox Disk Image",
	KindVirtualBoxSavedState:                "VirtualBox Saved State",
	KindVMwareDiskImage:                     "VMware Disk Image",
	KindVMwareNVRAM:                         "VMware NVRAM",
	KindVMwareSnapshot:                      "VMware Snapshot",
	KindVMwareSuspend:                       "VMware Suspend",
	KindVocaloidProject:                     "Vocaloid Project",
	KindVOCAudio:                            "VOC Audio",
	KindVRML3DModel:                         "VRML 3D Model",
	KindVulkanSPIRV:                         "Vulkan SPIR-V",
	KindWADArchive:                          "WAD Archive",
	KindWARCFile:                            "WARC File",
	KindWarcraftIIIReplay:                   "Warcraft III Replay",
	KindWavPackAudio:                        "WavPack Audio",
	KindWavPackCorrection:                   "WavPack Correction",
	KindWebAssemblyModule:                   "WebAssembly Module",
	KindWebVTT:                              "WebVTT",
	KindWiiBackupFileSystem:                 "Wii Backup File System",
	KindWiiROM:                              "Wii ROM",
	KindWiiUArchive:                         "Wii U Archive",
	KindWindowsEventLog:                     "Windows Event Log",
	KindWindowsHelp:                         "Windows Help File",
	KindWindowsHibernation:                  "Windows Hibernation File",
	KindWindowsImagingFormat:                "Windows Imaging Format",
	KindWindowsMinidump:                     "Windows Minidump",
	KindWindowsPrecompiledHeader:            "Windows Precompiled Header",
	KindWindowsPrefetch:                     "Windows Prefetch",
	KindWindowsRegistryHive:                 "Windows Registry Hive",
	KindWindowsResourceFile:                 "Windows Resource File",
	KindWindowsShortcut:                     "Windows Shortcut",
	KindWindowsTypeLibrary:                  "Windows Type Library (TLB)",
	KindWiredTigerDatabase:                  "WiredTiger Database",
	KindWOFF2Font:                           "WOFF2 Font",
	KindWOFFFont:                            "WOFF Font",
	KindWordPerfectGraphics:                 "WordPerfect Graphics",
	KindWTVVideo:                            "WTV Video",
	KindXARArchive:                          "XAR Archive",
	KindXBMImage:                            "X BitMap Image",
	KindXbox360Executable:                   "Xbox 360 Executable",
	KindXboxExecutable:                      "Xbox Executable",
	KindXboxISO:                             "Xbox ISO",
	KindXCOFFExecutable:                     "XCOFF Executable",
	KindXFSFilesystem:                       "XFS Filesystem",
	KindXilinxBitstream:                     "Xilinx Bitstream",
	KindXMLDocument:                         "XML Document",
	KindXPMImage:                            "XPM Image",
	KindXZArchive:                           "XZ Archive",
	KindYamahaSMAF:                          "Yamaha SMAF",
	KindYUV4MPEG2Video:                      "YUV4MPEG2 Video",
	KindZBackup:                             "ZBackup",
	KindZchunk:                              "Zchunk Archive",
	KindZFSFilesystem:                       "ZFS Filesystem",
	KindZIPArchive:                          "ZIP Archive",
	KindZlibData:                            "Zlib Data",
	KindZOOArchive:                          "ZOO Archive",
	KindZPAQArchive:                         "ZPAQ Archive",
	KindZstandardArchive:                    "Zstandard Archive",
	KindZstandardDictionary:                 "Zstandard Dictionary",
	KindZXTape:                              "ZX Spectrum Tape",
}

var typeNames = [...]string{
	TypeNone:                     "",
	Type32Bit:                    "32-bit",
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
	TypeAfterEffectsProjectAEP:   "After Effects Project (AEP)",
	TypeAIFCAudio:                "AIFC Audio",
	TypeAIFFAudio:                "AIFF Audio",
	TypeAndroidAppBundleAAB:      "Android App Bundle (AAB)",
	TypeAndroidArchiveAAR:        "Android Archive (AAR)",
	TypeAndroidPackageAPK:        "Android Package (APK)",
	TypeAndroidPackageXAPK:       "Android Package (XAPK)",
	TypeAndroidSplitAPKS:         "Android Split APK Set (APKS)",
	TypeAndroidSystemPackageAPEX: "Android System Package (APEX)",
	TypeAPPXPackage:              "APPX Package",
	TypeArchLinuxPackage:         "Arch Linux Package",
	TypeASCIIText:                "ASCII Text",
	TypeAVIFImage:                "AVIF Image",
	TypeAVIFImageSequence:        "AVIF Image Sequence",
	TypeAVIVideo:                 "AVI Video",
	TypeBGZF:                     "BGZF",
	TypeBigEndian:                "Big-Endian",
	TypeBinaryBigEndian:          "Binary Big-Endian",
	TypeBinaryLittleEndian:       "Binary Little-Endian",
	TypeBlackmagicRAW:            "Blackmagic RAW",
	TypeBlockDevice:              "Block Device",
	TypeByteSwapped:              "Byte-Swapped",
	TypeCanonRAW3CR3:             "Canon RAW 3 (CR3)",
	TypeCB7:                      "CB7",
	TypeCBR:                      "CBR",
	TypeCBT:                      "CBT",
	TypeCBZ:                      "CBZ",
	TypeCDAAudio:                 "CD Audio",
	TypeCharacterDevice:          "Character Device",
	TypeCodestream:               "Codestream",
	TypeCondaPackage:             "Conda Package",
	TypeContainer:                "Container",
	TypeCorelDRAWDocumentCDR:     "CorelDRAW Document (CDR)",
	TypeCRXVersion2:              "Version 2",
	TypeCRXVersion3:              "Version 3",
	TypeCubaseProjectCPR:         "Cubase Project (CPR)",
	TypeDEX035:                   "DEX 035",
	TypeDEX036:                   "DEX 036",
	TypeDEX037:                   "DEX 037",
	TypeDEX038:                   "DEX 038",
	TypeDEX039:                   "DEX 039",
	TypeDEX040:                   "DEX 040",
	TypeDEX041:                   "DEX 041",
	TypeDirectory:                "Directory",
	TypeDownloadableSounds:       "Downloadable Sounds",
	TypeDSAPrivateKey:            "DSA Private Key",
	TypeEAC3:                     "E-AC-3",
	TypeECPrivateKey:             "EC Private Key",
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
	TypeFabricMod:                "Fabric Mod",
	TypeFAT12:                    "FAT12",
	TypeFAT16:                    "FAT16",
	TypeFAT32:                    "FAT32",
	TypeFirefoxExtensionXPI:      "Firefox Extension (XPI)",
	TypeFLACAudio:                "FLAC Audio",
	TypeForgeMod:                 "Forge Mod",
	TypeGIF87a:                   "GIF87a",
	TypeGIF89a:                   "GIF89a",
	TypeHEIFImage:                "HEIF Image",
	TypeILBMImage:                "ILBM Image",
	TypeIOSApplicationArchiveIPA: "iOS Application Archive (IPA)",
	TypeIWAD:                     "IWAD",
	TypeJavaArchiveJAR:           "Java Archive (JAR)",
	TypeJavaEnterpriseArchiveEAR: "Java Enterprise Archive (EAR)",
	TypeJavaWebArchiveWAR:        "Java Web Archive (WAR)",
	TypeJMOD:                     "JMOD",
	TypeKDB:                      "KDB",
	TypeKDBX:                     "KDBX",
	TypeKerasModel:               "Keras Model",
	TypeKMZArchive:               "KMZ Archive",
	TypeKritaDocumentKRA:         "Krita Document (KRA)",
	TypeKTX:                      "KTX",
	TypeKTX2:                     "KTX2",
	TypeLittleEndian:             "Little-Endian",
	TypeLOVEGame:                 "LÖVE Game",
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
	TypeMicrosoftProjectDocumentMPP:                     "Microsoft Project Document (MPP)",
	TypeMicrosoftPublisherDocumentPUB:                   "Microsoft Publisher Document (PUB)",
	TypeMicrosoftVisioDrawingVSD:                        "Microsoft Visio Drawing (VSD)",
	TypeMicrosoftWordDocumentDOC:                        "Microsoft Word Document (DOC)",
	TypeMicrosoftWordDocumentDOCX:                       "Microsoft Word Document (DOCX)",
	TypeMicrosoftWordMacroEnabledDocumentDOCM:           "Microsoft Word Macro-Enabled Document (DOCM)",
	TypeMicrosoftWordMacroEnabledTemplateDOTM:           "Microsoft Word Macro-Enabled Template (DOTM)",
	TypeMicrosoftWordTemplateDOTX:                       "Microsoft Word Template (DOTX)",
	TypeMinecraftResourcePack:                           "Minecraft Resource Pack",
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
	TypeOpenDocumentChartODC:                            "OpenDocument Chart (ODC)",
	TypeOpenDocumentDatabaseODB:                         "OpenDocument Database (ODB)",
	TypeOpenDocumentFormulaODF:                          "OpenDocument Formula (ODF)",
	TypeOpenDocumentGraphicsODG:                         "OpenDocument Graphics (ODG)",
	TypeOpenDocumentImageODI:                            "OpenDocument Image (ODI)",
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
	TypePyTorchModel:                                    "PyTorch Model",
	TypeQCOW1:                                           "QCOW1",
	TypeQCOW2:                                           "QCOW2",
	TypeQCPAudio:                                        "QCP Audio",
	TypeQuickTimeMovie:                                  "QuickTime Movie",
	TypeRAR4:                                            "RAR4",
	TypeRAR5:                                            "RAR5",
	TypeRIFFMIDI:                                        "MIDI",
	TypeRIFFPalette:                                     "Palette",
	TypeRSAPrivateKey:                                   "RSA Private Key",
	TypeSketchDocument:                                  "Sketch Document",
	TypeSlackwarePackage:                                "Slackware Package",
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
	TypeVagrantBox:                                      "Vagrant Box",
	TypeVisualStudioExtensionVSIX:                       "Visual Studio Extension (VSIX)",
	TypeVMDK:                                            "VMDK",
	TypeVorbisAudio:                                     "Vorbis Audio",
	TypeWAVAudio:                                        "WAV Audio",
	TypeWebM:                                            "WebM",
	TypeWebPImage:                                       "WebP Image",
	TypeWindowsAnimatedCursor:                           "Windows Animated Cursor",
	TypeWindowsCursor:                                   "Windows Cursor",
	TypeWindowsIcon:                                     "Windows Icon",
	TypeWindowsLE:                                       "Linear Executable (LE)",
	TypeWindowsLX:                                       "OS/2 Linear Executable (LX)",
	TypeWindowsMetafileWMF:                              "Windows Metafile (WMF)",
	TypeWindowsNE:                                       "16-bit New Executable (NE)",
	TypeWrapper:                                         "Wrapper",
	TypeXMLPaperSpecificationXPS:                        "XML Paper Specification (XPS)",
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
