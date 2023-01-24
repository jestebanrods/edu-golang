package internal

import (
	"fmt"
	"time"
)

type multimedia interface {
	Name() string
	DurationStr() string
	AddTime(time.Duration)
}

type audioFile struct {
	name     string
	duration time.Duration
}

func (f audioFile) Name() string {
	return f.name
}

func (f audioFile) DurationStr() string {
	return f.duration.String()
}

func (f audioFile) AddTime(t time.Duration) {
	f.duration = f.duration + t
}

type videoFile struct {
	name     string
	duration time.Duration
	ext      string
}

func (f videoFile) Name() string {
	return f.name
}

func (f videoFile) DurationStr() string {
	return f.duration.String()
}

func (f *videoFile) AddTime(t time.Duration) {
	f.duration = f.duration + t
}

func (f videoFile) Extension() string {
	return f.ext
}

func Interfaces() {
	var af = audioFile{
		name:     "music",
		duration: 250 * time.Second,
	}

	var vf = &videoFile{
		name:     "video",
		duration: 15 * time.Minute,
		ext:      "mp4",
	}

	multimediaInfo(vf)

	// it works
	addMarketing(vf)
	multimediaInfo(vf)

	multimediaInfo(af)

	// it doesn't work
	addMarketing(af)
	multimediaInfo(af)
}

func multimediaInfo(m multimedia) {
	fmt.Printf("the %s is %s long\n", m.Name(), m.DurationStr())
}

func addMarketing(m multimedia) {
	m.AddTime(5 * time.Minute)
}
