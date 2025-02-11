package repository

import (
    entity "javifood-restify/internal/domain/entity"
)

type RestaurantRepository interface {
    New(restaurant *entity.Restaurant)
}
