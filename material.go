package tetra3d

import "github.com/hajimehoshi/ebiten/v2"

type Material struct {
	Name            string
	Image           *ebiten.Image
	Tags            *Tags
	BackfaceCulling bool
}

func NewMaterial(name string) *Material {
	return &Material{
		Name:            name,
		Tags:            NewTags(),
		BackfaceCulling: true,
	}
}

func (material *Material) Clone() *Material {
	newMat := NewMaterial(material.Name)
	newMat.Image = material.Image
	newMat.Tags = material.Tags.Clone()
	newMat.BackfaceCulling = material.BackfaceCulling
	return newMat
}