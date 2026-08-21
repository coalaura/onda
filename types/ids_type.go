package types

// TypeID values are generated in alphabetical order and are not stable across releases.
type TypeID uint16

const (
	TypeNone TypeID = iota
	Type32Bit
	Type32BitBigEndian
	Type32BitLittleEndian
	Type3G2Video
	Type3GPPVideo
	Type3MFDocument
	Type64Bit
	Type64BitBigEndian
	Type64BitLittleEndian
	TypeAC3Audio
	TypeAdobeDNG
	TypeAfterEffectsProject
	TypeAIFCAudio
	TypeAIFFAudio
	TypeAndroidAppBundle
	TypeAndroidArchive
	TypeAndroidPackage
	TypeAndroidPackageX
	TypeAndroidSplitAPKSet
	TypeAndroidSystemPackage
	TypeAndroidVendorBoot
	TypeAngularComponent
	TypeApacheLog
	TypeAPPXPackage
	TypeArchLinuxPackage
	TypeASCII
	TypeAsciiDoc
	TypeAtari800
	TypeAudioDataTransportStream
	TypeAVIFImage
	TypeAVIFImageSequence
	TypeAVIVideo
	TypeBashScript
	TypeBatchScript
	TypeBDAV
	TypeBGZF
	TypeBigEndian
	TypeBigTIFF
	TypeBinary
	TypeBinaryBigEndian
	TypeBinaryLittleEndian
	TypeBitwig
	TypeBlackmagicRAWVideo
	TypeBlockDevice
	TypeByteSwapped
	TypeCanonRAW3Image
	TypeCatalog
	TypeCDAAudio
	TypeCeltx
	TypeCharacterDevice
	TypeClojureScript
	TypeCMakeScript
	TypeCodestream
	TypeCOFFObject
	TypeComicBook
	TypeCommitMessage
	TypeCommodore64
	TypeCondaPackage
	TypeContainer
	TypeCopyOnWrite
	TypeCoqSource
	TypeCorelDRAWDocument
	TypeCPPSource
	TypeCSharpSource
	TypeCSource
	TypeCSS
	TypeCSVData
	TypeCubaseProject
	TypeDartSource
	TypeDescription
	TypeDiffPatch
	TypeDirectory
	TypeDjangoTemplate
	TypeDockerComposeConfiguration
	TypeDockerfile
	TypeDownloadableSounds
	TypeDSA
	TypeDSIKModule
	TypeEAC3Audio
	TypeEC
	TypeElixirScript
	TypeEmacsLispScript
	TypeEmpty
	TypeEncrypted
	TypeEnhancedMetafile
	TypeEPUBDocument
	TypeExt2
	TypeExt3
	TypeExt4
	TypeF4VVideo
	TypeFabricMod
	TypeFAT12
	TypeFAT16
	TypeFAT32
	TypeFirefoxExtension
	TypeFishScript
	TypeFLACAudio
	TypeFlatpak
	TypeForgeMod
	TypeFSharpSource
	TypeGameMakerStudio
	TypeGameMakerStudio2
	TypeGeoJSON
	TypeGeoPackage
	TypeGIF87a
	TypeGIF89a
	TypeGitConfig
	TypeGitDiff
	TypeGoSource
	TypeGoTemplate
	TypeGraphQL
	TypeGraphQLSchema
	TypeGuileScript
	TypeHalfLife1
	TypeHalfLife2
	TypeHandlebarsTemplate
	TypeHaskellSource
	TypeHDPhoto
	TypeHE
	TypeHEIFImage
	TypeHighEntropy
	TypeHTTPLog
	TypeIdrisSource
	TypeILBMImage
	TypeIMMMFormat
	TypeINIConfiguration
	TypeIOSApplicationArchive
	TypeIWAD
	TypeJavaArchive
	TypeJavaEnterpriseArchive
	TypeJavaScriptSource
	TypeJavaSource
	TypeJavaWebArchive
	TypeJinjaTemplate
	TypeJSON5Source
	TypeJSONCSource
	TypeKDMV
	TypeKerasModel
	TypeKMZArchive
	TypeKotlinSource
	TypeKritaDocument
	TypeLargeDocument
	TypeLaTeXDocument
	TypeLatin1
	TypeLeanSource
	TypeLegacy
	TypeList
	TypeLittleEndian
	TypeLogData
	TypeLOVEGame
	TypeLuaScript
	TypeLZ4Legacy
	TypeLZMACompressed
	TypeM4VVideo
	TypeMakefile
	TypeMarkdownDocument
	TypeMatroskaVideo
	TypeMBTiles
	TypeMicrosoftExcelAddIn
	TypeMicrosoftExcelMacroEnabledTemplate
	TypeMicrosoftExcelMacroEnabledWorkbook
	TypeMicrosoftExcelTemplate
	TypeMicrosoftExcelWorkbook
	TypeMicrosoftInstaller
	TypeMicrosoftInstallerPatch
	TypeMicrosoftOutlookMessage
	TypeMicrosoftPowerPointAddIn
	TypeMicrosoftPowerPointMacroEnabledPresentation
	TypeMicrosoftPowerPointMacroEnabledSlideshow
	TypeMicrosoftPowerPointMacroEnabledTemplate
	TypeMicrosoftPowerPointPresentation
	TypeMicrosoftPowerPointSlideshow
	TypeMicrosoftPowerPointTemplate
	TypeMicrosoftProjectDocument
	TypeMicrosoftPublisherDocument
	TypeMicrosoftVisioDrawing
	TypeMicrosoftWordDocument
	TypeMicrosoftWordMacroEnabledDocument
	TypeMicrosoftWordMacroEnabledTemplate
	TypeMicrosoftWordTemplate
	TypeMIDIAudio
	TypeMinecraftResourcePack
	TypeMotionJPEG2000Video
	TypeMP3ID3Tagged
	TypeMP4Video
	TypeMPEG12Video
	TypeMPEG4Audio
	TypeMPEGLayer1
	TypeMPEGLayer2
	TypeMPEGLayer3
	TypeMSIXPackage
	TypeNamedPipe
	TypeNanosecondBigEndian
	TypeNanosecondLittleEndian
	TypeNewASCII
	TypeNewASCIIWithCRC
	TypeNginxLog
	TypeNikonRAW
	TypeNimSource
	TypeNixExpression
	TypeNpmPackage
	TypeNuGetPackage
	TypeObjectiveCSource
	TypeOCIImageLayout
	TypeOggSkeleton3
	TypeOggSkeleton4
	TypeOldASCII
	TypeOlympusRAW
	TypeOpenAPISpecification
	TypeOpenDocumentChart
	TypeOpenDocumentDatabase
	TypeOpenDocumentFormula
	TypeOpenDocumentGraphics
	TypeOpenDocumentImage
	TypeOpenDocumentPresentation
	TypeOpenDocumentSpreadsheet
	TypeOpenDocumentText
	TypeOpenRasterImage
	TypeOpusAudio
	TypeOrgMode
	TypeOsuBeatmap
	TypeOsuSkin
	TypePAM
	TypePanasonicRAW
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
	TypePEC
	TypePentaxRAW
	TypePerlScript
	TypePES
	TypePGMASCII
	TypePGMBinary
	TypePHPScript
	TypePowerShellScript
	TypePPMASCII
	TypePPMBinary
	TypeProcreate
	TypePropertiesFile
	TypeProtocolBuffer
	TypePSVita
	TypePWAD
	TypePythonScript
	TypePythonSourceDistribution
	TypePythonWheel
	TypePyTorchModel
	TypeQCPAudio
	TypeQuickTimeMovie
	TypeReactComponent
	TypeReStructuredText
	TypeRF64Audio
	TypeRIFFPalette
	TypeRSA
	TypeRScript
	TypeRubyScript
	TypeRustSource
	TypeScalaSource
	TypeShellcheckDirective
	TypeSkinnableFrame
	TypeSlackwarePackage
	TypeSocket
	TypeSoliditySource
	TypeSonyRAW
	TypeSonyRAWSR2
	TypeSoundFont2
	TypeSpecial
	TypeSpeexAudio
	TypeSpineBinary
	TypeSQLScript
	TypeSSHConfig
	TypeStudioOne
	TypeSvelteComponent
	TypeSwiftSource
	TypeSymbolicLink
	TypeSystemVerilogSource
	TypeTerraformConfiguration
	TypeTerraformModule
	TypeTeXDocument
	TypeTheoraVideo
	TypeThriftInterface
	TypeTI83
	TypeTI83Plus
	TypeTI89
	TypeTI92
	TypeTI92Plus
	TypeTOMLConfiguration
	TypeTopoJSON
	TypeTSVData
	TypeTypeScriptSource
	TypeUncompressed
	TypeUTF16BE
	TypeUTF16LE
	TypeUTF32BE
	TypeUTF32LE
	TypeUTF8
	TypeVagrantBox
	TypeVerilogSource
	TypeVersion0
	TypeVersion035
	TypeVersion036
	TypeVersion037
	TypeVersion038
	TypeVersion039
	TypeVersion040
	TypeVersion041
	TypeVersion1
	TypeVersion2
	TypeVersion3
	TypeVersion4
	TypeVersion5
	TypeVersion6
	TypeVersion7
	TypeVersion8
	TypeVersionB4
	TypeVersionCE
	TypeVHDLSource
	TypeVimScript
	TypeVisualStudioExtension
	TypeVorbisAudio
	TypeVoyage200
	TypeVueComponent
	TypeWAVAudio
	TypeWebMVideo
	TypeWebPImage
	TypeWestwoodVQA
	TypeWindowsAnimatedCursor
	TypeWindowsLE
	TypeWindowsLX
	TypeWindowsMetafile
	TypeWindowsNE
	TypeWorkspace
	TypeWrapper
	TypeXMLPaperSpecification
	TypeXSLTStylesheet
	TypeYAMLConfiguration
	TypeZigSource
	TypeZlibCompressed
	TypeZshScript
	MaxTypes
)

