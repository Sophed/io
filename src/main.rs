use engine::{entity::Entity, hitbox::Hitbox, Vec2};
use raylib::prelude::*;

mod engine;

fn main() {

    let (mut rl, thread) = raylib::init()
        .size(640, 480)
        .title("Hello, World")
        .build();

    let player_texture = engine::texture::load("player", &mut rl, &thread);
    let mut player = engine::entity::create(
        engine::Vec2 { x: 100.0, y: 200.0 },
        player_texture
    );
    
    while !rl.window_should_close() {
        let mut d = rl.begin_drawing(&thread);
        d.clear_background(Color::DARKGRAY);

        let mouse_pos : Vec2 = Vec2 {
            x: d.get_mouse_x() as f32,
            y: d.get_mouse_y() as f32
        };
        player.pos = mouse_pos;
        Entity::update_hitbox(&mut player);

        Entity::draw(&player, &mut d);
        Hitbox::draw(&player.hitbox, &mut d);
        
        d.draw_fps(10, 10);
    }
}