extends CharacterBody3D

@export var max_health: float = 100.0
var health: float

# Tipo correto: ColorRect
var health_bar_fill: ColorRect

func _ready():
	health = max_health

	# Corpo do dummy
	var body_mesh = MeshInstance3D.new()
	body_mesh.mesh = BoxMesh.new()
	body_mesh.scale = Vector3(1,2,1)
	add_child(body_mesh)

	# Colisor
	var collider = CollisionShape3D.new()
	var box_shape = BoxShape3D.new()
	box_shape.extents = Vector3(0.5,1,0.5)
	collider.shape = box_shape
	add_child(collider)

	# Barra de vida 2D
	var canvas = CanvasLayer.new()
	add_child(canvas)

	var control = Control.new()
	control.anchor_left = 0.5
	control.anchor_right = 0.5
	control.anchor_top = 0.0
	control.anchor_bottom = 0.0
	control.position = Vector2(0,-50)  # acima da cabeça
	canvas.add_child(control)

	# Fundo da barra
	var health_bar_bg = ColorRect.new()
	health_bar_bg.color = Color(0.1,0.1,0.1)
	health_bar_bg.size = Vector2(100,10)
	control.add_child(health_bar_bg)

	# Preenchimento da barra
	health_bar_fill = ColorRect.new()
	health_bar_fill.color = Color(1,0,0)
	health_bar_fill.size = Vector2(100,10)
	health_bar_fill.position = Vector2.ZERO
	health_bar_bg.add_child(health_bar_fill)

func take_damage(amount: float):
	health -= amount
	health = clamp(health, 0, max_health)
	print(str(health) + " " + str(amount)) 
	_update_health_bar()
	if health <= 0:
		die()

func die():
	queue_free()

func _update_health_bar():
	var ratio = health / max_health
	health_bar_fill.size.x = 100 * ratio
