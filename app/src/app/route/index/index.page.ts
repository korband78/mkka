import { Component, OnInit, ViewChild, ChangeDetectorRef } from '@angular/core';
import { SwiperOptions } from 'swiper';
import { SwiperComponent } from 'ngx-useful-swiper';

declare var Kakao, kakao;

@Component({
  selector: 'route-index',
  templateUrl: './index.page.html',
  styleUrls: ['./index.page.scss'],
})
export class IndexPage implements OnInit {
  kakao_key: string = "618005443a12c1dd30373fcdae3b3a25";
  title: string = "민기♡경아 결혼합니다";
  desc: string = "11월 8일 오후 2시 30분, 민기♡경아 결혼합니다";
  url: string = "https://mkka.xxz.kr";

  is_show_image: boolean = false;

  is_show_contact: boolean = true;

  @ViewChild('swiper', { static: false }) swiper_elem: SwiperComponent;
  config: SwiperOptions = {
     autoHeight: false,
     allowTouchMove: true,
     navigation: {
       nextEl: '.swiper-button-next',
       prevEl: '.swiper-button-prev'
     },
     loop: true,
     observer: true,
     observeParents: true,
     effect: 'flip'
  };

  constructor(
    private changeDetector: ChangeDetectorRef
  ) { }

  ngAfterViewInit() {
    this.changeDetector.detectChanges();
  }

  ngOnInit() {
    setTimeout(() => {
      try {
        Kakao.init(this.kakao_key);
        // 카카오링크 버튼을 생성합니다. 처음 한번만 호출하면 됩니다.
        Kakao.Link.createDefaultButton({
            container: '#kakao_link', // 버튼 id
            objectType: 'location', // 카카오톡 링크 타입
            content: {
                title: this.title, // 타이틀
                description: this.desc, // 상세정보
                imageUrl: `${this.url}/assets/img/sns.jpg`, // 이미지
                link: {
                    mobileWebUrl: this.url, // 모바일 주소 걍 location.href
                    webUrl: this.url  // 웹 주소 걍 location.href
                },
                imageWidth: 800, // 이미지가로
                imageHeight: 600// 이미지 세로
            },
            address: "경기 성남시 분당구 판교역로226번길 16" // container 가 location 일때 위치보기 아이콘 이 갈 주소 (ex. 경기 고양시 덕양구 용두로47번안길 47)
        });

        var map_x = 37.40051272333002;
        var map_y = 127.11146888843061;
        var container = document.getElementById('map');
        var options = {
            center: new kakao.maps.LatLng(map_x, map_y),
            level: 3
        };

        var map = new kakao.maps.Map(container, options);
        var markerPosition  = new kakao.maps.LatLng(map_x, map_y);

        // 마커를 생성합니다
        var marker = new kakao.maps.Marker({
            position: markerPosition
        });

        // 마커가 지도 위에 표시되도록 설정합니다
        marker.setMap(map);  
      } catch(err) {}
    }, 1000);
  }

  kakao_story() {
    Kakao.Story.open({
      url: this.url,
      text: this.title,
    });
  }

  show_image(index: number) {
    this.swiper_elem.swiper.slideTo(index);
    this.is_show_image = true;
  }

  next_slide() {
    this.swiper_elem.swiper.slideNext();
  }

  prev_slide() {
    this.swiper_elem.swiper.slidePrev();
  }

  get_download_image_url(only_index: boolean = false) {
    let result: string = "";
    if (this.swiper_elem && this.swiper_elem.swiper) {
      if (only_index) {
        result = `${this.swiper_elem.swiper.activeIndex}`;
      } else {
        result = `assets/img/${this.swiper_elem.swiper.activeIndex}.jpg`;
      }
    }
    return result;
  }
}
