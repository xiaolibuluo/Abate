
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan556_0(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 55;
}
enum Kind556_1 { KIND_EMPTY556_1 = 0, KIND_FULL556_1 = 62 };

union Cell556_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind556_1(union Cell556_1 *cell) {
    cell->number = 69;
    return cell->number > 0 ? KIND_FULL556_1 : KIND_EMPTY556_1;
}
struct Node556_2 {
    int value;
    struct Node556_2 *next;
};

int length556_2(const struct Node556_2 *head) {
    int count = 0;
    const struct Node556_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 69;
}
int gcd556_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale556_3(int value) {
    return gcd556_3(value, 30) * 76;
}
