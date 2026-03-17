package types

type KindID uint16
type TypeID uint16

const (
	KindUnknown KindID = iota
	Kind7ZipArchive
	KindAACAudio
	KindAMRAudio
	KindAMRWBAudio
	KindAndroidBootImage
	KindAndroidOAT
	KindAndroidODEX
	KindAndroidSparseImage
	KindApacheArrowFile
	KindApacheParquet
	KindAppImage
	KindAppleBinaryPropertyList
	KindAppleDouble
	KindAppleiWorkDocument
	KindAppleSingle
	KindARArchive
	KindASFContainer
	KindAUAudio
	KindAvroObjectContainer
	KindBAMData
	KindBlenderFile
	KindBMPImage
	KindBPGImage
	KindBtrfsFilesystem
	KindBzip2Archive
	KindCabinetArchive
	KindCAFAudio
	KindCanonRAWImage
	KindCHMDocument
	KindCPIOArchive
	KindCRAMData
	KindCRXBrowserExtension
	KindDalvikExecutable
	KindDDSImage
	KindDebianPackage
	KindDICOMMedicalImage
	KindDjVuDocument
	KindDSDIFFAudio
	KindDSFAudio
	KindDTSAudio
	KindEBMLContainer
	KindEOTFont
	KindExecutableAndLinkableFormat
	KindExtFilesystem
	KindFarbfeldImage
	KindFastTrackerModule
	KindFilesystemEntry
	KindFITSAstronomicalImage
	KindFLACAudio
	KindFLIFImage
	KindFLVVideo
	KindFujifilmRAWImage
	KindGIFImage
	KindGIMPXCFImage
	KindGitIndex
	KindGitPack
	KindGLTFBinary
	KindGRIBData
	KindGzipArchive
	KindGzipData
	KindHDF5Data
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
	KindLZ4Frame
	KindLZIPArchive
	KindLZMAData
	KindLZOPArchive
	KindM3U8Playlist
	KindMacBinary
	KindMachOBinary
	KindMachOUniversalBinary
	KindMetafileImage
	KindMicrosoftOutlookEmailFolder
	KindMIDISequence
	KindMOBIDocument
	KindMonkeysAudio
	KindMPEG2TransportStream
	KindMPEGAudio
	KindMPEGAudioFrame
	KindMPEGProgramStream
	KindMPEGTransportStream
	KindMusepackAudio
	KindNESROM
	KindNetCDFData
	KindNetpbmImage
	KindOggContainer
	KindOLECompoundDocument
	KindOlympusRAWImage
	KindOpenEXRImage
	KindOpenTypeFont
	KindORCColumnarData
	KindPCAPCapture
	KindPCAPNGCapture
	KindPCXImage
	KindPDFDocument
	KindPEMCertificate
	KindPEMPrivateKey
	KindPGPMessage
	KindPhotoshopDocument
	KindPKCS12
	KindPNGImage
	KindPortableExecutable
	KindPostScriptDocument
	KindPythonBytecode
	KindQCOWDiskImage
	KindQOIImage
	KindRadianceHDRImage
	KindRARArchive
	KindRealMedia
	KindRichTextFormatDocument
	KindRIFFContainer
	KindRPMPackage
	KindRubyGemPackage
	KindScreamTrackerModule
	KindShebangScript
	KindShockwaveFlash
	KindSketchDocument
	KindSnappyFramedData
	KindSQLiteDatabase
	KindSquashFSFilesystem
	KindStuffItArchive
	KindSunRasterImage
	KindSVGImage
	KindTAKAudio
	KindTARArchive
	KindTextFile
	KindTIFFImage
	KindTorrentFile
	KindTrueTypeCollection
	KindTrueTypeFont
	KindTTAAudio
	KindUBootImage
	KindUnixCompressArchive
	KindVCard
	KindVHDDiskImage
	KindVHDXDiskImage
	KindVirtualBoxDiskImage
	KindVMwareDiskImage
	KindVOCAudio
	KindWARCFile
	KindWavPackAudio
	KindWebAssemblyModule
	KindWindowsEventLog
	KindWindowsImagingFormat
	KindWindowsRegistryHive
	KindWindowsShortcut
	KindWOFF2Font
	KindWOFFFont
	KindWTVVideo
	KindXARArchive
	KindXFSFilesystem
	KindXMLDocument
	KindXZArchive
	KindZIPArchive
	KindZlibData
	KindZstandardArchive
)

