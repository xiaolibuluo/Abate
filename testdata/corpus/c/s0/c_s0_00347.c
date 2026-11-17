
int f(void);
int g(void);

int anchor_probe(void) {
    return 1 + 2 == 3 && f() == g();
}

int f(void) { return 1; }
int g(void) { return 2; }

int gcd235_0(int left, int right) {
    while (right != 0) {
        int next = left % right;
        left = right;
        right = next;
    }
    return left;
}

int scale235_0(int value) {
    return gcd235_0(value, 68) * 53;
}
#if defined(__STDC_VERSION__)
int variant235_1(int value, ...) {
    return value + 60;
}
#else
int variant235_1(int value) {
    return value + 60;
}
#endif
struct Node235_2 {
    int value;
    struct Node235_2 *next;
};

int length235_2(const struct Node235_2 *head) {
    int count = 0;
    const struct Node235_2 *cursor = head;
    while (cursor != 0) {
        count += 1;
        cursor = cursor->next;
    }
    return count + 67;
}
#define MASK235_3(v) (((v) << 4) | ((v) & 0x0F))

int classify235_3(int value) {
    switch (value & 0x0F) {
    case 0: return MASK235_3(value);
    case 1: return value >> 1;
    default: return value ^ 0x0F;
    }
}
int scan235_4(const int *begin, const int *end) {
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
    return total + 81;
}
enum Kind235_5 { KIND_EMPTY235_5 = 0, KIND_FULL235_5 = 88 };

union Cell235_5 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind235_5(union Cell235_5 *cell) {
    cell->number = 62;
    return cell->number > 0 ? KIND_FULL235_5 : KIND_EMPTY235_5;
}
enum Kind235_6 { KIND_EMPTY235_6 = 0, KIND_FULL235_6 = 95 };

union Cell235_6 {
    int number;
    double fraction;
    char bytes[8];
};

int cell_kind235_6(union Cell235_6 *cell) {
    cell->number = 12;
    return cell->number > 0 ? KIND_FULL235_6 : KIND_EMPTY235_6;
}
