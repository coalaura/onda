package types

type KindID uint16
type TypeID uint16

const (
	KindUnknown KindID = iota
	KindFilesystemEntry
	KindZIPArchive
	KindUnixCompressArchive
	Kind7ZipArchive
	KindRARArchive
	KindGzipArchive
	KindBzip2Archive
	KindXZArchive
	KindZstandardArchive
	KindLZ4Frame
	KindLZIPArchive
	KindLZOPArchive
	KindSnappyFramedData
	KindCabinetArchive
	KindXARArchive
	KindWindowsImagingFormat
	KindARArchive
	KindDebianPackage
	KindSquashFSFilesystem
	KindCPIOArchive
	KindRPMPackage
	KindFLACAudio
	KindMIDISequence
	KindAUAudio
	KindAMRAudio
	KindAMRWBAudio
	KindAACAudio
	KindCAFAudio
	KindWavPackAudio
	KindMonkeysAudio
	KindDSFAudio
	KindTAKAudio
	KindVOCAudio
	KindMusepackAudio
	KindImpulseTrackerModule
	KindFastTrackerModule
	KindScreamTrackerModule
	KindPDFDocument
	KindPostScriptDocument
	KindRichTextFormatDocument
	KindSQLiteDatabase
	KindHDF5Data
	KindPGPMessage
	KindApacheParquet
	KindApacheArrowFile
	KindFITSAstronomicalImage
	KindDICOMMedicalImage
	KindWindowsRegistryHive
	KindAppleBinaryPropertyList
	KindKeePassDatabase
	KindAvroObjectContainer
	KindORCColumnarData
	KindQCOWDiskImage
	KindVHDXDiskImage
	KindVMwareDiskImage
	KindAndroidSparseImage
	KindVirtualBoxDiskImage
	KindISO9660Image
	KindTrueTypeFont
	KindOpenTypeFont
	KindWOFFFont
	KindWOFF2Font
	KindTrueTypeCollection
	KindPNGImage
	KindJPEGImage
	KindGIFImage
	KindBMPImage
	KindTIFFImage
	KindQOIImage
	KindPhotoshopDocument
	KindRadianceHDRImage
	KindOpenEXRImage
	KindDDSImage
	KindJPEGXLImage
	KindKTXTexture
	KindICNSIcon
	KindFarbfeldImage
	KindJPEG2000Image
	KindBPGImage
	KindJPEGXRImage
	KindSunRasterImage
	KindCRXBrowserExtension
	KindDalvikExecutable
	KindAndroidOAT
	KindAndroidODEX
	KindJavaClass
	KindJavaKeyStore
	KindWebAssemblyModule
	KindNESROM
	KindPCAPCapture
	KindPCAPNGCapture
	KindShockwaveFlash
	KindLLVMBitcode
	KindWindowsShortcut
	KindFLVVideo
	KindIVFVideo
	KindMPEGProgramStream
	KindASFContainer
	KindMPEGAudio
	KindMachOBinary
	KindMachOUniversalBinary
	KindExecutableAndLinkableFormat
	KindMetafileImage
	KindIFFContainer
	KindISOBaseMedia
	KindRIFFContainer
	KindOLECompoundDocument
	KindEBMLContainer
	KindPortableExecutable
	KindNetpbmImage
	KindTARArchive
	KindMPEGTransportStream
	KindICOCURImage
	KindOggContainer
	KindGIMPXCFImage
	KindCanonRAWImage
	KindOlympusRAWImage
	KindFujifilmRAWImage
	KindRealMedia
	KindWindowsEventLog
	KindAndroidBootImage
	KindUBootImage
	KindWARCFile
)

