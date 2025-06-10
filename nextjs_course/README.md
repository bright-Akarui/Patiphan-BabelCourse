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

            เช่น
                interface DataRow {
                    id:number | string;
                }
                interface DataGridColumn<T extends DataRow> {
                    field: keyof T;
                    hederName: string;
                }
                interface DataGridProps<T extends DataRow> {
                    title: string;
                    columns: DataGridColumn<T>[]
                    rows?: T[]
                }
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
- csr หรือ (client-site-rendering)
- การเปลี่ยน route หรือ path ที่เข้ามาให้เป็นตามที่เรากำหนด
    - แก้ไขในไฟล์ next.config.mjs เพิ่ม redirects() เข้ามา
        -   /** @type {import('next').NextConfig} */
            const nextConfig = {
            eslint: {
                dirs: ['.'],
                ignoreDuringBuilds: true,
            },
            redirects() {
                return [
                {
                    // source เมื่อวิ่งเข้ามาที่ / 
                    source: '/',
                    // destination ให้เปลี่ยนไปที่
                    destination: '/leaves',
                    // permanent true = ถาวร ,permanent false = ชั่วคราว
                    permanent: true

                }
                ]
            }
            };

            export default nextConfig;
- Next/Navigation
    - useRouter
        - เปลี่ยน path บนหน้าจอ คล้ายกับ link component
        -   const route = useRouter()
            // push เก็บเป็น stack เช่น push รอบที่ 1 เก็บ article push รอบที่ 2 เก็บ article/id ถ้ากด back จะกลับไปที่หน้า article
            route.push('/article')

            // ถ้ากดจะเกิดการแทนที่ ที่หน้าเดิมเลย
            route.replace('/leaves')

            // ถ้ากดจะเกิดการลบ stack บนสุด
            route.back()
    - usePathname
        - ดึงค่าของ path name ออกจาก url ได้
            - http://localhost:3000/api/leaves >>>>> /api/leaves
        -   const pathname = usePathname();
            // จได้ค่า path name ปัจจุบุนออกมา /api/leaves
            console.log(pathname)
    - useSearchParams
        - ดึงค่าของ query string ที่อยู่ใน param ออกมาได้
            - http://localhost:3000/api/leaves?reasonid=1 >>> reasonid=1
        -   const search = useSearchParams();
            // จได้ค่าของ reasonid ออกมาเป็น 1
            console.log(search.get('reasonid'))
    - useParams
        - ดึงค่า id 
            - http://localhost:3000/api/leaves/1 >>>>> 1
- ถ้า type ติด promise อยู่ให้ใส่ Awaited เช่น
    - export type ArticleItem = Awaited<ReturnType<typeof findAll>>[number];
- การเชื่อม database 
    - ใช้ prisma 
    - เริ่มต้นให้ติดตั้ง prisma
        - ออกคำสั่ง npm add -D prisma
    - สร้างการเชื่อมต่อ 
        - ออกคำสั่ง npx prisma init --datasource-provider postgresql
        - จะได้ folder สำหรับเชื่อมต่อมา
    - หลังจากนั้นในโฟลเดอร์ของ prisma ให้ใช้สร้างฐานข้อมูลขึ้นมาในไฟล์ schema.prisma เช่น
        generator client {
        provider = "prisma-client-js"
        output   = "../app/generated/prisma"
        }
        datasource db {
        provider = "postgresql"
        url      = env("DATABASE_URL")
        }
        enum Role {
        ADIMN
        MANAGER
        MEMBER
        }
        enum LeaveStatus {
        PENDING
        APPROVED
        REJECTED
        }
        model User {
        id           Int           @id @default(autoincrement())
        name         String
        email        String        @unique
        image        String?
        role         Role          @default(MEMBER)
        createdAt    DateTime      @default(now())
        updatedAt    DateTime      @default(now())
        leaves       Leave[]
        annoucements Annoucement[]
        }
        model Leave {
        Id              Int         @id @default(autoincrement())
        stats           LeaveStatus @default(PENDING)
        reason          String
        leaveDate       String
        rejectionReason String?
        userId          Int
        createdAt       DateTime    @default(now())
        updatedAt       DateTime    @default(now())
        user            User        @relation(fields: [userId], references: [id])

        // จะระบุได้ว่า user คนหนึ่งจะไม่สามารถมี leaveDate ซ้ำกันได้
        @@unique([userId, leaveDate])
        }
        model Annoucement {
        Id        Int      @id @default(autoincrement())
        title     String
        slug      String   @unique
        excerpt   String
        content   String
        userId    Int
        createdAt DateTime @default(now())
        updatedAt DateTime @default(now())
        user      User     @relation(fields: [userId], references: [id])
        }
    - หลังจากนั้นใช้คำสั่ง npx prisma db push เพื่อสร้าง db
    - และทดสอบการใช้งานโดยใช้คำสั่ง npx prisma db push
    - การสร้่าง migration
        - รันคำสั่ง npx prisma migrate dev --name ชื่อที่ต้องการ
            - npx prisma migrate dev --name init
