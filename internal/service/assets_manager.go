package service

import (
	"log"
	"log/slog"
	"os"
	"path"
	"slices"
	"strings"

	"github.com/veandco/go-sdl2/img"
	"github.com/veandco/go-sdl2/sdl"
)

const ASSETS_DIR = "assets"
var IMAGE_ASSET_EXTS = []string{
    "png", "jpg", "jpeg", "bmp",
}
var AUDIO_ASSET_EXTS = []string {
    "mp3", "wav", "ogg",
}

var manager AssetsManagerService

func InitAssetsManager() {
    manager = AssetsManagerService{
        Assets: make(map[string]Asset, 1),
    }

    err := img.Init(img.INIT_PNG | img.INIT_JPG)
    if err != nil {
        slog.Error("Error initializing SDL Image")
        panic(err)
    }

    manager.loadAssetsFromDir(ASSETS_DIR)
}

func DeinitAssetsManager() {}

func AssetsManager() *AssetsManagerService {
    return &manager
}

type AssetsManagerService struct {
    Assets map[string]Asset
}

type AssetKey = string
type AssetType = string

const AssetTypeAudio AssetType = "audio"
const AssetTypeSprite AssetType = "sprite"

type Asset struct {
    assetType AssetType
    sprite    *sdl.Texture
    audio     *sdl.AudioSpec
}

func (self *AssetsManagerService) NewAsset(
    sprite   *sdl.Surface,
    audio    *sdl.AudioSpec,
    fileName string,
) (Asset, error) {
    if sprite == nil && audio == nil {
        return Asset{}, EmptyAssetError
    }

    if sprite != nil {
        key := GetAssetKey(
            fileName, AssetTypeSprite,
        )
        log.Println("Saving asset as", key)

        spriteTexture, err := Renderer().CreateTextureFromSurface(
            sprite,
        )
        if err != nil {
            slog.Error("Error converting surface to texture:" + err.Error())
            return Asset{}, err
        }

        asset := Asset{
            sprite: spriteTexture,
            assetType: AssetTypeSprite,
        }

        self.Assets[key] = asset
        log.Println(self.Assets)
        
        return asset, nil
    }

    if audio != nil {
        key := GetAssetKey(
            fileName, AssetTypeAudio,
        )

        asset := Asset{
            audio: audio,
            assetType: AssetTypeAudio,
        }

        self.Assets[key] = asset
        
        return asset, nil
    }

    return Asset{}, nil
}

func (self *AssetsManagerService) GetSprite(
    fileName string,
) (*sdl.Texture, error) {
    key := GetAssetKey(fileName, AssetTypeSprite)

    if asset, exists := self.Assets[key]; exists == true {
        return asset.sprite, nil
    } else {
        return nil, AssetDoesNotExistError
    }
}

func (self *AssetsManagerService) GetAudio(
    fileName string,
) (*sdl.AudioSpec, error) {
    key := GetAssetKey(fileName, AssetTypeAudio)

    if asset, exists := self.Assets[key]; exists == true {
        return asset.audio, nil
    } else {
        return nil, AssetDoesNotExistError
    }
}

func (self *AssetsManagerService) loadAssetsFromDir(dir string) {
    log.Println("Loading assets from", dir)
    assets, err := os.ReadDir(dir)
    if err != nil {
        slog.Error("Can't load assets: " + err.Error())
        return
    }

    for _, asset := range assets {
        log.Println("Loading", asset.Name())
        if asset.IsDir() {
        log.Println(asset.Name(), "is a directory")
            self.loadAssetsFromDir(
                path.Join(dir, asset.Name()),
            )
            return
        }

        pcs := strings.Split(asset.Name(), ".")
        if len(pcs) < 2 {
            slog.Error(
                "Invalid file name, can't distinguish asset type: " +
                asset.Name(),
            )
            continue
        }

        ext := pcs[len(pcs) - 1]

        if slices.Contains(IMAGE_ASSET_EXTS, ext) {
            log.Println("Loading image", asset.Name())
            surface, err := img.Load(
                path.Join(dir, asset.Name()),
            )
            if err != nil {
                slog.Error("Error loading surface from " + asset.Name())
                slog.Error(err.Error())
                continue
            }

            log.Println("Creating asset", asset.Name())
            self.NewAsset(
                surface,
                nil,
                asset.Name(),
            )
        }
    }
}

func GetAssetKey(
    fileName string,
    assetType AssetType,
) string {
    var prefix string

    switch assetType {
    case AssetTypeAudio:
        prefix = "a"
    case AssetTypeSprite:
        prefix = "s"
    }

    filePcs := strings.Split(fileName, ".")

    return prefix + filePcs[0]
}

type EmptyAsset struct {}
func (EmptyAsset) Error() string {
    return "Given asset is empty";
}
var EmptyAssetError = EmptyAsset {}

type AssetDoesNotExist struct {}
func (AssetDoesNotExist) Error() string {
    return "Requested asset does not exist";
}
var AssetDoesNotExistError = AssetDoesNotExist {}
