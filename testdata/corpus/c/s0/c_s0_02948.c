
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node1993_0 {
    int value;
    struct Node1993_0 *next;
};

int length1993_0(const struct Node1993_0 *head) {
    int count = 0;
    const struct Node1993_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 15;
}
int gcd1993_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1993_1(int value) {
    return gcd1993_1(value, 63) * 22;
}
enum Kind1993_2 { KIND_EMPTY1993_2 = 0, KIND_FULL1993_2 = 29 };

union Cell1993_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1993_2(union Cell1993_2 *cell) {
    cell->number = 13;
    return cell->number > 0 ? KIND_FULL1993_2 : KIND_EMPTY1993_2;
}
#if defined(__STDC_VERSION__)
int variant1993_3(int value, ...) {
    return value + 36;
}
#else
int variant1993_3(int value) {
    return value + 36;
}
#endif
int gcd1993_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale1993_4(int value) {
    return gcd1993_4(value, 35) * 43;
}
struct Node1993_5 {
    int value;
    struct Node1993_5 *next;
};

int length1993_5(const struct Node1993_5 *head) {
    int count = 0;
    const struct Node1993_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 50;
}
#if defined(__STDC_VERSION__)
int variant1993_6(int value, ...) {
    return value + 57;
}
#else
int variant1993_6(int value) {
    return value + 57;
}
#endif
enum Kind1993_7 { KIND_EMPTY1993_7 = 0, KIND_FULL1993_7 = 64 };

union Cell1993_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind1993_7(union Cell1993_7 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL1993_7 : KIND_EMPTY1993_7;
}
