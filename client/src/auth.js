let code;

export function getCode() {
	return code;
}

export function setCode(value) {
	value = String(value);

	if (!/^\d{4}$/.test(value)) {
		throw Error("Provided code has an invalid format.");
	}

	code = value;
}