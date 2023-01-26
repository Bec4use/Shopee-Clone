<script>
    let second = 59
    let mins = 59
    let hours = 0
    setInterval(() => {
        second -= 1
      if(second==0){
        second=59
    }
    }, 1000)
    setInterval(() => {
        mins -= 1
      if(mins==0){
        mins=59
    }
    }, 60000)
    setInterval(() => {
        hours -= 1
    }, 3600000)
  </script>
<div class="container mt-4">
    <div class="flashmenu d-flex">
        <div class="flash-l d-flex">
            <div class="img me-4">
                <img src="https://deo.shopeemobile.com/shopee/shopee-pcmall-live-sg/fb1088de81e42c4e538967ec12cb5caa.png" alt="">
            </div>
            <div class="countdown d-flex">
                <div class="timeitem">{('00'+hours).slice(-2)}</div><h4>:</h4>
                <div class="timeitem">{('00'+mins).slice(-2)}</div><h4>:</h4>
                <div class="timeitem">{('00'+second).slice(-2)}</div>
            </div>
        </div>
        <div class="flash-r">
            <a href="/">ดูเพิ่มเติม &gt;</a>
        </div>
    </div>
    <div class="flashcontent">
        <div class="img-flashsale d-flex container">
            <a href="/" class="flashitem-1">
                <div class="flashsale1">
                    <div class="flashimg-1"></div>
                    <div class="flashimgcover"></div>
                    <div class="price-1">฿11,900
                        <div class="flashbar">
                            <img src="https://deo.shopeemobile.com/shopee/shopee-pcmall-live-sg/ac7f81d9ee062223753413ec98497a86.png" alt="">
                        </div>
                    </div>
                    <div class="discount-1">
                        <span>ลด</span>
                        <span>73%</span>
                    </div>
                    
                </div>  
            </a>
            <a href="/" class="flashitem-1">
                <div class="flashsale1">
                    <div class="flashimg-2"></div>
                    <div class="flashimgcover"></div>
                    <div class="price-1">฿18,290<div class="flashbar">
                        <img src="https://deo.shopeemobile.com/shopee/shopee-pcmall-live-sg/ac7f81d9ee062223753413ec98497a86.png" alt="">
                    </div> 
                    </div>
                    <div class="discount-1">
                        <span>ลด</span>
                        <span>80%</span>
                    </div>
                </div>  
            </a>
            <a href="/" class="flashitem-1">
                <div class="flashsale1">
                    <div class="flashimg-3"></div>
                    <div class="flashimgcover"></div>
                    <div class="price-1">฿2,710
                        <div class="flashbar">
                            <img src="https://deo.shopeemobile.com/shopee/shopee-pcmall-live-sg/ac7f81d9ee062223753413ec98497a86.png" alt="">
                        </div>
                    </div>
                    <div class="discount-1">
                        <span>ลด</span>
                        <span>73%</span>
                    </div>
                </div>  
            </a>
            <a href="/" class="flashitem-1">
                <div class="flashsale1">
                    <div class="flashimg-4"></div>
                    <div class="flashimgcover"></div>
                    <div class="price-1">฿1,680
                        <div class="flashbar">
                            <img src="https://deo.shopeemobile.com/shopee/shopee-pcmall-live-sg/ac7f81d9ee062223753413ec98497a86.png" alt="">
                        </div>
                    </div>
                    <div class="discount-1">
                        <span>ลด</span>
                        <span>73%</span>
                    </div>
                </div>  
            </a>
            <a href="/" class="flashitem-1">
                <div class="flashsale1">
                    <div class="flashimg-5"></div>
                    <div class="flashimgcover"></div>
                    <div class="price-1">฿25,900
                        <div class="flashbar">
                            <img src="https://deo.shopeemobile.com/shopee/shopee-pcmall-live-sg/ac7f81d9ee062223753413ec98497a86.png" alt="">
                        </div>
                    </div>
                    <div class="discount-1">
                        <span>ลด</span>
                        <span>73%</span>
                    </div>
                </div>  
            </a>
            <a href="/" class="flashitem-1">
                <div class="flashsale1">
                    <div class="flashimg-6"></div>
                    <div class="flashimgcover"></div>
                    <div class="price-1">฿14,900
                        <div class="flashbar">
                            <img src="https://deo.shopeemobile.com/shopee/shopee-pcmall-live-sg/ac7f81d9ee062223753413ec98497a86.png" alt="">
                        </div>
                    </div>
                    <div class="discount-1">
                        <span>ลด</span>
                        <span>73%</span>
                    </div>
                </div>  
            </a>
        </div>
    </div>
