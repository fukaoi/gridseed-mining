package main

type tomlConfig struct {
	Pool pool
}

type pool struct {
	Algo     string
	Option 	 string
	Host1 string
	Host2 string
	UserPassWorker string
	UsbDevice string
	UsbDeviceCount string
}
