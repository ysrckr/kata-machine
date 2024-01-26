export default function linear_search(
    haystack: number[],
    needle: number,
): boolean {
    let i = 0;
    while (i < haystack.length) {
        if (needle === haystack[i]) {
            return true;
        }
        i++;
    }
    return false;
}
