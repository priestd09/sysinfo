package sysinfo

import "testing"

func TestGetUptime(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Uptime: %s\n", sys.GetUptime())
}

func TestNumberOfProcs(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Number of procs: %d\n", sys.Procs)
}

func TestGetTotalRAMInKB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total RAM: %d KB\n", sys.GetTotalRAMInKB())
}

func TestGetTotalRAMInMB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total RAM: %d MB\n", sys.GetTotalRAMInMB())
}

func TestGetTotalRAMInGB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total RAM: %.1f GB\n", sys.GetTotalRAMInGB())
}

func TestGetTotalHighRAMInKB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total High RAM: %d KB\n", sys.GetTotalHighRAMInKB())
}

func TestGetTotalHighRAMInMB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total High RAM: %d MB\n", sys.GetTotalHighRAMInMB())
}

func TestGetTotalHighRAMInGB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total High RAM: %.1f GB\n", sys.GetTotalHighRAMInGB())
}

func TestGetFreeRAMInKB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free RAM: %d KB\n", sys.GetFreeRAMInKB())
}

func TestGetFreeRAMInMB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free RAM: %d MB\n", sys.GetFreeRAMInMB())
}

func TestGetFreeRAMInGB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free RAM: %.1f GB\n", sys.GetFreeRAMInGB())
}

func TestGetFreeHighRAMInKB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free High RAM: %d KB\n", sys.GetFreeHighRAMInKB())
}

func TestGetFreeHighRAMInMB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free High RAM: %d MB\n", sys.GetFreeHighRAMInMB())
}

func TestGetFreeHighRAMInGB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free High RAM: %.1f GB\n", sys.GetFreeHighRAMInGB())
}

func TestGetBufferRAMInKB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Buffer RAM: %d KB\n", sys.GetBufferRAMInKB())
}

func TestGetBufferRAMInMB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Buffer RAM: %d MB\n", sys.GetBufferRAMInMB())
}

func TestGetBufferRAMInGB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Buffer RAM: %.1f GB\n", sys.GetBufferRAMInGB())
}

func TestSharedRAMInKB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Shared RAM: %d KB\n", sys.GetSharedRAMInKB())
}

func TestSharedRAMInMB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Shared RAM: %d MB\n", sys.GetSharedRAMInMB())
}

func TestSharedRAMInGB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Shared RAM: %.1f GB\n", sys.GetSharedRAMInGB())
}

func TestTotalSwapInKB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total Swap: %d KB\n", sys.GetTotalSwapInKB())
}

func TestTotalSwapInMB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total Swap: %d MB\n", sys.GetTotalSwapInMB())
}

func TestTotalSwapInGB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Total Swap: %.1f GB\n", sys.GetTotalSwapInGB())
}

func TestFreeSwapInKB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free Swap: %d KB\n", sys.GetFreeSwapInKB())
}

func TestFreeSwapInMB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free Swap: %d MB\n", sys.GetFreeSwapInMB())
}

func TestFreeSwapInGB(t *testing.T) {
	sys, err := GetSysInfo()
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("Free Swap: %.1f GB\n", sys.GetFreeSwapInGB())
}
