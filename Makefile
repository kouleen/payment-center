thrift=thrift
idl_path = idl/payment-center

gen_service:
	kitex -module github.com/kouleen/payment-center -service payment.service  $(idl_path)/payment.thrift
