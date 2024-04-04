use raylib::{color::Color, drawing::RaylibDraw};

use super::{hitbox::Hitbox, texture::GameTexture, Vec2};

pub struct Entity {
    pub pos: Vec2,
    pub velocity: Vec2,
    pub direction: i32,
    pub hitbox: Hitbox,
    pub texture: GameTexture
}

pub fn create(pos: Vec2, texture: GameTexture) -> Entity {

    let height = texture.data.height as f32;
    let width = texture.data.width as f32;

    let top_left = Vec2 { x: pos.x, y: pos.y };
    let bottom_right = Vec2 {
        x: pos.x + width,
        y: pos.y + height
    };

    Entity {
        pos,
        velocity: Vec2 { x: 0.0, y: 0.0 },
        direction: 1,
        hitbox: Hitbox { top_left, bottom_right },
        texture
    }
}

impl Entity {
    
    pub fn update(entity: &mut Entity) {
        let new_pos = Vec2 {
            x: entity.pos.x + entity.velocity.x,
            y: entity.pos.y + entity.velocity.y
        };
        entity.pos = new_pos;
        Entity::update_hitbox(entity)
    }

    pub fn update_hitbox(entity: &mut Entity) {
        let height = entity.texture.data.height as f32;
        let width = entity.texture.data.width as f32;
        let bottom_right = Vec2 {
            x: entity.pos.x + width,
            y: entity.pos.y + height
        };
        entity.hitbox.top_left = Vec2 { x: entity.pos.x, y: entity.pos.y };
        entity.hitbox.bottom_right = bottom_right;
    }

    pub fn draw_direction(entity: &mut Entity, d: &mut raylib::prelude::RaylibDrawHandle) {
        d.draw_line(
            Hitbox::center(&entity.hitbox).x as i32,
            Hitbox::center(&entity.hitbox).y as i32,
            (Hitbox::center(&entity.hitbox).x + (100.0 * entity.direction as f32)) as i32,
            Hitbox::center(&entity.hitbox).y as i32,
            Color::GREEN
        )
    }

    pub fn draw(entity: &Entity, d: &mut raylib::prelude::RaylibDrawHandle) {

        let x = entity.pos.x as i32;
        let y = entity.pos.y as i32;

        d.draw_texture(
            &entity.texture.data,
            x,
            y,
            Color::WHITE
        );
    }

}