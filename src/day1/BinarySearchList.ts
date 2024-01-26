export default function bs_list(haystack: number[], needle: number): boolean {
    let low = 0;
    let high = haystack.length;

    while (low < high) {
        const mid = Math.floor(low + (high - low) / 2);
        const value = haystack[mid];
        if (value === needle) {
            return true;
        }

        if (value < needle) {
            low = mid + 1;
        } else {
            high = mid;
        }
    }

    return false;
}
