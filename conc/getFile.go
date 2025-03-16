package main

import (
	"context"
	"crypto/rand"
	"fmt"
	"log"
	"strings"
	"sync"
	"time"
)

func GetFile(ctx context.Context, name string) ([]byte, error) {
	if name == "" {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	select {
	case <-ctx.Done():
		return nil, ctx.Err()
	case <-ticker.C:

	}

	if strings.HasPrefix(name, "invalid") {
		return nil, fmt.Errorf("invalid name %q", name)
	}

	b := make([]byte, 10)
	n, err := rand.Read(b)
	if err != nil {
		return nil, fmt.Errorf("getting file %q: %w", name, err)
	}

	return b[:n], nil
}

// GetFilesOld пример функции, которую нужно оптимизировать.
// Менять эту функцию не нужно, она нужна чтоб
// сравнить поведение двух функций после оптимизации
func GetFilesOld(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	for _, name := range names {
		result[name], err = GetFile(ctx, name)
		if err != nil {
			return nil, err
		}
	}

	return result, nil
}

// Решение: Вывзвать в горутинах каждый GetFile при этом использовать WaitGroup чтобы дождаться получения всех файлов, иначе в конце не успеем ничего получить.
// Так же для вывода ошибки использовал доп. переменную getFileErr так как нельзя из горутины напрямую возвращать ошибку в родительскую функцию, и проверка ошибки будет после того как дождемся исполнения всех горутин после wg.Wait()
// Таким образом вместо 2 секунд получаем 1 секунду исполнения программы.

// GetFilesNew эту функцию можно менять, за исключением её
// сигнатуры
func GetFilesNew(ctx context.Context, names ...string) (result map[string][]byte, err error) {
	if len(names) == 0 {
		return nil, nil
	}

	result = make(map[string][]byte, len(names))
	var getFileErr error
	var wg sync.WaitGroup

	for _, name := range names {
		wg.Add(1)
		go func() {
			defer wg.Done()
			result[name], err = GetFile(ctx, name)
			if err != nil {
				getFileErr = err
			}
		}()
	}

	wg.Wait()

	if getFileErr != nil {
		return nil, getFileErr
	}

	return result, nil
}

func main() {
	start := time.Now()
	files, err := GetFilesNew(context.TODO(), "1", "2")
	if err != nil {
		log.Fatalln(time.Since(start), err)
	}

	fmt.Println(time.Since(start))
	for name := range files {
		fmt.Println(name)
	}
}
