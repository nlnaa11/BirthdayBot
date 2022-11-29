package storages

const _ = "../../../tests_data/"

// func Test_NewTransferFile_CreateTransferFile(t *testing.T) {
// 	path := dir + "test1.json"

// 	err := NewTransferFile(path, "", "", 2424)
// 	assert.Error(t, err)

// 	err = NewTransferFile(path, "formsLink", "databasePath", 2424)
// 	assert.NoError(t, err)
// }

// func Test_AddedRemovedEntry_CorrectFillTransferFile(t *testing.T) {
// 	path := dir + "test1.json"

// 	tr := New()

// 	added := []users.User{
// 		{
// 			UserID:    111111,
// 			UserName:  "@aaa",
// 			FirstName: "A",
// 			LastName:  "AAA",
// 		},
// 		{
// 			FirstName: "C",
// 			LastName:  "CCC",
// 		},
// 	}

// 	err := tr.AddedEntry(added, path, 2424)
// 	assert.NoError(t, err)

// 	removed := []users.User{
// 		{
// 			UserName:  "@bbb",
// 			FirstName: "B",
// 			LastName:  "BBB",
// 		},
// 	}

// 	err = RemovedEntry(removed, path, 2424)
// 	assert.NoError(t, err)
// }

// func Test_AddedEntry_CorrectFillTransferWithAdded(t *testing.T) {
// 	path := dir + "test1.json"

// 	testUsers := []users.User{
// 		{
// 			UserID:    111111,
// 			UserName:  "@aaa",
// 			FirstName: "A",
// 			LastName:  "AAA",
// 		},
// 	}

// 	// invalid chatID
// 	err := AddedEntry(testUsers, path, 2525)
// 	assert.Error(t, err)

// 	err = AddedEntry(testUsers, path, 2424)
// 	assert.NoError(t, err)

// 	other := []users.User{
// 		{
// 			UserName:  "@bbb",
// 			FirstName: "B",
// 			LastName:  "BBB",
// 		},
// 		{
// 			FirstName: "C",
// 			LastName:  "CCC",
// 		},
// 	}

// 	testUsers = append(testUsers, other...)

// 	err = AddedEntry(testUsers, path, 2424)
// 	assert.NoError(t, err)
// }

// func Test_RemovedEntry_CorrectFillTransferWithRemoved(t *testing.T) {
// 	path := dir + "test1.json"

// 	testUsers := []users.User{
// 		{
// 			UserID:    111111,
// 			UserName:  "@aaa",
// 			FirstName: "A",
// 			LastName:  "AAA",
// 		},
// 	}

// 	// invalid chatID
// 	err := AddedEntry(testUsers, path, 2525)
// 	assert.Error(t, err)

// 	err = RemovedEntry(testUsers, path, 2424)
// 	assert.NoError(t, err)

// 	other := []users.User{
// 		{
// 			UserName:  "@bbb",
// 			FirstName: "B",
// 			LastName:  "BBB",
// 		},
// 		{
// 			FirstName: "C",
// 			LastName:  "CCC",
// 		},
// 	}

// 	testUsers = append(testUsers, other...)

// 	err = RemovedEntry(testUsers, path, 2424)
// 	assert.NoError(t, err)
// }

// func Test_CheckBirthdays_CorrectExtractBirthdays(t *testing.T) {

// }
