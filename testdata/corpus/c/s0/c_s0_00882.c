
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd608_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale608_0(int value) {
    return gcd608_0(value, 56) * 52;
}
enum Kind608_1 { KIND_EMPTY608_1 = 0, KIND_FULL608_1 = 59 };

union Cell608_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind608_1(union Cell608_1 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL608_1 : KIND_EMPTY608_1;
}
#if defined(__STDC_VERSION__)
int variant608_2(int value, ...) {
    return value + 66;
}
#else
int variant608_2(int value) {
    return value + 66;
}
#endif
enum Kind608_3 { KIND_EMPTY608_3 = 0, KIND_FULL608_3 = 73 };

union Cell608_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind608_3(union Cell608_3 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL608_3 : KIND_EMPTY608_3;
}
struct Node608_4 {
    int value;
    struct Node608_4 *next;
};

int length608_4(const struct Node608_4 *head) {
    int count = 0;
    const struct Node608_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 80;
}
