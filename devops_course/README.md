```yml
DevOps (development และ operations)
    1  จุดเดน่คือ ทำแบบ automate และส่งให้ลูกค้าได้เร็ว
        1. การรวมตัว
            - รวมทุกทีมที่มีวางแผนด้วยกันทั้งทีม dev และ ops
        2. การทำงานร่วมกัน
            - รวมทีม dev และ op
        3. การทำงานแบบอัตโนมัติ หรือ automated
        4. อิทธิพลต่อ SDLC
            - เป็นการทำงานแบบ agile ส่งทีละ feature ให้ลูกค้าหรือส่งแบบซอยย่อยได้
    2 ขั้นตอนการทำงานแบบ infinity loop
        1.การวางแผน
            - กำหนดขอบเขตหรือจุดประสงค์ในการพัฒนา
        2. การพัฒนา
            - รวบรวมความต้องการและกำหนดขอบเขตของ project
        3. การทดสอบ
        4. การปรับใช้
        5. การตรวจสอบ
            - ในสภาพแวดล้อมสามารถทำงานได้จริงหรือเปล่า
    3. คำศัพท์ DevOps ที่ควรทราบ
        CI/CD(Continued intigetion / Continued Delivery,Continued Deployment)
        หลักการทำงาน
            CI(Code >>..push..>> Github(Server) >>..Test..>> Automate testing) 
            เมื่อtest ผ่านแล้วจะไปทำการ build ของ 
            CD(Code >>..build..>> >>..registerImage..>> Docker >>..Deploy..>> )
        CI(Continued intigetion)
            Conde
                1. โค้ดที่ทำให้ application รันได้
                2. โค้ดประเภท IaC(infarstructer as code)
                    - เขียนเพื่อสร้าง server
                    - เขียนในรูปของโค้ดว่าต้องการ ram เท่าไหร่ cpu เท่าไหร่ และอื่นๆ
                    - เราจะใช้ teraform
                3. โค้ดประเภท CaC(Comfiguration as code)
                    - การตั้งค่า config เพื่อปรับแต่งทรัพยากรบน server
            Github(Server)
                - การเก็บโค้ดบน Github(Server)
                - จัดเก็บโค้ดแต่ละ versions
                - ทำ workflow เมื่อมีเหตุการณ์นี้เกิดขึ้นจะให้ทำอะไรต่อไป
                    - เมื่อเกิดการ push จะทำการตรวจสอบว่ามันจะผ่านที่เราเขียน automate testing หรือไม่ (Push >> runtest)
        CD(Continued Delivery,Continued Deployment)
            Code
                - หลังจาก ren test ผ่าน
            Docker
                - ส่งโค้ดเข้ามา build image หลังจากนั้นเอาไปฝากไว้ที่ตัว server และทำการ deploy
        ทั้งหมดต้องมีการ monitoring and logging
            - ติดตามและเก็บ log เพื่อดูการทำงาน ว่าเกิดอะไรขึ้นบ้าง
    4. ประโยชน์ของ DevOps
        - ส่งมอบได้เร็ว
        - มีประสิทธิภาพ
    5. DevSecOps (dev security ops)
        - ความปลอดภัยในการส่ง

Clound computing
    - มี 2 รูปแบบ
        1. server ภายนอก หรือที่เราใช้บริการในการดูแล server
        2. server ภายใน gเราดูแลเอง
    - ประเภทของ cloud computing xประกอบด้วย 3 ประเภทดังนี้
        1. infrastructure as a service (IaaS)
            - คือบริการแบบ cloud computing ที่มีการเตรียมในเรื่องของ server ให้กับเรา ลงทุกอย่างเองได้
                - ยกตัวอย่างเช่น aws, azure, google clound
        2. platform as a service (PaaS)
            - คือ บริการที่เตรียม platform หรือพวก server พวกทรัพยากรให้กับเราเอง และโยนโค้ดของเราขึ้นไปอย่างเดียว
                - ยกตัวอย่างเช่น heroku
        3. Software as a service (SaaS)
            - software ที่รันอยู่บน cloud
                - ยกตัวอย่างเช่น office365
    - ประโยชน์
        - ลดค่าใช้จ่าย
        - เพิ่มความปลอดภัย
        - เพิ่มประสิทธิภาพ
        - ความน่าเชื่อถือ
    - ตัวอย่าง บริการของ AWS
        - Amazon Elastic Compute clound (EC2)
            - บริการ server เป็นการใช้ VM(vitual machine) server เสมือน เรียกกันว่า VPS(Virtual private server)
        - Amazon Simple Storage Service (S3)
            - บริการฝากไฟล์
        - Amazon Lamba
            - serverate
    - Clound native คือ
        -  cแนวทางที่ software ใช้ในการสร้าง ปรับใช้ และจัดการ application ที่ทันสมัยในสภาพแวดล้อมการประมวณผลบล cloud
Microservices
    - Monolithic คืออะไร
        - การออกแบบ application แบบ Monolithic หรือก็คือการที่เรารวมทุก feature ไว้ในแอพเดียว ซึ่งใน application จะมีการใช้ ฐานข้อมูลร่วมกัน
        - คล้ายกับของระบบ ERP ที่รวม order,stock,invent ต่างๆอยู่ด้วยกัน
    - ข้อดีของ Monolithic
        - ง่ายต่อการเขียนโค้ด
        - แก้ไขได้ง่าย
        - ง่ายต่อการ deploy
        - ไม่ยุ่งยากต่อการ scale ระบบ
    - ข้อเสียของ Monolithic
        - ผูกติดกับภาษาของโค้ดไปแล้ว สามารถใช้ได้คือ 1 ภาษา
        - ความยุ่งยากในการทำงานเป็นทีม
        - ความช้าในการ deploy
            - ยกตัวอย่างตอนที่ deploy ใน ax
        - การ Scale ระบบไร้ประสิทธิภาพ
            - ยกตัวอย่างถ้าเราอยากเพิ่มให้ order ลองรับจำนวนคนได้เราต้องก็อปโค้ดอีกตัวมาเพื่อรองรับคน แต่ในโค้ดนี้นกลับมี stock เข้ามาเกี่ยวด้วยทั้งที่เราไม่ได้อยากเพิ่มการรองรับคนของ stock
    - จากที่กล่าวมาทั้งหมด การแก้ปัญหาทั้งหมดจึงต้องใช้ Microservices
        - Microservices คืออะไร
            - Microservices จะถูกออกแบบตามหลักการของ Cloud Native โดยใน 1 ระบบจะประกอบไปด้วยหน่วยย่อยๆ หรือ service ที่เชื่อต่อกันแบบหลวมๆ แต่ละ services มีขนาดเล็กและแยก deploy ได้อย่างอิสระ
            - เช่น แยก order เป็น 1 service, แยก stock เป็น 1 service, แยก payment เป็น 1 service และใช้ database ของใครของมัน
        - ลักษณะพื้นฐานของ Microservices
            - Autonomy สามารถแยกพัฒนาได้ เช่น มีทีมสำหรับการพัฒนาเซอร์วิสของ User, Article, Order และ Payment แยกแตกต่างกันออกไป พร้อมทั้งสามารถแยก Deploy หรือ scale แยกแต่ละเซอร์วิสได้โดยไม่กระทบกับเซอร์วิสอื่น
            - Specialization แต่ละเซอร์วิสได้รับการออกแบบโดยอาศัยการแบ่งแยกตาม Business Capability เพื่อให้สามารถแก้ปัญหาทางธุรกิจนั้น ๆ ได้อย่างเฉพาะเจาะจง
            - Agility ความคล่องตัวนั้นเกิดจากการที่ Microservices ส่งเสริมให้องค์กรมีการออกแบบทีมขนาดเล็กเพื่อเป็นเจ้าของแต่ละเซอร์วิส
            - Flexible Scaling: Microservices ทำให้เกิดการ scale แต่ละเซอร์วิสได้อย่างอิสระโดยไม่กระทบกับเซอร์วิสอื่น ๆ
        - API Gateway คืออะไร
            - หน้าที่คือ ตั้งค่าว่าสามารถเข้าถึงอะไรได้บ้าง หรือสามารถตรวจสอบสิทธิ์ได้
                - สามารถทำ late limit ได้ โดยที่สามารถกัน bot ยิงได้ หรือกันคนยิงบ่อยๆได้ เช่น ใน 1 นาทีสามารถยิงได้ไม่เกิน 10 ครั้ง 
        - การสื่อสารระหว่าง Service
            - ยกตัวอย่างเช่น order ต้องการ ดู stock เพื่อสามารถเช็คได้ว่าของพอขายไหม
            - สามารถสร้างการสื่อสารได้ 2 รูปแบบคือ Synchronous Messages และ Asynchronous Messages
                1 การสื่อสารแบบ Synchronous Messages(ผสานจังหวะหรือแบบต่อเนื่อง)
                    - ยกตัวอย่าง ทำงานผ่าน Restful API
                        - order ตรวจสอบ stock ผ่าน api 
                    - -ข้อเสีย อาจจะประมวณผลช้า หรือต้องเข้าคิวในการประมวณผล
                2. การสื่อสารแบบ Asynchronous Messages(ไม่ผสานจังหวะหรือแบบไม่ต่อเนื่อง)
                    - วิธีการสื่อสารแบบไม่ต้องรอการตอบกลับ
                    - มี channel คั้นกลาง หรือท่อคั้นกลาง
                    - หรือก็คือเมื่อทำเสร็จเดี๋ยวมีการแจ้งเตือนไปเองไม่ต้องมารอที่หน้าเว็บ อาจจะผ่าน email หรืออื่นๆ
        - Saga Pattern
            - เป็นกระบวณการในการ rollback ของแต่ละ service ในกรณีที่บาง service อาจจะไม่ผ่านกระบวณการ หรือที่เรียกว่า Compensating Transaction
            - รูปแบบการทำงานของ Saga แบ่งออกเป็นส 2 ประเภท คือ Choreography-based Saga และ Orchestration-based Saga
                1. Choreography-based Saga
                    - รูปแบบของ Choreography-based Saga นี้จะเรียกเซอร์วิสต่าง ๆ ที่เข้ามาเกี่ยวข้องว่า Saga Participant มีหน้าที่ในการสอดส่องเหตุการณ์ที่เกี่ยวข้องกับการทำงานของตน เมื่อได้รับเหตุการณ์นั้นแล้วจึงทำการประมวลผลต่อและส่งต่อเหตุการณ์ใหม่ไปยัง channel เพื่อให้เซอร์วิสถัดไปมองเห็นและดำเนินการสำหรับขั้นตอนถัดไป
                2. Orchestration-based Saga
                    - Orchestration-based Saga เป็นรูปแบบที่มีเซอร์วิสหนึ่งทำหน้าที่เป็นผู้ควบคุมขั้นตอนการดำเนินงาน เรียกเซอร์วิสนี้ว่าเป็น Orchestrator โดยทำหน้าที่ดำเนินการผ่านการตรวจตราเหตุการณ์ต่าง ๆ ที่เกี่ยวข้องแล้วจึงประมวลผลด้วยการส่งเหตุการณ์ที่ควรเป็นไปยังเซอร์วิสปลายทางที่มีหน้าที่ในการจัดการสิ่งนั้น Orchestrator จะมีการบันทึกสถานะของการดำเนินงานต่าง ๆ ที่เกี่ยวข้องในรูปแบบของ Saga Log
            - Compensating Transaction
                - ธุรกรรมย่อยใด ๆ ในเซอร์วิสต้องมี Compensating Transaction ที่สอดคล้องกันเกิดขึ้นด้วยเพื่อใช้ในการย้อนกลับเมื่อการทำงานในเซอร์วิสนั้นล้มเหลว เช่น T1: Order requested จะมีธุรกรรมย้อนกลับที่เข้าคู่กันคือ C1: Order cancelled เป็นต้น
Docker 
    - การทำงานของ docker และ Docker compose 
        - สร้าง Docker file >>..docker build..>> Docker images >>..docker run..>> conntainer
        - คำสั่งที่ควรู้
            - docker run -p 5152:80 test/intro-to-decops-ui:1.0
            - docker ps
                - ตรวจสอบว่า docker ไหน run อยู่บ้าง
            - docker stop name-container
                - คำสั่งหยุดการทำงาน
            - docker start name-container
                - คำสั่งเริ่มการทำงาน
            - docker rm name-container
                - คำสั่งลบ
            - docker build -t patiphansak/devops-api:1.0 .
            - docker compose up
            - docker compose down
            - docker push patiphansak/devops-api:1.0
        - docker compose
            - ถ้ารเรามีสัก 10 service จะเสียเวลาในการ start ขึ้นมา เพื่อที่จะแก้ไขจุดนี้จึงต้องใช้ docker compose
            - เป็นการรวมทุก service ไว้ในไฟล์เดียว และทำการ boot ขึค้นมาทีเดียวได้
        - สร้างไฟล์ docker-compose.yml 
            - ใช้คำสั่ง docker compose up
    - docker build and push
        - การสร้าง Docker Images
            - การสร้าง Docker Image นั้นต้องอาศัยชุดคำสั่งซึ่งแสดงถึงขั้นตอนในการได้มาซึ่ง Image ชุดคำสั่งดังกล่าวจะเขียนไว้ในไฟล์ชื่อ Dockerfile คำสั่งแต่ละบรรทัดของ Dockerfile นั้นเมื่อแปลงเป็น Docker Image จะทำให้เกิดสิ่งที่เรียกว่า Layers ขึ้นมา เราจึงกล่าวได้ว่า Docker Image แท้ที่จริงแล้วก็คือ Layers หลาย ๆ Layers ที่เกิดจากชุดคำสั่งใน Dockerfile ประกอบรวมกันเป็น Image นั่นเอง
            - ภายหลังการสร้างชุดคำสั่งใน Dockerfile เราจะออกคำสั่ง docker build เพื่อทำการสร้าง Docker Image และเพื่อให้ Image ดังกล่าวสามารถถูกโหลดไปใช้งานได้ เราต้องส่ง Image ไปจัดเก็บไว้ในส่วนจัดเก็บบนคลาวด์ที่เรียกว่า Docker Registry โดยทั่วไปแล้ว Docker Registry ที่เป็นมาตรฐานและใช้งานโดยทั่วไปคือ Docker Hub เราใช้คำสั่ง docker push เพื่อนำส่ง Image ไปจัดเก็บยัง Docker Registry
        - ตัวอย่าง layer ที่ cash ได้ 
```yml
            FORM golang:1.22.1 #เป็นการบอกว่า Image ของเราจะต่อยอดการทำงานมาจาก Image ที่ชื่อ golang โดยมีเวอร์ชันของแท็กคือ 1.22.1
            WORKDIR /app #สร้างโฟลเดอร์คือ /app และย้ายการทำงานไปยังโฟลเดอร์ดังกล่าว
            COPY . . #ให้ทำการคัดลอกไฟล์ต่าง ๆ ทั้งหมดมาวางที่ /app จุดตัวแรกคือการก็อปไฟล์ทั้งหมดจาก dir ปัจจุบัน ไปไว้ที่ จุดตัวที่สองหรือก็แค่ที่ WORKDIR /app
            RUN go mod download #ออกคำสั่ง go mod download
            RUN CGO_ENABLED=0 go build -o api
            CMD ["/app/api"]
            # เมื่อคอนเทนเนอร์ได้รับการรัน ให้ออกคำสั่ง /app/api เพื่อรัน Executable File