const (
	TypeNone TypeID = iota
	Type32BitBigEndian
	Type32BitLittleEndian
	Type3G2Media
	Type3GPPMedia
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
	TypeCanonRAW3CR3
	TypeCDAAudio
	TypeCharacterDevice
	TypeCodestream
	TypeCondaPackage
	TypeContainer
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
	TypeJavaArchiveJAR
	TypeJavaEnterpriseArchiveEAR
	TypeJavaWebArchiveWAR
	TypeJMOD
	TypeKDBX
	TypeKMZArchive
	TypeKTX
	TypeKTX2
	TypeLittleEndian
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
)

var kindNames = [...]string{
	KindUnknown:                     "Unknown",
	Kind7ZipArchive:                 "7-Zip Archive",
	KindAACAudio:                    "AAC Audio",
	KindAMRAudio:                    "AMR Audio",
	KindAMRWBAudio:                  "AMR-WB Audio",
	KindAndroidBootImage:            "Android Boot Image",
	KindAndroidOAT:                  "Android OAT",
	KindAndroidODEX:                 "Android ODEX",
	KindAndroidSparseImage:          "Android Sparse Image",
	KindApacheArrowFile:             "Apache Arrow File",
	KindApacheParquet:               "Apache Parquet",
	KindAppImage:                    "AppImage",
	KindAppleBinaryPropertyList:     "Apple Binary Property List",
	KindAppleDouble:                 "AppleDouble File",
	KindAppleSingle:                 "AppleSingle File",
	KindARArchive:                   "AR Archive",
	KindASFContainer:                "ASF Container",
	KindAUAudio:                     "AU Audio",
	KindAvroObjectContainer:         "Avro Object Container",
	KindBAMData:                     "BAM Data",
	KindBlenderFile:                 "Blender File",
	KindBMPImage:                    "BMP Image",
	KindBPGImage:                    "BPG Image",
	KindBtrfsFilesystem:             "Btrfs Filesystem",
	KindBzip2Archive:                "Bzip2 Archive",
	KindCabinetArchive:              "Cabinet Archive",
	KindCAFAudio:                    "CAF Audio",
	KindCanonRAWImage:               "Canon RAW Image",
	KindCHMDocument:                 "CHM Document",
	KindCPIOArchive:                 "CPIO Archive",
	KindCRAMData:                    "CRAM Data",
	KindCRXBrowserExtension:         "CRX Browser Extension",
	KindDalvikExecutable:            "Dalvik Executable",
	KindDDSImage:                    "DDS Image",
	KindDebianPackage:               "Debian Package",
	KindDICOMMedicalImage:           "DICOM Medical Image",
	KindDjVuDocument:                "DjVu Document",
	KindDSDIFFAudio:                 "DSDIFF Audio",
	KindDSFAudio:                    "DSF Audio",
	KindDTSAudio:                    "DTS Audio",
	KindEBMLContainer:               "EBML Container",
	KindEOTFont:                     "EOT Font",
	KindExecutableAndLinkableFormat: "Executable and Linkable Format",
	KindExtFilesystem:               "ext Filesystem",
	KindFarbfeldImage:               "Farbfeld Image",
	KindFastTrackerModule:           "FastTracker Module",
	KindFilesystemEntry:             "Filesystem Entry",
	KindFITSAstronomicalImage:       "FITS Astronomical Image",
	KindFLACAudio:                   "FLAC Audio",
	KindFLIFImage:                   "FLIF Image",
	KindFLVVideo:                    "FLV Video",
	KindFujifilmRAWImage:            "Fujifilm RAW Image",
	KindGIFImage:                    "GIF Image",
	KindGIMPXCFImage:                "GIMP XCF Image",
	KindGitIndex:                    "Git Index",
	KindGitPack:                     "Git Pack",
	KindGRIBData:                    "GRIB Data",
	KindGzipArchive:                 "Gzip Archive",
	KindGzipData:                    "Gzip Data",
	KindHDF5Data:                    "HDF5 Data",
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
	KindJNGImage:                    "JNG Image",
	KindJPEG2000Image:               "JPEG 2000 Image",
	KindJPEGImage:                   "JPEG Image",
	KindJPEGXLImage:                 "JPEG XL Image",
	KindJPEGXRImage:                 "JPEG XR Image",
	KindKeePassDatabase:             "KeePass Database",
	KindKTXTexture:                  "KTX Texture",
	KindLLVMBitcode:                 "LLVM Bitcode",
	KindLuaBytecode:                 "Lua Bytecode",
	KindLZ4Frame:                    "LZ4 Frame",
	KindLZIPArchive:                 "LZIP Archive",
	KindLZMAData:                    "LZMA Data",
	KindLZOPArchive:                 "LZOP Archive",
	KindMacBinary:                   "MacBinary",
	KindMachOBinary:                 "Mach-O Binary",
	KindMachOUniversalBinary:        "Mach-O Universal Binary",
	KindMetafileImage:               "Metafile Image",
	KindMIDISequence:                "MIDI Sequence",
	KindMonkeysAudio:                "Monkey's Audio",
	KindMPEG2TransportStream:        "MPEG Transport Stream",
	KindMPEGAudio:                   "MPEG Audio",
	KindMPEGAudioFrame:              "MPEG Audio",
	KindMPEGProgramStream:           "MPEG Program Stream",
	KindMPEGTransportStream:         "MPEG Transport Stream",
	KindMusepackAudio:               "Musepack Audio",
	KindNESROM:                      "NES ROM",
	KindNetCDFData:                  "NetCDF Data",
	KindNetpbmImage:                 "Netpbm Image",
	KindOggContainer:                "Ogg Container",
	KindOLECompoundDocument:         "OLE Compound Document",
	KindOlympusRAWImage:             "Olympus RAW Image",
	KindOpenEXRImage:                "OpenEXR Image",
	KindOpenTypeFont:                "OpenType Font",
	KindORCColumnarData:             "ORC Columnar Data",
	KindPCAPCapture:                 "PCAP Capture",
	KindPCAPNGCapture:               "PCAPNG Capture",
	KindPCXImage:                    "PCX Image",
	KindPDFDocument:                 "PDF Document",
	KindPGPMessage:                  "PGP Message",
	KindPhotoshopDocument:           "Photoshop Document",
	KindPKCS12:                      "PKCS#12",
	KindPNGImage:                    "PNG Image",
	KindPortableExecutable:          "Portable Executable",
	KindPostScriptDocument:          "PostScript Document",
	KindPythonBytecode:              "Python Bytecode",
	KindQCOWDiskImage:               "QCOW Disk Image",
	KindQOIImage:                    "QOI Image",
	KindRadianceHDRImage:            "Radiance HDR Image",
	KindRARArchive:                  "RAR Archive",
	KindRealMedia:                   "RealMedia",
	KindRichTextFormatDocument:      "Rich Text Format Document",
	KindRIFFContainer:               "RIFF Container",
	KindRPMPackage:                  "RPM Package",
	KindRubyGemPackage:              "RubyGem Package",
	KindScreamTrackerModule:         "Scream Tracker Module",
	KindShockwaveFlash:              "Shockwave Flash",
	KindSketchDocument:              "Sketch Document",
	KindSnappyFramedData:            "Snappy Framed Data",
	KindSQLiteDatabase:              "SQLite Database",
	KindSquashFSFilesystem:          "SquashFS Filesystem",
	KindStuffItArchive:              "StuffIt Archive",
	KindSunRasterImage:              "Sun Raster Image",
	KindTAKAudio:                    "TAK Audio",
	KindTARArchive:                  "TAR Archive",
	KindTextFile:                    "Text File",
	KindTIFFImage:                   "TIFF Image",
	KindTorrentFile:                 "Torrent File",
	KindTrueTypeCollection:          "TrueType Collection",
	KindTrueTypeFont:                "TrueType Font",
	KindTTAAudio:                    "TTA Audio",
	KindUBootImage:                  "U-Boot Image",
	KindUnixCompressArchive:         "Unix Compress Archive",
	KindVHDDiskImage:                "VHD Disk Image",
	KindVHDXDiskImage:               "VHDX Disk Image",
	KindVirtualBoxDiskImage:         "VirtualBox Disk Image",
	KindVMwareDiskImage:             "VMware Disk Image",
	KindVOCAudio:                    "VOC Audio",
	KindWARCFile:                    "WARC File",
	KindWavPackAudio:                "WavPack Audio",
	KindWebAssemblyModule:           "WebAssembly Module",
	KindWindowsEventLog:             "Windows Event Log",
	KindWindowsImagingFormat:        "Windows Imaging Format",
	KindWindowsRegistryHive:         "Windows Registry Hive",
	KindWindowsShortcut:             "Windows Shortcut",
	KindWOFF2Font:                   "WOFF2 Font",
	KindWOFFFont:                    "WOFF Font",
	KindXARArchive:                  "XAR Archive",
	KindXFSFilesystem:               "XFS Filesystem",
	KindXZArchive:                   "XZ Archive",
	KindZIPArchive:                  "ZIP Archive",
	KindZlibData:                    "Zlib Data",
	KindZstandardArchive:            "Zstandard Archive",
	KindAppleiWorkDocument:          "Apple iWork Document",
	KindGLTFBinary:                  "glTF Binary",
	KindICalendar:                   "iCalendar",
	KindICCProfile:                  "ICC Profile",
	KindLHAArchive:                  "LHA Archive",
	KindM3U8Playlist:                "M3U8 Playlist",
	KindMicrosoftOutlookEmailFolder: "Microsoft Outlook Email Folder",
	KindMOBIDocument:                "MOBI Document",
	KindPEMCertificate:              "PEM Certificate",
	KindPEMPrivateKey:               "PEM Private Key",
	KindShebangScript:               "Shebang Script",
	KindSVGImage:                    "SVG Image",
	KindVCard:                       "vCard",
	KindWTVVideo:                    "WTV Video",
	KindXMLDocument:                 "XML Document",
}

