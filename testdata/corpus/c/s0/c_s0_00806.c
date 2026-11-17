
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node557_0 {
    int value;
    struct Node557_0 *next;
};

int length557_0(const struct Node557_0 *head) {
    int count = 0;
    const struct Node557_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 68;
}
enum Kind557_1 { KIND_EMPTY557_1 = 0, KIND_FULL557_1 = 75 };

union Cell557_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind557_1(union Cell557_1 *cell) {
    cell->number = 15;
    return cell->number > 0 ? KIND_FULL557_1 : KIND_EMPTY557_1;
}
#if defined(__STDC_VERSION__)
int variant557_2(int value, ...) {
    return value + 82;
}
#else
int variant557_2(int value) {
    return value + 82;
}
#endif
#if defined(__STDC_VERSION__)
int variant557_3(int value, ...) {
    return value + 89;
}
#else
int variant557_3(int value) {
    return value + 89;
}
#endif
int gcd557_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale557_4(int value) {
    return gcd557_4(value, 48) * 96;
}
