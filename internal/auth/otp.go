package auth

import (
	"context"
	"time"
)

type OTP struct {
	Key     string
	Created time.Time
}

// This map wil contain the one time password
type RetentionMap map[string]OTP

func NewRetentionMap(ctx context.Context, retentionPeriod time.Duration) RetentionMap {
	rm := make(RetentionMap)
	go rm.Retention(ctx, retentionPeriod)
	return rm
}

func (rm RetentionMap) NewOTP (otp string) error {
	o:= OTP {
		Key: otp,
		Created: time.Now(),
	}

	rm[o.Key] = o 
	return nil
}


func (rm RetentionMap) VerifyOTP (otp string) bool {
	if _, ok := rm[otp]; !ok {
		return false
	}
	delete(rm, otp)
	return true
}


func (rm RetentionMap) Retention (ctx context.Context, retentionPeriod time.Duration) {
	ticker := time.NewTicker(400 * time.Millisecond)
	for {
		select{

		case <- ticker.C :
			for _,op := range rm {
				if op.Created.Add(retentionPeriod).Before(time.Now()){
					delete(rm, op.Key)
				}
			}
		case <- ctx.Done():
			return
		}

	}
}
