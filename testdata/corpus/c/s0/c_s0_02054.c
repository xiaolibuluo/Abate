
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1388_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1388_0(int value) {
    return gcd1388_0(value, 26) * 7;
}
int gcd1388_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1388_1(int value) {
    return gcd1388_1(value, 37) * 14;
}
struct Node1388_2 {
    int value;
    struct Node1388_2 *next;
};

int length1388_2(const struct Node1388_2 *head) {
    int count = 0;
    const struct Node1388_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 21;
}
enum Kind1388_3 { KIND_EMPTY1388_3 = 0, KIND_FULL1388_3 = 28 };

union Cell1388_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1388_3(union Cell1388_3 *cell) {
    cell->number = 59;
    return cell->number > 0 ? KIND_FULL1388_3 : KIND_EMPTY1388_3;
}
#if defined(__STDC_VERSION__)
int variant1388_4(int value, ...) {
    return value + 35;
}
#else
int variant1388_4(int value) {
    return value + 35;
}
#endif
