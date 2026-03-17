//go:generate go run sort.go

package types

type KindID uint16
type TypeID uint16

const (
	KindUnknown KindID = iota
	Kind7ZipArchive
	KindAACAudio
	KindAMRAudio
	KindAMRWBAudio
	KindARArchive
	KindASFContainer
	KindAUAudio
	KindAndroidBootImage
	KindAndroidOAT
	KindAndroidODEX
	KindAndroidSparseImage
	KindApacheArrowFile
	KindApacheParquet
	KindAppImage
	KindAppleBinaryPropertyList
	KindAppleDouble
	KindAppleSingle
	KindAppleiWorkDocument
	KindAvroObjectContainer
	KindBAMData
	KindBMPImage
	KindBPGImage
	KindBlenderFile
	KindBtrfsFilesystem
	KindBzip2Archive
	KindCAFAudio
	KindCHMDocument
	KindCPIOArchive
	KindCRAMData
	KindCRXBrowserExtension
	KindCabinetArchive
	KindCanonRAWImage
	KindDDSImage
	KindDICOMMedicalImage
	KindDSDIFFAudio
	KindDSFAudio
	KindDTSAudio
	KindDalvikExecutable
	KindDebianPackage
	KindDjVuDocument
	KindEBMLContainer
	KindEOTFont
	KindExecutableAndLinkableFormat
	KindExtFilesystem
	KindFITSAstronomicalImage
	KindFLACAudio
	KindFLIFImage
	KindFLVVideo
	KindFarbfeldImage
	KindFastTrackerModule
	KindFilesystemEntry
	KindFujifilmRAWImage
	KindGIFImage
	KindGIMPXCFImage
	KindGLTFBinary
	KindGRIBData
	KindGitIndex
	KindGitPack
	KindGzipArchive
	KindGzipData
	KindHDF5Data
	KindICCProfile
	KindICNSIcon
	KindICOCURImage
	KindICalendar
	KindIFFContainer
	KindISO9660Image
	KindISOBaseMedia
	KindIVFVideo
	KindImpulseTrackerModule
	KindJNGImage
	KindJPEG2000Image
	KindJPEGImage
	KindJPEGXLImage
	KindJPEGXRImage
	KindJavaClass
	KindJavaKeyStore
	KindJavaModule
	KindKTXTexture
	KindKeePassDatabase
	KindLHAArchive
	KindLLVMBitcode
	KindLZ4Frame
	KindLZIPArchive
	KindLZMAData
	KindLZOPArchive
	KindLuaBytecode
	KindM3U8Playlist
	KindMIDISequence
	KindMOBIDocument
	KindMPEG2TransportStream
	KindMPEGAudio
	KindMPEGAudioFrame
	KindMPEGProgramStream
	KindMPEGTransportStream
	KindMacBinary
	KindMachOBinary
	KindMachOUniversalBinary
	KindMetafileImage
	KindMicrosoftOutlookEmailFolder
	KindMonkeysAudio
	KindMusepackAudio
	KindNESROM
	KindNetCDFData
	KindNetpbmImage
	KindOLECompoundDocument
	KindORCColumnarData
	KindOggContainer
	KindOlympusRAWImage
	KindOpenEXRImage
	KindOpenTypeFont
	KindPCAPCapture
	KindPCAPNGCapture
	KindPCXImage
	KindPDFDocument
	KindPEMCertificate
	KindPEMPrivateKey
	KindPGPMessage
	KindPKCS12
	KindPNGImage
	KindPhotoshopDocument
	KindPortableExecutable
	KindPostScriptDocument
	KindPythonBytecode
	KindQCOWDiskImage
	KindQOIImage
	KindRARArchive
	KindRIFFContainer
	KindRPMPackage
	KindRadianceHDRImage
	KindRealMedia
	KindRichTextFormatDocument
	KindRubyGemPackage
	KindSQLiteDatabase
	KindSVGImage
	KindScreamTrackerModule
	KindShebangScript
	KindShockwaveFlash
	KindSketchDocument
	KindSnappyFramedData
	KindSquashFSFilesystem
	KindStuffItArchive
	KindSunRasterImage
	KindTAKAudio
	KindTARArchive
	KindTIFFImage
	KindTTAAudio
	KindTextFile
	KindTorrentFile
	KindTrueTypeCollection
	KindTrueTypeFont
	KindUBootImage
	KindUnixCompressArchive
	KindVCard
	KindVHDDiskImage
	KindVHDXDiskImage
	KindVMwareDiskImage
	KindVOCAudio
	KindVirtualBoxDiskImage
	KindWARCFile
	KindWOFF2Font
	KindWOFFFont
	KindWTVVideo
	KindWavPackAudio
	KindWebAssemblyModule
	KindWindowsEventLog
	KindWindowsImagingFormat
	KindWindowsRegistryHive
	KindWindowsShortcut
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
	TypeADTS
	TypeAIFCAudio
	TypeAIFFAudio
	TypeAPPXPackage
	TypeASCIIText
	TypeAVIFImage
	TypeAVIFImageSequence
	TypeAVIVideo
	TypeAdobeDNGDNG
	TypeAndroidAppBundleAAB
	TypeAndroidArchiveAAR
	TypeAndroidPackageAPK
	TypeAndroidPackageXAPK
	TypeAndroidSplitAPKS
	TypeAndroidSystemPackageAPEX
	TypeBGZF
	TypeBigEndian
	TypeBinaryBigEndian
	TypeBinaryLittleEndian
	TypeBlockDevice
	TypeCDAAudio
	TypeCRXVersion2
	TypeCRXVersion3
	TypeCanonRAW3CR3
	TypeCharacterDevice
	TypeCodestream
	TypeCondaPackage
	TypeContainer
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
	TypeEPUBDocument
	TypeEmpty
	TypeEnhancedMetafileEMF
	TypeExt2
	TypeExt3
	TypeExt4
	TypeF4VVideo
	TypeFLACAudio
	TypeFirefoxExtensionXPI
	TypeGIF87a
	TypeGIF89a
	TypeHEIFImage
	TypeIOSApplicationArchiveIPA
	TypeJMOD
	TypeJavaArchiveJAR
	TypeJavaEnterpriseArchiveEAR
	TypeJavaWebArchiveWAR
	TypeKDBX
	TypeKMZArchive
	TypeKTX
	TypeKTX2
	TypeLZMACompressed
	TypeLittleEndian
	TypeM2TS
	TypeM2TSBDAV
	TypeM4VVideo
	TypeMP3
	TypeMP3ID3Tagged
	TypeMP4Video
	TypeMPEG4AudioM4AFamily
	TypeMPEGLayer2
	TypeMPEGLayer3
	TypeMSIXPackage
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
	TypePGMASCII
	TypePGMBinary
	TypePPMASCII
	TypePPMBinary
	TypePSB
	TypePSD
	TypePanasonicRAWRW2
	TypePentaxRAWPEF
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
	TypeTS
	TypeTheoraVideo
	TypeUTF8Text
	TypeUncompressed
	TypeVMDK
	TypeVisualStudioExtensionVSIX
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
	KindARArchive:                   "AR Archive",
	KindASFContainer:                "ASF Container",
	KindAUAudio:                     "AU Audio",
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
	KindAppleiWorkDocument:          "Apple iWork Document",
	KindAvroObjectContainer:         "Avro Object Container",
	KindBAMData:                     "BAM Data",
	KindBMPImage:                    "BMP Image",
	KindBPGImage:                    "BPG Image",
	KindBlenderFile:                 "Blender File",
	KindBtrfsFilesystem:             "Btrfs Filesystem",
	KindBzip2Archive:                "Bzip2 Archive",
	KindCAFAudio:                    "CAF Audio",
	KindCHMDocument:                 "CHM Document",
	KindCPIOArchive:                 "CPIO Archive",
	KindCRAMData:                    "CRAM Data",
	KindCRXBrowserExtension:         "CRX Browser Extension",
	KindCabinetArchive:              "Cabinet Archive",
	KindCanonRAWImage:               "Canon RAW Image",
	KindDDSImage:                    "DDS Image",
	KindDICOMMedicalImage:           "DICOM Medical Image",
	KindDSDIFFAudio:                 "DSDIFF Audio",
	KindDSFAudio:                    "DSF Audio",
	KindDTSAudio:                    "DTS Audio",
	KindDalvikExecutable:            "Dalvik Executable",
	KindDebianPackage:               "Debian Package",
	KindDjVuDocument:                "DjVu Document",
	KindEBMLContainer:               "EBML Container",
	KindEOTFont:                     "EOT Font",
	KindExecutableAndLinkableFormat: "Executable and Linkable Format",
	KindExtFilesystem:               "ext Filesystem",
	KindFITSAstronomicalImage:       "FITS Astronomical Image",
	KindFLACAudio:                   "FLAC Audio",
	KindFLIFImage:                   "FLIF Image",
	KindFLVVideo:                    "FLV Video",
	KindFarbfeldImage:               "Farbfeld Image",
	KindFastTrackerModule:           "FastTracker Module",
	KindFilesystemEntry:             "Filesystem Entry",
	KindFujifilmRAWImage:            "Fujifilm RAW Image",
	KindGIFImage:                    "GIF Image",
	KindGIMPXCFImage:                "GIMP XCF Image",
	KindGLTFBinary:                  "glTF Binary",
	KindGRIBData:                    "GRIB Data",
	KindGitIndex:                    "Git Index",
	KindGitPack:                     "Git Pack",
	KindGzipArchive:                 "Gzip Archive",
	KindGzipData:                    "Gzip Data",
	KindHDF5Data:                    "HDF5 Data",
	KindICCProfile:                  "ICC Profile",
	KindICNSIcon:                    "ICNS Icon",
	KindICOCURImage:                 "ICO/CUR Image",
	KindICalendar:                   "iCalendar",
	KindIFFContainer:                "IFF Container",
	KindISO9660Image:                "ISO 9660 Image",
	KindISOBaseMedia:                "ISO Base Media",
	KindIVFVideo:                    "IVF Video",
	KindImpulseTrackerModule:        "Impulse Tracker Module",
	KindJNGImage:                    "JNG Image",
	KindJPEG2000Image:               "JPEG 2000 Image",
	KindJPEGImage:                   "JPEG Image",
	KindJPEGXLImage:                 "JPEG XL Image",
	KindJPEGXRImage:                 "JPEG XR Image",
	KindJavaClass:                   "Java Class",
	KindJavaKeyStore:                "Java KeyStore",
	KindJavaModule:                  "Java Module",
	KindKTXTexture:                  "KTX Texture",
	KindKeePassDatabase:             "KeePass Database",
	KindLHAArchive:                  "LHA Archive",
	KindLLVMBitcode:                 "LLVM Bitcode",
	KindLZ4Frame:                    "LZ4 Frame",
	KindLZIPArchive:                 "LZIP Archive",
	KindLZMAData:                    "LZMA Data",
	KindLZOPArchive:                 "LZOP Archive",
	KindLuaBytecode:                 "Lua Bytecode",
	KindM3U8Playlist:                "M3U8 Playlist",
	KindMIDISequence:                "MIDI Sequence",
	KindMOBIDocument:                "MOBI Document",
	KindMPEG2TransportStream:        "MPEG Transport Stream",
	KindMPEGAudio:                   "MPEG Audio",
	KindMPEGAudioFrame:              "MPEG Audio",
	KindMPEGProgramStream:           "MPEG Program Stream",
	KindMPEGTransportStream:         "MPEG Transport Stream",
	KindMacBinary:                   "MacBinary",
	KindMachOBinary:                 "Mach-O Binary",
	KindMachOUniversalBinary:        "Mach-O Universal Binary",
	KindMetafileImage:               "Metafile Image",
	KindMicrosoftOutlookEmailFolder: "Microsoft Outlook Email Folder",
	KindMonkeysAudio:                "Monkey's Audio",
	KindMusepackAudio:               "Musepack Audio",
	KindNESROM:                      "NES ROM",
	KindNetCDFData:                  "NetCDF Data",
	KindNetpbmImage:                 "Netpbm Image",
	KindOLECompoundDocument:         "OLE Compound Document",
	KindORCColumnarData:             "ORC Columnar Data",
	KindOggContainer:                "Ogg Container",
	KindOlympusRAWImage:             "Olympus RAW Image",
	KindOpenEXRImage:                "OpenEXR Image",
	KindOpenTypeFont:                "OpenType Font",
	KindPCAPCapture:                 "PCAP Capture",
	KindPCAPNGCapture:               "PCAPNG Capture",
	KindPCXImage:                    "PCX Image",
	KindPDFDocument:                 "PDF Document",
	KindPEMCertificate:              "PEM Certificate",
	KindPEMPrivateKey:               "PEM Private Key",
	KindPGPMessage:                  "PGP Message",
	KindPKCS12:                      "PKCS#12",
	KindPNGImage:                    "PNG Image",
	KindPhotoshopDocument:           "Photoshop Document",
	KindPortableExecutable:          "Portable Executable",
	KindPostScriptDocument:          "PostScript Document",
	KindPythonBytecode:              "Python Bytecode",
	KindQCOWDiskImage:               "QCOW Disk Image",
	KindQOIImage:                    "QOI Image",
	KindRARArchive:                  "RAR Archive",
	KindRIFFContainer:               "RIFF Container",
	KindRPMPackage:                  "RPM Package",
	KindRadianceHDRImage:            "Radiance HDR Image",
	KindRealMedia:                   "RealMedia",
	KindRichTextFormatDocument:      "Rich Text Format Document",
	KindRubyGemPackage:              "RubyGem Package",
	KindSQLiteDatabase:              "SQLite Database",
	KindSVGImage:                    "SVG Image",
	KindScreamTrackerModule:         "Scream Tracker Module",
	KindShebangScript:               "Shebang Script",
	KindShockwaveFlash:              "Shockwave Flash",
	KindSketchDocument:              "Sketch Document",
	KindSnappyFramedData:            "Snappy Framed Data",
	KindSquashFSFilesystem:          "SquashFS Filesystem",
	KindStuffItArchive:              "StuffIt Archive",
	KindSunRasterImage:              "Sun Raster Image",
	KindTAKAudio:                    "TAK Audio",
	KindTARArchive:                  "TAR Archive",
	KindTIFFImage:                   "TIFF Image",
	KindTTAAudio:                    "TTA Audio",
	KindTextFile:                    "Text File",
	KindTorrentFile:                 "Torrent File",
	KindTrueTypeCollection:          "TrueType Collection",
	KindTrueTypeFont:                "TrueType Font",
	KindUBootImage:                  "U-Boot Image",
	KindUnixCompressArchive:         "Unix Compress Archive",
	KindVCard:                       "vCard",
	KindVHDDiskImage:                "VHD Disk Image",
	KindVHDXDiskImage:               "VHDX Disk Image",
	KindVMwareDiskImage:             "VMware Disk Image",
	KindVOCAudio:                    "VOC Audio",
	KindVirtualBoxDiskImage:         "VirtualBox Disk Image",
	KindWARCFile:                    "WARC File",
	KindWOFF2Font:                   "WOFF2 Font",
	KindWOFFFont:                    "WOFF Font",
	KindWTVVideo:                    "WTV Video",
	KindWavPackAudio:                "WavPack Audio",
	KindWebAssemblyModule:           "WebAssembly Module",
	KindWindowsEventLog:             "Windows Event Log",
	KindWindowsImagingFormat:        "Windows Imaging Format",
	KindWindowsRegistryHive:         "Windows Registry Hive",
	KindWindowsShortcut:             "Windows Shortcut",
	KindXARArchive:                  "XAR Archive",
	KindXFSFilesystem:               "XFS Filesystem",
	KindXMLDocument:                 "XML Document",
	KindXZArchive:                   "XZ Archive",
	KindZIPArchive:                  "ZIP Archive",
	KindZlibData:                    "Zlib Data",
	KindZstandardArchive:            "Zstandard Archive",
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
	TypeADTS:                     "ADTS",
	TypeAIFCAudio:                "AIFC Audio",
	TypeAIFFAudio:                "AIFF Audio",
	TypeAPPXPackage:              "APPX Package",
	TypeASCIIText:                "ASCII Text",
	TypeAVIFImage:                "AVIF Image",
	TypeAVIFImageSequence:        "AVIF Image Sequence",
	TypeAVIVideo:                 "AVI Video",
	TypeAdobeDNGDNG:              "Adobe DNG (DNG)",
	TypeAndroidAppBundleAAB:      "Android App Bundle (AAB)",
	TypeAndroidArchiveAAR:        "Android Archive (AAR)",
	TypeAndroidPackageAPK:        "Android Package (APK)",
	TypeAndroidPackageXAPK:       "Android Package (XAPK)",
	TypeAndroidSplitAPKS:         "Android Split APK Set (APKS)",
	TypeAndroidSystemPackageAPEX: "Android System Package (APEX)",
	TypeBGZF:                     "BGZF",
	TypeBigEndian:                "Big-Endian",
	TypeBinaryBigEndian:          "Binary Big-Endian",
	TypeBinaryLittleEndian:       "Binary Little-Endian",
	TypeBlockDevice:              "Block Device",
	TypeCDAAudio:                 "CD Audio",
	TypeCRXVersion2:              "Version 2",
	TypeCRXVersion3:              "Version 3",
	TypeCanonRAW3CR3:             "Canon RAW 3 (CR3)",
	TypeCharacterDevice:          "Character Device",
	TypeCodestream:               "Codestream",
	TypeCondaPackage:             "Conda Package",
	TypeContainer:                "Container",
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
	TypeEPUBDocument:             "EPUB Document",
	TypeEmpty:                    "Empty",
	TypeEnhancedMetafileEMF:      "Enhanced Metafile (EMF)",
	TypeExt2:                     "ext2",
	TypeExt3:                     "ext3",
	TypeExt4:                     "ext4",
	TypeF4VVideo:                 "F4V Video",
	TypeFLACAudio:                "FLAC Audio",
	TypeFirefoxExtensionXPI:      "Firefox Extension (XPI)",
	TypeGIF87a:                   "GIF87a",
	TypeGIF89a:                   "GIF89a",
	TypeHEIFImage:                "HEIF Image",
	TypeIOSApplicationArchiveIPA: "iOS Application Archive (IPA)",
	TypeJMOD:                     "JMOD",
	TypeJavaArchiveJAR:           "Java Archive (JAR)",
	TypeJavaEnterpriseArchiveEAR: "Java Enterprise Archive (EAR)",
	TypeJavaWebArchiveWAR:        "Java Web Archive (WAR)",
	TypeKDBX:                     "KDBX",
	TypeKMZArchive:               "KMZ Archive",
	TypeKTX:                      "KTX",
	TypeKTX2:                     "KTX2",
	TypeLZMACompressed:           "LZMA Compressed",
	TypeLittleEndian:             "Little-Endian",
	TypeM2TS:                     "M2TS",
	TypeM2TSBDAV:                 "M2TS/BDAV",
	TypeM4VVideo:                 "M4V Video",
	TypeMP3:                      "MP3",
	TypeMP3ID3Tagged:             "MP3 (ID3 Tagged)",
	TypeMP4Video:                 "MP4 Video",
	TypeMPEG4AudioM4AFamily:      "MPEG-4 Audio (M4A Family)",
	TypeMPEGLayer2:               "MPEG Layer II",
	TypeMPEGLayer3:               "MPEG Layer III",
	TypeMSIXPackage:              "MSIX Package",
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
	TypePGMASCII:                                        "PGM ASCII",
	TypePGMBinary:                                       "PGM binary",
	TypePPMASCII:                                        "PPM ASCII",
	TypePPMBinary:                                       "PPM binary",
	TypePSB:                                             "PSB",
	TypePSD:                                             "PSD",
	TypePanasonicRAWRW2:                                 "Panasonic RAW (RW2)",
	TypePentaxRAWPEF:                                    "Pentax RAW (PEF)",
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
	TypeSpanned:                                         "Spanned",
	TypeSpecial:                                         "Special",
	TypeSpeexAudio:                                      "Speex Audio",
	TypeStreamVersion7:                                  "Stream Version 7",
	TypeStreamVersion8:                                  "Stream Version 8",
	TypeSymbolicLink:                                    "Symbolic Link",
	TypeTS:                                              "TS",
	TypeTheoraVideo:                                     "Theora Video",
	TypeUTF8Text:                                        "UTF-8 Text",
	TypeUncompressed:                                    "Uncompressed",
	TypeVMDK:                                            "VMDK",
	TypeVisualStudioExtensionVSIX:                       "Visual Studio Extension (VSIX)",
	TypeVorbisAudio:                                     "Vorbis Audio",
	TypeWAVAudio:                                        "WAV Audio",
	TypeWebM:                                            "WebM",
	TypeWebPImage:                                       "WebP Image",
	TypeWindowsCursor:                                   "Windows Cursor",
	TypeWindowsIcon:                                     "Windows Icon",
	TypeWindowsMetafileWMF:                              "Windows Metafile (WMF)",
	TypeWrapper:                                         "Wrapper",
	TypeZlibCompressed:                                  "Zlib Compressed",
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
