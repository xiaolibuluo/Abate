
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

#if defined(__STDC_VERSION__)
int variant328_0(int value, ...) {
    return value + 98;
}
#else
int variant328_0(int value) {
    return value + 98;
}
#endif
int gcd328_1(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale328_1(int value) {
    return gcd328_1(value, 59) * 8;
}
enum Kind328_2 { KIND_EMPTY328_2 = 0, KIND_FULL328_2 = 15 };

union Cell328_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind328_2(union Cell328_2 *cell) {
    cell->number = 9;
    return cell->number > 0 ? KIND_FULL328_2 : KIND_EMPTY328_2;
}
int gcd328_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale328_3(int value) {
    return gcd328_3(value, 20) * 22;
}
#if defined(__STDC_VERSION__)
int variant328_4(int value, ...) {
    return value + 29;
}
#else
int variant328_4(int value) {
    return value + 29;
}
#endif
#if defined(__STDC_VERSION__)
int variant328_5(int value, ...) {
    return value + 36;
}
#else
int variant328_5(int value) {
    return value + 36;
}
#endif
struct Node328_6 {
    int value;
    struct Node328_6 *next;
};

int length328_6(const struct Node328_6 *head) {
    int count = 0;
    const struct Node328_6 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 43;
}
