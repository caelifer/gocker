// Code generated by github.com/gobuffalo/packr/v2. DO NOT EDIT.

// You can use the "packr2 clean" command to clean up this,
// and any other packr generated files.
package packrd

import (
	"github.com/gobuffalo/packr/v2"
	"github.com/gobuffalo/packr/v2/file/resolver"
)

var _ = func() error {
	const gk = "36a9088a71bc628df2e7a6b3d27a96a6"
	g := packr.New(gk, "")
	hgr, err := resolver.NewHexGzip(map[string]string{
		"4f59bf6e3a17584507bc2e40f6b668e7": "1f8b08000000000000ff64cf414fc23018c6f17b3fc5e361091cf62ebb1af0c270313a4a66807844fada35b2b5a1858bf1bb9b766a4838bd87ff2fcdd3c75636d0f6b81ff4fd2c9bc3bae0a9b65b3e796307640f422c575bd4b22ccb46569b97e5fc8a3556f96476b27dae9e5afcb7f53e74a92ce4fa0d04122236f381c95d02af9fc6d51ca6f88aecafa5b433a1abd8fda676b382b6d01ca04de8ceef74b07d314e2e14bbe2d0ab781354ecc0833f9f18f945cc327c838f9e6f5eca2fc803a820a251a50de34deeea938b5e4da631fc040000ffff87d1ab7b2c010000",
		"56a843f683c28af8d3acb35582da637c": "1f8b08000000000000ff3ccd31aec2301084e17e4f314da4e415cf1700ae0015fd926c8885b1adf56e1121ee8e808876a4ff9bcae38daf825db747e6bbfc9fbe433fa03b10d95a05c76ab1e48666eaa3e1414008e069c25a5cd1cce7198ba8d0938842c099539cd804b6b0a16c39abc09bf32509cd9e47f4a55ac3dfc60fbfac1f20aa453f4f2ae69a91637aebaf000000ffff53d53da3af000000",
		"67c2b59c8741c3fe3f9db52e4e4a72ed": "1f8b08000000000000fff248cdc9c957482bcacf55b051b555c84bcc4dd50bc92cc949cdac4ad5d05450b5e302040000ffff122b7d2a22000000",
		"91ab460f060e204b86863ca1392680ef": "1f8b08000000000000ff74ce314ff33010c6f1d9f7291e45eaabf81d9cbda24cb0360c116b65926b6215c7e17c01a1aadf1d25442c88d5ffc73fdde4db8bef1977bb03461fd93d7d3f9416bb7ba210a7248a924ca19c358c7d41648a3ee830bfb836c52aabb0b683546b3f7f56c26f73102ec8129de7b145c3594ff5a4218df9f4ec5f43e7954bc5ff4d748dc5958c607fc0f6d91df9a3544b64d2a47909ff36e07a23c3b26e97e47e3c4b46dca3489292452c11505568ea877a8fe82f8ca0785fc60efb055e90df0620ee98fe607cd721a7c88849183e6796f5260c2cece8465f010000ffff742fed224d010000",
		"bb890f3a8aa56c497012afb45b9738ad": "1f8b08000000000000ff9c514d4b2b3114dde757dc172824a524f096efd92e147457858aaed3f14e1a3a930c77eed411e97f9764fca0825db83d5f3937a773d5de79848bd912a26bd1dc4d80d2305b0911da2e11831200d207de0d5b53a5d6fab41deada35c97a8cf1459ea76d9b0e21fac63192ed9aa1df798c3f7b7225b287bf671439e33bddedbd45a244bd145a887a8815acf159a58e7b98df761c52ec35a879e9646e3022394eb480e2d2f02a003cfc5bc224c85e2d0440a8b3221339ca3cb8263c3946a5ff17fccf1262688a1b8090078ae0df437bf31878b76157ed15126901703c8df4e6328daa9c5c5e9479886d1ad7ae4598ade402a431f6739debd04cd358c6b6cb3fda4bfddb2200158fb944f9cffcfc558a8c232b3d716683ac64be5a2ecaf119f7e69e5ceceb442d92fa98f304ac78d459fa552286461c85780b0000ffff533e79ff6e020000",
		"be63988dd6aff95c2a1ff130f81faee9": "1f8b08000000000000ff548fc14a03311086cf9ba7080b9544966cbd8af522f524554ac18388a4eb6c1adc24dbc9045b4adf5d922ea2971cbef932ff3fa3eebeb4017e375b70af1da8970b1092cfee19b36e0c485cb0aa268864bda919ab6a636997b6aa0bae35619bfa5e0fa135e0fd31bfd9acff5b9110a8db615bb6f4c716619f2c42cd24637df21ddf40a48f157c0be2d75394da487e6215f2db059f7c550cc958651a0e5846195d3d8f64838fa7b36415aa555822061480985d4c3e8b53b5bc639dbc0714b2ccd4aba59d30c5fcfd9af93a79210b86584a6406310d14c5c57e022f10a27a08ce69ff191b3e977ff9a31d2036fc26dbfde58e09becddfb3b8dc273d881a0eda8d03283a50ddf05eadb4831c7d663f010000ffff4748513ba1010000",
	})
	if err != nil {
		panic(err)
	}
	g.DefaultResolver = hgr

	func() {
		b := packr.New("genny:genny:new", "../new/templates")
		b.SetResolver("-name-/-name-.go.plush", packr.Pointer{ForwardBox: gk, ForwardPath: "bb890f3a8aa56c497012afb45b9738ad"})
		b.SetResolver("-name-/-name-_test.go.plush", packr.Pointer{ForwardBox: gk, ForwardPath: "be63988dd6aff95c2a1ff130f81faee9"})
		b.SetResolver("-name-/options.go.plush", packr.Pointer{ForwardBox: gk, ForwardPath: "56a843f683c28af8d3acb35582da637c"})
		b.SetResolver("-name-/options_test.go.plush", packr.Pointer{ForwardBox: gk, ForwardPath: "91ab460f060e204b86863ca1392680ef"})
		b.SetResolver("-name-/templates/example.txt.plush", packr.Pointer{ForwardBox: gk, ForwardPath: "67c2b59c8741c3fe3f9db52e4e4a72ed"})
	}()

	func() {
		b := packr.New("gocker:genny:gocker", "../gocker/templates")
		b.SetResolver("Gockerfile.plush", packr.Pointer{ForwardBox: gk, ForwardPath: "4f59bf6e3a17584507bc2e40f6b668e7"})
	}()

	return nil
}()