const (
	TypeNone TypeID = iota
	TypeSymbolicLink
	TypeDirectory
	TypeNamedPipe
	TypeSocket
	TypeCharacterDevice
	TypeBlockDevice
	TypeSpecial
	TypeEmpty
	TypeSpanned
	TypeRAR4
	TypeRAR5
	TypeNewASCII
	TypeNewASCIIWithCRC
	TypeOldASCII
	TypeBinaryBigEndian
	TypeBinaryLittleEndian
	TypeADTS
	TypeStreamVersion7
	TypeStreamVersion8
	TypeKDBX
	TypeQCOW2
	TypeVMDK
	TypeGIF87a
	TypeGIF89a
	TypeLittleEndian
	TypeBigEndian
	TypePSD
	TypePSB
	TypeCodestream
	TypeContainer
	TypeKTX
	TypeKTX2
	TypeDEX035
	TypeDEX036
	TypeDEX037
	TypeDEX038
	TypeDEX039
	TypeDEX040
	TypeDEX041
	TypeNanosecondLittleEndian
	TypeNanosecondBigEndian
	TypeUncompressed
	TypeZlibCompressed
	TypeLZMACompressed
	TypeWrapper
	TypeMP3ID3Tagged
	TypeMP3
	Type32BitBigEndian
	Type32BitLittleEndian
	Type64BitBigEndian
	Type64BitLittleEndian
	Type64Bit
	TypeEnhancedMetafileEMF
	TypeWindowsMetafileWMF
	TypeAIFFAudio
	TypeAIFCAudio
	TypeAVIFImage
	TypeAVIFImageSequence
	TypeHEIFImage
	Type3GPPMedia
	TypeMPEG4AudioM4AFamily
	TypeQuickTimeMovie
	TypeMP4Video
	TypeWAVAudio
	TypeAVIVideo
	TypeWebPImage
	TypeWebM
	TypeMatroska
	TypeEPUBDocument
	TypeOpenDocumentTextODT
	TypeOpenDocumentSpreadsheetODS
	TypeOpenDocumentPresentationODP
	TypeMicrosoftWordDocumentDOCX
	TypeMicrosoftExcelSpreadsheetXLSX
	TypeMicrosoftPowerPointPresentationPPTX
	TypeAndroidPackageAPK
	TypeJavaArchiveJAR
	TypeIOSApplicationArchiveIPA
	TypeOpusAudio
	TypeVorbisAudio
	TypeSpeexAudio
	TypeTheoraVideo
	TypeFLACAudio
	TypePBMASCII
	TypePGMASCII
	TypePPMASCII
	TypePBMBinary
	TypePGMBinary
	TypePPMBinary
	TypePAM
	TypeTS
	TypeM2TS
	TypeWindowsIcon
	TypeWindowsCursor
	TypeELF
	TypeELF32
	TypeELF64
	TypeELF32LittleEndian
	TypeELF32BigEndian
	TypeELF64LittleEndian
	TypeELF64BigEndian
	TypePE32X86
	TypePE32X8664
	TypePE32ARM
	TypePE32ARMv7
	TypePE32ARM64
	TypePE32Itanium
	TypePE32Unknown
	TypePE32PlusX8664
	TypePE32PlusARM64
	TypePE32PlusUnknown
	TypeMicrosoftWordDocumentDOC
	TypeMicrosoftExcelWorkbookXLS
	TypeMicrosoftPowerPointPresentationPPT
	TypeAndroidAppBundleAAB
	TypeVisualStudioExtensionVSIX
	TypeKMZArchive
	TypeFirefoxExtensionXPI
)

