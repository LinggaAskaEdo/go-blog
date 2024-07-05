package entity

import "net/http"

var (
	ErrMsgBadRequest = Message{
		StatusCode: http.StatusBadRequest,
		EN:         `Invalid Input. Please Validate Your Input.`,
		ID:         `Kesalahan Input. Mohon Cek Kembali Masukkan Anda.`,
	}
	ErrMsgNotFound = Message{
		StatusCode: http.StatusNotFound,
		EN:         `Record Does Not Exist. Please Validate Your Input Or Contact Administrator.`,
		ID:         `Data Tidak Ditemukan. Mohon Cek Kembali Masukkan Anda Atau Hubungi Administrator.`,
	}
	ErrMsgUnauthorized = Message{
		StatusCode: http.StatusUnauthorized,
		EN:         `Unauthorized Access. You are not authorized to access this resource.`,
		ID:         `Akses Ditolak. Anda Belum Diijinkan Untuk Mengakses Aplikasi.`,
	}
	ErrMsgISE = Message{
		StatusCode: http.StatusInternalServerError,
		EN:         `Internal Server Error. Please Call Administrator.`,
		ID:         `Terjadi Kendala Pada Server. Mohon Hubungi Administrator.`,
	}
	ErrMsgConflict = Message{
		StatusCode: http.StatusConflict,
		EN:         `Record has existed and must be unique. Please Validate Your Input Or Contact Administrator.`,
		ID:         `Data sudah ada. Mohon Cek Kembali Masukkan Anda Atau Hubungi Administrator.`,
	}
	ErrMsgForbidden = Message{
		StatusCode: http.StatusForbidden,
		EN:         `Forbidden. You don't have permission to access this resource.`,
		ID:         `Terlarang. Anda tidak memiliki izin untuk mengakses aplikasi.`,
	}
	ErrMsgUnprocessable = Message{
		StatusCode: http.StatusUnprocessableEntity,
		EN:         `Not Able to Process This Entity. Please Validate Your Input and Try Again`,
		ID:         `Entitas Ini Tidak Dapat Diproses. Mohon Cek Kembali Masukkan Anda Dan Coba Kembali`,
	}
	ErrMsgTooManyRequest = Message{
		StatusCode: http.StatusTooManyRequests,
		EN:         `Too Many Request For This Entity. Please Wait And Try Again.`,
		ID:         `Permintaan Terlalu Banyak Untuk Entitas Ini. Mohon Tunggu Dan Coba Kembali.`,
	}
	ErrMsgServiceUnavailable = Message{
		StatusCode: http.StatusServiceUnavailable,
		EN:         `Service is unavailable.`,
		ID:         `Layanan sedang tidak tersedia.`,
	}
)
