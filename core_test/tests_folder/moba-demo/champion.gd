extends CharacterBody3D

@export var speed: float = 5.0
@export var cam: Camera3D
@export var projectile_q_scene: PackedScene
@export var projectile_w_scene: PackedScene
@export var projectile_e_scene: PackedScene   # se for usar projétil ou efeito diferente no “E”
@export var projectile_r_scene: PackedScene

# Cooldowns em segundos
var cooldowns = {
	"ability_1": 3.0,
	"ability_2": 5.0,
	"ability_3": 7.0,
	"ability_4": 10.0
}
var timers = {
	"ability_1": 0.0,
	"ability_2": 0.0,
	"ability_3": 0.0,
	"ability_4": 0.0
}

var bars: Dictionary = {}

# Movimento
var target_pos: Vector3 = Vector3.ZERO
var moving: bool = false

# Hold‑cast estados
var holding = {
	"ability_1": false,
	"ability_2": false,
	"ability_3": false,
	"ability_4": false
}
var hold_targets = {
	"ability_1": Vector3.ZERO,
	"ability_2": Vector3.ZERO,
	"ability_3": Vector3.ZERO,
	"ability_4": Vector3.ZERO
}
var hold_lines = {
	"ability_1": null,
	"ability_2": null,
	"ability_3": null,
	"ability_4": null
}

func _ready():
	# Configurar GUI de cooldown
	var canvas = CanvasLayer.new()
	add_child(canvas)
	var gui_root = Control.new()
	gui_root.anchor_left = 0.9
	gui_root.anchor_right = 1.0
	gui_root.anchor_top = 0.1
	gui_root.anchor_bottom = 0.9
	canvas.add_child(gui_root)
	
	var vbox = VBoxContainer.new()
	vbox.size_flags_horizontal = Control.SIZE_EXPAND_FILL
	vbox.size_flags_vertical = Control.SIZE_EXPAND_FILL
	gui_root.add_child(vbox)
	
	for ability in ["ability_1","ability_2","ability_3","ability_4"]:
		var bar = ProgressBar.new()
		bar.min_value = 0.0
		bar.max_value = 1.0
		bar.value = 1.0
		bar.set_custom_minimum_size(Vector2(30,150))
		bar.fill_mode = ProgressBar.FillMode.FILL_TOP_TO_BOTTOM
		
		var style_bg = StyleBoxFlat.new()
		style_bg.bg_color = Color(0.1, 0.1, 0.1)
		style_bg.set_border_color(Color(0,0,0))
		style_bg.set_border_width_all(2)
		bar.add_theme_stylebox_override("background", style_bg)
		
		var style_fill = StyleBoxFlat.new()
		style_fill.bg_color = Color(0.2, 0.8, 1.0)
		bar.add_theme_stylebox_override("fill", style_fill)
		
		vbox.add_child(bar)
		bars[ability] = bar
		
		var spacer = Control.new()
		spacer.set_custom_minimum_size(Vector2(0,10))
		vbox.add_child(spacer)
		
		var line_inst = MeshInstance3D.new()
		var imm = ImmediateMesh.new()
		line_inst.mesh = imm
		add_child(line_inst)
		line_inst.visible = false
		hold_lines[ability] = line_inst

func _input(event):
	if event is InputEventMouseButton and event.pressed and event.button_index == MOUSE_BUTTON_RIGHT:
		var mouse2D = get_viewport().get_mouse_position()
		var origin = cam.project_ray_origin(mouse2D)
		var dir = cam.project_ray_normal(mouse2D)
		var to = origin + dir * 2000.0
		var params = PhysicsRayQueryParameters3D.create(origin, to)
		var hit = get_world_3d().direct_space_state.intersect_ray(params)
		if hit:
			target_pos = hit.position
			moving = true
	
	# Processar pressed e released para cada habilidade
	for ability in ["ability_1","ability_2","ability_3","ability_4"]:
		if event.is_action_pressed(ability) and can_use(ability):
			holding[ability] = true
			hold_lines[ability].visible = true
		
		if event.is_action_released(ability) and holding[ability]:
			# Se for teleporte (ability_3)
			if ability == "ability_3":
				global_transform.origin = hold_targets[ability]
				moving = false
			else:
				var scene = null
				match ability:
					"ability_1":
						scene = projectile_q_scene
					"ability_2":
						scene = projectile_w_scene
					"ability_4":
						scene = projectile_r_scene
				if scene:
					shoot_projectile(scene, global_transform.origin, hold_targets[ability])
			start_cooldown(ability)
			holding[ability] = false
			hold_lines[ability].visible = false

func _physics_process(delta):
	position.y = 1
	if moving:
		var dir3 = target_pos - global_transform.origin
		dir3.y = 0
		if dir3.length() > 0.2:
			dir3 = dir3.normalized()
			velocity.x = dir3.x * speed
			velocity.z = dir3.z * speed
			move_and_slide()
		else:
			moving = false
			velocity = Vector3.ZERO
	

	for ability in timers.keys():
		if timers[ability] > 0.0:
			timers[ability] -= delta
			bars[ability].value = 1.0 - (timers[ability] / cooldowns[ability])
		else:
			bars[ability].value = 1.0

	for ability in holding.keys():
		if holding[ability]:
			var mouse2D = get_viewport().get_mouse_position()
			var origin = cam.project_ray_origin(mouse2D)
			var dir = cam.project_ray_normal(mouse2D)
			var to = origin + dir * 2000.0
			var params = PhysicsRayQueryParameters3D.create(origin, to)
			var hit = get_world_3d().direct_space_state.intersect_ray(params)
			if hit:
				hold_targets[ability] = hit.position
				hit.position.y = 1
				var line_inst = hold_lines[ability]
			# Mover a linha para a origem do mundo
				line_inst.global_transform.origin = Vector3.ZERO
				var imm = line_inst.mesh as ImmediateMesh
				imm.clear_surfaces()
				imm.surface_begin(Mesh.PRIMITIVE_LINES)
				imm.surface_set_color(Color(1,0,0))
				imm.surface_add_vertex(global_transform.origin)
				imm.surface_add_vertex(hit.position)
				imm.surface_end()
				line_inst.visible = true


func shoot_projectile(scene: PackedScene, origin: Vector3, target: Vector3):
	var proj = scene.instantiate()
	proj.global_transform.origin = origin
	if proj.has_method("setup"):
		proj.setup(target, origin)
	get_parent().add_child(proj)


# Cooldown helpers
func start_cooldown(ability: String):
	timers[ability] = cooldowns[ability]

func can_use(ability: String) -> bool:
	return timers[ability] <= 0.0
