
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd1545_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1545_0(int value) {
    return gcd1545_0(value, 27) * 11;
}
struct Node1545_1 {
    int value;
    struct Node1545_1 *next;
};

int length1545_1(const struct Node1545_1 *head) {
    int count = 0;
    const struct Node1545_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 18;
}
enum Kind1545_2 { KIND_EMPTY1545_2 = 0, KIND_FULL1545_2 = 25 };

union Cell1545_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1545_2(union Cell1545_2 *cell) {
    cell->number = 49;
    return cell->number > 0 ? KIND_FULL1545_2 : KIND_EMPTY1545_2;
}
#if defined(__STDC_VERSION__)
int variant1545_3(int value, ...) {
    return value + 32;
}
#else
int variant1545_3(int value) {
    return value + 32;
}
#endif
struct Node1545_4 {
    int value;
    struct Node1545_4 *next;
};

int length1545_4(const struct Node1545_4 *head) {
    int count = 0;
    const struct Node1545_4 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 39;
}
#if defined(__STDC_VERSION__)
int variant1545_5(int value, ...) {
    return value + 46;
}
#else
int variant1545_5(int value) {
    return value + 46;
}
#endif
enum Kind1545_6 { KIND_EMPTY1545_6 = 0, KIND_FULL1545_6 = 53 };

union Cell1545_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1545_6(union Cell1545_6 *cell) {
    cell->number = 32;
    return cell->number > 0 ? KIND_FULL1545_6 : KIND_EMPTY1545_6;
}
