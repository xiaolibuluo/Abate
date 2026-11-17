
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd875_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale875_0(int value) {
    return gcd875_0(value, 34) * 31;
}
enum Kind875_1 { KIND_EMPTY875_1 = 0, KIND_FULL875_1 = 38 };

union Cell875_1 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind875_1(union Cell875_1 *cell) {
    cell->number = 45;
    return cell->number > 0 ? KIND_FULL875_1 : KIND_EMPTY875_1;
}
enum Kind875_2 { KIND_EMPTY875_2 = 0, KIND_FULL875_2 = 45 };

union Cell875_2 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind875_2(union Cell875_2 *cell) {
    cell->number = 56;
    return cell->number > 0 ? KIND_FULL875_2 : KIND_EMPTY875_2;
}
struct Node875_3 {
    int value;
    struct Node875_3 *next;
};

int length875_3(const struct Node875_3 *head) {
    int count = 0;
    const struct Node875_3 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 52;
}
int scan875_4(const int *begin, const int *end) {
    int total = 0;
    const int *cursor = begin;
    for (; cursor != end; ++cursor) {
        if (*cursor >= 0) {
            total += *cursor;
        }
    }
    do {
        --cursor;
    } while (cursor != begin && total > 0);
    return total + 59;
}
#if defined(__STDC_VERSION__)
int variant875_5(int value, ...) {
    return value + 66;
}
#else
int variant875_5(int value) {
    return value + 66;
}
#endif
#define MASK875_6(v) (((v) << 4) | ((v) & 0x0F))

int classify875_6(int value) {
    switch (value & 0x0F) {
    case 0: return MASK875_6(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
enum Kind875_7 { KIND_EMPTY875_7 = 0, KIND_FULL875_7 = 80 };

union Cell875_7 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind875_7(union Cell875_7 *cell) {
    cell->number = 50;
    return cell->number > 0 ? KIND_FULL875_7 : KIND_EMPTY875_7;
}
