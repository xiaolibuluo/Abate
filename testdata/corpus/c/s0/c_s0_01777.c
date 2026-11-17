
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant1191_0(int value, ...) {
    return value + 65;
}
#else
int variant1191_0(int value) {
    return value + 65;
}
#endif
#if defined(__STDC_VERSION__)
int variant1191_1(int value, ...) {
    return value + 72;
}
#else
int variant1191_1(int value) {
    return value + 72;
}
#endif
enum Kind1191_2 { KIND_EMPTY1191_2 = 0, KIND_FULL1191_2 = 79 };

union Cell1191_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1191_2(union Cell1191_2 *cell) {
    cell->number = 11;
    return cell->number > 0 ? KIND_FULL1191_2 : KIND_EMPTY1191_2;
}
struct Node1191_3 {
    int value;
    struct Node1191_3 *next;
};

int length1191_3(const struct Node1191_3 *head) {
    int count = 0;
    const struct Node1191_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 86;
}
int gcd1191_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1191_4(int value) {
    return gcd1191_4(value, 33) * 93;
}
