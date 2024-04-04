pub mod hitbox;
pub mod entity;
pub mod texture;

pub const ASSETS_DIRECTORY: &str = "assets";
pub const PIXEL_SCALE: i32 = 4;

pub struct Vec2 {
    pub x: f32,
    pub y: f32,
}