var kindNames = [...]string{
	KindUnknown:                     "Unknown",
	KindFilesystemEntry:             "Filesystem Entry",
	KindZIPArchive:                  "ZIP Archive",
	KindUnixCompressArchive:         "Unix Compress Archive",
	Kind7ZipArchive:                 "7-Zip Archive",
	KindRARArchive:                  "RAR Archive",
	KindGzipArchive:                 "Gzip Archive",
	KindBzip2Archive:                "Bzip2 Archive",
	KindXZArchive:                   "XZ Archive",
	KindZstandardArchive:            "Zstandard Archive",
	KindLZ4Frame:                    "LZ4 Frame",
	KindLZIPArchive:                 "LZIP Archive",
	KindLZOPArchive:                 "LZOP Archive",
	KindSnappyFramedData:            "Snappy Framed Data",
	KindCabinetArchive:              "Cabinet Archive",
	KindXARArchive:                  "XAR Archive",
	KindWindowsImagingFormat:        "Windows Imaging Format",
	KindARArchive:                   "AR Archive",
	KindDebianPackage:               "Debian Package",
	KindSquashFSFilesystem:          "SquashFS Filesystem",
	KindCPIOArchive:                 "CPIO Archive",
	KindRPMPackage:                  "RPM Package",
	KindFLACAudio:                   "FLAC Audio",
	KindMIDISequence:                "MIDI Sequence",
	KindAUAudio:                     "AU Audio",
	KindAMRAudio:                    "AMR Audio",
	KindAMRWBAudio:                  "AMR-WB Audio",
	KindAACAudio:                    "AAC Audio",
	KindCAFAudio:                    "CAF Audio",
	KindWavPackAudio:                "WavPack Audio",
	KindMonkeysAudio:                "Monkey's Audio",
	KindDSFAudio:                    "DSF Audio",
	KindTAKAudio:                    "TAK Audio",
	KindVOCAudio:                    "VOC Audio",
	KindMusepackAudio:               "Musepack Audio",
	KindImpulseTrackerModule:        "Impulse Tracker Module",
	KindFastTrackerModule:           "FastTracker Module",
	KindScreamTrackerModule:         "Scream Tracker Module",
	KindPDFDocument:                 "PDF Document",
	KindPostScriptDocument:          "PostScript Document",
	KindRichTextFormatDocument:      "Rich Text Format Document",
	KindSQLiteDatabase:              "SQLite Database",
	KindHDF5Data:                    "HDF5 Data",
	KindPGPMessage:                  "PGP Message",
	KindApacheParquet:               "Apache Parquet",
	KindApacheArrowFile:             "Apache Arrow File",
	KindFITSAstronomicalImage:       "FITS Astronomical Image",
	KindDICOMMedicalImage:           "DICOM Medical Image",
	KindWindowsRegistryHive:         "Windows Registry Hive",
	KindAppleBinaryPropertyList:     "Apple Binary Property List",
	KindKeePassDatabase:             "KeePass Database",
	KindAvroObjectContainer:         "Avro Object Container",
	KindORCColumnarData:             "ORC Columnar Data",
	KindQCOWDiskImage:               "QCOW Disk Image",
	KindVHDXDiskImage:               "VHDX Disk Image",
	KindVMwareDiskImage:             "VMware Disk Image",
	KindAndroidSparseImage:          "Android Sparse Image",
	KindVirtualBoxDiskImage:         "VirtualBox Disk Image",
	KindISO9660Image:                "ISO 9660 Image",
	KindTrueTypeFont:                "TrueType Font",
	KindOpenTypeFont:                "OpenType Font",
	KindWOFFFont:                    "WOFF Font",
	KindWOFF2Font:                   "WOFF2 Font",
	KindTrueTypeCollection:          "TrueType Collection",
	KindPNGImage:                    "PNG Image",
	KindJPEGImage:                   "JPEG Image",
	KindGIFImage:                    "GIF Image",
	KindBMPImage:                    "BMP Image",
	KindTIFFImage:                   "TIFF Image",
	KindQOIImage:                    "QOI Image",
	KindPhotoshopDocument:           "Photoshop Document",
	KindRadianceHDRImage:            "Radiance HDR Image",
	KindOpenEXRImage:                "OpenEXR Image",
	KindDDSImage:                    "DDS Image",
	KindJPEGXLImage:                 "JPEG XL Image",
	KindKTXTexture:                  "KTX Texture",
	KindICNSIcon:                    "ICNS Icon",
	KindFarbfeldImage:               "Farbfeld Image",
	KindJPEG2000Image:               "JPEG 2000 Image",
	KindBPGImage:                    "BPG Image",
	KindJPEGXRImage:                 "JPEG XR Image",
	KindSunRasterImage:              "Sun Raster Image",
	KindCRXBrowserExtension:         "CRX Browser Extension",
	KindDalvikExecutable:            "Dalvik Executable",
	KindAndroidOAT:                  "Android OAT",
	KindAndroidODEX:                 "Android ODEX",
	KindJavaClass:                   "Java Class",
	KindJavaKeyStore:                "Java KeyStore",
	KindWebAssemblyModule:           "WebAssembly Module",
	KindNESROM:                      "NES ROM",
	KindPCAPCapture:                 "PCAP Capture",
	KindPCAPNGCapture:               "PCAPNG Capture",
	KindShockwaveFlash:              "Shockwave Flash",
	KindLLVMBitcode:                 "LLVM Bitcode",
	KindWindowsShortcut:             "Windows Shortcut",
	KindFLVVideo:                    "FLV Video",
	KindIVFVideo:                    "IVF Video",
	KindMPEGProgramStream:           "MPEG Program Stream",
	KindASFContainer:                "ASF Container",
	KindMPEGAudio:                   "MPEG Audio",
	KindMachOBinary:                 "Mach-O Binary",
	KindMachOUniversalBinary:        "Mach-O Universal Binary",
	KindExecutableAndLinkableFormat: "Executable and Linkable Format",
	KindMetafileImage:               "Metafile Image",
	KindIFFContainer:                "IFF Container",
	KindISOBaseMedia:                "ISO Base Media",
	KindRIFFContainer:               "RIFF Container",
	KindOLECompoundDocument:         "OLE Compound Document",
	KindEBMLContainer:               "EBML Container",
	KindPortableExecutable:          "Portable Executable",
	KindNetpbmImage:                 "Netpbm Image",
	KindTARArchive:                  "TAR Archive",
	KindMPEGTransportStream:         "MPEG Transport Stream",
	KindICOCURImage:                 "ICO/CUR Image",
	KindOggContainer:                "Ogg Container",
	KindGIMPXCFImage:                "GIMP XCF Image",
	KindCanonRAWImage:               "Canon RAW Image",
	KindOlympusRAWImage:             "Olympus RAW Image",
	KindFujifilmRAWImage:            "Fujifilm RAW Image",
	KindRealMedia:                   "RealMedia",
	KindWindowsEventLog:             "Windows Event Log",
	KindAndroidBootImage:            "Android Boot Image",
	KindUBootImage:                  "U-Boot Image",
	KindWARCFile:                    "WARC File",
}

