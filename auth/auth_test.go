package auth

import "testing"

func TestGetAccessToken(t *testing.T) {
	resp, err := GetAccessToken("wx383399d1bd40c153", "42871018c8ada710632305a8bf68b244", "ticket@@@zady9BxMrWlOrMZ4SQT5TbRu8n_hhGLdGN27BVz0tpYkyvgbh2PtRuPe3komWYIKI0QpdznAzpXVm0L-RPPybw")
	t.Logf("resp(%+v) err(%+v)", resp, err)
}