```
        - ตัวอย่าง layer ที่ไม่ cash
```yml
            FROM golang:1.22.1 AS builder
            WORKDIR /app
            COPY go.mod go.sum ./
            RUN go mod download
            COPY . .
            RUN CGO_ENABLED=0 go build -o api
```
```yml
        - คำสั่งสร้าง docler image
            - docker build -t ชื่อในdockerHub:version .(dot ตรงนี้คือ dir ปัจจุบันของ Dockerfile)
             - docker build -t patiphansak/devops-api:1.0 .
    - api gateway and kong
        - API Gateway หน้าที่
            - Filter: กรอง traffic หรือ request ที่เข้ามาจากทั้ง web, mobile, web service และอื่น ๆ โดยแก้ปัญหาได้หลายหลาก เช่น การมีหลาย Entry point (URL) ของแต่ละเซอร์วิส ทำให้ยากต่อการจดจำและเข้าถึง และปัญหาด้านความปลอดภัย เป็นต้น
            - Single Entry Point: เปิดเผย Endpoint ของ APIs ต่าง ๆ โดยเปิดเผยเพียงแค่ endpoint เดียว แล้วให้ request ที่เข้ามาถูกพิจารณาส่วนของ Path ก่อน เช่น เมื่อพาธเป็น /articles ให้วิ่งต่อไปยังเซอร์วิสคือ Article ส่วนนี้จึงกล่าวได้ว่า API Gateway ทำหน้าที่ Routing
            - API Management: การจัดการกับ API ต่าง ๆ ที่จะถูกส่งต่อไป เช่น rate limit เป็นต้น
            - Security mechanism: กลไกของเรื่องความปลอดภัย ไม่ว่าจะเป็นการเข้าถึงบางเซอร์วิสที่ต้องผ่านการลอคอินก่อน หรือการทำ logging เป็นต้น
    - Terraform
        - ใช้บริการ digital ocean
        - อ่าน doc และทำตาม
