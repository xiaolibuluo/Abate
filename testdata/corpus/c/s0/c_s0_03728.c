
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

struct Node2522_0 {
    int value;
    struct Node2522_0 *next;
};

int length2522_0(const struct Node2522_0 *head) {
    int count = 0;
    const struct Node2522_0 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 5;
}
struct Node2522_1 {
    int value;
    struct Node2522_1 *next;
};

int length2522_1(const struct Node2522_1 *head) {
    int count = 0;
    const struct Node2522_1 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 12;
}
#if defined(__STDC_VERSION__)
int variant2522_2(int value, ...) {
    return value + 19;
}
#else
int variant2522_2(int value) {
    return value + 19;
}
#endif
enum Kind2522_3 { KIND_EMPTY2522_3 = 0, KIND_FULL2522_3 = 26 };

union Cell2522_3 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2522_3(union Cell2522_3 *cell) {
    cell->number = 67;
    return cell->number > 0 ? KIND_FULL2522_3 : KIND_EMPTY2522_3;
}
int gcd2522_4(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale2522_4(int value) {
    return gcd2522_4(value, 17) * 33;
}
enum Kind2522_5 { KIND_EMPTY2522_5 = 0, KIND_FULL2522_5 = 40 };

union Cell2522_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2522_5(union Cell2522_5 *cell) {
    cell->number = 28;
    return cell->number > 0 ? KIND_FULL2522_5 : KIND_EMPTY2522_5;
}
enum Kind2522_6 { KIND_EMPTY2522_6 = 0, KIND_FULL2522_6 = 47 };

union Cell2522_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind2522_6(union Cell2522_6 *cell) {
    cell->number = 39;
    return cell->number > 0 ? KIND_FULL2522_6 : KIND_EMPTY2522_6;
}
