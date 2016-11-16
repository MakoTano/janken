package main

import "testing"

func Test_doJanken(t *testing.T) {
	tests := []struct {
		me, com         Hand
		expectWantError bool
		expectResult    Result
	}{
		// valid input
		{rock, rock, false, even},
		{rock, paper, false, lose},
		{rock, scissors, false, win},
		{scissors, rock, false, lose},
		{scissors, paper, false, win},
		{scissors, scissors, false, even},
		{paper, rock, false, win},
		{paper, paper, false, even},
		{paper, scissors, false, lose},

		// invalid input
		{4, rock, true, 0},
		{rock, 4, true, 0},
	}
	for _, tt := range tests {
		result, err := doJanken(tt.me, tt.com)
		if (err != nil) != tt.expectWantError {
			t.Errorf("should have error %+v", tt)
			continue
		}

		// エラーのときはresultの検査は必要ないのでcontinueします
		if err != nil {
			continue
		}
		if result != tt.expectResult {
			t.Errorf("not match result: %v, expect: %v", result, tt.expectResult)
		}
	}
}
