use raylib::{color::Color, drawing::RaylibDraw};

use super::Vec2;

pub struct Hitbox {
    pub top_left: Vec2,
    pub bottom_right: Vec2,
}

impl Hitbox {
    
    pub fn width(&self) -> f32 {
        self.bottom_right.x - self.top_left.x
    }
    pub fn height(&self) -> f32 {
        self.bottom_right.y - self.top_left.y
    }
    pub fn center(&self) -> Vec2 {
        Vec2 {
            x: self.top_left.x + (self.width() / 2.0),
            y: self.top_left.y + (self.height() / 2.0),
        }
    }
    pub fn draw(&self, d: &mut raylib::prelude::RaylibDrawHandle) {
        d.draw_rectangle_lines(
            self.top_left.x as i32,
            self.top_left.y as i32,
            self.width() as i32,
            self.height() as i32,
            Color::RED
        )
    }
    
}