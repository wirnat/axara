package conf

var DefaultAppConf = AppConf{
	System: System{
		Env:                 "",
		Address:             "",
		OssType:             "local",
		DbType:              "mysql",
		MigrationSourcePath: "./infrastructure/migration/source",
	},
	Mysql: Mysql{
		Path:              "localhost",
		Port:              "3301",
		Config:            "charset=utf8mb4&parseTime=True&loc=Local",
		DbName:            "msglow_customer_app",
		Username:          "app",
		Password:          "123456",
		MaxIdleConnection: 0,
		MaxOpenConnection: 0,
		LogMode:           "",
		LogZap:            false,
	},
	Pgsql: Pgsql{
		Path:              "localhost",
		Port:              "",
		Config:            "",
		DbName:            "",
		Username:          "",
		Password:          "",
		MaxIdleConnection: 0,
		MaxOpenConnection: 0,
		LogMode:           "",
		LogZap:            false,
	},
	Log: Log{
    	MaxSize:   0,
    	Plugin:    "",
    	Compress:  false,
    	Location:  "",
    	MaxAge:    0,
    	MaxBackup: 0,
    },
}
