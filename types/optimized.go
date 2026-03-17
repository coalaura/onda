// code generated, don't edit
package types

func detectOptimized(b Buffer) *Metadata {
	if b.Len() == 0 {
		return nil
	}

	if b.Len() > 0 {
		_ = b[0] // BCE hint
		switch b[0] {
		case 0x00:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x00:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x00:
							if b.Len() > 3 {
								_ = b[3] // BCE hint
								switch b[3] {
								case 0x01:
									if b.Len() >= 8 && string(b[:8]) == "\x00\x00\x00\x01Bud1" {
										return &Metadata{Kind: KindAppleDesktopServicesStore}
									}
								case 0x0c:
									if b.Len() > 4 {
										_ = b[4] // BCE hint
										switch b[4] {
										case 0x4a:
											if b.Len() >= 12 && string(b[:12]) == "\x00\x00\x00\fJXL \r\n\x87\n" {
												return &Metadata{Kind: KindJPEGXLImage, Type: TypeContainer}
											}
										case 0x6a:
											if b.Len() >= 12 && string(b[:12]) == "\x00\x00\x00\fjP  \r\n\x87\n" {
												return &Metadata{Kind: KindJPEG2000Image}
											}
										}
									}
								}
							}
						case 0x01:
							if b.Len() > 3 {
								_ = b[3] // BCE hint
								switch b[3] {
								case 0x00:
									if b.Len() >= 4 && string(b[:4]) == "\x00\x00\x01\x00" {
										return &Metadata{Kind: KindICOCURImage, Type: TypeWindowsIcon}
									}
								case 0xb3:
									if b.Len() >= 4 && string(b[:4]) == "\x00\x00\x01\xb3" {
										return &Metadata{Kind: KindMPEGVideo, Type: TypeMPEG12Video}
									}
								case 0xba:
									if b.Len() >= 4 && string(b[:4]) == "\x00\x00\x01\xba" {
										return &Metadata{Kind: KindMPEGProgramStream}
									}
								}
							}
						case 0x02:
							if b.Len() >= 4 && string(b[:4]) == "\x00\x00\x02\x00" {
								return &Metadata{Kind: KindICOCURImage, Type: TypeWindowsCursor}
							}
						case 0x27:
							if b.Len() >= 4 && string(b[:4]) == "\x00\x00'\n" {
								return &Metadata{Kind: KindESRIShapefile}
							}
						}
					}
				case 0x01:
					if b.Len() >= 19 && string(b[:19]) == "\x00\x01\x00\x00Standard Jet DB" {
						return &Metadata{Kind: KindMicrosoftAccessDatabase}
					}
					if b.Len() >= 4 && string(b[:4]) == "\x00\x01\x00\x00" {
						return &Metadata{Kind: KindTrueTypeFont}
					}
				case 0x05:
					if b.Len() > 3 {
						_ = b[3] // BCE hint
						switch b[3] {
						case 0x00:
							if b.Len() >= 4 && string(b[:4]) == "\x00\x05\x16\x00" {
								return &Metadata{Kind: KindAppleSingle}
							}
						case 0x07:
							if b.Len() >= 4 && string(b[:4]) == "\x00\x05\x16\a" {
								return &Metadata{Kind: KindAppleDouble}
							}
						}
					}
				case 0x4d:
					if b.Len() >= 4 && string(b[:4]) == "\x00MRM" {
						return &Metadata{Kind: KindMinoltaRAWImage}
					}
				case 0x50:
					if b.Len() >= 4 && string(b[:4]) == "\x00PBP" {
						return &Metadata{Kind: KindPlayStationPortableExecutable}
					}
				case 0x61:
					if b.Len() >= 4 && string(b[:4]) == "\x00asm" {
						return &Metadata{Kind: KindWebAssemblyModule}
					}
				case 0xba:
					if b.Len() >= 4 && string(b[:4]) == "\x00\xba\xb1\f" {
						return &Metadata{Kind: KindZFSFilesystem}
					}
				}
			}
		case 0x01:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x00:
					if b.HasMask(0, "\x01\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00 EMF", "\xff\xff\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindMetafileImage, Type: TypeEnhancedMetafileEMF}
					}
				case 0xda:
					if b.Len() >= 2 && string(b[:2]) == "\x01\xda" {
						return &Metadata{Kind: KindSGIImage}
					}
				}
			}
		case 0x02:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x00:
					if b.Len() >= 4 && string(b[:4]) == "\x02\x00\f\x00" {
						return &Metadata{Kind: KindAndroidCompiledResources}
					}
				case 0x21:
					if b.Len() >= 4 && string(b[:4]) == "\x02!L\x18" {
						return &Metadata{Kind: KindLZ4Frame, Type: TypeLZ4Legacy}
					}
				}
			}
		case 0x03:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x00:
					if b.Len() >= 4 && string(b[:4]) == "\x03\x00\b\x00" {
						return &Metadata{Kind: KindAndroidBinaryXML}
					}
				case 0xd9:
					if b.Len() > 4 {
						_ = b[4] // BCE hint
						switch b[4] {
						case 0x65:
							if b.Len() >= 8 && string(b[:8]) == "\x03٢\x9ae\xfbK\xb5" {
								return &Metadata{Kind: KindKeePassDatabase, Type: TypeKDB}
							}
						case 0x67:
							if b.Len() >= 8 && string(b[:8]) == "\x03٢\x9ag\xfbK\xb5" {
								return &Metadata{Kind: KindKeePassDatabase, Type: TypeKDBX}
							}
						}
					}
				}
			}
		case 0x04:
			if b.Len() >= 4 && string(b[:4]) == "\x04\"M\x18" {
				return &Metadata{Kind: KindLZ4Frame}
			}
		case 0x06:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x06:
					if b.Len() >= 16 && string(b[:16]) == "\x06\x06\xed\xf5\xd8\x1dF\xe5\xbd1\xef\xe7\xfet\xb7\x1d" {
						return &Metadata{Kind: KindAdobeInDesignDocument}
					}
				case 0x0e:
					if b.Len() >= 14 && string(b[:14]) == "\x06\x0e+4\x02\x05\x01\x01\r\x01\x02\x01\x01\x02" {
						return &Metadata{Kind: KindMaterialExchangeFormat}
					}
				}
			}
		case 0x0a:
			if b.Len() >= 4 && string(b[:4]) == "\n\r\r\n" {
				return &Metadata{Kind: KindPCAPNGCapture}
			}
		case 0x0c:
			if b.Len() >= 4 && string(b[:4]) == "\f\xb1\xba\x00" {
				return &Metadata{Kind: KindZFSFilesystem}
			}
		case 0x0e:
			if b.Len() >= 4 && string(b[:4]) == "\x0e\x03\x13\x01" {
				return &Metadata{Kind: KindHDF4Data}
			}
		case 0x13:
			if b.Len() >= 4 && string(b[:4]) == "\x13\xab\xa1\\" {
				return &Metadata{Kind: KindASTCTexture}
			}
		case 0x1b:
			if b.Len() >= 4 && string(b[:4]) == "\x1bLua" {
				return &Metadata{Kind: KindLuaBytecode}
			}
		case 0x1f:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x8b:
					if b.HasMask(0, "\x1f\x8b\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00BC\x02\x00", "\xff\xff\x00\x00\x00\x00\x00\x00\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindGzipData, Type: TypeBGZF}
					}
					if b.Len() >= 2 && string(b[:2]) == "\x1f\x8b" {
						return &Metadata{Kind: KindGzipArchive}
					}
				case 0x9d:
					if b.Len() >= 2 && string(b[:2]) == "\x1f\x9d" {
						return &Metadata{Kind: KindUnixCompressArchive}
					}
				case 0xff:
					if b.Len() >= 4 && string(b[:4]) == "\x1f\xff\xe8\x00" {
						return &Metadata{Kind: KindDTSAudio}
					}
				}
			}
		case 0x21:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x3c:
					if b.Len() >= 8 && string(b[:8]) == "!<arch>\n" {
						return &Metadata{Kind: KindARArchive}
					}
				case 0x42:
					if b.Len() >= 4 && string(b[:4]) == "!BDN" {
						return &Metadata{Kind: KindMicrosoftOutlookEmailFolder}
					}
				}
			}
		case 0x23:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x21:
					if b.Len() > 2 && b[2] == 0x41 {
						if b.Len() > 5 {
							_ = b[5] // BCE hint
							switch b[5] {
							case 0x0a:
								if b.Len() >= 6 && string(b[:6]) == "#!AMR\n" {
									return &Metadata{Kind: KindAMRAudio}
								}
							case 0x2d:
								if b.Len() >= 9 && string(b[:9]) == "#!AMR-WB\n" {
									return &Metadata{Kind: KindAMRWBAudio}
								}
							}
						}
					}
					if b.Len() >= 2 && string(b[:2]) == "#!" {
						return &Metadata{Kind: KindShebangScript}
					}
				case 0x3f:
					if b.Len() > 3 {
						_ = b[3] // BCE hint
						switch b[3] {
						case 0x41:
							if b.Len() >= 10 && string(b[:10]) == "#?RADIANCE" {
								return &Metadata{Kind: KindRadianceHDRImage}
							}
						case 0x47:
							if b.Len() >= 6 && string(b[:6]) == "#?RGBE" {
								return &Metadata{Kind: KindRadianceHDRImage}
							}
						}
					}
				case 0x45:
					if b.Len() >= 7 && string(b[:7]) == "#EXTM3U" {
						return &Metadata{Kind: KindM3U8Playlist}
					}
				}
			}
		case 0x24:
			if b.Len() >= 4 && string(b[:4]) == "$FL2" {
				return &Metadata{Kind: KindSPSSData}
			}
		case 0x25:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x21:
					if b.Len() >= 11 && string(b[:11]) == "%!PS-Adobe-" {
						return &Metadata{Kind: KindPostScriptDocument}
					}
				case 0x46:
					if b.Len() >= 5 && string(b[:5]) == "%FDF-" {
						return &Metadata{Kind: KindAcrobatFormsDataFormat}
					}
				case 0x50:
					if b.Len() >= 5 && string(b[:5]) == "%PDF-" {
						return &Metadata{Kind: KindPDFDocument}
					}
				}
			}
		case 0x27:
			if b.Len() >= 4 && string(b[:4]) == "'\x05\x19V" {
				return &Metadata{Kind: KindUBootImage}
			}
		case 0x28:
			if b.Len() >= 4 && string(b[:4]) == "(\xb5/\xfd" {
				return &Metadata{Kind: KindZstandardArchive}
			}
		case 0x2d:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x2d:
					if b.Len() > 11 {
						_ = b[11] // BCE hint
						switch b[11] {
						case 0x43:
							if b.Len() >= 27 && string(b[:27]) == "-----BEGIN CERTIFICATE-----" {
								return &Metadata{Kind: KindPEMCertificate}
							}
						case 0x4f:
							if b.Len() >= 35 && string(b[:35]) == "-----BEGIN OPENSSH PRIVATE KEY-----" {
								return &Metadata{Kind: KindPEMPrivateKey}
							}
						case 0x50:
							if b.Len() > 12 {
								_ = b[12] // BCE hint
								switch b[12] {
								case 0x47:
									if b.Len() >= 14 && string(b[:14]) == "-----BEGIN PGP" {
										return &Metadata{Kind: KindPGPMessage}
									}
								case 0x52:
									if b.Len() >= 27 && string(b[:27]) == "-----BEGIN PRIVATE KEY-----" {
										return &Metadata{Kind: KindPEMPrivateKey}
									}
								}
							}
						case 0x52:
							if b.Len() >= 31 && string(b[:31]) == "-----BEGIN RSA PRIVATE KEY-----" {
								return &Metadata{Kind: KindPEMPrivateKey, Type: TypeRSAPrivateKey}
							}
						}
					}
				case 0x72:
					if b.Len() >= 8 && string(b[:8]) == "-rom1fs-" {
						return &Metadata{Kind: KindROMFS}
					}
				}
			}
		case 0x2e:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x52:
					if b.Len() >= 4 && string(b[:4]) == ".RMF" {
						return &Metadata{Kind: KindRealMedia}
					}
				case 0x73:
					if b.Len() >= 4 && string(b[:4]) == ".snd" {
						return &Metadata{Kind: KindAUAudio}
					}
				}
			}
		case 0x2f:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x2a:
					if b.Len() >= 9 && string(b[:9]) == "/* XPM */" {
						return &Metadata{Kind: KindXPMImage}
					}
				case 0x2f:
					if b.Len() >= 6 && string(b[:6]) == "//Maya" {
						return &Metadata{Kind: KindMayaASCII}
					}
				}
			}
		case 0x30:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x26:
					if b.Len() >= 16 && string(b[:16]) == "0&\xb2u\x8ef\xcf\x11\xa6\xd9\x00\xaa\x00b\xcel" {
						return &Metadata{Kind: KindASFContainer}
					}
				case 0x37:
					if b.Len() > 5 {
						_ = b[5] // BCE hint
						switch b[5] {
						case 0x31:
							if b.Len() >= 6 && string(b[:6]) == "070701" {
								return &Metadata{Kind: KindCPIOArchive, Type: TypeNewASCII}
							}
						case 0x32:
							if b.Len() >= 6 && string(b[:6]) == "070702" {
								return &Metadata{Kind: KindCPIOArchive, Type: TypeNewASCIIWithCRC}
							}
						case 0x37:
							if b.Len() >= 6 && string(b[:6]) == "070707" {
								return &Metadata{Kind: KindCPIOArchive, Type: TypeOldASCII}
							}
						}
					}
				}
			}
		case 0x31:
			if b.Len() >= 4 && string(b[:4]) == "1\x18\x10\x06" {
				return &Metadata{Kind: KindUBIFSFilesystem}
			}
		case 0x34:
			if b.Len() >= 4 && string(b[:4]) == "4\x12\xaaU" {
				return &Metadata{Kind: KindValvePak}
			}
		case 0x37:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x6b:
					if b.Len() >= 4 && string(b[:4]) == "7kSt" {
						return &Metadata{Kind: KindZPAQArchive}
					}
				case 0x7a:
					if b.Len() >= 6 && string(b[:6]) == "7z\xbc\xaf'\x1c" {
						return &Metadata{Kind: Kind7ZipArchive}
					}
				case 0x7f:
					if b.Len() >= 4 && string(b[:4]) == "7\x7f\x06Q" {
						return &Metadata{Kind: KindSQLite3WriteAheadLog, Type: TypeLittleEndian}
					}
				case 0x80:
					if b.Len() >= 4 && string(b[:4]) == "7\x80@\x12" {
						return &Metadata{Kind: KindNintendo64ROM, Type: TypeByteSwapped}
					}
				case 0xa4:
					if b.Len() >= 4 && string(b[:4]) == "7\xa40\xec" {
						return &Metadata{Kind: KindZstandardDictionary}
					}
				}
			}
		case 0x38:
			if b.Len() > 3 {
				_ = b[3] // BCE hint
				switch b[3] {
				case 0x42:
					if b.Len() >= 4 && string(b[:4]) == "8BPB" {
						return &Metadata{Kind: KindPhotoshopDocument, Type: TypePSB}
					}
				case 0x53:
					if b.Len() >= 4 && string(b[:4]) == "8BPS" {
						return &Metadata{Kind: KindPhotoshopDocument, Type: TypePSD}
					}
				}
			}
		case 0x3a:
			if b.Len() >= 4 && string(b[:4]) == ":\xff&\xed" {
				return &Metadata{Kind: KindAndroidSparseImage}
			}
		case 0x3c:
			if b.HasMask(0, "<HTML ", "\xff\xdf\xdf\xdf\xdf\xff") {
				return &Metadata{Kind: KindHTMLDocument}
			}
			if b.HasMask(0, "<HTML>", "\xff\xdf\xdf\xdf\xdf\xff") {
				return &Metadata{Kind: KindHTMLDocument}
			}
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x21:
					if b.HasMask(0, "<!DOCTYPE HTML", "\xff\xff\xdf\xdf\xdf\xdf\xdf\xdf\xdf\xff\xdf\xdf\xdf\xdf") {
						return &Metadata{Kind: KindHTMLDocument}
					}
				case 0x3f:
					if b.Len() >= 5 && string(b[:5]) == "<?xml" {
						return &Metadata{Kind: KindXMLDocument}
					}
				case 0x73:
					if b.Len() >= 11 && string(b[:11]) == "<stata_dta>" {
						return &Metadata{Kind: KindStataData}
					}
				}
			}
		case 0x40:
			if b.Len() >= 4 && string(b[:4]) == "@\x127\x80" {
				return &Metadata{Kind: KindNintendo64ROM, Type: TypeLittleEndian}
			}
		case 0x41:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x43:
					if b.Len() >= 4 && string(b[:4]) == "AC10" {
						return &Metadata{Kind: KindAutoCADDrawing}
					}
				case 0x46:
					if b.Len() >= 4 && string(b[:4]) == "AFF\x00" {
						return &Metadata{Kind: KindAdvancedForensicFormat}
					}
				case 0x4c:
					if b.Len() >= 4 && string(b[:4]) == "ALZ\x01" {
						return &Metadata{Kind: KindALZArchive}
					}
				case 0x4e:
					if b.Len() >= 8 && string(b[:8]) == "ANDROID!" {
						return &Metadata{Kind: KindAndroidBootImage}
					}
				case 0x52:
					if b.Len() >= 6 && string(b[:6]) == "ARROW1" {
						return &Metadata{Kind: KindApacheArrowFile}
					}
				case 0x54:
					if b.HasMask(0, "AT&TFORM\x00\x00\x00\x00DJVI", "\xff\xff\xff\xff\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindDjVuDocument}
					}
					if b.HasMask(0, "AT&TFORM\x00\x00\x00\x00DJVM", "\xff\xff\xff\xff\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindDjVuDocument}
					}
					if b.HasMask(0, "AT&TFORM\x00\x00\x00\x00DJVU", "\xff\xff\xff\xff\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindDjVuDocument}
					}
					if b.HasMask(0, "AT&TFORM\x00\x00\x00\x00THUM", "\xff\xff\xff\xff\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindDjVuDocument}
					}
				}
			}
		case 0x42:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x41:
					if b.Len() >= 4 && string(b[:4]) == "BAM\x01" {
						return &Metadata{Kind: KindBAMData}
					}
				case 0x43:
					if b.Len() >= 4 && string(b[:4]) == "BC\xc0\xde" {
						return &Metadata{Kind: KindLLVMBitcode}
					}
				case 0x45:
					if b.Len() > 9 {
						_ = b[9] // BCE hint
						switch b[9] {
						case 0x4c:
							if b.Len() >= 15 && string(b[:15]) == "BEGIN:VCALENDAR" {
								return &Metadata{Kind: KindICalendar}
							}
						case 0x52:
							if b.Len() >= 11 && string(b[:11]) == "BEGIN:VCARD" {
								return &Metadata{Kind: KindVCard}
							}
						}
					}
				case 0x49:
					if b.Len() > 3 {
						_ = b[3] // BCE hint
						switch b[3] {
						case 0x62:
							if b.Len() >= 4 && string(b[:4]) == "BIKb" {
								return &Metadata{Kind: KindBinkVideo}
							}
						case 0x66:
							if b.Len() >= 4 && string(b[:4]) == "BIKf" {
								return &Metadata{Kind: KindBinkVideo}
							}
						case 0x67:
							if b.Len() >= 4 && string(b[:4]) == "BIKg" {
								return &Metadata{Kind: KindBinkVideo}
							}
						case 0x68:
							if b.Len() >= 4 && string(b[:4]) == "BIKh" {
								return &Metadata{Kind: KindBinkVideo}
							}
						case 0x69:
							if b.Len() >= 4 && string(b[:4]) == "BIKi" {
								return &Metadata{Kind: KindBinkVideo}
							}
						}
					}
				case 0x4c:
					if b.Len() >= 7 && string(b[:7]) == "BLENDER" {
						return &Metadata{Kind: KindBlenderFile}
					}
				case 0x4d:
					if b.Len() >= 2 && string(b[:2]) == "BM" {
						return &Metadata{Kind: KindBMPImage}
					}
				case 0x50:
					if b.Len() >= 4 && string(b[:4]) == "BPG\xfb" {
						return &Metadata{Kind: KindBPGImage}
					}
				case 0x5a:
					if b.Len() >= 3 && string(b[:3]) == "BZh" {
						return &Metadata{Kind: KindBzip2Archive}
					}
				}
			}
		case 0x43:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x44:
					if b.Len() > 3 {
						_ = b[3] // BCE hint
						switch b[3] {
						case 0x01:
							if b.Len() >= 4 && string(b[:4]) == "CDF\x01" {
								return &Metadata{Kind: KindNetCDFData}
							}
						case 0x02:
							if b.Len() >= 4 && string(b[:4]) == "CDF\x02" {
								return &Metadata{Kind: KindNetCDFData}
							}
						}
					}
				case 0x49:
					if b.Len() >= 4 && string(b[:4]) == "CISO" {
						return &Metadata{Kind: KindPlayStationPortableISO}
					}
				case 0x52:
					if b.Len() >= 4 && string(b[:4]) == "CRAM" {
						return &Metadata{Kind: KindCRAMData}
					}
				case 0x57:
					if b.Len() >= 3 && string(b[:3]) == "CWS" {
						return &Metadata{Kind: KindShockwaveFlash, Type: TypeZlibCompressed}
					}
				case 0x72:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x32:
							if b.Len() > 4 {
								_ = b[4] // BCE hint
								switch b[4] {
								case 0x02:
									if b.Len() >= 8 && string(b[:8]) == "Cr24\x02\x00\x00\x00" {
										return &Metadata{Kind: KindCRXBrowserExtension, Type: TypeCRXVersion2}
									}
								case 0x03:
									if b.Len() >= 8 && string(b[:8]) == "Cr24\x03\x00\x00\x00" {
										return &Metadata{Kind: KindCRXBrowserExtension, Type: TypeCRXVersion3}
									}
								}
							}
							if b.Len() >= 4 && string(b[:4]) == "Cr24" {
								return &Metadata{Kind: KindCRXBrowserExtension}
							}
						case 0x65:
							if b.Len() >= 20 && string(b[:20]) == "Creative Voice File\x1a" {
								return &Metadata{Kind: KindVOCAudio}
							}
						}
					}
				}
			}
		case 0x44:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x41:
					if b.Len() >= 4 && string(b[:4]) == "DAA\x1a" {
						return &Metadata{Kind: KindPowerISODAA}
					}
				case 0x44:
					if b.Len() >= 4 && string(b[:4]) == "DDS " {
						return &Metadata{Kind: KindDDSImage}
					}
				case 0x49:
					if b.Len() >= 4 && string(b[:4]) == "DIRC" {
						return &Metadata{Kind: KindGitIndex}
					}
				case 0x4b:
					if b.Len() >= 4 && string(b[:4]) == "DKIF" {
						return &Metadata{Kind: KindIVFVideo}
					}
				case 0x53:
					if b.Len() >= 4 && string(b[:4]) == "DSD " {
						return &Metadata{Kind: KindDSFAudio}
					}
				}
			}
		case 0x45:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x3d:
					if b.Len() >= 4 && string(b[:4]) == "E=\xcd(" {
						return &Metadata{Kind: KindCramFS}
					}
				case 0x56:
					if b.Len() >= 8 && string(b[:8]) == "EVF\t\r\n\xff\x00" {
						return &Metadata{Kind: KindEnCaseImage}
					}
				case 0x6c:
					if b.Len() >= 8 && string(b[:8]) == "ElfFile\x00" {
						return &Metadata{Kind: KindWindowsEventLog}
					}
				case 0x78:
					if b.Len() >= 17 && string(b[:17]) == "Extended Module: " {
						return &Metadata{Kind: KindFastTrackerModule}
					}
				}
			}
		case 0x46:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x4c:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x49:
							if b.Len() >= 4 && string(b[:4]) == "FLIF" {
								return &Metadata{Kind: KindFLIFImage}
							}
						case 0x56:
							if b.Len() >= 3 && string(b[:3]) == "FLV" {
								return &Metadata{Kind: KindFLVVideo}
							}
						}
					}
				case 0x4f:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x52:
							if b.Len() > 3 {
								_ = b[3] // BCE hint
								switch b[3] {
								case 0x34:
									if b.Len() >= 4 && string(b[:4]) == "FOR4" {
										return &Metadata{Kind: KindMayaBinary}
									}
								case 0x4d:
									if b.HasMask(0, "FORM\x00\x00\x00\x00FRM8", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
										return &Metadata{Kind: KindDSDIFFAudio}
									}
									if b.HasMask(0, "FORM\x00\x00\x00\x00AIFC", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
										return &Metadata{Kind: KindIFFContainer, Type: TypeAIFCAudio}
									}
									if b.HasMask(0, "FORM\x00\x00\x00\x00AIFF", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
										return &Metadata{Kind: KindIFFContainer, Type: TypeAIFFAudio}
									}
									if b.HasMask(0, "FORM\x00\x00\x00\x00DJVI", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
										return &Metadata{Kind: KindDjVuDocument}
									}
									if b.HasMask(0, "FORM\x00\x00\x00\x00DJVM", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
										return &Metadata{Kind: KindDjVuDocument}
									}
									if b.HasMask(0, "FORM\x00\x00\x00\x00DJVU", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
										return &Metadata{Kind: KindDjVuDocument}
									}
									if b.HasMask(0, "FORM\x00\x00\x00\x00THUM", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
										return &Metadata{Kind: KindDjVuDocument}
									}
									if b.Len() >= 4 && string(b[:4]) == "FORM" {
										return &Metadata{Kind: KindIFFContainer}
									}
								}
							}
						case 0x56:
							if b.Len() >= 4 && string(b[:4]) == "FOVb" {
								return &Metadata{Kind: KindSigmaRAWImage}
							}
						}
					}
				case 0x55:
					if b.Len() >= 16 && string(b[:16]) == "FUJIFILMCCD-RAW " {
						return &Metadata{Kind: KindFujifilmRAWImage}
					}
				case 0x57:
					if b.Len() >= 3 && string(b[:3]) == "FWS" {
						return &Metadata{Kind: KindShockwaveFlash, Type: TypeUncompressed}
					}
				}
			}
		case 0x47:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x49:
					if b.Len() > 4 {
						_ = b[4] // BCE hint
						switch b[4] {
						case 0x37:
							if b.Len() >= 6 && string(b[:6]) == "GIF87a" {
								return &Metadata{Kind: KindGIFImage, Type: TypeGIF87a}
							}
						case 0x39:
							if b.Len() >= 6 && string(b[:6]) == "GIF89a" {
								return &Metadata{Kind: KindGIFImage, Type: TypeGIF89a}
							}
						}
					}
				case 0x52:
					if b.Len() >= 4 && string(b[:4]) == "GRIB" {
						return &Metadata{Kind: KindGRIBData}
					}
				}
			}
		case 0x49:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x44:
					if b.Len() >= 3 && string(b[:3]) == "ID3" {
						return &Metadata{Kind: KindMPEGAudio, Type: TypeMP3ID3Tagged}
					}
				case 0x49:
					if b.Len() >= 4 && string(b[:4]) == "II\xbc\x01" {
						return &Metadata{Kind: KindJPEGXRImage, Type: TypeLittleEndian}
					}
				case 0x4d:
					if b.Len() >= 4 && string(b[:4]) == "IMPM" {
						return &Metadata{Kind: KindImpulseTrackerModule}
					}
				case 0x54:
					if b.Len() >= 4 && string(b[:4]) == "ITSF" {
						return &Metadata{Kind: KindCHMDocument}
					}
				case 0x57:
					if b.Len() >= 4 && string(b[:4]) == "IWAD" {
						return &Metadata{Kind: KindWADArchive, Type: TypeIWAD}
					}
				}
			}
		case 0x4a:
			if b.Len() >= 4 && string(b[:4]) == "JMOD" {
				return &Metadata{Kind: KindJavaModule, Type: TypeJMOD}
			}
		case 0x4b:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x44:
					if b.Len() >= 4 && string(b[:4]) == "KDMV" {
						return &Metadata{Kind: KindVMwareDiskImage, Type: TypeVMDK}
					}
				case 0x61:
					if b.Len() >= 21 && string(b[:21]) == "Kaydara FBX Binary  \x00" {
						return &Metadata{Kind: KindFBXModel}
					}
				}
			}
		case 0x4c:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x00:
					if b.Len() >= 20 && string(b[:20]) == "L\x00\x00\x00\x01\x14\x02\x00\x00\x00\x00\x00\xc0\x00\x00\x00\x00\x00\x00F" {
						return &Metadata{Kind: KindWindowsShortcut}
					}
				case 0x55:
					if b.Len() >= 6 && string(b[:6]) == "LUKS\xba\xbe" {
						return &Metadata{Kind: KindLUKSDiskEncryption}
					}
				case 0x5a:
					if b.Len() >= 4 && string(b[:4]) == "LZIP" {
						return &Metadata{Kind: KindLZIPArchive}
					}
				}
			}
		case 0x4d:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x3c:
					if b.Len() >= 4 && string(b[:4]) == "M<\xb2\xa1" {
						return &Metadata{Kind: KindPCAPCapture, Type: TypeNanosecondLittleEndian}
					}
				case 0x41:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x43:
							if b.Len() >= 4 && string(b[:4]) == "MAC " {
								return &Metadata{Kind: KindMonkeysAudio}
							}
						case 0x54:
							if b.Len() >= 19 && string(b[:19]) == "MATLAB 5.0 MAT-file" {
								return &Metadata{Kind: KindMATLABData}
							}
						}
					}
				case 0x43:
					if b.Len() >= 8 && string(b[:8]) == "MComprHD" {
						return &Metadata{Kind: KindMAMECHD}
					}
				case 0x45:
					if b.Len() >= 16 && string(b[:16]) == "MEDIA DESCRIPTOR" {
						return &Metadata{Kind: KindMediaDescriptor}
					}
				case 0x4d:
					if b.Len() >= 4 && string(b[:4]) == "MM\x01\xbc" {
						return &Metadata{Kind: KindJPEGXRImage, Type: TypeBigEndian}
					}
				case 0x50:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x2b:
							if b.Len() >= 3 && string(b[:3]) == "MP+" {
								return &Metadata{Kind: KindMusepackAudio, Type: TypeStreamVersion7}
							}
						case 0x43:
							if b.Len() >= 4 && string(b[:4]) == "MPCK" {
								return &Metadata{Kind: KindMusepackAudio, Type: TypeStreamVersion8}
							}
						}
					}
				case 0x53:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x43:
							if b.Len() >= 4 && string(b[:4]) == "MSCF" {
								return &Metadata{Kind: KindCabinetArchive}
							}
						case 0x57:
							if b.Len() >= 8 && string(b[:8]) == "MSWIM\x00\x00\x00" {
								return &Metadata{Kind: KindWindowsImagingFormat}
							}
						}
					}
				case 0x54:
					if b.Len() >= 4 && string(b[:4]) == "MThd" {
						return &Metadata{Kind: KindMIDISequence}
					}
				}
			}
		case 0x4e:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x45:
					if b.Len() >= 4 && string(b[:4]) == "NES\x1a" {
						return &Metadata{Kind: KindNESROM}
					}
				case 0x53:
					if b.Len() >= 4 && string(b[:4]) == "NSO0" {
						return &Metadata{Kind: KindNintendoSwitchNSO}
					}
				}
			}
		case 0x4f:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x46:
					if b.Len() >= 4 && string(b[:4]) == "OFR " {
						return &Metadata{Kind: KindOptimFROGAudio}
					}
				case 0x52:
					if b.Len() >= 3 && string(b[:3]) == "ORC" {
						return &Metadata{Kind: KindORCColumnarData}
					}
				case 0x54:
					if b.Len() >= 4 && string(b[:4]) == "OTTO" {
						return &Metadata{Kind: KindOpenTypeFont}
					}
				case 0x62:
					if b.Len() >= 4 && string(b[:4]) == "Obj\x01" {
						return &Metadata{Kind: KindAvroObjectContainer}
					}
				}
			}
		case 0x50:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x2a:
					if b.Len() >= 4 && string(b[:4]) == "P*M\x18" {
						return &Metadata{Kind: KindZstandardArchive, Type: TypeZstandardSkinnableFrame}
					}
				case 0x41:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x43:
							if b.Len() >= 4 && string(b[:4]) == "PACK" {
								return &Metadata{Kind: KindGitPack}
							}
						case 0x52:
							if b.Len() >= 4 && string(b[:4]) == "PAR1" {
								return &Metadata{Kind: KindApacheParquet}
							}
						}
					}
				case 0x46:
					if b.Len() >= 4 && string(b[:4]) == "PFS0" {
						return &Metadata{Kind: KindNintendoSwitchPackage}
					}
				case 0x47:
					if b.Len() >= 5 && string(b[:5]) == "PGDMP" {
						return &Metadata{Kind: KindPostgreSQLCustomDump}
					}
				case 0x4b:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x05:
							if b.Len() >= 4 && string(b[:4]) == "PK\x05\x06" {
								return &Metadata{Kind: KindZIPArchive, Type: TypeEmpty}
							}
						case 0x07:
							if b.Len() >= 4 && string(b[:4]) == "PK\a\b" {
								return &Metadata{Kind: KindZIPArchive, Type: TypeSpanned}
							}
						}
					}
				case 0x56:
					if b.Len() >= 4 && string(b[:4]) == "PVR\x03" {
						return &Metadata{Kind: KindPVRTexture}
					}
				case 0x57:
					if b.Len() >= 4 && string(b[:4]) == "PWAD" {
						return &Metadata{Kind: KindWADArchive, Type: TypePWAD}
					}
				case 0x75:
					if b.Len() >= 20 && string(b[:20]) == "PuTTY-User-Key-File-" {
						return &Metadata{Kind: KindPuttyPrivateKey}
					}
				}
			}
		case 0x51:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x06:
					if b.Len() >= 4 && string(b[:4]) == "Q\x06\x7f7" {
						return &Metadata{Kind: KindSQLite3WriteAheadLog, Type: TypeBigEndian}
					}
				case 0x46:
					if b.Len() >= 4 && string(b[:4]) == "QFI\xfb" {
						return &Metadata{Kind: KindQCOWDiskImage, Type: TypeQCOW2}
					}
				}
			}
		case 0x52:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x44:
					if b.Len() > 3 {
						_ = b[3] // BCE hint
						switch b[3] {
						case 0x32:
							if b.Len() >= 5 && string(b[:5]) == "RDX2\n" {
								return &Metadata{Kind: KindRData}
							}
						case 0x33:
							if b.Len() >= 5 && string(b[:5]) == "RDX3\n" {
								return &Metadata{Kind: KindRData}
							}
						}
					}
				case 0x45:
					if b.Len() >= 5 && string(b[:5]) == "REDIS" {
						return &Metadata{Kind: KindRedisDatabase}
					}
				case 0x49:
					if b.HasMask(0, "RIFF\x00\x00\x00\x00CDDA", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindRIFFContainer, Type: TypeCDAAudio}
					}
					if b.HasMask(0, "RIFF\x00\x00\x00\x00DLS ", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindRIFFContainer, Type: TypeDownloadableSounds}
					}
					if b.HasMask(0, "RIFF\x00\x00\x00\x00QLCM", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindRIFFContainer, Type: TypeQCPAudio}
					}
					if b.HasMask(0, "RIFF\x00\x00\x00\x00sfbk", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindRIFFContainer, Type: TypeSoundFont2}
					}
					if b.HasMask(0, "RIFF\x00\x00\x00\x00WAVE", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindRIFFContainer, Type: TypeWAVAudio}
					}
					if b.HasMask(0, "RIFF\x00\x00\x00\x00CDR\x00", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\x00") {
						return &Metadata{Kind: KindRIFFContainer, Type: TypeCorelDRAWDocumentCDR}
					}
					if b.HasMask(0, "RIFF\x00\x00\x00\x00WEBP", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindRIFFContainer, Type: TypeWebPImage}
					}
					if b.HasMask(0, "RIFF\x00\x00\x00\x00AVI ", "\xff\xff\xff\xff\x00\x00\x00\x00\xff\xff\xff\xff") {
						return &Metadata{Kind: KindRIFFContainer, Type: TypeAVIVideo}
					}
					if b.Len() >= 4 && string(b[:4]) == "RIFF" {
						return &Metadata{Kind: KindRIFFContainer}
					}
				case 0x4b:
					if b.Len() >= 4 && string(b[:4]) == "RKA " {
						return &Metadata{Kind: KindRKAudio}
					}
				case 0x56:
					if b.Len() >= 4 && string(b[:4]) == "RVZ\x01" {
						return &Metadata{Kind: KindDolphinRVZ}
					}
				case 0x61:
					if b.Len() > 6 {
						_ = b[6] // BCE hint
						switch b[6] {
						case 0x00:
							if b.Len() >= 7 && string(b[:7]) == "Rar!\x1a\a\x00" {
								return &Metadata{Kind: KindRARArchive, Type: TypeRAR4}
							}
						case 0x01:
							if b.Len() >= 8 && string(b[:8]) == "Rar!\x1a\a\x01\x00" {
								return &Metadata{Kind: KindRARArchive, Type: TypeRAR5}
							}
						}
					}
				case 0x75:
					if b.Len() >= 8 && string(b[:8]) == "RubyGems" {
						return &Metadata{Kind: KindRubyGemPackage}
					}
				}
			}
		case 0x53:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x43:
					if b.Len() >= 4 && string(b[:4]) == "SCE\x00" {
						return &Metadata{Kind: KindPlayStationExecutable}
					}
				case 0x44:
					if b.Len() >= 4 && string(b[:4]) == "SDPX" {
						return &Metadata{Kind: KindDPXImage, Type: TypeBigEndian}
					}
				case 0x49:
					if b.Len() > 2 {
						_ = b[2] // BCE hint
						switch b[2] {
						case 0x4d:
							if b.Len() >= 9 && string(b[:9]) == "SIMPLE  =" {
								return &Metadata{Kind: KindFITSAstronomicalImage}
							}
						case 0x54:
							if b.Len() > 3 {
								_ = b[3] // BCE hint
								switch b[3] {
								case 0x21:
									if b.Len() >= 4 && string(b[:4]) == "SIT!" {
										return &Metadata{Kind: KindStuffItArchive}
									}
								case 0x44:
									if b.Len() >= 4 && string(b[:4]) == "SITD" {
										return &Metadata{Kind: KindStuffItArchive}
									}
								}
							}
						}
					}
				case 0x4d:
					if b.Len() > 3 {
						_ = b[3] // BCE hint
						switch b[3] {
						case 0x32:
							if b.Len() >= 4 && string(b[:4]) == "SMK2" {
								return &Metadata{Kind: KindSmackerVideo}
							}
						case 0x34:
							if b.Len() >= 4 && string(b[:4]) == "SMK4" {
								return &Metadata{Kind: KindSmackerVideo}
							}
						}
					}
				case 0x51:
					if b.Len() >= 16 && string(b[:16]) == "SQLite format 3\x00" {
						return &Metadata{Kind: KindSQLiteDatabase}
					}
				case 0x6b:
					if b.Len() >= 14 && string(b[:14]) == "SketchUp Model" {
						return &Metadata{Kind: KindSketchUpModel}
					}
				case 0x74:
					if b.Len() >= 16 && string(b[:16]) == "StuffIt (c)1997-" {
						return &Metadata{Kind: KindStuffItArchive}
					}
				}
			}
		case 0x54:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x54:
					if b.Len() >= 4 && string(b[:4]) == "TTA1" {
						return &Metadata{Kind: KindTTAAudio}
					}
				case 0x5a:
					if b.Len() >= 4 && string(b[:4]) == "TZif" {
						return &Metadata{Kind: KindTimezoneData}
					}
				}
			}
		case 0x55:
			if b.Len() >= 4 && string(b[:4]) == "U3D\x00" {
				return &Metadata{Kind: KindU3DModel}
			}
		case 0x56:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x42:
					if b.Len() >= 4 && string(b[:4]) == "VBSP" {
						return &Metadata{Kind: KindSourceEngineBSP}
					}
				case 0x54:
					if b.Len() >= 4 && string(b[:4]) == "VTF\x00" {
						return &Metadata{Kind: KindValveTextureFormat}
					}
				}
			}
		case 0x57:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x41:
					if b.Len() >= 5 && string(b[:5]) == "WARC/" {
						return &Metadata{Kind: KindWARCFile}
					}
				case 0x42:
					if b.Len() >= 4 && string(b[:4]) == "WBFS" {
						return &Metadata{Kind: KindWiiBackupFileSystem}
					}
				case 0x45:
					if b.Len() >= 6 && string(b[:6]) == "WEBVTT" {
						return &Metadata{Kind: KindWebVTT}
					}
				case 0x55:
					if b.Len() >= 4 && string(b[:4]) == "WUA\x01" {
						return &Metadata{Kind: KindWiiUArchive}
					}
				}
			}
		case 0x58:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x42:
					if b.Len() >= 4 && string(b[:4]) == "XBEH" {
						return &Metadata{Kind: KindXboxExecutable}
					}
				case 0x45:
					if b.Len() >= 4 && string(b[:4]) == "XEX2" {
						return &Metadata{Kind: KindXbox360Executable}
					}
				case 0x46:
					if b.Len() >= 4 && string(b[:4]) == "XFSB" {
						return &Metadata{Kind: KindXFSFilesystem}
					}
				case 0x49:
					if b.Len() >= 4 && string(b[:4]) == "XISO" {
						return &Metadata{Kind: KindXboxISO}
					}
				case 0x50:
					if b.Len() >= 4 && string(b[:4]) == "XPDS" {
						return &Metadata{Kind: KindDPXImage, Type: TypeLittleEndian}
					}
				}
			}
		case 0x59:
			if b.Len() >= 4 && string(b[:4]) == "Y\xa6j\x95" {
				return &Metadata{Kind: KindSunRasterImage}
			}
		case 0x5a:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x4f:
					if b.Len() >= 4 && string(b[:4]) == "ZOO " {
						return &Metadata{Kind: KindZOOArchive}
					}
				case 0x57:
					if b.Len() >= 3 && string(b[:3]) == "ZWS" {
						return &Metadata{Kind: KindShockwaveFlash, Type: TypeLZMACompressed}
					}
				case 0x58:
					if b.Len() >= 7 && string(b[:7]) == "ZXTape!" {
						return &Metadata{Kind: KindZXTape}
					}
				}
			}
		case 0x5b:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x43:
					if b.Len() >= 9 && string(b[:9]) == "[CloneCD]" {
						return &Metadata{Kind: KindCloneCDControl}
					}
				case 0x70:
					if b.Len() >= 10 && string(b[:10]) == "[playlist]" {
						return &Metadata{Kind: KindPlaylistFile}
					}
				}
			}
		case 0x60:
			if b.Len() >= 2 && string(b[:2]) == "`\xea" {
				return &Metadata{Kind: KindARJArchive}
			}
		case 0x61:
			if b.Len() >= 4 && string(b[:4]) == "art\n" {
				return &Metadata{Kind: KindAndroidART}
			}
		case 0x62:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x6f:
					if b.Len() >= 16 && string(b[:16]) == "book\x00\x00\x00\x00mark\x00\x00\x00\x00" {
						return &Metadata{Kind: KindMacOSAlias}
					}
				case 0x70:
					if b.Len() >= 8 && string(b[:8]) == "bplist00" {
						return &Metadata{Kind: KindAppleBinaryPropertyList}
					}
				case 0x76:
					if b.Len() >= 4 && string(b[:4]) == "bvx2" {
						return &Metadata{Kind: KindLZFSEData}
					}
				}
			}
		case 0x63:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x61:
					if b.Len() >= 4 && string(b[:4]) == "caff" {
						return &Metadata{Kind: KindCAFAudio}
					}
				case 0x78:
					if b.Len() >= 8 && string(b[:8]) == "cxsparse" {
						return &Metadata{Kind: KindVHDDiskImage}
					}
				}
			}
		case 0x64:
			if b.Len() > 2 {
				_ = b[2] // BCE hint
				switch b[2] {
				case 0x78:
					if b.Len() > 5 {
						_ = b[5] // BCE hint
						switch b[5] {
						case 0x33:
							if b.Len() > 6 {
								_ = b[6] // BCE hint
								switch b[6] {
								case 0x35:
									if b.Len() >= 8 && string(b[:8]) == "dex\n035\x00" {
										return &Metadata{Kind: KindDalvikExecutable, Type: TypeDEX035}
									}
								case 0x36:
									if b.Len() >= 8 && string(b[:8]) == "dex\n036\x00" {
										return &Metadata{Kind: KindDalvikExecutable, Type: TypeDEX036}
									}
								case 0x37:
									if b.Len() >= 8 && string(b[:8]) == "dex\n037\x00" {
										return &Metadata{Kind: KindDalvikExecutable, Type: TypeDEX037}
									}
								case 0x38:
									if b.Len() >= 8 && string(b[:8]) == "dex\n038\x00" {
										return &Metadata{Kind: KindDalvikExecutable, Type: TypeDEX038}
									}
								case 0x39:
									if b.Len() >= 8 && string(b[:8]) == "dex\n039\x00" {
										return &Metadata{Kind: KindDalvikExecutable, Type: TypeDEX039}
									}
								}
							}
						case 0x34:
							if b.Len() > 6 {
								_ = b[6] // BCE hint
								switch b[6] {
								case 0x30:
									if b.Len() >= 8 && string(b[:8]) == "dex\n040\x00" {
										return &Metadata{Kind: KindDalvikExecutable, Type: TypeDEX040}
									}
								case 0x31:
									if b.Len() >= 8 && string(b[:8]) == "dex\n041\x00" {
										return &Metadata{Kind: KindDalvikExecutable, Type: TypeDEX041}
									}
								}
							}
						}
					}
				case 0x79:
					if b.Len() >= 4 && string(b[:4]) == "dey\n" {
						return &Metadata{Kind: KindAndroidODEX}
					}
				}
			}
		case 0x66:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x4c:
					if b.Len() >= 4 && string(b[:4]) == "fLaC" {
						return &Metadata{Kind: KindFLACAudio}
					}
				case 0x61:
					if b.Len() >= 8 && string(b[:8]) == "farbfeld" {
						return &Metadata{Kind: KindFarbfeldImage}
					}
				}
			}
		case 0x67:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x69:
					if b.Len() >= 9 && string(b[:9]) == "gimp xcf " {
						return &Metadata{Kind: KindGIMPXCFImage}
					}
				case 0x6c:
					if b.Len() >= 4 && string(b[:4]) == "glTF" {
						return &Metadata{Kind: KindGLTFBinary}
					}
				}
			}
		case 0x68:
			if b.Len() >= 4 && string(b[:4]) == "hsqs" {
				return &Metadata{Kind: KindSquashFSFilesystem}
			}
		case 0x69:
			if b.Len() >= 4 && string(b[:4]) == "icns" {
				return &Metadata{Kind: KindICNSIcon}
			}
		case 0x6f:
			if b.Len() >= 4 && string(b[:4]) == "oat\n" {
				return &Metadata{Kind: KindAndroidOAT}
			}
		case 0x70:
			if b.Len() >= 4 && string(b[:4]) == "ply\n" {
				return &Metadata{Kind: KindPLYModel}
			}
		case 0x71:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x6f:
					if b.Len() >= 4 && string(b[:4]) == "qoif" {
						return &Metadata{Kind: KindQOIImage}
					}
				case 0xc7:
					if b.Len() >= 2 && string(b[:2]) == "q\xc7" {
						return &Metadata{Kind: KindCPIOArchive, Type: TypeBinaryBigEndian}
					}
				}
			}
		case 0x72:
			if b.Len() >= 4 && string(b[:4]) == "regf" {
				return &Metadata{Kind: KindWindowsRegistryHive}
			}
		case 0x74:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x42:
					if b.Len() >= 4 && string(b[:4]) == "tBaK" {
						return &Metadata{Kind: KindTAKAudio}
					}
				case 0x74:
					if b.Len() >= 4 && string(b[:4]) == "ttcf" {
						return &Metadata{Kind: KindTrueTypeCollection}
					}
				}
			}
		case 0x76:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x2f:
					if b.Len() >= 4 && string(b[:4]) == "v/1\x01" {
						return &Metadata{Kind: KindOpenEXRImage}
					}
				case 0x64:
					if b.Len() >= 4 && string(b[:4]) == "vdex" {
						return &Metadata{Kind: KindAndroidVDEX}
					}
				case 0x68:
					if b.Len() >= 8 && string(b[:8]) == "vhdxfile" {
						return &Metadata{Kind: KindVHDXDiskImage}
					}
				}
			}
		case 0x77:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x4f:
					if b.Len() > 3 {
						_ = b[3] // BCE hint
						switch b[3] {
						case 0x32:
							if b.Len() >= 4 && string(b[:4]) == "wOF2" {
								return &Metadata{Kind: KindWOFF2Font}
							}
						case 0x46:
							if b.Len() >= 4 && string(b[:4]) == "wOFF" {
								return &Metadata{Kind: KindWOFFFont}
							}
						}
					}
				case 0x76:
					if b.Len() >= 4 && string(b[:4]) == "wvpk" {
						return &Metadata{Kind: KindWavPackAudio}
					}
				}
			}
		case 0x78:
			if b.Len() >= 4 && string(b[:4]) == "xar!" {
				return &Metadata{Kind: KindXARArchive}
			}
		case 0x7a:
			if b.Len() >= 4 && string(b[:4]) == "z\x1a \x10" {
				return &Metadata{Kind: KindSymbianInstallationFormat}
			}
		case 0x7b:
			if b.Len() >= 5 && string(b[:5]) == "{\\rtf" {
				return &Metadata{Kind: KindRichTextFormatDocument}
			}
		case 0x7f:
			if b.Len() >= 4 && string(b[:4]) == "\x7f\xfe\x80\x01" {
				return &Metadata{Kind: KindDTSAudio}
			}
		case 0x80:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x2a:
					if b.Len() >= 4 && string(b[:4]) == "\x80*_\xd7" {
						return &Metadata{Kind: KindCineonImage}
					}
				case 0x37:
					if b.Len() >= 4 && string(b[:4]) == "\x807\x12@" {
						return &Metadata{Kind: KindNintendo64ROM, Type: TypeBigEndian}
					}
				}
			}
		case 0x85:
			if b.Len() >= 2 && string(b[:2]) == "\x85\x19" {
				return &Metadata{Kind: KindJFFS2Filesystem}
			}
		case 0x89:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x48:
					if b.Len() >= 8 && string(b[:8]) == "\x89HDF\r\n\x1a\n" {
						return &Metadata{Kind: KindHDF5Data}
					}
				case 0x4c:
					if b.Len() >= 9 && string(b[:9]) == "\x89LZO\x00\r\n\x1a\n" {
						return &Metadata{Kind: KindLZOPArchive}
					}
				case 0x50:
					if b.Len() >= 8 && string(b[:8]) == "\x89PNG\r\n\x1a\n" {
						return &Metadata{Kind: KindPNGImage}
					}
				}
			}
		case 0x8a:
			if b.Len() >= 8 && string(b[:8]) == "\x8aMNG\r\n\x1a\n" {
				return &Metadata{Kind: KindMNGImage}
			}
		case 0x8b:
			if b.Len() >= 8 && string(b[:8]) == "\x8bJNG\r\n\x1a\n" {
				return &Metadata{Kind: KindJNGImage}
			}
		case 0x93:
			if b.Len() >= 6 && string(b[:6]) == "\x93NUMPY" {
				return &Metadata{Kind: KindNumPyArray}
			}
		case 0xa1:
			if b.Len() > 2 {
				_ = b[2] // BCE hint
				switch b[2] {
				case 0x3c:
					if b.Len() >= 4 && string(b[:4]) == "\xa1\xb2<M" {
						return &Metadata{Kind: KindPCAPCapture, Type: TypeNanosecondBigEndian}
					}
				case 0xc3:
					if b.Len() >= 4 && string(b[:4]) == "\xa1\xb2\xc3\xd4" {
						return &Metadata{Kind: KindPCAPCapture, Type: TypeBigEndian}
					}
				}
			}
		case 0xab:
			if b.Len() > 5 {
				_ = b[5] // BCE hint
				switch b[5] {
				case 0x31:
					if b.Len() >= 12 && string(b[:12]) == "\xabKTX 11\xbb\r\n\x1a\n" {
						return &Metadata{Kind: KindKTXTexture, Type: TypeKTX}
					}
				case 0x32:
					if b.Len() >= 12 && string(b[:12]) == "\xabKTX 20\xbb\r\n\x1a\n" {
						return &Metadata{Kind: KindKTXTexture, Type: TypeKTX2}
					}
				}
			}
		case 0xb7:
			if b.Len() >= 16 && string(b[:16]) == "\xb7\xd8\x00 7I\xda\x11\xa6N\x00\a\xe9^\xad\x8d" {
				return &Metadata{Kind: KindWTVVideo}
			}
		case 0xc5:
			if b.Len() >= 4 && string(b[:4]) == "\xc5\xd0\xd3\xc6" {
				return &Metadata{Kind: KindEncapsulatedPostScript}
			}
		case 0xc7:
			if b.Len() >= 2 && string(b[:2]) == "\xc7q" {
				return &Metadata{Kind: KindCPIOArchive, Type: TypeBinaryLittleEndian}
			}
		case 0xcf:
			if b.Len() >= 4 && string(b[:4]) == "ϭ\x12\xfe" {
				return &Metadata{Kind: KindOutlookExpressDatabase}
			}
		case 0xd4:
			if b.Len() >= 4 && string(b[:4]) == "\xd4ò\xa1" {
				return &Metadata{Kind: KindPCAPCapture, Type: TypeLittleEndian}
			}
		case 0xd7:
			if b.Len() >= 4 && string(b[:4]) == "\xd7\xcdƚ" {
				return &Metadata{Kind: KindMetafileImage, Type: TypeWindowsMetafileWMF}
			}
		case 0xd9:
			if b.Len() >= 3 && string(b[:3]) == "\xd9\xd9\xf7" {
				return &Metadata{Kind: KindCBORData}
			}
		case 0xde:
			if b.Len() >= 4 && string(b[:4]) == "\xde\xc0\x17\v" {
				return &Metadata{Kind: KindLLVMBitcode, Type: TypeWrapper}
			}
		case 0xed:
			if b.Len() >= 4 && string(b[:4]) == "\xed\xab\xee\xdb" {
				return &Metadata{Kind: KindRPMPackage}
			}
		case 0xfd:
			if b.Len() >= 6 && string(b[:6]) == "\xfd7zXZ\x00" {
				return &Metadata{Kind: KindXZArchive}
			}
		case 0xfe:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x7f:
					if b.Len() >= 4 && string(b[:4]) == "\xfe\x7f\x01\x80" {
						return &Metadata{Kind: KindDTSAudio}
					}
				case 0xed:
					if b.Len() >= 4 && string(b[:4]) == "\xfe\xed\xfe\xed" {
						return &Metadata{Kind: KindJavaKeyStore}
					}
				}
			}
		case 0xff:
			if b.Len() > 1 {
				_ = b[1] // BCE hint
				switch b[1] {
				case 0x06:
					if b.Len() >= 10 && string(b[:10]) == "\xff\x06\x00\x00sNaPpY" {
						return &Metadata{Kind: KindSnappyFramedData}
					}
				case 0x0a:
					if b.Len() >= 2 && string(b[:2]) == "\xff\n" {
						return &Metadata{Kind: KindJPEGXLImage, Type: TypeCodestream}
					}
				case 0x1f:
					if b.Len() >= 4 && string(b[:4]) == "\xff\x1f\x00\xe8" {
						return &Metadata{Kind: KindDTSAudio}
					}
				case 0x4f:
					if b.Len() >= 4 && string(b[:4]) == "\xffO\xffQ" {
						return &Metadata{Kind: KindJPEG2000Image, Type: TypeCodestream}
					}
				case 0xd8:
					if b.Len() >= 3 && string(b[:3]) == "\xff\xd8\xff" {
						return &Metadata{Kind: KindJPEGImage}
					}
				case 0xf1:
					if b.Len() >= 2 && string(b[:2]) == "\xff\xf1" {
						return &Metadata{Kind: KindAACAudio, Type: TypeADTS}
					}
				case 0xf9:
					if b.Len() >= 2 && string(b[:2]) == "\xff\xf9" {
						return &Metadata{Kind: KindAACAudio, Type: TypeADTS}
					}
				case 0xfe:
					if b.Len() >= 4 && string(b[:4]) == "\xff\xfe\xff\x0e" {
						return &Metadata{Kind: KindSketchUpModel}
					}
				}
			}
		}
	}

	if b.Len() > 2 && b[2] == 0x2d {
		if b.Len() > 4 {
			_ = b[4] // BCE hint
			switch b[4] {
			case 0x68:
				if b.Len() > 5 {
					_ = b[5] // BCE hint
					switch b[5] {
					case 0x30:
						if b.Len() >= 7 && string(b[2:7]) == "-lh0-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x31:
						if b.Len() >= 7 && string(b[2:7]) == "-lh1-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x32:
						if b.Len() >= 7 && string(b[2:7]) == "-lh2-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x33:
						if b.Len() >= 7 && string(b[2:7]) == "-lh3-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x34:
						if b.Len() >= 7 && string(b[2:7]) == "-lh4-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x35:
						if b.Len() >= 7 && string(b[2:7]) == "-lh5-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x36:
						if b.Len() >= 7 && string(b[2:7]) == "-lh6-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x37:
						if b.Len() >= 7 && string(b[2:7]) == "-lh7-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x64:
						if b.Len() >= 7 && string(b[2:7]) == "-lhd-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					}
				}
			case 0x7a:
				if b.Len() > 5 {
					_ = b[5] // BCE hint
					switch b[5] {
					case 0x34:
						if b.Len() >= 7 && string(b[2:7]) == "-lz4-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x35:
						if b.Len() >= 7 && string(b[2:7]) == "-lz5-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					case 0x73:
						if b.Len() >= 7 && string(b[2:7]) == "-lzs-" {
							return &Metadata{Kind: KindLHAArchive}
						}
					}
				}
			}
		}
	}

	if b.Len() > 3 {
		_ = b[3] // BCE hint
		switch b[3] {
		case 0x2d:
			if b.Len() >= 11 && string(b[3:11]) == "-FVE-FS-" {
				return &Metadata{Kind: KindBitLockerDiskEncryption}
			}
		case 0x45:
			if b.Len() >= 11 && string(b[3:11]) == "EXFAT   " {
				return &Metadata{Kind: KindExFATFilesystem}
			}
		case 0x4e:
			if b.Len() >= 11 && string(b[3:11]) == "NTFS    " {
				return &Metadata{Kind: KindNTFSFilesystem}
			}
		}
	}

	if b.Len() >= 12 && string(b[4:12]) == "$\xff\xaeQi\x9a\xa2!" {
		return &Metadata{Kind: KindGameBoyAdvanceROM}
	}

	if b.Len() >= 14 && string(b[7:14]) == "**ACE**" {
		return &Metadata{Kind: KindACEArchive}
	}

	if b.Len() >= 21 && string(b[8:21]) == "debian-binary" {
		return &Metadata{Kind: KindDebianPackage}
	}

	if b.Len() >= 32 && string(b[12:32]) == "\xc2\xea\x81`\xb3\x14\x11Ͻ\x92\b\x00\t\xc71\x8c\x18\x1f\x10\x11" {
		return &Metadata{Kind: KindSASData}
	}

	if b.Len() >= 20 && string(b[16:20]) == "NRO0" {
		return &Metadata{Kind: KindNintendoSwitchNRO}
	}

	if b.Len() >= 28 && string(b[24:28]) == "]\x1c\x9e\xa3" {
		return &Metadata{Kind: KindWiiROM}
	}

	if b.Len() >= 32 && string(b[28:32]) == "\xc23\x9f=" {
		return &Metadata{Kind: KindGameCubeROM}
	}

	if b.Len() >= 36 && string(b[32:36]) == "NXSB" {
		return &Metadata{Kind: KindAPFSFilesystem}
	}

	if b.Len() >= 36 && string(b[34:36]) == "LP" {
		return &Metadata{Kind: KindEOTFont}
	}

	if b.Len() >= 40 && string(b[36:40]) == "acsp" {
		return &Metadata{Kind: KindICCProfile}
	}

	if b.Len() >= 48 && string(b[44:48]) == "SCRM" {
		return &Metadata{Kind: KindScreamTrackerModule}
	}

	if b.Len() >= 68 && string(b[60:68]) == "BOOKMOBI" {
		return &Metadata{Kind: KindMOBIDocument}
	}

	if b.Len() >= 104 && string(b[64:104]) == "<<< Oracle VM VirtualBox Disk Image >>>\n" {
		return &Metadata{Kind: KindVirtualBoxDiskImage}
	}

	if b.Len() >= 106 && string(b[102:106]) == "mBIN" {
		return &Metadata{Kind: KindMacBinary}
	}

	if b.Len() >= 132 && string(b[128:132]) == "DICM" {
		return &Metadata{Kind: KindDICOMMedicalImage}
	}

	if b.Len() >= 200 && string(b[192:200]) == "$\xff\xaeQi\x9a\xa2!" {
		return &Metadata{Kind: KindNintendoDSROM}
	}

	if b.Len() > 256 {
		_ = b[256] // BCE hint
		switch b[256] {
		case 0x48:
			if b.Len() >= 260 && string(b[256:260]) == "HEAD" {
				return &Metadata{Kind: KindNintendoSwitchROM}
			}
		case 0x4e:
			if b.Len() >= 260 && string(b[256:260]) == "NCSD" {
				return &Metadata{Kind: KindNintendo3DSROM}
			}
		case 0x53:
			if b.Len() >= 260 && string(b[256:260]) == "SEGA" {
				return &Metadata{Kind: KindSegaMegaDriveROM}
			}
		}
	}

	if b.Len() >= 268 && string(b[260:268]) == "\xce\xedff\xcc\r\x00\v" {
		return &Metadata{Kind: KindGameBoyROM}
	}

	if b.Len() > 1024 {
		_ = b[1024] // BCE hint
		switch b[1024] {
		case 0x10:
			if b.Len() >= 1028 && string(b[1024:1028]) == "\x10 \xf5\xf2" {
				return &Metadata{Kind: KindF2FSFilesystem}
			}
		case 0x48:
			if b.Len() > 1025 {
				_ = b[1025] // BCE hint
				switch b[1025] {
				case 0x2b:
					if b.Len() >= 1026 && string(b[1024:1026]) == "H+" {
						return &Metadata{Kind: KindHFSPlusFilesystem}
					}
				case 0x58:
					if b.Len() >= 1026 && string(b[1024:1026]) == "HX" {
						return &Metadata{Kind: KindHFSPlusFilesystem}
					}
				}
			}
		}
	}

	if b.Len() >= 1036 && string(b[1032:1036]) == "\x02\t\x01\x12" {
		return &Metadata{Kind: KindNILFS2Filesystem}
	}

	if b.Len() >= 4112 && string(b[4096:4112]) == "ƅs\xf6N\x1aEʂe\xf5\x7fH\xbam\x81" {
		return &Metadata{Kind: KindBcachefsFilesystem}
	}

	if b.Len() > 32769 {
		_ = b[32769] // BCE hint
		switch b[32769] {
		case 0x43:
			if b.Len() >= 32774 && string(b[32769:32774]) == "CD001" {
				return &Metadata{Kind: KindISO9660Image}
			}
		case 0x4e:
			if b.Len() > 32773 {
				_ = b[32773] // BCE hint
				switch b[32773] {
				case 0x32:
					if b.Len() >= 32774 && string(b[32769:32774]) == "NSR02" {
						return &Metadata{Kind: KindUniversalDiskFormat}
					}
				case 0x33:
					if b.Len() >= 32774 && string(b[32769:32774]) == "NSR03" {
						return &Metadata{Kind: KindUniversalDiskFormat}
					}
				}
			}
		}
	}

	if b.Len() >= 65608 && string(b[65600:65608]) == "_BHRfS_M" {
		return &Metadata{Kind: KindBtrfsFilesystem}
	}

	return nil
}
