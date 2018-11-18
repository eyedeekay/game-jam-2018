package config

/*

 */

import (
//"log"
//"strings"
)

import (
	"github.com/eyedeekay/sam-forwarder/config"
	"github.com/zieckey/goini"
)

type Conf struct {
	config *goini.INI

	i2ptunconf.Conf
}

// WholeGameINILoad loads variables from an ini file into the Conf data structure.
func (c *Conf) WholeGameINILoad(iniFile string) error {
	var err error
	if iniFile != "none" {
		c.config = goini.New()
		err = c.config.ParseFile(iniFile)
		if err != nil {
			return err
		}

		/*		if v, ok := c.config.GetInt("WholeGame.idleTimout"); ok {
					c.IdleTimeout = v
				} else {
					c.IdleTimeout = 360000
				}*/

	}
	return nil
}

// NewWholeGameConf returns a Conf structure from an ini file, for modification
// before starting the tunnel
func NewWholeGameConf(iniFile string) (*Conf, error) {
	var err error
	var c Conf
	if err = c.I2PINILoad(iniFile); err != nil {
		return nil, err
	}
	if err = c.WholeGameINILoad(iniFile); err != nil {
		return nil, err
	}
	return &c, nil
}

func NewBlankWholeGameConf() *Conf {
	var c Conf
	return &c
}