var typeNames = [...]string{
	TypeNone:                                "",
	TypeSymbolicLink:                        "Symbolic Link",
	TypeDirectory:                           "Directory",
	TypeNamedPipe:                           "Named Pipe",
	TypeSocket:                              "Socket",
	TypeCharacterDevice:                     "Character Device",
	TypeBlockDevice:                         "Block Device",
	TypeSpecial:                             "Special",
	TypeEmpty:                               "Empty",
	TypeSpanned:                             "Spanned",
	TypeRAR4:                                "RAR4",
	TypeRAR5:                                "RAR5",
	TypeNewASCII:                            "New ASCII",
	TypeNewASCIIWithCRC:                     "New ASCII with CRC",
	TypeOldASCII:                            "Old ASCII",
	TypeBinaryBigEndian:                     "Binary Big-Endian",
	TypeBinaryLittleEndian:                  "Binary Little-Endian",
	TypeADTS:                                "ADTS",
	TypeStreamVersion7:                      "Stream Version 7",
	TypeStreamVersion8:                      "Stream Version 8",
	TypeKDBX:                                "KDBX",
	TypeQCOW2:                               "QCOW2",
	TypeVMDK:                                "VMDK",
	TypeGIF87a:                              "GIF87a",
	TypeGIF89a:                              "GIF89a",
	TypeLittleEndian:                        "Little-Endian",
	TypeBigEndian:                           "Big-Endian",
	TypePSD:                                 "PSD",
	TypePSB:                                 "PSB",
	TypeCodestream:                          "Codestream",
	TypeContainer:                           "Container",
	TypeKTX:                                 "KTX",
	TypeKTX2:                                "KTX2",
	TypeDEX035:                              "DEX 035",
	TypeDEX036:                              "DEX 036",
	TypeDEX037:                              "DEX 037",
	TypeDEX038:                              "DEX 038",
	TypeDEX039:                              "DEX 039",
	TypeDEX040:                              "DEX 040",
	TypeDEX041:                              "DEX 041",
	TypeNanosecondLittleEndian:              "Nanosecond Little-Endian",
	TypeNanosecondBigEndian:                 "Nanosecond Big-Endian",
	TypeUncompressed:                        "Uncompressed",
	TypeZlibCompressed:                      "Zlib Compressed",
	TypeLZMACompressed:                      "LZMA Compressed",
	TypeWrapper:                             "Wrapper",
	TypeMP3ID3Tagged:                        "MP3 (ID3 Tagged)",
	TypeMP3:                                 "MP3",
	Type32BitBigEndian:                      "32-bit Big-Endian",
	Type32BitLittleEndian:                   "32-bit Little-Endian",
	Type64BitBigEndian:                      "64-bit Big-Endian",
	Type64BitLittleEndian:                   "64-bit Little-Endian",
	Type64Bit:                               "64-bit",
	TypeEnhancedMetafileEMF:                 "Enhanced Metafile (EMF)",
	TypeWindowsMetafileWMF:                  "Windows Metafile (WMF)",
	TypeAIFFAudio:                           "AIFF Audio",
	TypeAIFCAudio:                           "AIFC Audio",
	TypeAVIFImage:                           "AVIF Image",
	TypeAVIFImageSequence:                   "AVIF Image Sequence",
	TypeHEIFImage:                           "HEIF Image",
	Type3GPPMedia:                           "3GPP Media",
	TypeMPEG4AudioM4AFamily:                 "MPEG-4 Audio (M4A Family)",
	TypeQuickTimeMovie:                      "QuickTime Movie",
	TypeMP4Video:                            "MP4 Video",
	TypeWAVAudio:                            "WAV Audio",
	TypeAVIVideo:                            "AVI Video",
	TypeWebPImage:                           "WebP Image",
	TypeWebM:                                "WebM",
	TypeMatroska:                            "Matroska",
	TypeEPUBDocument:                        "EPUB Document",
	TypeOpenDocumentTextODT:                 "OpenDocument Text (ODT)",
	TypeOpenDocumentSpreadsheetODS:          "OpenDocument Spreadsheet (ODS)",
	TypeOpenDocumentPresentationODP:         "OpenDocument Presentation (ODP)",
	TypeMicrosoftWordDocumentDOCX:           "Microsoft Word Document (DOCX)",
	TypeMicrosoftExcelSpreadsheetXLSX:       "Microsoft Excel Spreadsheet (XLSX)",
	TypeMicrosoftPowerPointPresentationPPTX: "Microsoft PowerPoint Presentation (PPTX)",
	TypeAndroidPackageAPK:                   "Android Package (APK)",
	TypeJavaArchiveJAR:                      "Java Archive (JAR)",
	TypeIOSApplicationArchiveIPA:            "iOS Application Archive (IPA)",
	TypeOpusAudio:                           "Opus Audio",
	TypeVorbisAudio:                         "Vorbis Audio",
	TypeSpeexAudio:                          "Speex Audio",
	TypeTheoraVideo:                         "Theora Video",
	TypeFLACAudio:                           "FLAC Audio",
	TypePBMASCII:                            "PBM ASCII",
	TypePGMASCII:                            "PGM ASCII",
	TypePPMASCII:                            "PPM ASCII",
	TypePBMBinary:                           "PBM binary",
	TypePGMBinary:                           "PGM binary",
	TypePPMBinary:                           "PPM binary",
	TypePAM:                                 "PAM",
	TypeTS:                                  "TS",
	TypeM2TS:                                "M2TS",
	TypeWindowsIcon:                         "Windows Icon",
	TypeWindowsCursor:                       "Windows Cursor",
	TypeELF:                                 "ELF",
	TypeELF32:                               "ELF32",
	TypeELF64:                               "ELF64",
	TypeELF32LittleEndian:                   "ELF32 Little-Endian",
	TypeELF32BigEndian:                      "ELF32 Big-Endian",
	TypeELF64LittleEndian:                   "ELF64 Little-Endian",
	TypeELF64BigEndian:                      "ELF64 Big-Endian",
	TypePE32X86:                             "PE32 x86",
	TypePE32X8664:                           "PE32 x86-64",
	TypePE32ARM:                             "PE32 ARM",
	TypePE32ARMv7:                           "PE32 ARMv7",
	TypePE32ARM64:                           "PE32 ARM64",
	TypePE32Itanium:                         "PE32 Itanium",
	TypePE32Unknown:                         "PE32 Unknown",
	TypePE32PlusX8664:                       "PE32+ x86-64",
	TypePE32PlusARM64:                       "PE32+ ARM64",
	TypePE32PlusUnknown:                     "PE32+ Unknown",
	TypeMicrosoftWordDocumentDOC:            "Microsoft Word Document (DOC)",
	TypeMicrosoftExcelWorkbookXLS:           "Microsoft Excel Workbook (XLS)",
	TypeMicrosoftPowerPointPresentationPPT:  "Microsoft PowerPoint Presentation (PPT)",
	TypeAndroidAppBundleAAB:                 "Android App Bundle (AAB)",
	TypeVisualStudioExtensionVSIX:           "Visual Studio Extension (VSIX)",
	TypeKMZArchive:                          "KMZ Archive",
	TypeFirefoxExtensionXPI:                 "Firefox Extension (XPI)",
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
