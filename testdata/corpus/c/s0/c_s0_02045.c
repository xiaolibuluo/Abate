
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1381_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1381_0(int value) {
    return gcd1381_0(value, 38) * 13;
}
enum Kind1381_1 { KIND_EMPTY1381_1 = 0, KIND_FULL1381_1 = 20 };

union Cell1381_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1381_1(union Cell1381_1 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1381_1 : KIND_EMPTY1381_1;
}
struct Node1381_2 {
    int value;
    struct Node1381_2 *next;
};

int length1381_2(const struct Node1381_2 *head) {
    int count = 0;
    const struct Node1381_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 27;
}
#if defined(__STDC_VERSION__)
int variant1381_3(int value, ...) {
    return value + 34;
}
#else
int variant1381_3(int value) {
    return value + 34;
}
#endif
