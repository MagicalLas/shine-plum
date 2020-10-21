# shine-plum

Web-UI for Golang Clean Architecture

## How To Use

### Install

```
go get github.com/MagicalLas/shine-plum@latest
```

### Usage

```go
package main

import (
	shine_plum "github.com/MagicalLas/shine-plum"
	"github.com/MagicalLas/shine-plum/internal/example"
)

func main() {
	shine := shine_plum.NewShine()
    u := &example.UseCase{} // Using Your UseCase or Application Service Struct
	plum := shine_plum.NewPlum(u)

	sp := shine_plum.Bright(plum, shine)

	<-sp.TurnOn()
    // Wait during web server execute.
}
```

HTTP 요청을 보내서 Method를 호출하거나 설명을 볼 수 있습니다. 현재는`HOST/<UseCaseName>/<MethodName>` 이런 패턴만 지원하며, 여러 UseCase가 있어도 작동합니다.
Get Method는 메서드에 대한 설명을 볼 수 있습니다. 이 설명은 Method의 주석에서 GoDoc과 같은 형식으로 가져옵니다. 추후에는 Example이나 파라미터에 대한 설명도 제공할 예정입니다.
```
GET: localhost:8080/UseCase/GetAd


Response
> {
>    "Description": "GetAd is get ad.",
>    "UseCase": "GetAd"
>  }

POST: localhost:8080/UseCase/ListAd
body:
{
    "command": {
        "LineitemID": "l",
        "ADNID: 1,
        "UnitID": 2,
        "User": {
            "IFA": "asd",
            "IP":  "ww"
        }
    }
}

Response
> [
>   {
>      "Title":"tt",
>      "Description":"dd",
>      "Creatives":[
>         {
>            "Image":"http://magical.dev",
>            "Icon":""
>         },
>         {
>            "Image":"",
>            "Icon":"http://magical.dev"
>         }
>      ]
>   }
> ]
```

## 설계 테마

이 부분에서는 프로젝트가 어떻게 구상되었고, 어떤 은유들을 사용하고 있는지 설명합니다. 맛있고 빛나는 Plum을 드셔보세요!

### Plum

Plum은 UseCase를 감싼 General한 Golang Object입니다. 맛있는 자두와 같이 UseCase구현체를 씨앗으로 더욱 풍부한 기능들을 제공합니다.

### Shine

Shine은 기존에 존재하지 않는 인터페이스를 정의합니다. Plum을 웹 API로 연결시켜주는 역할을 합니다. Plum을 어떻게 Web API를 통해 사용할 것인지 정의하고 있습니다.
굳이 WebAPI에 국한되지 않고 swagger형식대로 문서를 만드는 것과 같이 새로운 외부와의 interaction을 정의합니다. 단순히 API뿐만 아니라 FrontEnd작업이 있을 때, 1차적으로 Shine과 상호작용하게 됩니다. 

### ShinePlum

빛(Shine)과 열매(Plum)만 있다고 해서 열매에 빛이 나지 않습니다. 열매에 빛을 넣어준 상태가 바로 ShinePlum입니다. ShinePlum을 TuenOn하면 실제로 Web Server가 동작하며 Shine을 통해 Plum을 조작할 수 있습니다.

## Why Las made this project?

버즈빌에서는 클린 아키텍처를 사용해서 MicroService를 구현하고 있습니다. 그 과정에서 API는 gRPC를 사용하고 있습니다. 그러나 gRPC는 HTTP에 비하여는 클라이언트가 불편하고 정리가 잘 되지 않는 점이 있습니다.
또한 웹서버가 아니다보니 System Operation을 위한 API를 어디선가 따로 구현할 필요가 있습니다.

ShinePlum을 이용하면 이러한 Web Admin API Tool을 자동으로 만들어줍니다. 또한 추후에 제공될 기능으로 UseCase객체에 대한 example과 자세한 설명들을 통해서 gRPC가 아닌 API를 쉽게 보낼 수 있습니다.
PostMan이나 OpenAPI형식으로 api에 대한 내용을 공유하기도 편하고 직접 frontend를 만들 수 있어 더 좋은 유저경험을 만들 수도 있습니다.
굳이 PostMan이나 웹페이지가 아니더라도 kiki나 slack command로 바로 요청을 보낼 수 있다는 것도 큰 장점입니다.

## LICENSE

GNU Affero General Public License