</div>

<style>
    .container {
        height: 100%;
        max-width: 75rem;
        background: #fff;
    }
    .flash-l{
        margin: 1rem 1rem;
        /* padding: 1rem 0; */
    }
    .timeitem{
        background-color: #000;
        color: #fff;
        padding: 0.1rem 0.425rem;
        height: 1.875rem;
        font-size: 1.2rem;
        border-radius: 5px;
        margin:0 0.2rem;
    }
    .flashmenu{
        justify-content: space-between;
    }
    .flash-r{
        margin: 1.5rem 1rem;
    }
    .flash-r a{
        text-decoration: none;
        color: #ee4d2d;
    }
    .img-flashsale{
        margin: 0 auto;
    }
    .flashsale1{
        background-color: #fff;
        height: 14.8rem;
        width: 12.125rem;
        margin: 0 1px;
        text-align: center;
        justify-content: center;
    }
    .flashimg-1{
        background-image: url("https://cf.shopee.co.th/file/a8d55f5c20e1501e76ff32db6045e572_tn");
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: relative;
        background-size: cover; 
    }
    .flashitem-1 {
        text-decoration: none;
        color: black;
    }
    .discount-1 {
        position: relative;
        background-color: rgba(255,212,36,.9);
        width: 2.25rem;
        height: 2rem;
        transform: scale(1.4);
        border-radius: 0 0 0 50%;
        left: 150px;
        top: -215px;
        z-index: 99;
    }
    .discount-1 span {
        display: flex;
        font-size: 0.625rem;
        text-align: center;
        justify-content: center;
    }
    .discount-1 span+span {
        color: red;
    }
    .price-1{
        color: #ee4d2d;
    }
    .flashimg-2{
        background-image: url("https://cf.shopee.co.th/file/9d1485e950870f8e5cea2cee9c439de4_tn");
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: relative;
        background-size: cover;
    }
    .flashimgcover{
        background-image: url("https://cf.shopee.co.th/file/7ccd657497baadbe425c2ac1dd464c7c_tn");
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: absolute;
        background-size: cover;
        top: 93.5rem;
    }
    .flashimg-3{
        background-image: url("https://cf.shopee.co.th/file/17839b72cde60b4d10375b77d50cc643_tn");
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: relative;
        background-size: cover;
    }
    .flashimg-4{
        background-image: url("https://cf.shopee.co.th/file/d8fc110bea3962e9d8c06454ab36390e_tn");
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: relative;
        background-size: cover;
    }
    .flashimg-5{
        background-image: url("https://cf.shopee.co.th/file/7be30879de0dd95195369cc190e78759_tn");
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: relative;
        background-size: cover;
    }
    .flashimg-6{
        background-image: url("https://cf.shopee.co.th/file/1826524b22408a10c4c43daa9da1ddcb_tn");
        height: 10.125rem;
        width: 10.125rem;
        margin: 1rem auto;
        position: relative;
        background-size: cover;
    }
    .flashbar {
        height: 1rem;
        background-image: url("https://deo.shopeemobile.com/shopee/shopee-pcmall-live-sg/ac7f81d9ee062223753413ec98497a86.png");
        width: 50%;
        background-size: cover;
        text-align: center;
        margin: 0;
        margin-left: 1rem;
        position: relative;
    }
    .flashbar img{
        opacity: 1;
        height: 1rem;
        width: 9.375rem;
        opacity: 0.3;
        margin-bottom: 10px;
        position: relative;

    }
</style>