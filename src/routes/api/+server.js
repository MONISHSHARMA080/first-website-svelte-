export function GET() {
	const number = Math.floor(Math.random() * 10) + 1;
    /**
     * @type {any}
     */
    const a = process.cwd()
        

	return new Response(a, {
		headers: {
			'Content-Type': 'application/json'
		}
	});
}