#include "config_core.h"
#include "config_if.h"

bool ConfigInterface::getConfigFromFile(const std::string& configFile)
{

	ConfigCore configCore;
	if (configCore.openFile(configFile.c_str(), "r"))
	{
        group1Field1= configCore.getStr("group1", "field1");
		group1Field2= configCore.getInt("group1", "field2");
		group1Field3= configCore.getInt("group1", "field3");

        group2Field1= configCore.getStr("group2", "field1");
		group2Field2= configCore.getInt("group2", "field2");
		group2Field3= configCore.getInt("group2", "field3");

        return true;
	}
    return false;
}
