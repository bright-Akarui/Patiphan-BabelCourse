Next.js
- คำสั่งสร้าง next
    - npm create next-app@latest
- Static site generation
    - การที่ next.js สร้าง api และเก็บข้อมูลนั้นเอาไว้ตลอด หรือ render ทางฝั่ง server แล้วค่อยส่ง html ออกไป
    - จะทำให้ค่าที่ได้จาก api เป็นค่าเดิมเสมอ
- Incremental Static Regeneration
########################
```ts
-typescript
    - const theme = {
        colors: {
            primary: '#ee22ee',
            secondary: '#ffee66'
        }
    } as cont // กำหนดเพื่อให้เห็นค่าสีใน theme.colors.primary

    theme.colors.primary
    - การใช้งาน interface{}
        - กำหนดเพื่อ ไม่ให้เราใส่ชนิดข้อมูลผิด กันข้อผิดพลาด
        - ยกตัวอย่างการใช้งาน วิธีที่ 1 
            - interface Amimal {
                age: number;
            };
            interface Person extends Animal{ // extends การสืบทอด
                name: string;
                gender: 'male' | 'female';
                socials?: { // เครื่องหมาย ? ใส่เพื่อให้เป็น optional
                    line?: string;
                    facebook?: string;
                }
            };
            const somchai: Person = {
                name: 'Somchai',
                age: 24;
                gender: 'male' | 'female',
                socials?: {
                    facebook?: string,
                }
            };
        - ยกตัวอย่างการใช้งาน วิธีที่ 2
            - type Animal {
                age: number;
            };
            - type Address {
                lat: number;
                lng: number;
            };
            type Person = Animal & {
                name: string;
                gender: 'male' | 'female';
                address: Address[];
                socials?: { // เครื่องหมาย ? ใส่เพื่อให้เป็น optional
                    line?: string;
                    facebook?: string;
                }
            };
        - กรณีใช้งานกับ function 1;
            - interface Options{
                x: number;
                y: number;
            };

            function foo(bar: number, options?: Options){

            }
            foo(1, {x: 1, y: 2})
        - กรณีใช้งานกับ function 2;
            function findById<T extends { id: number }>(items: T[], id: T['id']){
                return items.find ((item) => item.id === id);
            }
            cons products = [
                { id: 1, tile:"title#1"},
                { id: 2, tile:"title#2"},
                { id: 3, tile:"title#3"},
            ]

            findById(products, 2);

        - utility type
            - type Animal {
                age: number;
            };
            // - type Address {
            //     lat: number;
            //     lng: number;
            // };
            type Address = Record<'lat' | 'lng', number>;

            type Person = Animal & {
                name: string;
                gender: 'male' | 'female';
                address: Record<'lat' | 'lng', number>[];
                socials?: Partial<{
                    line?: string;
                    facebook?: string;
                }>
            }

            type BasicInfo = Pick<Person, "name" | "gender"> // เอาแค่ name, gender ใน person
            type OtherInfo = Omit<Person, "name" | "gender"> // ไม่เอาแค่ name, gender ใน person

        การใช้งาน
            - function getTheme() {
                return {
                    colos: {
                        primary: 'eeeeee',
                        secondary: 'fffff'
                    }
                }
            }
            // type GetThemeReturn = ReturnType<getTheme> //ไม่สามารถใช้ typescript ได้
            type GetThemeReturn = ReturnType<typeof getTheme> //สามารถใช้ typescript ได้ จะได้{colos: {primary: 'eeeeee',secondary: 'fffff'}}
            type Colors = GetThemeReturn['colors'] //เข้าถึงชนิดข้อมูล colos หรือ type ของมัน จะได้ primary: string,secondary: string
            type ColorKeys = keyof ;// จะได้ "primary","secondary"

- React และ Next.js
    - npm create next-app@latest
        - code project-name
    - สำหรับตรวจสอบว่าโค้ดได้มาตรฐานไหม
        - npm add -D eslint-config-next eslint-config-prettier eslint-plugin-prettier prettier prettier-plugin-tailwindcss @typescript-eslint/eslint-plugin
- การใช้งาน useEffect แบบ connect และ close
    - useEffect(() => {
        connect(id)

        //Unmounting or Cleanup
        return () => close(id);
    }, [id])
- การดึงข้อมูลเข้ามาตรงๆ หรือเชื่อมกับฐานข้อมูลตรงๆ และการจัดการข้อมูลจากการดึงข้อมูลฃ
    - เราสามารถระบุในไฟล์ได้ว่า ไฟล์นี้จะถูกกระทำที่ฝั่ง server หรือฝั่ง client โดยการระบุดังนี้
        - 'use client'
        - 'use server'
    - ssg,ssr,isr
        - ssg คือ การที่ next.js สร้างไฟล์ html ขึ้นมาหลังจาก build และส่งให้ทางฝั่ง client และจะไม่ถูกสร้าง html ใหม่อีก
        - ssr คือ การที่ next.js สร้างไฟล์ html ขึ้นมาหลังจาก build และส่งให้ทางฝั่ง client และจะถูกสร้างใหม่ทุกครั้งเมื่อมีการร้องขอเข้ามา
        - isr คือ การที่ next.js สร้างไฟล์ html ขึ้นมาหลังจาก build และส่งให้ทางฝั่ง client และจะถูกสร้างใหม่ทุกๆ กี่นาทีตามทำเรากำหนด
    - ถ้าไม่ใส่อะไรเลยคือ "ssg"
    - ถ้าใช้วิธีการดึงข้อมูลตครงๆโดยไม่ผ่านการ fetch เพื่อที่จะให้ได้ข้อมูลไหมต้องใส่ดังตัวอย่างเพื่อให้ข้อมูลเป็น dynamic หรือการทำ "ssr"
        - export const dynamic = 'force-dynamic'
    - การทำ static หรือ การ revalidate ทุกๆ 15 วินาที หรือการทำ "isr"
        - export const revalidate = 15
- การจัดการการดึงข้อมูลผ่าน API หรือ การ fetch
    - หาต้องการให้ build ใหม่ทุกครั้ง หรือใช้เป็นวิธีของ "ssr" ต้องใช้ {cache: 'no-store'} จะได้ข้อมูลใหม่เสมอ
        - ยกตัวอย่างเช่น
        - const res = awwait fetch(`http://localhost:5151/articles/${id}`,{cache: 'no-store'})
    - หาต้องการให้ build ทุกๆกี่นาทีตามที่เรากำหนด หรือใช้เป็นวิธีของ "isr" ต้องใช้ {next: {revalidate: 15}} จะได้ข้อมูลใหม่เสมอ
        - ยกตัวอย่างเช่น
        - const res = awwait fetch(`http://localhost:5151/articles/${id}`,{next: {revalidate: 15}})
        - ถ้าใช้จะต้องกำหนดข้อมูลตั้งต้นให้ด้วย
            - export const generateStaticParams = () =>{
                return [{id:'1'}]
            }
-Next.js API
    - CRUD Example
        - การรับ parameter จะรับได้ 2 ตัวแปร 1.Request 2.id (req: Request,{ params: { id }}: Params)
            - Request คือ body ที่ถูกส่งเข้ามาดำเนินการ
            - ยกตัวอย่างเช่น export const PATCH = async (req: Request,{ params: { id }}: Params)
        - Read or GET
            export const GET = () => {
            const articles = [{id:1},{id:2}]
            return new Response(JSON.stringify(articles),{
                status: 200,
                headers: {
                'Content-Type': 'application/json'
                }
            })
            }
        - Create or POST
            export const POST = async (req: Request) => {
                const form = await (req.json() as Promise<CreateArticleInput>);
                const formValidation = await add.safeParseAsync(form);

                if(!formValidation.success) {
                    return new Response(JSON.stringify(formValidation.error),{
                    status: 422,
                    headers: {
                        'Content-Type': 'application/json'
                    }
                    })
                }
                const article = await create(formValidation.data);
                return new Response(JSON.stringify(article),{
                    status: 201,
                    headers: {
                    'Content-Type': 'application/json'
                    }
                })
            }
