package loraserver

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"strings"
	"time"

	"github.com/hashicorp/terraform/helper/logging"
	"github.com/hashicorp/terraform/helper/resource"
	"github.com/hashicorp/terraform/terraform"
	"golang.org/x/oauth2"
)

type Config struct {
	Address  string
	Port     string
	Username string
	Password string
}


