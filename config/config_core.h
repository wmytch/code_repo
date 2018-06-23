#ifndef __CONFIG_CORE_H__
#define __CONFIG_CORE_H__
#include <stdlib.h>
#include <stdio.h>
#include <string.h>
#include <map>
#include <string>

#define CONFIGLEN           256 

typedef std::map<std::string, std::string> FieldMap;
typedef std::map<std::string, FieldMap> GroupMap;

class ConfigCore
{
public:
	ConfigCore();
	virtual ~ConfigCore();
	int  getInt(const char* groupKey, const char* fieldKey);
	char *getStr(const char* groupKey, const char* fieldKey);
	bool openFile(const char* pathName, const char* type);
protected:
	bool getKey(const char* groupKey, const char* fieldKey);
	FILE* fp;
	char  fieldValue[CONFIGLEN];
	GroupMap groupMap;
};

#endif // !__CONFIG_CORE_H__