var typeNames = [...]string{
	TypeNone:                     "",
	Type32BitBigEndian:           "32-bit Big-Endian",
	Type32BitLittleEndian:        "32-bit Little-Endian",
	Type3G2Media:                 "3G2 Media",
	Type3GPPMedia:                "3GPP Media",
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
	TypeCharacterDevice:          "Character Device",
	TypeCodestream:               "Codestream",
	TypeCondaPackage:             "Conda Package",
	TypeContainer:                "Container",
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
	TypeJavaArchiveJAR:           "Java Archive (JAR)",
	TypeJavaEnterpriseArchiveEAR: "Java Enterprise Archive (EAR)",
	TypeJavaWebArchiveWAR:        "Java Web Archive (WAR)",
	TypeJMOD:                     "JMOD",
	TypeKDBX:                     "KDBX",
	TypeKMZArchive:               "KMZ Archive",
	TypeKTX:                      "KTX",
	TypeKTX2:                     "KTX2",
	TypeLittleEndian:             "Little-Endian",
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
	TypePythonSourceDistributionSDist:                   "Python Source Distribution (sdist)",
	TypePythonWheelWHL:                                  "Python Wheel (WHL)",
	TypeQCOW2:                                           "QCOW2",
	TypeQuickTimeMovie:                                  "QuickTime Movie",
	TypeRAR4:                                            "RAR4",
	TypeRAR5:                                            "RAR5",
	TypeSketchDocument:                                  "Sketch Document",
	TypeSocket:                                          "Socket",
	TypeSonyRAWARW:                                      "Sony RAW (ARW)",
	TypeSonyRAWSR2:                                      "Sony RAW (SR2)",
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
	TypeCanonRAW3CR3:                                    "Canon RAW 3 (CR3)",
	TypeCDAAudio:                                        "CD Audio",
	TypeQCPAudio:                                        "QCP Audio",
	TypeRSAPrivateKey:                                   "RSA Private Key",
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
