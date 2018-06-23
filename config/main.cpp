#include <stdio.h>
#include <stdlib.h>
#include <unistd.h>
#include "config_if.h"
#include <iostream>

int main(int argc, char * argv[])
{
    std::string configFile;
    if(getopt(argc,argv,"c:")=='c')
    {
        printf("Using configuration file: %s\n",optarg);
	    configFile.assign(optarg);
    }else {
        printf("Using default configuration.\n");
	    configFile.assign("./config.ini");
    }
	if(!ConfigInterface::getInstance()->getConfigFromFile(configFile))
    {
        printf("Init configuration failed.\n");
        exit(-1);
    }

    std::cout<<"group1:"<<std::endl;
    std::cout<<"field1:"<<ConfigInterface::getInstance()->getGroup1Field1()<<std::endl;
    std::cout<<"field2:"<<ConfigInterface::getInstance()->getGroup1Field2()<<std::endl;
    std::cout<<"field3:"<<ConfigInterface::getInstance()->getGroup1Field3()<<std::endl;
    std::cout<<"group2:"<<std::endl;
    std::cout<<"field1:"<<ConfigInterface::getInstance()->getGroup2Field1()<<std::endl;
    std::cout<<"field2:"<<ConfigInterface::getInstance()->getGroup2Field2()<<std::endl;
    std::cout<<"field3:"<<ConfigInterface::getInstance()->getGroup2Field3()<<std::endl;
    return 0;
}
