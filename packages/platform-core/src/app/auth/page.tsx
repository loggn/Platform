import Image from 'next/image'

const LineInput = ({ placeholder }: { placeholder: string }) => {
	return (
		<input
			type="text"
			className="border-b-2 border-black focus:border-b-2 transition-all"
			placeholder={placeholder}
		/>
	)
}

const Auth = () => {
	return (
		<div className="w-screen h-screen flex justify-center items-center">
			<div className="bg fixed right-[200px] top-40 text-[400px] -z-10">
				Hello
			</div>
			<div className="bg fixed right-[600px] bottom-40 text-[400px] -z-10">
				World
			</div>
			<Image
				width={400}
				height={600}
				className="fixed left-0 bottom-0"
				src="/Kotone_Fujita.webp"
				alt="琴音"
			/>
			<div className="login-box w-[750px] h-[450px] flex items-center rounded overflow-hidden shadow-xl hover:shadow-2xl hover:scale-[102%] transition-all z-10">
				<div className="left h-full">
					<Image
						width={300}
						height={400}
						// layout="responsive"
						src="/美游.jpg"
						alt="美游"
						style={{
							width: 'auto',
							height: '100%',
						}}
					/>
				</div>
				<div className="right flex flex-col justify-center flex-grow h-full items-center gap-4 px-8 bg-white">
					<div className="self-start text-2xl mb-8 ml-10 font-black">
						Login
					</div>
					<LineInput placeholder="Username" />
					<LineInput placeholder="Password" />
					<div className="self-end text-sm gap-2 flex">
						<a href="/register">忘记密码？</a>
						<a href="/register">注册账号</a>
					</div>
					<button className="bg-black text-white rounded p-2 w-3/5">
						登录
					</button>
				</div>
			</div>
		</div>
	)
}

export default Auth