GitOps
Kubernetes
    - Container Orchestration
        - Container Orchestration เป็นกระบวนการที่ใช้เทคโนโลยีอัตโนมัติในการจัดการวงจรชีวิตของคอนเทนเนอร์ ซึ่งรวมถึงการจัดสรรทรัพยากร การปรับขนาด (ขยายและย่อ) การจัดการเครือข่าย และการดูแลความพร้อมใช้งาน
        - ปํญหาที่ทำไมเราต้องใช้ตัวนี้ คือ ทำอย่างไรจึงจะสามารถ Scale คอนเทนเนอร์ขึ้นมาหลายตัวได้อย่างอัตโนมัติเพื่อรองรับโหลดของผู้ใช้งานเมื่อทรัพยากรถึงขีดจำกัดที่กำหนด หรือเราจะสามารถจัดการการเชื่อมต่อเน็ตเวิร์กของคอนเทนเนอร์ที่หลากหลายได้อย่างไร เป็นต้น ปัญหาต่าง ๆ เหล่านี้สามารถแก้ไขได้ด้วย Container Orchestration
    - Control Plane: หัวใจของ Kubernetes Cluster ที่ควบคุมการทำงานทั้งหมด ส่วนนี้ประกอบด้วย:
        - kube-apiserver: เป็นหน้าตาของ Kubernetes Control Plane รวม API ให้เรียกใช้งานเพื่อจัดการคำขอจากภายในระบบ Cluster และภายนอกระบบ Cluster
        - kube-scheduler: ตรวจสอบสภาพความพร้อมของ Cluster และกำหนดการทำงานของ Pod ให้กับ Node ที่เหมาะสม
        - kube-controller-manager: มีการตรวจสอบทรัพยากรปัจจุบันว่ามีสถานะเป็นเช่นไร (Current State) Controller Manager จะพยายามเปลี่ยนแปลงสถานะปัจจุบันให้เป็นสถานะที่ต้องการ (Desired State)
        - Cloud Controller Manager: ส่วนนี้ใช้เพื่อติดต่อกับ Cloud Provider API เพื่อจัดการให้เกิดการสร้างหรือใช้งานทรัพยากรบน Cloud Provider นั้น