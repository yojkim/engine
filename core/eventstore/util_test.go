/*
 * Copyright 2018 It-chain
 *
 * Licensed under the Apache License, Version 2.0 (the "License");
 * you may not use this file except in compliance with the License.
 * You may obtain a copy of the License at
 *
 * https://www.apache.org/licenses/LICENSE-2.0
 *
 * Unless required by applicable law or agreed to in writing, software
 * distributed under the License is distributed on an "AS IS" BASIS,
 * WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
 * See the License for the specific language governing permissions and
 * limitations under the License.
 */

package eventstore_test

import (
	"errors"
	"fmt"

	"github.com/it-chain/midgard"
)

// aggregate
type User struct {
	Name string
	midgard.AggregateModel
}

func (u *User) On(event midgard.Event) error {

	switch v := event.(type) {

	case *UserCreatedEvent:
		u.ID = v.ID

	case *UserNameUpdatedEvent:
		u.Name = v.Name

	default:
		return errors.New(fmt.Sprintf("unhandled event [%s]", v))
	}

	return nil
}

// Command
type UserCreateCommand struct {
	midgard.CommandModel
}

type UserNameUpdateCommand struct {
	midgard.CommandModel
	Name string
}

// Event
type UserCreatedEvent struct {
	midgard.EventModel
}

type UserNameUpdatedEvent struct {
	midgard.EventModel
	Name string
}
