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
            type ColorKeys = keyof // จะได้ "primary","secondary"
- React และ Next.js
    - npm create next-app@latest
        - code project-name
    - สำหรับตรวจสอบว่าโค้ดได้มาตรฐานไหม
        - npm add -D eslint-config-next eslint-config-prettier eslint-plugin-prettier prettier prettier-plugin-tailwindcss @typescript-eslint/eslint-plugin