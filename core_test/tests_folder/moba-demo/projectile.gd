extends Area3D

@export var speed: float = 20.0
@export var damage: float = 20.0
var target_pos: Vector3

func setup(target: Vector3, origin: Vector3):
	global_transform.origin = origin
	target_pos = target

func _physics_process(delta):
	var dir = (target_pos - global_transform.origin).normalized()
	global_translate(dir * speed * delta)
	
	# Se passou do alvo, remove
	if global_transform.origin.distance_to(target_pos) < 0.5:
		queue_free()

func _ready():
	# Adiciona CollisionShape
	var shape = CollisionShape3D.new()
	shape.shape = SphereShape3D.new()
	shape.shape.radius = 0.25
	add_child(shape)
	
	connect("body_entered", Callable(self, "_on_body_entered"))

func _on_body_entered(body):
	if body.has_method("take_damage"):
		body.take_damage(damage)
		queue_free()
