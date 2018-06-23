PROJECT_DIR=.
INCLUDE=
CPP=g++
CC=gcc
CFLAGS=-g
CPPFLAGS=-std=c++11

OBJECTS=main.o config_if.o config_core.o 

TARGET= configTest

$(TARGET): $(OBJECTS)
	g++ $(CFLAGS) $(CPPFLAGS) $(OBJECTS) -o $(TARGET)

all:$(TARGET) 

include $(OBJECTS:.o=.d)

%.o: %.cpp
	$(CPP) -c $(CFLAGS) $(CPPFLAGS) $(INCLUDE) $< -o $@

%.d: %.cpp
	set -e; rm -f $@; \
	$(CPP) -MM $(CFLAGS) $(CPPFLAGS) $(INCLUDE) $< > $@.$$$$; \
	sed 's,\($*\)\.o[ :]*,\1.o $@ : ,g' < $@.$$$$ > $@; \
	rm -f $@.$$$$

.PHONY:clean
clean:
	rm -f $(OBJECTS) *.d