var typeNames = [...]string{
	TypeNone:                               "",
	Type32Bit:                              "32-bit",
	Type32BitBigEndian:                     "32-bit Big-Endian",
	Type32BitLittleEndian:                  "32-bit Little-Endian",
	Type3G2Video:                           "3G2 Video",
	Type3GPPVideo:                          "3GPP Video",
	Type3MFDocument:                        "3MF Document",
	Type64Bit:                              "64-bit",
	Type64BitBigEndian:                     "64-bit Big-Endian",
	Type64BitLittleEndian:                  "64-bit Little-Endian",
	TypeAC3Audio:                           "AC-3 Audio",
	TypeAdobeDNG:                           "Adobe DNG",
	TypeAfterEffectsProject:                "After Effects Project",
	TypeAIFCAudio:                          "AIFC Audio",
	TypeAIFFAudio:                          "AIFF Audio",
	TypeAndroidAppBundle:                   "AAB Package",
	TypeAndroidArchive:                     "AAR Package",
	TypeAndroidPackage:                     "APK Package",
	TypeAndroidPackageX:                    "XAPK Package",
	TypeAndroidSplitAPKSet:                 "APKS Package",
	TypeAndroidSystemPackage:               "Android System Package",
	TypeAndroidVendorBoot:                  "Android Vendor Boot Image",
	TypeAngularComponent:                   "Angular Component",
	TypeApacheLog:                          "Apache Log",
	TypeAPPXPackage:                        "APPX Package",
	TypeArchLinuxPackage:                   "Arch Linux Package",
	TypeASCII:                              "ASCII",
	TypeAsciiDoc:                           "AsciiDoc",
	TypeAtari800:                           "Atari 800",
	TypeAudioDataTransportStream:           "Audio Data Transport Stream",
	TypeAVIFImage:                          "AVIF Image",
	TypeAVIFImageSequence:                  "AVIF Image Sequence",
	TypeAVIVideo:                           "AVI Video",
	TypeBashScript:                         "Bash Script",
	TypeBatchScript:                        "Batch Script",
	TypeBDAV:                               "BDAV",
	TypeBGZF:                               "BGZF",
	TypeBigEndian:                          "Big-Endian",
	TypeBigTIFF:                            "BigTIFF",
	TypeBinary:                             "Binary",
	TypeBinaryBigEndian:                    "Binary Big-Endian",
	TypeBinaryLittleEndian:                 "Binary Little-Endian",
	TypeBitwig:                             "Bitwig Studio Project",
	TypeBlackmagicRAWVideo:                 "Blackmagic RAW Video",
	TypeBlockDevice:                        "Block Device",
	TypeByteSwapped:                        "Byte-Swapped",
	TypeCanonRAW3Image:                     "Canon RAW 3 Image",
	TypeCatalog:                            "Catalog",
	TypeCDAAudio:                           "CD Audio",
	TypeCeltx:                              "Celtx Project",
	TypeCharacterDevice:                    "Character Device",
	TypeClojureScript:                      "Clojure Script",
	TypeCMakeScript:                        "CMake Script",
	TypeCodestream:                         "Codestream",
	TypeCOFFObject:                         "COFF Object",
	TypeComicBook:                          "Comic Book",
	TypeCommitMessage:                      "Commit Message",
	TypeCommodore64:                        "Commodore 64",
	TypeCondaPackage:                       "Conda Package",
	TypeContainer:                          "Container",
	TypeCopyOnWrite:                        "Copy-On-Write",
	TypeCoqSource:                          "Coq Source",
	TypeCorelDRAWDocument:                  "CorelDRAW Document",
	TypeCPPSource:                          "C++ Source",
	TypeCSharpSource:                       "C# Source",
	TypeCSource:                            "C Source",
	TypeCSS:                                "Cascading Style Sheets",
	TypeCSVData:                            "CSV Data",
	TypeCubaseProject:                      "Cubase Project",
	TypeDartSource:                         "Dart Source",
	TypeDescription:                        "Description",
	TypeDiffPatch:                          "Diff/Patch",
	TypeDirectory:                          "Directory",
	TypeDjangoTemplate:                     "Django Template",
	TypeDockerComposeConfiguration:         "Docker Compose Configuration",
	TypeDockerfile:                         "Dockerfile",
	TypeDownloadableSounds:                 "Downloadable Sounds",
	TypeDSA:                                "DSA",
	TypeDSIKModule:                         "DSIK Module",
	TypeEAC3Audio:                          "E-AC-3 Audio",
	TypeEC:                                 "EC",
	TypeElixirScript:                       "Elixir Script",
	TypeEmacsLispScript:                    "Emacs Lisp Script",
	TypeEmpty:                              "Empty",
	TypeEncrypted:                          "Encrypted",
	TypeEnhancedMetafile:                   "Enhanced Metafile",
	TypeEPUBDocument:                       "EPUB Document",
	TypeExt2:                               "ext2",
	TypeExt3:                               "ext3",
	TypeExt4:                               "ext4",
	TypeF4VVideo:                           "F4V Video",
	TypeFabricMod:                          "Fabric Mod",
	TypeFAT12:                              "FAT12",
	TypeFAT16:                              "FAT16",
	TypeFAT32:                              "FAT32",
	TypeFirefoxExtension:                   "Firefox Extension",
	TypeFishScript:                         "Fish Script",
	TypeFLACAudio:                          "FLAC Audio",
	TypeFlatpak:                            "Flatpak Bundle",
	TypeForgeMod:                           "Forge Mod",
	TypeFSharpSource:                       "F# Source",
	TypeGameMakerStudio:                    "GameMaker Studio",
	TypeGameMakerStudio2:                   "GameMaker Studio 2",
	TypeGeoJSON:                            "GeoJSON",
	TypeGeoPackage:                         "GeoPackage",
	TypeGIF87a:                             "GIF87a",
	TypeGIF89a:                             "GIF89a",
	TypeGitConfig:                          "Git Configuration",
	TypeGitDiff:                            "Git Diff",
	TypeGoSource:                           "Go Source",
	TypeGoTemplate:                         "Go Template",
	TypeGraphQL:                            "GraphQL",
	TypeGraphQLSchema:                      "GraphQL Schema",
	TypeGuileScript:                        "Guile Script",
	TypeHalfLife1:                          "Half-Life 1",
	TypeHalfLife2:                          "Half-Life 2",
	TypeHandlebarsTemplate:                 "Handlebars Template",
	TypeHaskellSource:                      "Haskell Source",
	TypeHDPhoto:                            "HD Photo",
	TypeHE:                                 "HE",
	TypeHEIFImage:                          "HEIF Image",
	TypeHighEntropy:                        "High Entropy",
	TypeHTTPLog:                            "HTTP Log",
	TypeIdrisSource:                        "Idris Source",
	TypeILBMImage:                          "ILBM Image",
	TypeIMMMFormat:                         "IMMM Format",
	TypeINIConfiguration:                   "INI Configuration",
	TypeIOSApplicationArchive:              "iOS Application Archive",
	TypeIWAD:                               "IWAD",
	TypeJavaArchive:                        "Java Archive",
	TypeJavaEnterpriseArchive:              "Java Enterprise Archive",
	TypeJavaScriptSource:                   "JavaScript Source",
	TypeJavaSource:                         "Java Source",
	TypeJavaWebArchive:                     "Java Web Archive",
	TypeJinjaTemplate:                      "Jinja Template",
	TypeJSON5Source:                        "JSON5 Source",
	TypeJSONCSource:                        "JSONC Source",
	TypeKDMV:                               "KDMV",
	TypeKerasModel:                         "Keras Model",
	TypeKMZArchive:                         "KMZ Archive",
	TypeKotlinSource:                       "Kotlin Source",
	TypeKritaDocument:                      "Krita Document",
	TypeLargeDocument:                      "Large Document Format",
	TypeLaTeXDocument:                      "LaTeX Document",
	TypeLatin1:                             "Latin-1",
	TypeLeanSource:                         "Lean Source",
	TypeLegacy:                             "Legacy",
	TypeList:                               "List",
	TypeLittleEndian:                       "Little-Endian",
	TypeLogData:                            "Log Data",
	TypeLOVEGame:                           "LÖVE Game",
	TypeLuaScript:                          "Lua Script",
	TypeLZ4Legacy:                          "LZ4 Legacy",
	TypeLZMACompressed:                     "LZMA Compressed",
	TypeM4VVideo:                           "M4V Video",
	TypeMakefile:                           "Makefile",
	TypeMarkdownDocument:                   "Markdown Document",
	TypeMatroskaVideo:                      "Matroska Video",
	TypeMBTiles:                            "MBTiles",
	TypeMicrosoftExcelAddIn:                "Microsoft Excel Add-In",
	TypeMicrosoftExcelMacroEnabledTemplate: "Microsoft Excel Macro-Enabled Template",
	TypeMicrosoftExcelMacroEnabledWorkbook: "Microsoft Excel Macro-Enabled Workbook",
	TypeMicrosoftExcelTemplate:             "Microsoft Excel Template",
	TypeMicrosoftExcelWorkbook:             "Microsoft Excel Workbook",
	TypeMicrosoftInstaller:                 "Microsoft Installer",
	TypeMicrosoftInstallerPatch:            "Microsoft Installer Patch",
	TypeMicrosoftOutlookMessage:            "Microsoft Outlook Message",
	TypeMicrosoftPowerPointAddIn:           "Microsoft PowerPoint Add-In",
	TypeMicrosoftPowerPointMacroEnabledPresentation: "Microsoft PowerPoint Macro-Enabled Presentation",
	TypeMicrosoftPowerPointMacroEnabledSlideshow:    "Microsoft PowerPoint Macro-Enabled Slideshow",
	TypeMicrosoftPowerPointMacroEnabledTemplate:     "Microsoft PowerPoint Macro-Enabled Template",
	TypeMicrosoftPowerPointPresentation:             "Microsoft PowerPoint Presentation",
	TypeMicrosoftPowerPointSlideshow:                "Microsoft PowerPoint Slideshow",
	TypeMicrosoftPowerPointTemplate:                 "Microsoft PowerPoint Template",
	TypeMicrosoftProjectDocument:                    "Microsoft Project Document",
	TypeMicrosoftPublisherDocument:                  "Microsoft Publisher Document",
	TypeMicrosoftVisioDrawing:                       "Microsoft Visio Drawing",
	TypeMicrosoftWordDocument:                       "Microsoft Word Document",
	TypeMicrosoftWordMacroEnabledDocument:           "Microsoft Word Macro-Enabled Document",
	TypeMicrosoftWordMacroEnabledTemplate:           "Microsoft Word Macro-Enabled Template",
	TypeMicrosoftWordTemplate:                       "Microsoft Word Template",
	TypeMIDIAudio:                                   "MIDI Audio",
	TypeMinecraftResourcePack:                       "Minecraft Resource Pack",
	TypeMotionJPEG2000Video:                         "Motion JPEG 2000 Video",
	TypeMP3ID3Tagged:                                "MP3 (ID3 Tagged)",
	TypeMP4Video:                                    "MP4 Video",
	TypeMPEG12Video:                                 "MPEG-1/2 Video",
	TypeMPEG4Audio:                                  "MPEG-4 Audio",
	TypeMPEGLayer1:                                  "MPEG Layer I",
	TypeMPEGLayer2:                                  "MPEG Layer II",
	TypeMPEGLayer3:                                  "MPEG Layer III",
	TypeMSIXPackage:                                 "MSIX Package",
	TypeNamedPipe:                                   "Named Pipe",
	TypeNanosecondBigEndian:                         "Nanosecond Big-Endian",
	TypeNanosecondLittleEndian:                      "Nanosecond Little-Endian",
	TypeNewASCII:                                    "New ASCII",
	TypeNewASCIIWithCRC:                             "New ASCII with CRC",
	TypeNginxLog:                                    "Nginx Log",
	TypeNikonRAW:                                    "Nikon RAW",
	TypeNimSource:                                   "Nim Source",
	TypeNixExpression:                               "Nix Expression",
	TypeNpmPackage:                                  "npm Package",
	TypeNuGetPackage:                                "NuGet Package",
	TypeObjectiveCSource:                            "Objective-C Source",
	TypeOCIImageLayout:                              "OCI Image Layout",
	TypeOggSkeleton3:                                "Ogg Skeleton 3.0",
	TypeOggSkeleton4:                                "Ogg Skeleton 4.0",
	TypeOldASCII:                                    "Old ASCII",
	TypeOlympusRAW:                                  "Olympus RAW",
	TypeOpenAPISpecification:                        "OpenAPI Specification",
	TypeOpenDocumentChart:                           "OpenDocument Chart",
	TypeOpenDocumentDatabase:                        "OpenDocument Database",
	TypeOpenDocumentFormula:                         "OpenDocument Formula",
	TypeOpenDocumentGraphics:                        "OpenDocument Graphics",
	TypeOpenDocumentImage:                           "OpenDocument Image",
	TypeOpenDocumentPresentation:                    "OpenDocument Presentation",
	TypeOpenDocumentSpreadsheet:                     "OpenDocument Spreadsheet",
	TypeOpenDocumentText:                            "OpenDocument Text",
	TypeOpenRasterImage:                             "OpenRaster Image",
	TypeOpusAudio:                                   "Opus Audio",
	TypeOrgMode:                                     "Org Mode",
	TypeOsuBeatmap:                                  "osu! Beatmap",
	TypeOsuSkin:                                     "osu! Skin",
	TypePAM:                                         "PAM",
	TypePanasonicRAW:                                "Panasonic RAW",
	TypePBMASCII:                                    "PBM ASCII",
	TypePBMBinary:                                   "PBM binary",
	TypePE32ARM:                                     "PE32 ARM",
	TypePE32ARM64:                                   "PE32 ARM64",
	TypePE32ARMv7:                                   "PE32 ARMv7",
	TypePE32Itanium:                                 "PE32 Itanium",
	TypePE32PlusARM64:                               "PE32+ ARM64",
	TypePE32PlusUnknown:                             "PE32+ Unknown",
	TypePE32PlusX8664:                               "PE32+ x86-64",
	TypePE32Unknown:                                 "PE32 Unknown",
	TypePE32X86:                                     "PE32 x86",
	TypePE32X8664:                                   "PE32 x86-64",
	TypePEC:                                         "PEC",
	TypePentaxRAW:                                   "Pentax RAW",
	TypePerlScript:                                  "Perl Script",
	TypePES:                                         "PES",
	TypePGMASCII:                                    "PGM ASCII",
	TypePGMBinary:                                   "PGM binary",
	TypePHPScript:                                   "PHP Script",
	TypePowerShellScript:                            "PowerShell Script",
	TypePPMASCII:                                    "PPM ASCII",
	TypePPMBinary:                                   "PPM binary",
	TypeProcreate:                                   "Procreate Artwork",
	TypePropertiesFile:                              "Properties File",
	TypeProtocolBuffer:                              "Protocol Buffer",
	TypePSVita:                                      "PS Vita Package",
	TypePWAD:                                        "PWAD",
	TypePythonScript:                                "Python Script",
	TypePythonSourceDistribution:                    "Python Source Distribution",
	TypePythonWheel:                                 "Python Wheel",
	TypePyTorchModel:                                "PyTorch Model",
	TypeQCPAudio:                                    "QCP Audio",
	TypeQuickTimeMovie:                              "QuickTime Movie",
	TypeReactComponent:                              "React Component",
	TypeReStructuredText:                            "reStructuredText",
	TypeRF64Audio:                                   "RF64 Audio",
	TypeRIFFPalette:                                 "Palette",
	TypeRSA:                                         "RSA",
	TypeRScript:                                     "R Script",
	TypeRubyScript:                                  "Ruby Script",
	TypeRustSource:                                  "Rust Source",
	TypeScalaSource:                                 "Scala Source",
	TypeShellcheckDirective:                         "Shellcheck Directive",
	TypeSkinnableFrame:                              "Skinnable Frame",
	TypeSlackwarePackage:                            "Slackware Package",
	TypeSocket:                                      "Socket",
	TypeSoliditySource:                              "Solidity Source",
	TypeSonyRAW:                                     "Sony RAW",
	TypeSonyRAWSR2:                                  "Sony SR2 RAW",
	TypeSoundFont2:                                  "SoundFont 2",
	TypeSpecial:                                     "Special",
	TypeSpeexAudio:                                  "Speex Audio",
	TypeSpineBinary:                                 "Spine Binary",
	TypeSQLScript:                                   "SQL Script",
	TypeSSHConfig:                                   "SSH Configuration",
	TypeStudioOne:                                   "Studio One Song",
	TypeSvelteComponent:                             "Svelte Component",
	TypeSwiftSource:                                 "Swift Source",
	TypeSymbolicLink:                                "Symbolic Link",
	TypeSystemVerilogSource:                         "SystemVerilog Source",
	TypeTerraformConfiguration:                      "Terraform Configuration",
	TypeTerraformModule:                             "Terraform Module",
	TypeTeXDocument:                                 "TeX Document",
	TypeTheoraVideo:                                 "Theora Video",
	TypeThriftInterface:                             "Thrift Interface",
	TypeTI83:                                        "TI-83",
	TypeTI83Plus:                                    "TI-83 Plus",
	TypeTI89:                                        "TI-89",
	TypeTI92:                                        "TI-92",
	TypeTI92Plus:                                    "TI-92 Plus",
	TypeTOMLConfiguration:                           "TOML Configuration",
	TypeTopoJSON:                                    "TopoJSON",
	TypeTSVData:                                     "TSV Data",
	TypeTypeScriptSource:                            "TypeScript Source",
	TypeUncompressed:                                "Uncompressed",
	TypeUTF16BE:                                     "UTF-16 Big-Endian",
	TypeUTF16LE:                                     "UTF-16 Little-Endian",
	TypeUTF32BE:                                     "UTF-32 Big-Endian",
	TypeUTF32LE:                                     "UTF-32 Little-Endian",
	TypeUTF8:                                        "UTF-8",
	TypeVagrantBox:                                  "Vagrant Box",
	TypeVerilogSource:                               "Verilog Source",
	TypeVersion0:                                    "Version 0",
	TypeVersion035:                                  "Version 035",
	TypeVersion036:                                  "Version 036",
	TypeVersion037:                                  "Version 037",
	TypeVersion038:                                  "Version 038",
	TypeVersion039:                                  "Version 039",
	TypeVersion040:                                  "Version 040",
	TypeVersion041:                                  "Version 041",
	TypeVersion1:                                    "Version 1",
	TypeVersion2:                                    "Version 2",
	TypeVersion3:                                    "Version 3",
	TypeVersion4:                                    "Version 4",
	TypeVersion5:                                    "Version 5",
	TypeVersion6:                                    "Version 6",
	TypeVersion7:                                    "Version 7",
	TypeVersion8:                                    "Version 8",
	TypeVersionB4:                                   "Version B4",
	TypeVersionCE:                                   "Version CE",
	TypeVHDLSource:                                  "VHDL Source",
	TypeVimScript:                                   "Vim Script",
	TypeVisualStudioExtension:                       "Visual Studio Extension",
	TypeVorbisAudio:                                 "Vorbis Audio",
	TypeVoyage200:                                   "Voyage 200",
	TypeVueComponent:                                "Vue Component",
	TypeWAVAudio:                                    "WAV Audio",
	TypeWebMVideo:                                   "WebM Video",
	TypeWebPImage:                                   "WebP Image",
	TypeWestwoodVQA:                                 "Westwood VQA Video",
	TypeWindowsAnimatedCursor:                       "Windows Animated Cursor",
	TypeWindowsLE:                                   "Linear Executable (LE)",
	TypeWindowsLX:                                   "OS/2 Linear Executable (LX)",
	TypeWindowsMetafile:                             "Windows Metafile",
	TypeWindowsNE:                                   "16-bit New Executable (NE)",
	TypeWorkspace:                                   "Workspace",
	TypeWrapper:                                     "Wrapper",
	TypeXMLPaperSpecification:                       "XML Paper Specification",
	TypeXSLTStylesheet:                              "XSLT Stylesheet",
	TypeYAMLConfiguration:                           "YAML Configuration",
	TypeZigSource:                                   "Zig Source",
	TypeZlibCompressed:                              "Zlib Compressed",
	TypeZshScript:                                   "Zsh Script",
}

func (t TypeID) String() string {
	if int(t) >= 0 && int(t) < len(typeNames) {
		return typeNames[t]
	}

	return ""
}
