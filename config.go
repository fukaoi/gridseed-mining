package main

type tomlConfig struct {
	Pool pool
}

type pool struct {
	Algo     string
	Option 	 string
	Host string
	UserPassWorker string
	UsbDevice string
	UsbDeviceCount string
}
