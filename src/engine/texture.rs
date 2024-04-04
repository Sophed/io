use raylib::texture::Texture2D;

use super::{ASSETS_DIRECTORY, PIXEL_SCALE};

pub struct GameTexture {
    pub name: String,
    pub data: Texture2D
}

pub fn load(name: &str, handle: &mut raylib::RaylibHandle, thread: &raylib::RaylibThread) -> GameTexture {
    let path = format!("{0}/{1}.png", ASSETS_DIRECTORY, name);
    
    let mut texture: Texture2D = handle
        .load_texture(&thread, &path)
        .expect(&format!("Failed to load texture: {}", path));

    texture.width = texture.width * PIXEL_SCALE;
    texture.height = texture.height * PIXEL_SCALE;
    GameTexture {
        name: name.to_string(),
        data: texture
    }
}