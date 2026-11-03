package main

import (
	"math"
	"strings"
	"testing"
)

func TestCalculateBMI(t *testing.T) {
	// Arrange: определяем тестовые сценарии
	tests := []struct {
		name             string
		weight           float64
		height           float64
		expectedBMI      float64
		expectedCategory string
		wantErr          bool
		errSubstr        string
	}{
		// --- Ошибки валидации: weight ---
		{
			name:      "invalid weight: zero",
			weight:    0,
			height:    1.75,
			wantErr:   true,
			errSubstr: "weight must be greater than 0",
		},
		{
			name:      "invalid weight: negative",
			weight:    -5.0,
			height:    1.75,
			wantErr:   true,
			errSubstr: "weight must be greater than 0",
		},
		{
			name:      "invalid weight: upper boundary limit (1000)",
			weight:    1000.0,
			height:    1.75,
			wantErr:   true,
			errSubstr: "weight must be less than 1000",
		},
		{
			name:      "invalid weight: exceeds upper limit",
			weight:    1050.0,
			height:    1.75,
			wantErr:   true,
			errSubstr: "weight must be less than 1000",
		},

		// --- Ошибки валидации: height ---
		{
			name:      "invalid height: zero",
			weight:    70.0,
			height:    0,
			wantErr:   true,
			errSubstr: "height must be greater than 0",
		},
		{
			name:      "invalid height: negative",
			weight:    70.0,
			height:    -1.75,
			wantErr:   true,
			errSubstr: "height must be greater than 0",
		},
		{
			name:      "invalid height: upper boundary limit (3.0)",
			weight:    70.0,
			height:    3.0,
			wantErr:   true,
			errSubstr: "height must be less than 3",
		},
		{
			name:      "invalid height: exceeds upper limit",
			weight:    70.0,
			height:    3.2,
			wantErr:   true,
			errSubstr: "height must be less than 3",
		},

		// --- Граничные валидные значения параметров weight и height ---
		{
			name:             "valid lower boundary: weight 0.1",
			weight:           0.1,
			height:           1.0,
			expectedBMI:      0.1,
			expectedCategory: "underweight",
			wantErr:          false,
		},
		{
			name:             "valid upper boundary: weight 999.9",
			weight:           999.9,
			height:           2.0,
			expectedBMI:      250.0,
			expectedCategory: "obese",
			wantErr:          false,
		},
		{
			name:             "valid lower boundary: height 0.5",
			weight:           10.0,
			height:           0.5,
			expectedBMI:      40.0,
			expectedCategory: "obese",
			wantErr:          false,
		},
		{
			name:             "valid upper boundary: height 2.9",
			weight:           84.1,
			height:           2.9,
			expectedBMI:      10.0,
			expectedCategory: "underweight",
			wantErr:          false,
		},

		// --- Граничные переходы категорий BMI (для height=2.0, height^2=4.0) ---
		{
			name:             "category underweight: just below 18.5 (BMI 18.4)",
			weight:           73.6,
			height:           2.0,
			expectedBMI:      18.4,
			expectedCategory: "underweight",
			wantErr:          false,
		},
		{
			name:             "category boundary: exact 18.5 (normal)",
			weight:           74.0,
			height:           2.0,
			expectedBMI:      18.5,
			expectedCategory: "normal",
			wantErr:          false,
		},
		{
			name:             "category normal: just below 25.0 (BMI 24.9)",
			weight:           99.6,
			height:           2.0,
			expectedBMI:      24.9,
			expectedCategory: "normal",
			wantErr:          false,
		},
		{
			name:             "category boundary: exact 25.0 (overweight)",
			weight:           100.0,
			height:           2.0,
			expectedBMI:      25.0,
			expectedCategory: "overweight",
			wantErr:          false,
		},
		{
			name:             "category overweight: just below 30.0 (BMI 29.9)",
			weight:           119.6,
			height:           2.0,
			expectedBMI:      29.9,
			expectedCategory: "overweight",
			wantErr:          false,
		},
		{
			name:             "category boundary: exact 30.0 (obese)",
			weight:           120.0,
			height:           2.0,
			expectedBMI:      30.0,
			expectedCategory: "obese",
			wantErr:          false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Act: вызываем тестируемую функцию
			gotBMI, gotCategory, err := calculateBMI(tt.weight, tt.height)

			// Assert: проверяем результаты
			if tt.wantErr {
				if err == nil {
					t.Fatalf("expected error containing %q, got nil", tt.errSubstr)
				}
				if !strings.Contains(err.Error(), tt.errSubstr) {
					t.Errorf("error %q does not contain expected substring %q", err.Error(), tt.errSubstr)
				}
				return
			}

			if err != nil {
				t.Fatalf("unexpected error: %v", err)
			}

			if math.Abs(gotBMI-tt.expectedBMI) > 0.05 {
				t.Errorf("BMI = %.2f; expected %.2f (diff > 0.05)", gotBMI, tt.expectedBMI)
			}

			if gotCategory != tt.expectedCategory {
				t.Errorf("category = %q; expected %q", gotCategory, tt.expectedCategory)
			}
		})
	}
}