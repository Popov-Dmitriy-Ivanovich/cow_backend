package load

import (
	"cow_backend/models"
	"strconv"
	"time"

	"github.com/gin-gonic/gin"
)

const EXTERIOR_PICTURE_PATH = "./static/exterior/"

var exteriorPictureUniqueIndex uint64 = 0

func (l *Load) Exterior() func(*gin.Context) {
	return func(c *gin.Context) {
		form, err := c.MultipartForm()
		if err != nil {
			c.JSON(500, err)
			return
		}
		picture, ok := form.File["PicturePath"]
		if !ok || len(picture) == 0 {
			c.JSON(500, "not found field csv")
			return
		}

		now := time.Now()
		fileName := "exterior" + strconv.FormatInt(now.Unix(), 16) + "_" + strconv.FormatUint(exteriorPictureUniqueIndex, 16) + picture[0].Filename
		uploadedName := EXTERIOR_PICTURE_PATH + fileName

		if err := c.SaveUploadedFile(picture[0], uploadedName); err != nil {
			c.JSON(500, err)
			return
		}
		exteriorPictureUniqueIndex++

		cow := models.Cow{}
		db := models.GetDb()
		if err := db.Preload("Exterior").First(&cow, form.Value["CowID"][0]).Error; err != nil {
			c.JSON(422, "No CowID")
		}
		if cow.Exterior != nil {
			c.JSON(422, "Cow already has exterior data added")
		}
		cow.Exterior = &models.Exterior{}
		cow.Exterior.PicturePath = &fileName

		if r, err := strconv.ParseFloat(form.Value["Rating"][0], 64); err == nil {
			cow.Exterior.Rating = r
		} else {
			c.JSON(422, "Exterior not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ChestWidth"][0], 64); err == nil {
			cow.Exterior.ChestWidth = v
		} else {
			c.JSON(422, "ChestWidth not float")
		}

		if v, err := strconv.ParseFloat(form.Value["PelvicWidth"][0], 64); err == nil {
			cow.Exterior.PelvicWidth = v
		} else {
			c.JSON(422, "PelvicWidth not float")
		}

		if v, err := strconv.ParseFloat(form.Value["SacrumHeight"][0], 64); err == nil {
			cow.Exterior.SacrumHeight = v
		} else {
			c.JSON(422, "SacrumHeight not float")
		}

		if v, err := strconv.ParseFloat(form.Value["BodyDepth"][0], 64); err == nil {
			cow.Exterior.BodyDepth = v
		} else {
			c.JSON(422, "BodyDepth not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ExteriorType"][0], 64); err == nil {
			cow.Exterior.ExteriorType = v
		} else {
			c.JSON(422, "ExteriorType not float")
		}

		if v, err := strconv.ParseFloat(form.Value["BoneQHockJointRear"][0], 64); err == nil {
			cow.Exterior.BoneQHockJointRear = v
		} else {
			c.JSON(422, "BoneQHockJointRear not float")
		}

		if v, err := strconv.ParseFloat(form.Value["SacrumAngle"][0], 64); err == nil {
			cow.Exterior.SacrumAngle = v
		} else {
			c.JSON(422, "SacrumAngle not float")
		}

		if v, err := strconv.ParseFloat(form.Value["TopLine"][0], 64); err == nil {
			cow.Exterior.TopLine = v
		} else {
			c.JSON(422, "TopLine not float")
		}

		if v, err := strconv.ParseFloat(form.Value["HoofAngle"][0], 64); err == nil {
			cow.Exterior.HoofAngle = v
		} else {
			c.JSON(422, "HoofAngle not float")
		}

		if v, err := strconv.ParseFloat(form.Value["HindLegPosSide"][0], 64); err == nil {
			cow.Exterior.HindLegPosSide = v
		} else {
			c.JSON(422, "HindLegPosSide not float")
		}

		if v, err := strconv.ParseFloat(form.Value["HindLegPosRead"][0], 64); err == nil {
			cow.Exterior.HindLegPosRead = v
		} else {
			c.JSON(422, "HindLegPosRead not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ForeLegPosFront"][0], 64); err == nil {
			cow.Exterior.ForeLegPosFront = v
		} else {
			c.JSON(422, "ForeLegPosFront not float")
		}

		if v, err := strconv.ParseFloat(form.Value["UdderDepth"][0], 64); err == nil {
			cow.Exterior.UdderDepth = v
		} else {
			c.JSON(422, "UdderDepth not float")
		}

		if v, err := strconv.ParseFloat(form.Value["UdderBalance"][0], 64); err == nil {
			cow.Exterior.UdderBalance = v
		} else {
			c.JSON(422, "UdderBalance not float")
		}

		if v, err := strconv.ParseFloat(form.Value["HeightOfUdderAttach"][0], 64); err == nil {
			cow.Exterior.HeightOfUdderAttach = v
		} else {
			c.JSON(422, "HeightOfUdderAttach not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ForeUdderAttach"][0], 64); err == nil {
			cow.Exterior.ForeUdderAttach = v
		} else {
			c.JSON(422, "ForeUdderAttach not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ForeUdderPlcRear"][0], 64); err == nil {
			cow.Exterior.ForeUdderPlcRear = v
		} else {
			c.JSON(422, "ForeUdderPlcRear not float")
		}

		if v, err := strconv.ParseFloat(form.Value["HindTeatPlc"][0], 64); err == nil {
			cow.Exterior.HindTeatPlc = v
		} else {
			c.JSON(422, "HindTeatPlc not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ForeTeatLendth"][0], 64); err == nil {
			cow.Exterior.ForeTeatLendth = v
		} else {
			c.JSON(422, "ForeTeatLendth not float")
		}

		if v, err := strconv.ParseFloat(form.Value["RearTeatLength"][0], 64); err == nil {
			cow.Exterior.RearTeatLength = v
		} else {
			c.JSON(422, "RearTeatLength not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ForeTeatDiameter"][0], 64); err == nil {
			cow.Exterior.ForeTeatDiameter = v
		} else {
			c.JSON(422, "ForeTeatDiameter not float")
		}

		if v, err := strconv.ParseFloat(form.Value["RearTeatDiameter"][0], 64); err == nil {
			cow.Exterior.RearTeatDiameter = v
		} else {
			c.JSON(422, "RearTeatDiameter not float")
		}

		if v, err := strconv.ParseFloat(form.Value["CenterLigamentDepth"][0], 64); err == nil {
			cow.Exterior.CenterLigamentDepth = v
		} else {
			c.JSON(422, "CenterLigamentDepth not float")
		}

		if v, err := strconv.ParseFloat(form.Value["HarmonyOfMovement"][0], 64); err == nil {
			cow.Exterior.HarmonyOfMovement = v
		} else {
			c.JSON(422, "HarmonyOfMovement not float")
		}

		if v, err := strconv.ParseFloat(form.Value["Conditioning"][0], 64); err == nil {
			cow.Exterior.Conditioning = v
		} else {
			c.JSON(422, "Conditioning not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ProminenceOfMilkVeins"][0], 64); err == nil {
			cow.Exterior.ProminenceOfMilkVeins = v
		} else {
			c.JSON(422, "ProminenceOfMilkVeins not float")
		}

		if v, err := strconv.ParseFloat(form.Value["MilkStrength"][0], 64); err == nil {
			cow.Exterior.MilkStrength = v
		} else {
			c.JSON(422, "MilkStrength not float")
		}

		if v, err := strconv.ParseFloat(form.Value["BodyStructure"][0], 64); err == nil {
			cow.Exterior.BodyStructure = v
		} else {
			c.JSON(422, "BodyStructure not float")
		}

		if v, err := strconv.ParseFloat(form.Value["Limbs"][0], 64); err == nil {
			cow.Exterior.Limbs = v
		} else {
			c.JSON(422, "Limbs not float")
		}

		if v, err := strconv.ParseFloat(form.Value["Udder"][0], 64); err == nil {
			cow.Exterior.Udder = v
		} else {
			c.JSON(422, "Udder not float")
		}

		if v, err := strconv.ParseFloat(form.Value["ForeUdderWidth"][0], 64); err == nil {
			cow.Exterior.ForeUdderWidth = v
		} else {
			c.JSON(422, "ForeUdderWidth not float")
		}

		if v, err := strconv.ParseFloat(form.Value["HindUdderWidth"][0], 64); err == nil {
			cow.Exterior.HindUdderWidth = v
		} else {
			c.JSON(422, "HindUdderWidth not float")
		}

		if v, err := strconv.ParseFloat(form.Value["AcrumLength"][0], 64); err == nil {
			cow.Exterior.AcrumLength = v
		} else {
			c.JSON(422, "AcrumLength not float")
		}

		if err := db.Save(&cow).Error; err != nil {
			c.JSON(500, err.Error())
		}
		c.JSON(200, "OK")
	}
}
