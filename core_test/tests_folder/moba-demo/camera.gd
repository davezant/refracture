extends Camera3D

@export var follow_target: Node3D
@export var follow_height := 15.0
@export var follow_distance := 18.0
@export var follow_speed := 10.0
@export var move_speed := 20.0
@export var acceleration := 6.0
@export var deceleration := 8.0
@export var screen_edge_margin := 30
@export var vertical_sensitivity := 0.5

var cam_velocity := Vector3.ZERO
var following := false

func _input(event):
	# Quando aperta o SPACE, a câmera "teleporta" para o jogador
	if event.is_action_pressed("camera_follow") and follow_target:
		var p = follow_target.global_transform.origin
		global_transform.origin = Vector3(p.x, p.y + follow_height, p.z + follow_distance)
	if event is InputEventMouseButton:
		if event.is_pressed():
		# zoom in
			if event.button_index == MOUSE_BUTTON_WHEEL_UP:
				fov -= 1.0
			if event.button_index == MOUSE_BUTTON_WHEEL_DOWN:
				fov += 1.0
			# call the zoom function
func _physics_process(delta):
	# Segura SPACE para seguir o player
	following = Input.is_action_pressed("camera_follow")

	if following and follow_target:
		follow_player(delta)
	else:
		free_camera_edge_move(delta)

func follow_player(delta):
	var p = follow_target.global_transform.origin
	var desired = Vector3(p.x, p.y + follow_height, p.z + follow_distance)
	global_transform.origin = global_transform.origin.lerp(desired, follow_speed * delta)

func free_camera_edge_move(delta):
	var viewport := get_viewport()
	var mouse := viewport.get_mouse_position()
	var screen := viewport.get_visible_rect().size

	var dir := Vector3.ZERO

	# horizontal
	if mouse.x > screen.x - screen_edge_margin:
		dir.x += 1
	elif mouse.x < screen_edge_margin:
		dir.x -= 1

	# vertical (frente/trás)
	if mouse.y < screen_edge_margin:
		dir.z -= vertical_sensitivity
	elif mouse.y > screen.y - screen_edge_margin:
		dir.z += vertical_sensitivity

	# suavização
	if dir != Vector3.ZERO:
		dir.y = 0
		dir = dir.normalized()
		cam_velocity = cam_velocity.lerp(dir * move_speed, acceleration * delta)
	else:
		cam_velocity = cam_velocity.lerp(Vector3.ZERO, deceleration * delta)

	global_transform.origin += cam_velocity * delta
