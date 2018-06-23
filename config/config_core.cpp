#include "config_core.h"

ConfigCore::ConfigCore()
{
	memset(fieldValue, 0, sizeof(fieldValue));
	fp = nullptr;
}


ConfigCore::~ConfigCore()
{
	groupMap.clear();
	if (fp)
	{
		fclose(fp);
	}
}

bool ConfigCore::openFile(const char* pathName, const char* type)
{
	std::string configLine{""}, groupKey{""}, currentGroup{""}, fieldKey{""};
	char tmpLine[CONFIGLEN] = { 0 };
	FieldMap currentFieldMap;
	size_t  indexPos = 0;
	size_t  leftPos = 0;
	size_t  rightPos = 0;

	fp = fopen(pathName, type);

	if (!fp )
	{
		printf("open inifile %s error!\n", pathName);
		return false;
	}

	groupMap.clear();

	while (fgets(tmpLine, CONFIGLEN, fp))
	{
		configLine.assign(tmpLine);
		leftPos = configLine.find("\n");
		if (std::string::npos != leftPos )
		{
			configLine.erase(leftPos, 1);
		}
		leftPos = configLine.find("\r");
		if (std::string::npos != leftPos )
		{
			configLine.erase(leftPos, 1);
		}
		leftPos = configLine.find("[");
		rightPos = configLine.find("]");
		if (leftPos != std::string::npos 
			&& rightPos != std::string::npos )
		{
			configLine.erase(leftPos, 1);
			rightPos--;
			configLine.erase(rightPos, 1);
            if(!currentGroup.empty())
			    groupMap[currentGroup] = currentFieldMap;
			currentFieldMap.clear();
			currentGroup = configLine;
		}
		else
		{
			indexPos = configLine.find("=");
			if (std::string::npos != indexPos )
			{
				std::string key, value;
				key = configLine.substr(0, indexPos);
				value = configLine.substr(indexPos + 1, configLine.length() - indexPos - 1);
				currentFieldMap[key] = value;
			}
		}

	}
    if(!currentGroup.empty())
    {
	    groupMap[currentGroup] = currentFieldMap;
        return true;
    }
	return false;
}

bool ConfigCore::getKey(const char* groupKey, const char* fieldKey)
{

	FieldMap key = groupMap[groupKey];

	strcpy(fieldValue, key[fieldKey].c_str());

	return true;
}

int ConfigCore::getInt(const char* groupKey, const char* fieldKey)
{
	int res = 0;

	memset(fieldValue, 0, sizeof(fieldValue));
	if (getKey(groupKey, fieldKey))
	{
		res = atoi(fieldValue);
	}
	return res;
}

char *ConfigCore::getStr(const char* groupKey, const char* fieldKey)
{
	memset(fieldValue, 0, sizeof(fieldValue));
	getKey(groupKey, fieldKey);
	return fieldValue;

}
