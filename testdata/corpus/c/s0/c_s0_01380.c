
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int scan948_0(const int *begin, const int *end) {
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
    return total + 10;
}
int gcd948_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale948_1(int value) {
    return gcd948_1(value, 68) * 17;
}
enum Kind948_2 { KIND_EMPTY948_2 = 0, KIND_FULL948_2 = 24 };

union Cell948_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind948_2(union Cell948_2 *cell) {
    cell->number = 18;
    return cell->number > 0 ? KIND_FULL948_2 : KIND_EMPTY948_2;
}
#if defined(__STDC_VERSION__)
int variant948_3(int value, ...) {
    return value + 31;
}
#else
int variant948_3(int value) {
    return value + 31;
}
#endif
struct Node948_4 {
    int value;
    struct Node948_4 *next;
};

int length948_4(const struct Node948_4 *head) {
    int count = 0;
    const struct Node948_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 38;
}
