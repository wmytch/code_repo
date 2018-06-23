#ifndef __CONFIG_INTERFACE__
#define __CONFIG_INTERFACE__

#include <string>

class ConfigInterface
{
public:
	static ConfigInterface* getInstance()
    {
        static ConfigInterface instance;
	    return &instance;
    }

    bool getConfigFromFile(const std::string& configFile);

	inline const std::string& getGroup1Field1()
    {
	    return group1Field1;
    }
	inline unsigned short getGroup1Field2()
    {
	    return group1Field2;
    }
	inline unsigned int getGroup1Field3()
    {
	    return group1Field3;
    }
	inline const std::string& getGroup2Field1()
    {
	    return group2Field1;
    }
	inline unsigned short getGroup2Field2()
    {
	    return group2Field2;
    }
	inline unsigned int getGroup2Field3()
    {
	    return group2Field3;
    }
private:
	ConfigInterface(){}
	~ConfigInterface(){}
    
    std::string group1Field1;
    unsigned short group1Field2;
    unsigned int group1Field3;

    std::string group2Field1;
    unsigned short group2Field2;
    unsigned int group2Field3;
};
#endif 