- ถ้าอยากให้การเรียกใช้ path เวลาเราสร้าง folder แล้วอยากให้ไม่สนใจ Folder ที่สร้างให้ใส่ () ครอบไว้ เช่น (site)
- react query
    - หลักการคือเก็บ cash ไปให้ผู้ใช้เพื่อแสดงผล
    - ดูวิธีติดตั้งได้ที่ tanstcak query 
    - ตัวอย่างการใช้งาน
        'use client'
        import {
            useQuery,
            useMutation,
            useQueryClient,
            QueryClient,
            QueryClientProvider,
        } from '@tanstack/react-query'
        import { ReactNode } from 'react';
        import { ReactQueryDevtools } from '@tanstack/react-query-devtools'
        const queryClient = new QueryClient()

        interface ClientProvidersProps {
            children: ReactNode;
        }

        const ClientProviders = ({ children }: ClientProvidersProps) => {
            return <QueryClientProvider client={queryClient}>{children}
            <ReactQueryDevtools initialIsOpen={false}></ReactQueryDevtools>
            </QueryClientProvider>
        }

        export default ClientProviders;
    - ยกตัวอย่างโค้ดก่อนใช้ react query
        import { Leave } from "@/features/leaves/types"
        import { useEffect, useState } from "react"

        export const useGetLeaves = () => {
        const [leaves, setLeaves] = useState<Leave[]>([]);
        const [loading,setLoading] = useState(true)

        
        const fetchLeaves = async () => {
            const res = await fetch('http://localhost:3000/api/leaves')
            const leave = await (res.json() as Promise<Leave[]>)

            setLeaves(leave)
            setLoading(false)
        }
        useEffect(() => {
            fetchLeaves();
        }, [])
        return {loading:loading , leaves:leaves }
        }
    - โค้ดหลังใช้ react query
        import { Leave } from "@/features/leaves/types"
        import { useQuery } from "@tanstack/react-query";
        export const useGetLeaves = () => {
        return useQuery({
            queryKey: ['leaves'],
            queryFn: async () => {
            const res = await fetch('http://localhost:3000/api/leaves')
            const leaves = await (res.json() as Promise<Leave[]>)
            return leaves
            }
        })
        }
        - 'use client'
            import LevesList from '@/features/leaves/components/levesList'
            import { useGetLeaves } from '@/features/leaves/hooks/api'
            import React from 'react'

            const LeavePage = () => {
            const { data, status } = useGetLeaves();
            console.log(data)
                if (status === 'pending') return <div>Loading...</div>
                if (!data) return <div>No leave found</div>
                return <LevesList leaves={data}></LevesList>
            }

            export default LeavePage

            export const dynamic = 'force-dynamic'
    react-query
        - useQuery = ดึงข้อมูล (GET)
        - useMutation = เปลี่ยนข้อมูล (CREATE / UPDATE / DELETE)
            - useMutation ทำอะไรบ้าง
                - จัดการสถานะ: isLoading, isSuccess, isError

                - มี callback: onSuccess, onError, onSettled

                - ใช้คู่กับ form submission

                - ช่วย revalidate ข้อมูล (ผ่าน queryClientinvalidateQueries() ถ้าต้อง refresh)
            - ตัวอย่างการใช้งาน
                export const useCreateLeave = () => {
                return useMutation({
                    mutationFn: async (input: AddLeaveInput) => {
                    const res = await fetch('/api/leaves', {
                        method: 'POST',
                        body: JSON.stringify(input),
                    });
                    const leave = await (res.json() as Promise<LeaveDetails>);
                    return leave;
                    },
                });
                };
                // เรียกใช้งาน
                'use client';
                import LeaveForm from '@/features/leaves/components/LeaveForm';
                import { useCreateLeave } from '@/features/leaves/hooks/api';
                import { type AddLeaveInput } from '@/features/leaves/types';
                import { useRouter } from 'next/navigation';
                const CreateLeave = () => {
                const router = useRouter();
                const { mutateAsync } = useCreateLeave();
                const createLeave = async (form: AddLeaveInput) => {
                    await mutateAsync(form);
                    router.push('/leaves');
                };
                return <LeaveForm kind="create" onSubmit={createLeave}></LeaveForm>;
                };
                export default CreateLeave;
    Zustand
        - zustand ใช้สำหรับจัดการ state managment
        - หลัหกการทำงานของ zustand จะเก็บ state ไว้ใน store
        - การใช้งาน
            - ใช้คำสั่ง npm add zustand immer เพื่อติดตั้ง
            สามารถใช้งานได้ 2 รูปแบบ
            1.multiple store
            2.single store
        - ตัวอย่าง func ที่สร้างเพื่อเรียกใช้งาน toast
            import { create } from 'zustand';
            import { immer } from 'zustand/middleware/immer';
            import { devtools } from 'zustand/middleware';
            export interface UiState {
            toast: {
                type: 'Success' | 'Error';
                message: string;
            } | null;
            setToast: (toast: UiState['toast']) => void;
            clearToast: () => void;
            }
            export const useUiStore = create<UiState>()(
            immer(
                devtools((set, get) => {
                return {
                    toast: null,
                    setToast(toast) {
                    set(
                        (state) => {
                        state.toast = toast;
                        },
                        false,
                        { type: 'ui/setToast', toast },
                    );
                    setTimeout(() => {
                        get().clearToast();
                    }, 3_000);
                    },
                    clearToast() {
                    set(
                        (state) => {
                        state.toast = null;
                        },
                        false,
                        { type: 'ui/clearToast' },
                    );
                    },
                };
                }),
            ),
            );

            // การเรียกใช้งาน func 
            const [toast, clearToast] = useUiStore(
                useShallow((state) => [state.toast, state.clearToast]),
            );
            return <div>;
                {toast.message}
                <button type="button" onClick={clearToast}>
            </div>
    -Auth.js
        - ตัว libary นี้ จะช่วยการล็อคอินหรือล็อคเอ้าท์ให้ และจะสร้าง cookie ไว้ให้เรียกใช้งาน ข้อมูลผู้ใช้
        - การติดตั้งดูได้ที่ doc Auth.js
        - การใช้งาน
            1. สร้างการตั้งค่า
                export const authOptions = {
                session: {
                    strategy: 'jwt',
                },
                callbacks: {
                    jwt({ token, user, session, trigger }) {
                    if (trigger === 'update' && isUpdateSessionData(session)) {
                        if (session.image) token.picture = session.image;
                        if (session.name) token.name = session.name;
                        if (session.email) token.email = session.email;
                    }

                    if (user) {
                        token.sub = user.id;
                        token.email = user.email;
                        token.role = user.role;
                        token.name = user.name;
                        token.picture = user.image;
                    }

                    return token;
                    },
                    session({ session, token }) {
                    return {
                        ...session,
                        user: {
                        ...session.user,
                        id: token.sub,
                        role: token.role,
                        name: token.name,
                        email: token.email,
                        image: token.picture,
                        },
                    };
                    },
                },
                providers: [
                    CredentialsProvider({
                    name: 'Credentials',
                    credentials: {
                        email: { label: 'Email', type: 'email ' },
                        password: { label: 'Password', type: 'password' },
                    },
                    async authorize(credentials) {
                        const user = await db.user.findUnique({
                        where: {
                            email: credentials?.email,
                        },
                        });

                        if (!user) return null;
                        if (!credentials?.password) return null;
                        if (!(await bcrypt.compare(credentials.password, user.password))) {
                        return null;
                        }

                        return { ...user, id: user.id.toString() };
                    },
                    }),
                ],
                } satisfies NextAuthOptions;

            1. เรียกใช้การตั้งค่า
                import { authOptions } from '@/features/auth/auth';
                import NextAuth from 'next-auth';
                const handler = NextAuth(authOptions);
                export { handler as GET, handler as POST };
            2. การใช้งาน signIn 
                "use client"
                import { signIn } from "next-auth/react"
                
                export function SignIn() {
                return (
                    <button onClick={() => signIn("github", { redirectTo: "/dashboard" })}>
                    Sign In
                    </button>
                )
                }
            3. การดึงค่าสถานะมาจาก auth.js
                ใช้ตัว hook ที่ชื่อว่า useSeession()
                ยกตัวอย่างเช่น
                    "use client"
                    import { useSession } from "next-auth/react"
                    
                    export default function Dashboard() {
                    const { data: session } = useSession()
                    
                    if (session?.user?.role === "admin") {
                        return <p>You are an admin, welcome!</p>
                    }
                    
                    return <p>You are not authorized to view this page!</p>
                    }
    - Pararel route หรือ steaming ข้อมูล
        - ปัญหาที่เราจะต้องใช้การใช้งานของตัว parararl นั้น จะใช้ในกรณีที่ ในหนึ่งหน้า page จะต้องโหลดข้อมูลหลายอย่างพร้อมๆกัน
        - การทำงานแบบแยกการทำงานทำได้ หรือการทำงานแบบขนาน
            - สร้าง folder ที่ขึ้นต้นด้วย @ ตัว react จะมองว่ามันคือ Parallel 
        - ตัวอย่าง มันจะมองเป็น ReactNode ตัวหนึ่งและเรียกใช้งานได้เลย
            import { type ReactNode } from 'react';
            interface HomeLayoutProps {
            children: ReactNode;
            articles: ReactNode;
            announcements: ReactNode;
            }
            const HomeLayout = ({ children, articles, announcements }: HomeLayoutProps) => {
            return (
                <>
                {children}
                {articles}
                {announcements}
                </>
            );
            };
            export default HomeLayout;
        - ตัว react สามารถสร้างไฟล์ error.tsx หรือ loading.tsx หรือ not-found.tsx ได้เพื่อจัดการการโหลดข้อมูล โดยที่ไม่ต้อง import เข้าหน้าที่ต้องการ
    - intersect route 
        - ปัญหาที่เราจะต้องใช้คือ 
        - การแทรกกลางของการแสดงผล
        . (.)[id]
    - SEO
        - กำหนด title
            -meta data หรือที่ใส่ title ใน layout.tsx จะเป็นตัวบอกว่าหน้าเว็บเราชื่ออะไร
        - กำหนด favicon
        - กำหนดการแสดงผลจากการที่เราส่ง url ไปและต้องแสดง preview
        - path ที่หาตาม id ควรหาคา slug มากกว่า
        - set robot and sitemap เพื่อบ่งบอกว่า link ต่างๆที่อยู่หกน้าเว็บของเรามใีลักษณะเป็นอย่างไร