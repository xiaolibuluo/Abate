
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

enum Kind2370_0 { KIND_EMPTY2370_0 = 0, KIND_FULL2370_0 = 66 };

union Cell2370_0 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2370_0(union Cell2370_0 *cell) {
    cell->number = 68;
    return cell->number > 0 ? KIND_FULL2370_0 : KIND_EMPTY2370_0;
}
struct Node2370_1 {
    int value;
    struct Node2370_1 *next;
};

int length2370_1(const struct Node2370_1 *head) {
    int count = 0;
    const struct Node2370_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 73;
}
enum Kind2370_2 { KIND_EMPTY2370_2 = 0, KIND_FULL2370_2 = 80 };

union Cell2370_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2370_2(union Cell2370_2 *cell) {
    cell->number = 29;
    return cell->number > 0 ? KIND_FULL2370_2 : KIND_EMPTY2370_2;
}
int gcd2370_3(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2370_3(int value) {
    return gcd2370_3(value, 40) * 87;
}
#if defined(__STDC_VERSION__)
int variant2370_4(int value, ...) {
    return value + 94;
}
#else
int variant2370_4(int value) {
    return value + 94;
}
#endif
struct Node2370_5 {
    int value;
    struct Node2370_5 *next;
};

int length2370_5(const struct Node2370_5 *head) {
    int count = 0;
    const struct Node2370_5 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 101;
}
